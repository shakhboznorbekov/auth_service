CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY NOT NULL UNIQUE,
    first_name VARCHAR NOT NULL,
    last_name VARCHAR NOT NULL,
    login VARCHAR NOT NULL,
    password VARCHAR NOT NULL,
    phone_number VARCHAR NOT NULL,
    user_type int REFERENCES user_types(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS user_types(id int, title VARCHAR(90));
INSERT INTO user_types VALUES(0, "unspecified"), (1, "seller"), (2, "customer");

CREATE TABLE IF NOT EXISTS categories(
    id UUID PRIMARY KEY NOT NULL,
    name VARCHAR(255) NOT NULL UNIQUE,
    parent_id UUID REFERENCES categories(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS products (
    id UUID PRIMARY KEY,
    name VARCHAR(50) UNIQUE,
    price DECIMAL(16, 2) NOT NULL,
    category_id UUID NOT NULL REFERENCES categories(id),
    expiry_date TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS order_statuses(id int, title VARCHAR(90));
INSERT INTO order_statuses VALUES(0, "unspecified"), (1, "finished");


CREATE TABLE IF NOT EXISTS orders (
    id UUID PRIMARY KEY,
    status_id int REFERENCES order_statuses(id),
    total_price DECIMAL(16, 2),
    user_id UUID REFERENCES users(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS order_items (
    order_id UUID REFERENCES orders(id),
    product_id UUID REFERENCES products(id),
    quantity DECIMAL(16, 2),
    total_price DECIMAL(16, 2),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    PRIMARY KEY(order_id, product_id)
);
