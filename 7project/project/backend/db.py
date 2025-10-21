import mysql.connector
import os

db_config = {
    'host': os.getenv('DB_HOST', 'database'),
    'user': os.getenv('DB_USER', 'webuser'),
    'password': os.getenv('DB_PASSWORD', 'webpass'),
    'database': os.getenv('DB_NAME', 'webapp')
}

def get_connection():
    return mysql.connector.connect(**db_config)