from db import get_connection

# Idempotent, safe migrations to keep prod in sync without manual SQL
# - Adds password_hash column to users if missing
# - Ensures unique index on users(name)
# - Adds optional location columns to bikes (location_name, latitude, longitude)

def ensure_schema():
    conn = get_connection()
    cur = conn.cursor()
    try:
        cur.execute("ALTER TABLE users ADD COLUMN IF NOT EXISTS password_hash TEXT;")
    except Exception:
        # Column may already exist or table absent; ignore
        pass
    try:
        cur.execute("CREATE UNIQUE INDEX IF NOT EXISTS users_name_key ON users(name);")
    except Exception:
        # Index may exist or duplicates may prevent creation; ignore here
        pass
    # Location fields for bikes (nullable, non-breaking)
    try:
        cur.execute("ALTER TABLE bikes ADD COLUMN IF NOT EXISTS location_name TEXT;")
    except Exception:
        pass
    try:
        cur.execute("ALTER TABLE bikes ADD COLUMN IF NOT EXISTS latitude DOUBLE PRECISION;")
    except Exception:
        pass
    try:
        cur.execute("ALTER TABLE bikes ADD COLUMN IF NOT EXISTS longitude DOUBLE PRECISION;")
    except Exception:
        pass
    conn.commit()
    conn.close()
