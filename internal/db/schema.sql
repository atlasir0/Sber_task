CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    date DATE NOT NULL, 
    completed BOOLEAN NOT NULL DEFAULT FALSE
);

