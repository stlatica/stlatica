-- +goose Up
-- +goose StatementBegin
CREATE TABLE `timeline_users` (
  `timeline_id` char(26) NOT NULL COMMENT 'ulid',
  `user_id`     char(26) NOT NULL COMMENT 'user_id',
  `created_at`  bigint UNSIGNED NOT NULL COMMENT 'Unix time',
  PRIMARY KEY (`timeline_id`, `user_id`),
  CONSTRAINT `fk_timeline_users_timelines` FOREIGN KEY (`timeline_id`) REFERENCES `timelines` (`timeline_id`),
  CONSTRAINT `fk_timeline_users_users` FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`)
)
DEFAULT CHARACTER SET utf8mb4
DEFAULT COLLATE utf8mb4_ja_0900_as_cs
;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `timeline_users`;
-- +goose StatementEnd
