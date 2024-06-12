DROP TABLE IF EXISTS users CASCADE;
CREATE TABLE users(
  id UUID PRIMARY KEY,
  name VARCHAR(50) NOT NULL,
  password VARCHAR(200) NOT NULL
);

DROP TYPE IF EXISTS ROLE CASCADE;
CREATE TYPE ROLE AS ENUM ('CLIENT');

DROP TABLE IF EXISTS roles CASCADE;
CREATE TABLE roles(
  id SERIAL PRIMARY KEY,
  type ROLE UNIQUE NOT NULL 
);

DROP TYPE IF EXISTS PERMISSION CASCADE;
CREATE TYPE PERMISSION AS ENUM ('1', '2', '4', '8', '16', '32', '64');

DROP TABLE IF EXISTS permissions CASCADE;
CREATE TABLE permissions(
  permission PERMISSION NOT NULL,
  role_id SERIAL NOT NULL,
  UNIQUE(permission, role_id),
  CONSTRAINT fk_permissons_role FOREIGN KEY(role_id) REFERENCES roles(id)
);

DROP TABLE IF EXISTS accounts_roles CASCADE;
CREATE TABLE accounts_roles(
    user_id UUID,
    role_id INTEGER,
    UNIQUE(user_id, role_id),
    CONSTRAINT fk_users_roles_account FOREIGN KEY(user_id) REFERENCES users(id),
    CONSTRAINT fk_users_roles_role FOREIGN KEY(role_id) REFERENCES roles(id)
);

DROP TABLE IF EXISTS bonds CASCADE;
CREATE TABLE bonds (
    id UUID PRIMARY KEY,
    name VARCHAR(40) NOT NULL CHECK (char_length(name) >= 3 AND char_length(name) <= 40),
    quantity_sale INT NOT NULL CHECK (quantity_sale >= 1 AND quantity_sale <= 10000),
    sales_price DECIMAL(14, 4) NOT NULL CHECK (sales_price >= 0 AND sales_price <= 100000000),
    is_bought BOOLEAN DEFAULT FALSE,
    creator_user_id UUID NOT NULL,
    current_owner_id UUID NOT NULL,
    UNIQUE(id, creator_user_id),
    UNIQUE(id, current_owner_id),
    CONSTRAINT fk_creator_user_user FOREIGN KEY (creator_user_id) REFERENCES users (id) ON DELETE CASCADE,
    CONSTRAINT fk_current_owner_user FOREIGN KEY (current_owner_id) REFERENCES users (id) ON DELETE CASCADE
);

-- inserts
INSERT INTO users(id, name, password) VALUES
  ('580b87da-e389-4290-acbf-f6191467f401', 'Erik Sostenes Simon', '12345'),
  ('1148ab29-132b-4df7-9acc-b42a32c42a9f', 'Estefany Sostenes Simon', '12345');
