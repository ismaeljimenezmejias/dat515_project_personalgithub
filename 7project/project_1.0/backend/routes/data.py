from flask import Blueprint, jsonify, request
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
            # Accept either { title, price, rentable } or legacy { message: "..." }
            payload = request.get_json(silent=True) or {}

            title = None
            price = None
            rentable = False

            if isinstance(payload, dict) and 'title' in payload:
                title = (payload.get('title') or '').strip()
                price = payload.get('price', 0)
                rentable = bool(payload.get('rentable', False))
                if not title:
                    return jsonify({'error': 'title is required'}), 400
                content_to_store = json.dumps({'title': title, 'price': price, 'rentable': rentable})
            else:
                # legacy path: payload.message may be plain text or a JSON string
                msg = (payload.get('message') if isinstance(payload, dict) else None) or ''
                msg = str(msg).strip()
                if not msg:
                    return jsonify({'error': 'message/title is required'}), 400
                content_to_store = msg

            cursor.execute(
                "INSERT INTO messages (content, created_at) VALUES (%s, %s)",
                (content_to_store, datetime.now())
            )
            conn.commit()

        cursor.execute("SELECT * FROM messages ORDER BY created_at DESC LIMIT 10")


        rows = cursor.fetchall()
        messages = []
        for row in rows:
            mid = row[0]
            content = row[1]
            created = row[2]
            # Try to parse structured content
            item = {'id': mid, 'created_at': created.isoformat()}
            try:
                parsed = json.loads(content)
                if isinstance(parsed, dict) and 'title' in parsed:
                    item['title'] = parsed.get('title')
                    item['price'] = parsed.get('price', 0)
                    item['rentable'] = bool(parsed.get('rentable', False))
                else:
                    item['title'] = str(content)
                    item['price'] = None
                    item['rentable'] = False
            except Exception:
                item['title'] = str(content)
                item['price'] = None
                item['rentable'] = False

            messages.append(item)

        conn.close()
        return jsonify({'messages': messages})
    except Exception as e:
        return jsonify({'error': str(e)}), 500
