-- +goose Up
-- +goose StatementBegin
CREATE TABLE `user_refresh_tokens` (
  `hash_session_id_access_token`     varchar(64) NOT NULL COMMENT 'Hash(SessionID, AccessToken)',
  `refresh_token`   varchar(64) NOT NULL COMMENT 'RefreshToken',
  PRIMARY KEY(`hash_session_id_access_token`)
)
DEFAULT CHARACTER SET utf8mb4
DEFAULT COLLATE utf8mb4_ja_0900_as_cs
;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `user_refresh_tokens`;
-- +goose StatementEnd
