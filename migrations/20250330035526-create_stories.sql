
-- +migrate Up

CREATE TABLE categories(
    id BINARY(16) PRIMARY KEY COMMENT "カテゴリーID",
    name VARCHAR(255) NOT NULL COMMENT "カテゴリー名",
    description VARCHAR(255) NOT NULL COMMENT "カテゴリ概要",
    image_url VARCHAR(255) COMMENT "カテゴリ画像",
    created_at DATETIME NOT NULL COMMENT "作成日時",
    updated_at DATETIME NOT NULL COMMENT "更新日時"
);

CREATE TABLE stories(
    id BINARY(16) PRIMARY KEY COMMENT "ストーリーID",
    category_id BINARY(16) NOT NULL COMMENT "カテゴリーID",
    title VARCHAR(255) NOT NULL COMMENT "ストーリー題名",
    episode VARCHAR(255) NOT NULL COMMENT "ストーリー区分",
    image_url VARCHAR(255) COMMENT "ストーリー画像",
    created_at DATETIME NOT NULL COMMENT "作成日時",
    updated_at DATETIME NOT NULL COMMENT "更新日時",
    UNIQUE(category_id, episode),
    FOREIGN KEY (category_id) REFERENCES categories(id)
);

-- +migrate Down
DROP TABLE stories;
DROP TABLE categories;

