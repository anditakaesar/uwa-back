CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    "guid" UUID NOT NULL DEFAULT uuid_generate_v4(),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,
    username TEXT NOT NULL,
    password TEXT NOT NULL,
    email TEXT NOT NULL
);

CREATE INDEX idx_users_guid ON users("guid");
CREATE INDEX idx_users_deleted_at ON users(deleted_at);
CREATE INDEX idx_users_username ON users(username);
CREATE INDEX idx_users_email ON users(email);