#Database in local
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

#Postgres in Railway Database
import psycopg2
import os

def get_connection():
    db_url = os.getenv("DB_URL")
    if not db_url:
        raise RuntimeError(
            "DB_URL environment variable not set. Set it in your environment or docker-compose before starting the app."
        )
    # db_url = "postgresql://postgres:eKgRIQIGpJrdHjTKSyUNeqBpbhlWnTZM@nozomi.proxy.rlwy.net:15863/railway"
    conn = psycopg2.connect(db_url)
    return conn
