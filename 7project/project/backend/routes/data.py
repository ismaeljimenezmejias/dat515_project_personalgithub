from flask import Blueprint, jsonify, request, session
from db import get_connection
from datetime import datetime
from werkzeug.security import generate_password_hash, check_password_hash

data_bp = Blueprint('data', __name__)


def _require_valid_user():
    """
    Ensure the session user exists in DB. This handles the case where the
    database was reset after the user logged in and the session still holds
    a stale user_id. If a session name exists, try to (re)create the user.
    Returns: (user_id, error_response)
      - user_id: int when valid, else None
      - error_response: a Flask response tuple (json, status) when invalid
    """
    uid = session.get('user_id')
    name = session.get('user_name')
    if not uid:
        return None, (jsonify({'error': 'Authentication required'}), 401)
    try:
        conn = get_connection()
        cur = conn.cursor()
        cur.execute('SELECT id FROM users WHERE id=%s', (uid,))
        row = cur.fetchone()
        if row:
            conn.close()
            return uid, None
        # user id missing in DB -> try recreate if we have the name
        if name:
            cur.execute('INSERT INTO users (name, created_at) VALUES (%s, %s) RETURNING id', (name, datetime.now()))
            new_id = cur.fetchone()[0]
            conn.commit()
            conn.close()
            session['user_id'] = new_id
            return new_id, None
        conn.close()
    except Exception:
        # fall through to auth required
        pass
    session.clear()
    return None, (jsonify({'error': 'Authentication required'}), 401)


@data_bp.route('/api/data', methods=['GET', 'POST'])
@data_bp.route('/api/bikes', methods=['GET', 'POST'])
def bikes():
    try:
        conn = get_connection()
        cursor = conn.cursor()

        if request.method == 'POST':
            # Require login
            user_id, err = _require_valid_user()
            if err:
                return err
            payload = request.get_json(silent=True) or {}

            # Required/optional fields
            title = (payload.get('title') or '').strip()
            sale_price = payload.get('sale_price')
            rental_price = payload.get('rental_price')
            sale_type = payload.get('sale_type', 'venta')
            model = (payload.get('model') or '').strip()
            description = (payload.get('description') or '').strip()
            bike_condition = (payload.get('condition') or '').strip()
            image_url = (payload.get('image_url') or '').strip() or None
            location_name = (payload.get('location_name') or '').strip() or None
            latitude = payload.get('latitude')
            longitude = payload.get('longitude')

            if not title:
                return jsonify({'error': 'Title is required'}), 400

            # Auto-calc rental price when needed
            if sale_type in ['alquiler', 'ambos'] and sale_price and not rental_price:
                rental_price = float(sale_price) * 0.15

            cursor.execute(
                     """INSERT INTO bikes (title, sale_price, rental_price, sale_type, model, description, bike_condition, image_url, owner_id, created_at, location_name, latitude, longitude)
                         VALUES (%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)""",
                     (title, sale_price, rental_price, sale_type, model, description, bike_condition, image_url, user_id, datetime.now(), location_name, latitude, longitude)
            )
            conn.commit()
            return jsonify({'success': True}), 201

        # GET bikes with optional filters
        def _to_float(val):
            try:
                return float(val)
            except (TypeError, ValueError):
                return None

        filter_type = request.args.get('type')
        min_sale = _to_float(request.args.get('minSalePrice'))
        max_sale = _to_float(request.args.get('maxSalePrice'))
        min_rent = _to_float(request.args.get('minRentalPrice'))
        max_rent = _to_float(request.args.get('maxRentalPrice'))
        exclude_mine = request.args.get('excludeMine') in ('1', 'true', 'True')
        user_id_ex = session.get('user_id') if exclude_mine else None
        search_text = (request.args.get('search') or '').strip()

        base_sql = (
            "SELECT b.id, b.title, b.sale_price, b.rental_price, b.sale_type, b.model, "
            "b.description, b.bike_condition, b.created_at, b.owner_id, u.name as owner_name, b.image_url, b.location_name, b.latitude, b.longitude "
            "FROM bikes b LEFT JOIN users u ON b.owner_id = u.id"
        )
        where = []
        params = []
        if filter_type and filter_type in ['venta', 'alquiler']:
            where.append("(b.sale_type = %s OR b.sale_type = 'ambos')")
            params.append(filter_type)
        if user_id_ex:
            where.append("(b.owner_id IS NULL OR b.owner_id <> %s)")
            params.append(user_id_ex)

        if search_text:
            where.append("(b.title LIKE %s OR b.description LIKE %s OR u.name LIKE %s)")
            like = f"%{search_text}%"
            params.extend([like, like, like])

        # Price filters
        if filter_type == 'venta':
            if min_sale is not None:
                where.append("b.sale_price IS NOT NULL AND b.sale_price >= %s")
                params.append(min_sale)
            if max_sale is not None:
                where.append("b.sale_price IS NOT NULL AND b.sale_price <= %s")
                params.append(max_sale)
        elif filter_type == 'alquiler':
            if min_rent is not None:
                where.append("b.rental_price IS NOT NULL AND b.rental_price >= %s")
                params.append(min_rent)
            if max_rent is not None:
                where.append("b.rental_price IS NOT NULL AND b.rental_price <= %s")
                params.append(max_rent)
        else:
            price_clauses = []
            price_params = []
            sale_parts = []
            if min_sale is not None:
                sale_parts.append("b.sale_price >= %s")
                price_params.append(min_sale)
            if max_sale is not None:
                sale_parts.append("b.sale_price <= %s")
                price_params.append(max_sale)
            if sale_parts:
                price_clauses.append("(b.sale_price IS NOT NULL AND " + " AND ".join(sale_parts) + ")")

            rent_parts = []
            if min_rent is not None:
                rent_parts.append("b.rental_price >= %s")
                price_params.append(min_rent)
            if max_rent is not None:
                rent_parts.append("b.rental_price <= %s")
                price_params.append(max_rent)
            if rent_parts:
                price_clauses.append("(b.rental_price IS NOT NULL AND " + " AND ".join(rent_parts) + ")")

            if price_clauses:
                where.append("(" + " OR ".join(price_clauses) + ")")
                params.extend(price_params)

        if where:
            base_sql += " WHERE " + " AND ".join(where)
        base_sql += " ORDER BY b.created_at DESC"

        cursor.execute(base_sql, tuple(params))
        rows = cursor.fetchall()
        bikes = []
        for row in rows:
            created_val = row[8]
            created_str = created_val.isoformat() if hasattr(created_val, 'isoformat') else str(created_val) if created_val else None
            bikes.append({
                'id': row[0],
                'title': row[1],
                'sale_price': float(row[2]) if row[2] is not None else None,
                'rental_price': float(row[3]) if row[3] is not None else None,
                'sale_type': row[4],
                'model': row[5],
                'description': row[6],
                'condition': row[7],
                'created_at': created_str,
                'owner_id': row[9],
                'owner_name': row[10],
                'image_url': row[11],
                'location_name': row[12],
                'latitude': float(row[13]) if row[13] is not None else None,
                'longitude': float(row[14]) if row[14] is not None else None,
            })

        conn.close()
        return jsonify({'bikes': bikes, 'count': len(bikes)})
    except Exception as e:
        return jsonify({'error': str(e)}), 500


@data_bp.route('/api/me', methods=['GET'])
def me():
    return jsonify({
        'authenticated': bool(session.get('user_id')),
        'user': {
            'id': session.get('user_id'),
            'name': session.get('user_name')
        } if session.get('user_id') else None
    })


@data_bp.route('/api/signup', methods=['POST'])
def signup():
    payload = request.get_json(silent=True) or {}
    name = (payload.get('name') or '').strip()
    password = (payload.get('password') or '').strip()
    if not name or not password:
        return jsonify({'error': 'name and password are required'}), 400
    try:
        conn = get_connection()
        cur = conn.cursor()
        cur.execute('SELECT id FROM users WHERE name=%s', (name,))
        if cur.fetchone():
            conn.close()
            return jsonify({'error': 'User already exists'}), 409
        pwd_hash = generate_password_hash(password)
        cur.execute('INSERT INTO users (name, password_hash, created_at) VALUES (%s, %s, %s) RETURNING id', (name, pwd_hash, datetime.now()))
        user_id = cur.fetchone()[0]
        conn.commit()
        conn.close()
        session['user_id'] = user_id
        session['user_name'] = name
        return jsonify({'ok': True, 'user': {'id': user_id, 'name': name}}), 201
    except Exception as e:
        return jsonify({'error': str(e)}), 500


@data_bp.route('/api/login', methods=['POST'])
def login():
    payload = request.get_json(silent=True) or {}
    name = (payload.get('name') or '').strip()
    password = (payload.get('password') or '').strip()
    if not name:
        return jsonify({'error': 'Name is required'}), 400
    if not password:
        return jsonify({'error': 'Password is required'}), 400
    try:
        conn = get_connection()
        cur = conn.cursor()
        cur.execute('SELECT id, password_hash FROM users WHERE name=%s', (name,))
        row = cur.fetchone()
        if not row:
            conn.close()
            return jsonify({'error': 'User not found. Please sign up.'}), 404
        user_id, pwd_hash = row[0], row[1]
        if pwd_hash:
            if not check_password_hash(pwd_hash, password):
                return jsonify({'error': 'Invalid credentials'}), 401
        else:
            # Legacy user with no password: require one now and set it
            new_hash = generate_password_hash(password)
            cur2 = conn.cursor()
            cur2.execute('UPDATE users SET password_hash=%s WHERE id=%s', (new_hash, user_id))
            conn.commit()
        conn.close()
        session['user_id'] = user_id
        session['user_name'] = name
        return jsonify({'ok': True, 'user': {'id': user_id, 'name': name}})
    except Exception as e:
        return jsonify({'error': str(e)}), 500


@data_bp.route('/api/logout', methods=['POST'])
def logout():
    session.clear()
    return jsonify({'ok': True})


@data_bp.route('/api/bikes/<int:bike_id>', methods=['DELETE'])
def delete_bike(bike_id):
    user_id, err = _require_valid_user()
    if err:
        return err
    try:
        conn = get_connection()
        cursor = conn.cursor()
        cursor.execute('SELECT owner_id FROM bikes WHERE id = %s', (bike_id,))
        row = cursor.fetchone()
        if not row:
            return jsonify({'error': 'Bike not found'}), 404
        if row[0] != user_id:
            return jsonify({'error': 'Not authorized'}), 403
        cursor.execute('DELETE FROM bikes WHERE id = %s', (bike_id,))
        conn.commit()
        conn.close()
        return jsonify({'ok': True})
    except Exception as e:
        return jsonify({'error': str(e)}), 500


@data_bp.route('/api/bikes/<int:bike_id>', methods=['PUT'])
def update_bike(bike_id):
    user_id, err = _require_valid_user()
    if err:
        return err
    try:
        payload = request.get_json(silent=True) or {}
        title = (payload.get('title') or '').strip()
        sale_price = payload.get('sale_price')
        rental_price = payload.get('rental_price')
        sale_type = payload.get('sale_type', 'venta')
        model = (payload.get('model') or '').strip()
        description = (payload.get('description') or '').strip()
        bike_condition = (payload.get('condition') or '').strip()
    image_url = (payload.get('image_url') or '').strip() or None
    location_name = (payload.get('location_name') or '').strip() or None
    latitude = payload.get('latitude')
    longitude = payload.get('longitude')

        if not title:
            return jsonify({'error': 'Title is required'}), 400

        if sale_type in ['alquiler', 'ambos'] and sale_price and not rental_price:
            rental_price = float(sale_price) * 0.15

        conn = get_connection()
        cursor = conn.cursor()
        cursor.execute(
                """UPDATE bikes SET title=%s, sale_price=%s, rental_price=%s, sale_type=%s,
                    model=%s, description=%s, bike_condition=%s, image_url=%s, location_name=%s, latitude=%s, longitude=%s WHERE id=%s""",
                (title, sale_price, rental_price, sale_type, model, description, bike_condition, image_url, location_name, latitude, longitude, bike_id)
        )
        conn.commit()
        conn.close()
        return jsonify({'ok': True})
    except Exception as e:
        return jsonify({'error': str(e)}), 500


@data_bp.route('/api/messages', methods=['POST'])
def send_message():
    user_id, err = _require_valid_user()
    if err:
        return err
    try:
        payload = request.get_json(silent=True) or {}
        bike_id = payload.get('bike_id')
        receiver_id = payload.get('receiver_id')
        content = (payload.get('content') or '').strip()
        if not bike_id or not receiver_id or not content:
            return jsonify({'error': 'bike_id, receiver_id, and content are required'}), 400
        conn = get_connection()
        cursor = conn.cursor()
        # Validate receiver exists
        cursor.execute('SELECT id FROM users WHERE id=%s', (receiver_id,))
        if not cursor.fetchone():
            conn.close()
            return jsonify({'error': 'Receiver not found'}), 400
        cursor.execute('SELECT id FROM bikes WHERE id = %s', (bike_id,))
        if not cursor.fetchone():
            return jsonify({'error': 'Bike not found'}), 404
        cursor.execute(
            """INSERT INTO messages (bike_id, sender_id, receiver_id, content, created_at)
               VALUES (%s, %s, %s, %s, %s) RETURNING id""",
            (bike_id, user_id, receiver_id, content, datetime.now())
        )
        message_id = cursor.fetchone()[0]
        conn.commit()
        conn.close()
        return jsonify({'ok': True, 'message_id': message_id}), 201
    except Exception as e:
        return jsonify({'error': str(e)}), 500


@data_bp.route('/api/messages/bike/<int:bike_id>', methods=['GET'])
def get_bike_messages(bike_id):
    user_id = session.get('user_id')
    if not user_id:
        return jsonify({'error': 'Authentication required'}), 401
    try:
        conn = get_connection()
        cursor = conn.cursor()
        cursor.execute(
            """SELECT m.id, m.bike_id, m.sender_id, m.receiver_id, m.content, m.created_at,
                      sender.name as sender_name, receiver.name as receiver_name
               FROM messages m
               JOIN users sender ON m.sender_id = sender.id
               JOIN users receiver ON m.receiver_id = receiver.id
               WHERE m.bike_id = %s AND (m.sender_id = %s OR m.receiver_id = %s)
               ORDER BY m.created_at ASC""",
            (bike_id, user_id, user_id)
        )
        rows = cursor.fetchall()
        messages = []
        for row in rows:
            created_val = row[5]
            created_str = created_val.isoformat() if hasattr(created_val, 'isoformat') else str(created_val)
            messages.append({
                'id': row[0], 'bike_id': row[1], 'sender_id': row[2], 'receiver_id': row[3],
                'content': row[4], 'created_at': created_str, 'sender_name': row[6], 'receiver_name': row[7],
                'is_mine': row[2] == user_id
            })
        conn.close()
        return jsonify({'messages': messages})
    except Exception as e:
        return jsonify({'error': str(e)}), 500


@data_bp.route('/api/messages/conversations', methods=['GET'])
def get_conversations():
    user_id = session.get('user_id')
    if not user_id:
        return jsonify({'error': 'Authentication required'}), 401
    try:
        conn = get_connection()
        cursor = conn.cursor()
        cursor.execute(
            """SELECT DISTINCT 
                   m.bike_id,
                   b.title as bike_title,
                   b.image_url,
                   CASE WHEN m.sender_id = %s THEN m.receiver_id ELSE m.sender_id END as other_user_id,
                   u.name as other_user_name,
                   (SELECT content FROM messages m2 
                    WHERE m2.bike_id = m.bike_id 
                      AND (m2.sender_id = %s OR m2.receiver_id = %s)
                    ORDER BY m2.created_at DESC LIMIT 1) as last_message,
                   (SELECT created_at FROM messages m2 
                    WHERE m2.bike_id = m.bike_id 
                      AND (m2.sender_id = %s OR m2.receiver_id = %s)
                    ORDER BY m2.created_at DESC LIMIT 1) as last_message_at
               FROM messages m
               JOIN bikes b ON m.bike_id = b.id
               JOIN users u ON CASE WHEN m.sender_id = %s THEN m.receiver_id ELSE m.sender_id END = u.id
               WHERE m.sender_id = %s OR m.receiver_id = %s
               ORDER BY last_message_at DESC""",
            (user_id, user_id, user_id, user_id, user_id, user_id, user_id, user_id)
        )
        rows = cursor.fetchall()
        conversations = []
        for row in rows:
            last_msg_at = row[6]
            last_msg_str = last_msg_at.isoformat() if hasattr(last_msg_at, 'isoformat') else str(last_msg_at) if last_msg_at else None
            conversations.append({
                'bike_id': row[0],
                'bike_title': row[1],
                'bike_image': row[2],
                'other_user_id': row[3],
                'other_user_name': row[4],
                'last_message': row[5],
                'last_message_at': last_msg_str
            })
        conn.close()
        return jsonify({'conversations': conversations})
    except Exception as e:
        return jsonify({'error': str(e)}), 500
