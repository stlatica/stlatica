-- +goose Up
-- +goose StatementBegin
CREATE TABLE `plats` (
  `plat_id`    char(26) NOT NULL COMMENT 'ulid',
  `user_id`    char(26) NOT NULL COMMENT 'user_id',
  `content`    varchar(255) NOT NULL COMMENT 'body text',
  `created_at` bigint UNSIGNED NOT NULL COMMENT 'Unix time',
  `updated_at` bigint UNSIGNED NOT NULL COMMENT 'Unix time',
  PRIMARY KEY (`plat_id`),
  CONSTRAINT `fk_plats_users` FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`)
)
DEFAULT CHARACTER SET utf8mb4
DEFAULT COLLATE utf8mb4_ja_0900_as_cs
;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `plats`;
-- +goose StatementEnd
