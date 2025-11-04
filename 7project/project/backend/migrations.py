try:
    from mysql.connector import Error as MySQLError
except ImportError:  # pragma: no cover - fallback for environments without mysqlclient
    MySQLError = Exception

from db import get_connection

# Idempotent, safe migrations to keep prod in sync without manual SQL
# - Adds password_hash column to users if missing
# - Ensures unique index on users(name)
# - Adds optional location columns to bikes (location_name, latitude, longitude)


def _is_postgres_connection(conn) -> bool:
    return conn.__class__.__module__.startswith("psycopg2")


def ensure_schema():
    conn = get_connection()
    cur = conn.cursor()
    try:
        if _is_postgres_connection(conn):
            cur.execute("ALTER TABLE users ADD COLUMN IF NOT EXISTS password_hash TEXT;")
            cur.execute("CREATE UNIQUE INDEX IF NOT EXISTS users_name_key ON users(name);")
            cur.execute("ALTER TABLE bikes ADD COLUMN IF NOT EXISTS location_name TEXT;")
            cur.execute("ALTER TABLE bikes ADD COLUMN IF NOT EXISTS latitude DOUBLE PRECISION;")
            cur.execute("ALTER TABLE bikes ADD COLUMN IF NOT EXISTS longitude DOUBLE PRECISION;")
        else:
            # MySQL path – inspect information_schema instead of IF NOT EXISTS
            cur.execute(
                """
                SELECT COUNT(*) FROM INFORMATION_SCHEMA.COLUMNS
                WHERE TABLE_SCHEMA = DATABASE() AND TABLE_NAME = 'users' AND COLUMN_NAME = 'password_hash'
                """
            )
            result = cur.fetchone()
            column_missing = not result or result[0] == 0
            if column_missing:
                try:
                    cur.execute("ALTER TABLE users ADD COLUMN password_hash TEXT;")
                except Exception as exc:
                    if getattr(exc, "errno", None) != 1060:  # duplicate column
                        raise

            cur.execute("SHOW INDEX FROM users WHERE Key_name = 'users_name_key'")
            if not cur.fetchone():
                try:
                    cur.execute("CREATE UNIQUE INDEX users_name_key ON users(name);")
                except Exception as exc:
                    if getattr(exc, "errno", None) not in (1061, 1840):  # duplicate key / index exists
                        raise

            cur.execute(
                """
                SELECT COUNT(*) FROM INFORMATION_SCHEMA.COLUMNS
                WHERE TABLE_SCHEMA = DATABASE() AND TABLE_NAME = 'bikes' AND COLUMN_NAME = 'location_name'
                """
            )
            result = cur.fetchone()
            if not result or result[0] == 0:
                try:
                    cur.execute("ALTER TABLE bikes ADD COLUMN location_name TEXT;")
                except Exception as exc:
                    if getattr(exc, "errno", None) != 1060:
                        raise

            cur.execute(
                """
                SELECT COUNT(*) FROM INFORMATION_SCHEMA.COLUMNS
                WHERE TABLE_SCHEMA = DATABASE() AND TABLE_NAME = 'bikes' AND COLUMN_NAME = 'latitude'
                """
            )
            result = cur.fetchone()
            if not result or result[0] == 0:
                try:
                    cur.execute("ALTER TABLE bikes ADD COLUMN latitude DOUBLE;")
                except Exception as exc:
                    if getattr(exc, "errno", None) != 1060:
                        raise

            cur.execute(
                """
                SELECT COUNT(*) FROM INFORMATION_SCHEMA.COLUMNS
                WHERE TABLE_SCHEMA = DATABASE() AND TABLE_NAME = 'bikes' AND COLUMN_NAME = 'longitude'
                """
            )
            result = cur.fetchone()
            if not result or result[0] == 0:
                try:
                    cur.execute("ALTER TABLE bikes ADD COLUMN longitude DOUBLE;")
                except Exception as exc:
                    if getattr(exc, "errno", None) != 1060:
                        raise
    finally:
        conn.commit()
        conn.close()
