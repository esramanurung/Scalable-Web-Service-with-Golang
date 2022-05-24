-- DDL
-- CREATE TABLE <nama_table>

CREATE TABLE employees(
    id SERIAL PRIMARY KEY,
    full_name varchar(50) NOT NULL,
    email TEXT UNIQUE NOT NULL,
    age INT NOT NUll,
    division varchar(20) NOT NULL
);