-- Initialize database for jelastic-golang-hello
-- This file will be executed when the PostgreSQL container starts for the first time

-- Create additional databases if needed
-- CREATE DATABASE development;
-- CREATE DATABASE test;

-- Create additional users if needed
-- CREATE USER app_user WITH PASSWORD 'app_password';
-- GRANT ALL PRIVILEGES ON DATABASE testdb TO app_user;

-- You can add any initial data or schema here
-- Example:
-- CREATE TABLE IF NOT EXISTS sample_table (
--     id SERIAL PRIMARY KEY,
--     name VARCHAR(100) NOT NULL,
--     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
-- );

SELECT 'PostgreSQL initialized successfully for jelastic-golang-hello' as message;