ALTER TABLE users ADD COLUMN role_id INT;
ALTER TABLE users ADD CONSTRAINT fk_users_roles FOREIGN KEY (role_id) REFERENCES roles(id);