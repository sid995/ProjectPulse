-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TYPE user_role AS ENUM (
  'admin',
  'project_manager',
  'developer',
  'viewer'
);
CREATE TABLE IF NOT EXISTS users (
  id SERIAL PRIMARY KEY,
  email VARCHAR(255) NOT NULL UNIQUE,
  password VARCHAR(255) NOT NULL,
  first_name VARCHAR(100),
  last_name VARCHAR(100),
  role user_role NOT NULL DEFAULT 'viewer',
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP WITH TIME ZONE
);
CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_deleted_at ON users(deleted_at);
-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP INDEX IF EXISTS idx_users_deleted_at;
DROP INDEX IF EXISTS idx_users_email;
DROP TABLE IF EXISTS users;
DROP TYPE IF EXISTS user_role;