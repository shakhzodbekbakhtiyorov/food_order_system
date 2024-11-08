CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY NOT NULL,
    name VARCHAR(255),
    login VARCHAR(255),
    password_hash VARCHAR(255),
    type_id int
);

CREATE TABLE IF NOT EXISTS categories (
    id SERIAL PRIMARY KEY NOT NULL,
    name VARCHAR(255),
    image VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY NOT NULL, 
    name VARCHAR(255),
    price DOUBLE PRECISION,
    category_id INT REFERENCES categories(id),
    image VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS orders (
    id SERIAL PRIMARY KEY NOT NULL,
    order_time INT,
    status INT,
    order_number INT,
    amount DOUBLE PRECISION
);

CREATE TABLE IF NOT EXISTS order_items (
    id SERIAL PRIMARY KEY NOT NULL,
    count INT,
    status INT,
    product_id INT REFERENCES products(id),
    order_id INT REFERENCES orders(id)
);

CREATE TABLE IF NOT EXISTS payments (
    id SERIAL PRIMARY KEY NOT NULL,
    order_id INT REFERENCES orders(id),
    payment_type INT,
    amount DOUBLE PRECISION,
    change DOUBLE PRECISION
);

