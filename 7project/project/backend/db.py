import os
import mysql.connector
import psycopg2


def _connect_mysql():
    db_config = {
        "host": os.getenv("DB_HOST", "database"),
        "user": os.getenv("DB_USER", "webuser"),
        "password": os.getenv("DB_PASSWORD", "webpass"),
        "database": os.getenv("DB_NAME", "webapp"),
    }
    return mysql.connector.connect(**db_config)


def _connect_postgres(db_url: str):
    return psycopg2.connect(db_url)


def get_connection():
    db_url = os.getenv("DB_URL")
    if db_url:
        return _connect_postgres(db_url)
    return _connect_mysql()
