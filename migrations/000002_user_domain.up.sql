CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    "guid" UUID NOT NULL DEFAULT uuid_generate_v4(),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP,
    created_by TEXT,
    updated_by TEXT,
    deleted_by TEXT,
    username TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    phone_number TEXT NULL UNIQUE
);

CREATE INDEX idx_users_guid ON users("guid");
CREATE INDEX idx_users_deleted_at ON users(deleted_at);
CREATE INDEX idx_users_username ON users(username);
CREATE INDEX idx_users_email ON users(email);