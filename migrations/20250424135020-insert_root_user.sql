
-- +migrate Up
ALTER TABLE users MODIFY COLUMN created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP;
ALTER TABLE users MODIFY COLUMN updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;

-- magl-root:xFD5n3PasjOHgih4gmrwNGs2G2mEsyFO
INSERT INTO users (id, email, name, password, role)
VALUES (
    UNHEX(REPLACE(UUID(), '-', '')),
    'magl-rootuser@example.com',
    'magl-root',
    '$2a$12$KV36qS63rXzXLsczOTWzB.6lT8Er8WtoXGA6e/F.L6uvYt.1.L5oO',
    3
);

-- +migrate Down
