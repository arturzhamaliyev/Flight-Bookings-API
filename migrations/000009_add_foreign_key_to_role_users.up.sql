ALTER TABLE users
ADD CONSTRAINT role_fk FOREIGN KEY (role) REFERENCES roles(id);