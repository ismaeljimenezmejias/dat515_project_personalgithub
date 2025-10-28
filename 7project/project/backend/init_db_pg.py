from db import get_connection

# Destructive recreate: drop existing tables and recreate schema exactly
DDL = """
DROP TABLE IF EXISTS messages CASCADE;
DROP TABLE IF EXISTS bikes CASCADE;
DROP TABLE IF EXISTS users CASCADE;

-- Users
CREATE TABLE users (
  id          BIGSERIAL PRIMARY KEY,
  name        TEXT NOT NULL,
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

if __name__ == '__main__':
    conn = get_connection()
    cur = conn.cursor()
    cur.execute(DDL)
    # optional: seed some sample users/bikes for development
    try:
        cur.execute(SEED)
    except Exception:
        # ignore seed errors (e.g., duplicates)
        pass
    conn.commit()
    cur.close()
    conn.close()
    print('✅ Tables recreated and seeded (destructive).')
