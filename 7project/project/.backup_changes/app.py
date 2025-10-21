from flask import Flask
from routes.health import health_bp
from routes.data import data_bp
from routes.main import main_bp
# from db import get_connection
from db import init_db, db
import time
import os
app = Flask(__name__)
# Configure DB: use DATABASE_URL env var (e.g. postgres://...) or fallback to a local sqlite file.
# Changing to PostgreSQL later will be as simple as setting DATABASE_URL.
app.config['SQLALCHEMY_DATABASE_URI'] = os.getenv('DATABASE_URL', 'sqlite:///data.db')
app.config['SQLALCHEMY_TRACK_MODIFICATIONS'] = False
# Initialize SQLAlchemy and create tables
init_db(app)
# Registrar blueprints
app.register_blueprint(health_bp)
app.register_blueprint(data_bp)
app.register_blueprint(main_bp) # For serving index/about/publish templates.
if __name__ == '__main__':
    # Esperar a que la DB esté lista (intenta SELECT 1)
    for _ in range(30):
        try:
            with app.app_context():
                db.session.execute('SELECT 1')
            break
        except Exception:
            time.sleep(1)
    # In-memory fallback (kept commented):
    # try:
    #     conn = get_connection()
    #     conn.close()
    #     break
    # except:
    #     time.sleep(1)
    app.run(host='0.0.0.0', port=5000)
