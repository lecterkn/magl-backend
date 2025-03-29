
-- +migrate Up
ALTER TABLE users ADD COLUMN role TINYINT NOT NULL COMMENT "役職";

-- +migrate Down
