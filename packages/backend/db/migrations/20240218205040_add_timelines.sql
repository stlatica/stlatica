-- +goose Up
-- +goose StatementBegin
CREATE TABLE `timelines` (
  `timeline_id` varchar(64) NOT NULL COMMENT 'ulid',
  `actor_id`    varchar(64) NOT NULL COMMENT 'actor_id',
  `name`        varchar(64) NOT NULL COMMENT 'name',
  `description` varchar(255) NOT NULL COMMENT 'description',
  `created_at`  bigint UNSIGNED NOT NULL COMMENT 'Unix time',
  PRIMARY KEY (`timeline_id`),
  CONSTRAINT `fk_timelines_actors` FOREIGN KEY (`actor_id`) REFERENCES `actors` (`actor_id`)
)
DEFAULT CHARACTER SET utf8mb4
DEFAULT COLLATE utf8mb4_ja_0900_as_cs
;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `timelines`;
-- +goose StatementEnd
