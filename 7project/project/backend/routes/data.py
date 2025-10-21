from flask import Blueprint, jsonify, request, session
from db import get_connection
from datetime import datetime
import json

data_bp = Blueprint('data', __name__)


@data_bp.route('/api/data', methods=['GET', 'POST'])
@data_bp.route('/api/bikes', methods=['GET', 'POST'])
def data():
    try:
        conn = get_connection()
        cursor = conn.cursor()

        if request.method == 'POST':
            # Require login
            user_id = session.get('user_id')
            if not user_id:
                return jsonify({'error': 'Authentication required'}), 401
            payload = request.get_json(silent=True) or {}
            
            # Validar campos requeridos
            title = (payload.get('title') or '').strip()
            sale_price = payload.get('sale_price')
            rental_price = payload.get('rental_price')
            sale_type = payload.get('sale_type', 'venta')
            model = (payload.get('model') or '').strip()
            description = (payload.get('description') or '').strip()
            bike_condition = (payload.get('condition') or '').strip()
            image_url = (payload.get('image_url') or '').strip() or None
            
            if not title:
                return jsonify({'error': 'Title is required'}), 400
            
            # Calcular rental_price automáticamente si no se proporciona pero hay sale_price
            if sale_type in ['alquiler', 'ambos'] and sale_price and not rental_price:
                rental_price = float(sale_price) * 0.15
            
            # Insertar en la tabla bikes
            cursor.execute(
                """INSERT INTO bikes (title, sale_price, rental_price, sale_type, model, description, bike_condition, image_url, owner_id, created_at) 
                   VALUES (%s, %s, %s, %s, %s, %s, %s, %s, %s, %s)""",
                (title, sale_price, rental_price, sale_type, model, description, bike_condition, image_url, user_id, datetime.now())
            )
            conn.commit()
            return jsonify({'success': True, 'id': cursor.lastrowid}), 201

        # GET: obtener todas las bicis
        filter_type = request.args.get('type')  # 'venta', 'alquiler', o None para todas
        # Optional price filters
        def _to_float(val):
            try:
                return float(val)
            except (TypeError, ValueError):
                return None
        min_sale = _to_float(request.args.get('minSalePrice'))
        max_sale = _to_float(request.args.get('maxSalePrice'))
        min_rent = _to_float(request.args.get('minRentalPrice'))
        max_rent = _to_float(request.args.get('maxRentalPrice'))
        exclude_mine = request.args.get('excludeMine') in ('1', 'true', 'True')
        user_id = session.get('user_id') if exclude_mine else None
        search_text = request.args.get('search', '').strip()

        base_sql = (
            "SELECT b.id, b.title, b.sale_price, b.rental_price, b.sale_type, b.model, "
            "b.description, b.bike_condition, b.created_at, b.owner_id, u.name as owner_name, b.image_url "
            "FROM bikes b LEFT JOIN users u ON b.owner_id = u.id"
        )
        where = []
        params = []
        if filter_type and filter_type in ['venta', 'alquiler']:
            where.append("(b.sale_type = %s OR b.sale_type = 'ambos')")
            params.append(filter_type)
        if user_id:
            where.append("(b.owner_id IS NULL OR b.owner_id <> %s)")
            params.append(user_id)
        
        # Text search in title, description, and owner name
        if search_text:
            where.append("(b.title LIKE %s OR b.description LIKE %s OR u.name LIKE %s)")
            search_pattern = f"%{search_text}%"
            params.extend([search_pattern, search_pattern, search_pattern])

        # Price filtering logic
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
            # When viewing all, allow either sale OR rental to match if filters provided
            price_clauses = []
            price_params = []
            sale_clause_parts = []
            if min_sale is not None:
                sale_clause_parts.append("b.sale_price >= %s")
                price_params.append(min_sale)
            if max_sale is not None:
                sale_clause_parts.append("b.sale_price <= %s")
                price_params.append(max_sale)
            if sale_clause_parts:
                price_clauses.append("(b.sale_price IS NOT NULL AND " + " AND ".join(sale_clause_parts) + ")")

            rent_clause_parts = []
            if min_rent is not None:
                rent_clause_parts.append("b.rental_price >= %s")
                price_params.append(min_rent)
            if max_rent is not None:
                rent_clause_parts.append("b.rental_price <= %s")
                price_params.append(max_rent)
            if rent_clause_parts:
                price_clauses.append("(b.rental_price IS NOT NULL AND " + " AND ".join(rent_clause_parts) + ")")

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
            created_str = None
            if created_val:
                if hasattr(created_val, 'isoformat'):
                    created_str = created_val.isoformat()
                else:
                    created_str = str(created_val)
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
                'image_url': row[11]
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


@data_bp.route('/api/login', methods=['POST'])
def login():
    name = (request.json or {}).get('name')
    if not name or not str(name).strip():
        return jsonify({'error': 'Name is required'}), 400
    name = str(name).strip()
    conn = get_connection()
    cur = conn.cursor()
    # find or create
    cur.execute('SELECT id FROM users WHERE name=%s', (name,))
    row = cur.fetchone()
    if row:
        user_id = row[0]
    else:
        cur.execute('INSERT INTO users (name, created_at) VALUES (%s, %s)', (name, datetime.now()))
        conn.commit()
        user_id = cur.lastrowid
    conn.close()
    session['user_id'] = user_id
    session['user_name'] = name
    return jsonify({'ok': True, 'user': {'id': user_id, 'name': name}})


@data_bp.route('/api/logout', methods=['POST'])
def logout():
    session.clear()
    return jsonify({'ok': True})


@data_bp.route('/api/bikes/<int:bike_id>', methods=['DELETE'])
def delete_bike(bike_id):
    """Delete a bike (owner only)"""
    user_id = session.get('user_id')
    if not user_id:
        return jsonify({'error': 'Authentication required'}), 401
    
    try:
        conn = get_connection()
        cursor = conn.cursor()
        
        # Verify ownership
        cursor.execute('SELECT owner_id FROM bikes WHERE id = %s', (bike_id,))
        row = cursor.fetchone()
        if not row:
            return jsonify({'error': 'Bike not found'}), 404
        if row[0] != user_id:
            return jsonify({'error': 'Not authorized'}), 403
        
        # Delete
        cursor.execute('DELETE FROM bikes WHERE id = %s', (bike_id,))
        conn.commit()
        conn.close()
        return jsonify({'ok': True})
    except Exception as e:
        return jsonify({'error': str(e)}), 500


@data_bp.route('/api/bikes/<int:bike_id>', methods=['PUT'])
def update_bike(bike_id):
    """Update a bike (owner only)"""
    user_id = session.get('user_id')
    if not user_id:
        return jsonify({'error': 'Authentication required'}), 401
    
    try:
        conn = get_connection()
        cursor = conn.cursor()
        
        # Verify ownership
        cursor.execute('SELECT owner_id FROM bikes WHERE id = %s', (bike_id,))
        row = cursor.fetchone()
        if not row:
            return jsonify({'error': 'Bike not found'}), 404
        if row[0] != user_id:
            return jsonify({'error': 'Not authorized'}), 403
        
        # Update fields
        payload = request.get_json(silent=True) or {}
        title = (payload.get('title') or '').strip()
        sale_price = payload.get('sale_price')
        rental_price = payload.get('rental_price')
        sale_type = payload.get('sale_type', 'venta')
        model = (payload.get('model') or '').strip()
        description = (payload.get('description') or '').strip()
        bike_condition = (payload.get('condition') or '').strip()
        image_url = (payload.get('image_url') or '').strip() or None
        
        if not title:
            return jsonify({'error': 'Title is required'}), 400
        
        # Auto-calculate rental if needed
        if sale_type in ['alquiler', 'ambos'] and sale_price and not rental_price:
            rental_price = float(sale_price) * 0.15
        
        cursor.execute(
            """UPDATE bikes SET title=%s, sale_price=%s, rental_price=%s, sale_type=%s, 
               model=%s, description=%s, bike_condition=%s, image_url=%s 
               WHERE id=%s""",
            (title, sale_price, rental_price, sale_type, model, description, bike_condition, image_url, bike_id)
        )
        conn.commit()
        conn.close()
        return jsonify({'ok': True})
    except Exception as e:
        return jsonify({'error': str(e)}), 500
