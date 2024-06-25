-- inserts
INSERT INTO roles(id, type) VALUES(1, 'USER');
INSERT INTO permissions(permission, role_id) VALUES('1', 1), ( '2', 1), ('4', 1);
