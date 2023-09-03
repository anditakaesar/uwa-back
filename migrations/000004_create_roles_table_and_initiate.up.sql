CREATE TABLE IF NOT EXISTS roles (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    description VARCHAR(255),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP,
    created_by TEXT,
    updated_by TEXT,
    deleted_by TEXT
);

CREATE INDEX idx_roles_deleted_at ON roles(deleted_at);