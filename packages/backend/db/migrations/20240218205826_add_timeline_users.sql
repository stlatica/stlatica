-- +goose Up
-- +goose StatementBegin
CREATE TABLE `timeline_users` (
  `timeline_id` varchar(64) NOT NULL COMMENT 'ulid',
  `actor_id`    varchar(64) NOT NULL COMMENT 'actor_id',
  `created_at`  bigint UNSIGNED NOT NULL COMMENT 'Unix time',
  PRIMARY KEY (`timeline_id`, `actor_id`),
  CONSTRAINT `fk_timeline_users_timelines` FOREIGN KEY (`timeline_id`) REFERENCES `timelines` (`timeline_id`),
  CONSTRAINT `fk_timeline_users_actors` FOREIGN KEY (`actor_id`) REFERENCES `actors` (`actor_id`)
)
DEFAULT CHARACTER SET utf8mb4
DEFAULT COLLATE utf8mb4_ja_0900_as_cs
;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `timeline_users`;
-- +goose StatementEnd
