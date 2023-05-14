CREATE TABLE IF NOT EXISTS iplogs (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,
    "address" varchar(20),
    "count" INT
);

CREATE INDEX idx_iplogs_address ON iplogs("address");
CREATE INDEX idx_iplogs_deleted_at ON iplogs(deleted_at);