-- Simple users table
CREATE TABLE IF NOT EXISTS users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Bikes table with owner
CREATE TABLE IF NOT EXISTS bikes (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    sale_price DECIMAL(10, 2),
    rental_price DECIMAL(10, 2),
    sale_type ENUM('venta', 'alquiler', 'ambos') NOT NULL DEFAULT 'venta',
    model VARCHAR(255),
    description TEXT,
    bike_condition VARCHAR(100),
    image_url VARCHAR(255),
    owner_id INT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Seed users
INSERT INTO users (name) VALUES
('Alice'),
('Bob'),
('Charlie');

-- Sample bikes (rental_price = 15% of sale when applicable)
INSERT INTO bikes (title, sale_price, rental_price, sale_type, model, description, bike_condition, image_url, owner_id) VALUES
('Trek mountain bike', 350.00, 52.50, 'ambos', 'Trek Marlin 5', 'Great condition, barely used', 'Like new', 'https://images.unsplash.com/photo-1518650820938-039f23f89c79?w=800&q=80', 1),
('City bike for rent', NULL, 45.00, 'alquiler', 'Btwin City 500', 'Perfect for campus commute', 'Good', 'https://images.unsplash.com/photo-1520975922284-5f5733a998b9?w=800&q=80', 2),
('Road bike', 500.00, 75.00, 'ambos', 'Giant TCR', 'High-end road bike', 'Excellent', 'https://images.unsplash.com/photo-1509395176047-4a66953fd231?w=800&q=80', 3),
('E-bike', 1200.00, 180.00, 'ambos', 'Specialized Turbo', 'Long-lasting battery', 'New', 'https://images.unsplash.com/photo-1606229365484-9f4b31dadd4e?w=800&q=80', 1),
('Folding bike', 180.00, NULL, 'venta', 'Dahon Mariner', 'Ideal for public transport', 'Good', 'https://images.unsplash.com/photo-1595433707802-6b2626ef96f0?w=800&q=80', 2);

-- Messages table for buyer-seller communication
CREATE TABLE IF NOT EXISTS messages (
    id INT AUTO_INCREMENT PRIMARY KEY,
    bike_id INT NOT NULL,
    sender_id INT NOT NULL,
    receiver_id INT NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (bike_id) REFERENCES bikes(id) ON DELETE CASCADE,
    FOREIGN KEY (sender_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (receiver_id) REFERENCES users(id) ON DELETE CASCADE,
    INDEX idx_bike_messages (bike_id, created_at),
    INDEX idx_user_messages (sender_id, receiver_id)
);
