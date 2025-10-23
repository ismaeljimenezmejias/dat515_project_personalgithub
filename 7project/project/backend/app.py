# Copiar esto en la terminal local para copiar los archivos en la carpeta del escritorio
# scp -r -o ProxyJump=jumphost ubuntu@VM-dat515:/home/ubuntu/ismaeljimenezmejias-labs/7project/project "C:\Users\ismae\OneDrive - UAM\Escritorio\dat515-project"


from flask import Flask, jsonify
import os
from routes.health import health_bp
from routes.data import data_bp
from routes.main import main_bp
from db import get_connection
import time

app = Flask(__name__)
app.secret_key = os.getenv('SECRET_KEY', 'dev-secret-change-me')

# Disable template caching in development
#esto era por el flickering
app.config['TEMPLATES_AUTO_RELOAD'] = True
app.config['SEND_FILE_MAX_AGE_DEFAULT'] = 0

# Registrar blueprints
app.register_blueprint(health_bp)
app.register_blueprint(data_bp)
app.register_blueprint(main_bp) #For now saving index and about htmls.

if __name__ == '__main__':
    # Esperar a que la DB esté lista
    for _ in range(30):
        try:
            conn = get_connection()
            conn.close()
            break
        except:
            time.sleep(1)
    app.run(host='0.0.0.0', port=5000)
