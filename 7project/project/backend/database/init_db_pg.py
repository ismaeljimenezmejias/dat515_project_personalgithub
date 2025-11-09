import sys
import time

from db import get_connection

DDL = """
DROP TABLE IF EXISTS messages CASCADE;
DROP TABLE IF EXISTS bikes CASCADE;
DROP TABLE IF EXISTS users CASCADE;

-- Users
CREATE TABLE users (
  id          BIGSERIAL PRIMARY KEY,
  name        TEXT NOT NULL UNIQUE,
  password_hash TEXT,
  created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Bikes
CREATE TABLE bikes (
  id             BIGSERIAL PRIMARY KEY,
  title          TEXT NOT NULL,
  model          TEXT,
  image_url      TEXT,
  sale_type      TEXT NOT NULL CHECK (sale_type IN ('venta','alquiler','ambos')),
  sale_price     NUMERIC(12,2),
  rental_price   NUMERIC(12,2),
  bike_condition TEXT,
  description    TEXT,
  owner_id       BIGINT,
  created_at     TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  location_name  TEXT,
  latitude       NUMERIC,
  longitude      NUMERIC,
  FOREIGN KEY (owner_id) REFERENCES users(id) ON DELETE SET NULL
);
CREATE INDEX idx_bikes_created_at ON bikes(created_at);
CREATE INDEX idx_bikes_sale_type ON bikes(sale_type);

-- Messages
CREATE TABLE messages (
  id          BIGSERIAL PRIMARY KEY,
  bike_id     BIGINT NOT NULL REFERENCES bikes(id) ON DELETE CASCADE,
  sender_id   BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  receiver_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  content     TEXT NOT NULL,
  created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
CREATE INDEX idx_messages_bike_created_at ON messages(bike_id, created_at);
"""

SEED = """
INSERT INTO users (name) VALUES ('Alice'), ('Bob'), ('Charlie');

INSERT INTO bikes (title, sale_price, rental_price, sale_type, model, description, bike_condition, image_url, owner_id)
VALUES
('Trek mountain bike', 350.00, 52.50, 'ambos', 'Trek Marlin 5', 'Great condition, barely used', 'Like new', 'https://images.unsplash.com/photo-1518650820938-039f23f89c79?w=800&q=80', 1),
('City bike for rent', NULL, 45.00, 'alquiler', 'Btwin City 500', 'Perfect for campus commute', 'Good', 'https://images.unsplash.com/photo-1520975922284-5f5733a998b9?w=800&q=80', 2),
('Road bike', 500.00, 75.00, 'ambos', 'Giant TCR', 'High-end road bike', 'Excellent', 'https://images.unsplash.com/photo-1509395176047-4a66953fd231?w=800&q=80', 3);
"""

def wait_for_db(max_attempts: int = 20, delay_seconds: int = 3):
    """Try to grab a connection, backing off between attempts."""
    last_error = None
    for attempt in range(1, max_attempts + 1):
        try:
            return get_connection()
        except Exception as exc:  # pragma: no cover - logged for operators
            last_error = exc
            print(f"🕒 Waiting for database (attempt {attempt}/{max_attempts})…", flush=True)
            time.sleep(delay_seconds)
    print("❌ Gave up waiting for PostgreSQL to be ready.", file=sys.stderr)
    raise last_error  # type: ignore[operator]


def main():
    conn = wait_for_db()
    cur = conn.cursor()
    cur.execute(DDL)
    try:
        cur.execute(SEED)
    except Exception:  # pragma: no cover
        pass
    conn.commit()
    cur.close()
    conn.close()
    print('✅ Tables recreated and seeded (destructive).')


if __name__ == '__main__':
    main()
