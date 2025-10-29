# Copiar esto en la terminal local para copiar los archivos en la carpeta del escritorio
# scp -r -o ProxyJump=jumphost ubuntu@VM-dat515:/home/ubuntu/ismaeljimenezmejias-labs/7project/project "C:\Users\ismae\OneDrive - UAM\Escritorio\dat515-project"


from flask import Flask, jsonify, request
import os
from routes.health import health_bp
from routes.data import data_bp
from routes.main import main_bp
from db import get_connection
from migrations import ensure_schema
import time
import logging
from redis_client import get_redis
from flask_session import Session

app = Flask(__name__)
app.secret_key = os.getenv('SECRET_KEY', 'dev-secret-change-me')

# Disable template caching in development
#esto era por el flickering
app.config['TEMPLATES_AUTO_RELOAD'] = True
app.config['SEND_FILE_MAX_AGE_DEFAULT'] = 0

# Ensure minimal schema adjustments (idempotent) - gated by env
if os.getenv('MIGRATE_ON_START', 'true').lower() in ('1', 'true', 'yes'):
    try:
        ensure_schema()
    except Exception as _e:
        logging.warning("Schema ensure failed or not required: %s", _e)

# Registrar blueprints
app.register_blueprint(health_bp)
app.register_blueprint(data_bp)
app.register_blueprint(main_bp) #For now saving index and about htmls.

# Optional server-side sessions via Redis when configured
try:
    redis_client = get_redis()
    if redis_client is not None:
        app.config['SESSION_TYPE'] = 'redis'
        app.config['SESSION_REDIS'] = redis_client
        app.config['SESSION_PERMANENT'] = False
        app.config['SESSION_COOKIE_SAMESITE'] = 'Lax'
        app.config['SESSION_COOKIE_SECURE'] = os.getenv('SESSION_COOKIE_SECURE', 'true').lower() in ('1','true','yes')
        app.config['SESSION_KEY_PREFIX'] = os.getenv('SESSION_KEY_PREFIX', 'sess:')
        Session(app)
        logging.info('Sessions: Redis backend ENABLED')
    else:
        logging.info('Sessions: Redis not configured or unhealthy -> using signed cookies')
except Exception as _e:
    logging.warning('Sessions: Redis setup error -> using signed cookies: %s', _e)


# Basic structured logging
logging.basicConfig(level=logging.INFO)


# Unified error handlers (JSON for /api/*, simple text otherwise)
def _is_api_request() -> bool:
    try:
        return request.path.startswith('/api')
    except Exception:
        return False


@app.errorhandler(400)
def handle_400(e):
    if _is_api_request():
        return jsonify(error='Bad Request', detail=str(e)), 400
    return 'Bad Request', 400


@app.errorhandler(401)
def handle_401(e):
    if _is_api_request():
        return jsonify(error='Unauthorized', detail=str(e)), 401
    return 'Unauthorized', 401


@app.errorhandler(403)
def handle_403(e):
    if _is_api_request():
        return jsonify(error='Forbidden', detail=str(e)), 403
    return 'Forbidden', 403


@app.errorhandler(404)
def handle_404(e):
    if _is_api_request():
        return jsonify(error='Not Found'), 404
    return 'Not Found', 404


@app.errorhandler(Exception)
def handle_500(e):
    logging.exception('Unhandled error: %s', e)
    if _is_api_request():
        return jsonify(error='Internal Server Error'), 500
    return 'Internal Server Error', 500

if __name__ == '__main__':
    # Esperar a que la DB esté lista (solo para local/dev)
    for _ in range(30):
        try:
            conn = get_connection()
            conn.close()
            break
        except Exception:
            time.sleep(1)
    # Bind al puerto de entorno (Cloud Run establece PORT)
    app.run(host='0.0.0.0', port=int(os.getenv('PORT', 8080)))
