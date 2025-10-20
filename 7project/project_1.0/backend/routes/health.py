from flask import Blueprint, jsonify
from db import get_connection
from redis_client import get_redis
from datetime import datetime

health_bp = Blueprint('health', __name__)

@health_bp.route('/health')
def health():
    try:
        conn = get_connection()
        conn.close()
        db_status = 'healthy'
    except:
        db_status = 'unhealthy'

    try:
        r = get_redis()
        r.ping()
        cache_status = 'healthy'
    except:
        cache_status = 'unhealthy'

    status = 'healthy' if db_status == 'healthy' and cache_status == 'healthy' else 'unhealthy'
    return jsonify({
        'status': status,
        'database': db_status,
        'cache': cache_status,
        'timestamp': datetime.now().isoformat()
    })