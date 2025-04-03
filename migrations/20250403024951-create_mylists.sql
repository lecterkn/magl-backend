
-- +migrate Up
CREATE TABLE mylists(
    user_id BINARY(16) NOT NULL COMMENT "ユーザーID",
    story_id BINARY(16) NOT NULL COMMENT "ストーリーID",
    score TINYINT NOT NULL COMMENT "スコア",
    UNIQUE(user_id, story_id),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (story_id) REFERENCES stories(id)
);

-- +migrate Down
DROP TABLE mylists;
