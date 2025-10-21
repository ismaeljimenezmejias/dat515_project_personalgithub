from flask_sqlalchemy import SQLAlchemy

# Create the SQLAlchemy db instance. Call init_db(app) from app.py to bind it.
db = SQLAlchemy()

def init_db(app):
    db.init_app(app)
    # Create tables if they don't exist
    with app.app_context():
        from models import Message  # import models so tables are registered
        db.create_all()
    return db


#In memory database connection (commented out)
# import mysql.connector
# import os

# db_config = {
#     'host': os.getenv('DB_HOST', 'database'),
#     'user': os.getenv('DB_USER', 'webuser'),
#     'password': os.getenv('DB_PASSWORD', 'webpass'),
#     'database': os.getenv('DB_NAME', 'webapp')
# }

# def get_connection():
#     return mysql.connector.connect(**db_config)