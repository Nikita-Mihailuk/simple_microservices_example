CREATE TABLE IF NOT EXISTS users
(
    id SERIAL PRIMARY KEY,
    name TEXT,
    age INTEGER
);

INSERT INTO users(name, age) VALUES ('Ivan', 15)
