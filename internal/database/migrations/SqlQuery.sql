CREATE TABLE songs (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    "group" VARCHAR(255) NOT NULL,
    release_date DATE NOT NULL,
    text TEXT NOT NULL,
    link VARCHAR(255) NOT NULL
);