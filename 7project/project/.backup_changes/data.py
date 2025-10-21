from flask import Blueprint, jsonify, request, current_app
from db import db
from models import Message
from datetime import datetime
import json

data_bp = Blueprint('data', __name__)


@data_bp.route('/api/data', methods=['GET', 'POST'])
@data_bp.route('/api/bikes', methods=['GET', 'POST'])
def data():
    try:
        if request.method == 'POST':
            # Accept either { title, price, rentable } or legacy { message: "..." }
            payload = request.get_json(silent=True) or {}

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

            m = Message(content=content_to_store, created_at=datetime.utcnow())
            db.session.add(m)
            db.session.commit()

        # Query latest 10 messages
        rows = Message.query.order_by(Message.created_at.desc()).limit(10).all()
        messages = []
        for row in rows:
            mid = row.id
            content = row.content
            created = row.created_at
            item = {'id': mid, 'created_at': created.isoformat() if created else None}
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

        return jsonify({'messages': messages})
    except Exception as e:
        current_app.logger.exception('Error in data route')
        return jsonify({'error': str(e)}), 500
