-- +goose Up
-- +goose StatementBegin
CREATE TABLE `users` (
  `user_id`             char(26) NOT NULL COMMENT 'ulid',
  `preferred_user_id`   varchar(64) NOT NULL COMMENT 'preferred user id',
  `preferred_user_name` varchar(64) NOT NULL COMMENT 'preferred user name',
  `registered_at`       bigint UNSIGNED NOT NULL COMMENT 'Unix time',
  `is_public`           boolean NOT NULL COMMENT 'user is public',
  `mail_address`        varchar(256) UNIQUE NOT NULL COMMENT 'mail address',
  PRIMARY KEY (`user_id`),
  UNIQUE KEY `idx_users_preferred_user_id` (`preferred_user_id`),
  UNIQUE KEY `idx_users_mail_address` (`mail_address`)
)
DEFAULT CHARACTER SET utf8mb4
DEFAULT COLLATE utf8mb4_ja_0900_as_cs
;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `users`;
-- +goose StatementEnd
