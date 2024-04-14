-- +goose Up
-- +goose StatementBegin
CREATE TABLE `timelines` (
  `timeline_id` char(26) NOT NULL COMMENT 'ulid',
  `user_id`     char(26) NOT NULL COMMENT 'user_id',
  `name`        varchar(64) NOT NULL COMMENT 'name',
  `description` varchar(255) NOT NULL COMMENT 'description',
  `created_at`  bigint UNSIGNED NOT NULL COMMENT 'Unix time',
  `updated_at`  bigint UNSIGNED NOT NULL COMMENT 'Unix time',
  PRIMARY KEY (`timeline_id`),
  CONSTRAINT `fk_timelines_users` FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`)
)
DEFAULT CHARACTER SET utf8mb4
DEFAULT COLLATE utf8mb4_ja_0900_as_cs
;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `timelines`;
-- +goose StatementEnd
