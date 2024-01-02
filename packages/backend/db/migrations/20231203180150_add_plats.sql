-- +goose Up
-- +goose StatementBegin
CREATE TABLE `plats` (
  `plat_id`    varchar(64) NOT NULL COMMENT 'ulid',
  `actor_id`   varchar(64) NOT NULL COMMENT 'actor_id',
  `content`    varchar(255) NOT NULL COMMENT 'body text',
  `created_at` bigint UNSIGNED NOT NULL COMMENT 'Unix time',
  PRIMARY KEY (`plat_id`)
)
DEFAULT CHARACTER SET utf8mb4
DEFAULT COLLATE utf8mb4_ja_0900_as_cs
;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `plats`;
-- +goose StatementEnd
