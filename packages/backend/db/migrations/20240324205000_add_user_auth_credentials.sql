-- +goose Up
-- +goose StatementBegin
CREATE TABLE `user_auth_credentials` (
  `user_id`            char(26) NOT NULL COMMENT 'UserID',
  `encrypted_password` varchar(64) NOT NULL COMMENT 'Encrypted password',
  `created_at`         bigint UNSIGNED NOT NULL COMMENT 'Unix time',
  `updated_at`         bigint UNSIGNED NOT NULL COMMENT 'Unix time',
  PRIMARY KEY(`user_id`),
  CONSTRAINT `fk_user_auth_credentials_users` FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`)
)
DEFAULT CHARACTER SET utf8mb4
DEFAULT COLLATE utf8mb4_ja_0900_as_cs
;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `user_auth_credentials`;
-- +goose StatementEnd
