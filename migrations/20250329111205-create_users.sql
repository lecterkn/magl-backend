
-- +migrate Up
CREATE TABLE users(
    id BINARY(16) PRIMARY KEY COMMENT "ユーザーID",
    name VARCHAR(255) NOT NULL COMMENT "ユーザー名",
    email VARCHAR(255) NOT NULL COMMENT "メールアドレス",
    password BINARY(60) NOT NULL COMMENT "パスワード",
    created_at DATETIME NOT NULL COMMENT "作成日時",
    updated_at DATETIME NOT NULL COMMENT "更新日時",
    UNIQUE (name),
    UNIQUE (email)
);

-- +migrate Down
DROP TABLE users;
