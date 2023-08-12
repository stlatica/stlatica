-- +goose Up
CREATE TABLE `sample_users` (
  `user_id`      varchar(64) NOT NULL COMMENT 'ulid',
  `created_at`   bigint UNSIGNED NOT NULL COMMENT 'Unix time',
  `updated_at`   bigint UNSIGNED NOT NULL COMMENT 'Unix time',
  PRIMARY KEY (`user_id`)
)
  DEFAULT CHARACTER SET utf8mb4
  DEFAULT COLLATE utf8mb4_ja_0900_as_cs
;
-- +goose Down
SET FOREIGN_KEY_CHECKS = 0;

DROP TABLE `sample_users`;

SET FOREIGN_KEY_CHECKS = 1;
