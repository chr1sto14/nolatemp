CREATE DATABASE xxx;

CREATE TABLE xxx.temp (
    ts TIMESTAMP PRIMARY KEY,
    intemp DECIMAL,
    outtemp DECIMAL
);

CREATE TABLE xxx.images (
    id SERIAL PRIMARY KEY,
    img BYTES
);
