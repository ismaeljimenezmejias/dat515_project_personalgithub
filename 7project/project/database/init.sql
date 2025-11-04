-- Destructive MySQL schema setup for local development
SET FOREIGN_KEY_CHECKS = 0;
DROP TABLE IF EXISTS messages;
DROP TABLE IF EXISTS bikes;
DROP TABLE IF EXISTS users;
SET FOREIGN_KEY_CHECKS = 1;

CREATE TABLE users (
  id BIGINT PRIMARY KEY AUTO_INCREMENT,
  name VARCHAR(255) NOT NULL UNIQUE,
  password_hash VARCHAR(255),
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB;

CREATE TABLE bikes (
  id BIGINT PRIMARY KEY AUTO_INCREMENT,
  title VARCHAR(255) NOT NULL,
  model VARCHAR(255),
  image_url TEXT,
  sale_type ENUM('venta', 'alquiler', 'ambos') NOT NULL,
  sale_price DECIMAL(12,2),
  rental_price DECIMAL(12,2),
  bike_condition VARCHAR(255),
  description TEXT,
  owner_id BIGINT,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  location_name VARCHAR(255),
  latitude DECIMAL(10,8),
  longitude DECIMAL(11,8),
  CONSTRAINT fk_bikes_owner FOREIGN KEY (owner_id) REFERENCES users(id) ON DELETE SET NULL
) ENGINE=InnoDB;

CREATE INDEX idx_bikes_created_at ON bikes(created_at);
CREATE INDEX idx_bikes_sale_type ON bikes(sale_type);

CREATE TABLE messages (
  id BIGINT PRIMARY KEY AUTO_INCREMENT,
  bike_id BIGINT NOT NULL,
  sender_id BIGINT NOT NULL,
  receiver_id BIGINT NOT NULL,
  content TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT fk_messages_bike FOREIGN KEY (bike_id) REFERENCES bikes(id) ON DELETE CASCADE,
  CONSTRAINT fk_messages_sender FOREIGN KEY (sender_id) REFERENCES users(id) ON DELETE CASCADE,
  CONSTRAINT fk_messages_receiver FOREIGN KEY (receiver_id) REFERENCES users(id) ON DELETE CASCADE
) ENGINE=InnoDB;

CREATE INDEX idx_messages_bike_created_at ON messages(bike_id, created_at);

INSERT INTO users (name)
VALUES ('Alice'), ('Bob'), ('Charlie');

INSERT INTO bikes (
  title,
  sale_price,
  rental_price,
  sale_type,
  model,
  description,
  bike_condition,
  image_url,
  owner_id
) VALUES
  (
    'Trek mountain bike',
    350.00,
    52.50,
    'ambos',
    'Trek Marlin 5',
    'Great condition, barely used',
    'Like new',
    'https://images.unsplash.com/photo-1518650820938-039f23f89c79?w=800&q=80',
    1
  ),
  (
    'City bike for rent',
    NULL,
    45.00,
    'alquiler',
    'Btwin City 500',
    'Perfect for campus commute',
    'Good',
    'https://images.unsplash.com/photo-1520975922284-5f5733a998b9?w=800&q=80',
    2
  ),
  (
    'Road bike',
    500.00,
    75.00,
    'ambos',
    'Giant TCR',
    'High-end road bike',
    'Excellent',
    'https://images.unsplash.com/photo-1509395176047-4a66953fd231?w=800&q=80',
    3
  );
