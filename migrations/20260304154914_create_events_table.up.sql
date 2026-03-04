CREATE TABLE events (
    id SERIAL PRIMARY KEY,
    title VARCHAR(200) NOT NULL,
    description TEXT,
    datetime TIMESTAMP NOT NULL,
    location VARCHAR(200),
    price NUMERIC(10,2),
    image_url TEXT,
    category_id INT REFERENCES categories(id),
    priority INT DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);