-- +goose Up
-- +goose StatementBegin
CREATE TABLE `user_access_tokens` (
  `session_id`     varchar(64) NOT NULL COMMENT 'SessionID',
  `access_token`   varchar(64) NOT NULL COMMENT 'AccessToken',
  PRIMARY KEY(`session_id`)
)
DEFAULT CHARACTER SET utf8mb4
DEFAULT COLLATE utf8mb4_ja_0900_as_cs
;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `user_access_tokens`;
-- +goose StatementEnd
