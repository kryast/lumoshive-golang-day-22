CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    username VARCHAR(100),
    password VARCHAR(100),
    status VARCHAR(100)
);

CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
    description VARCHAR(256),
    status VARCHAR(100)
);