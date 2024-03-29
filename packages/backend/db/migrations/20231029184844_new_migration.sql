-- +goose Up
-- +goose StatementBegin
CREATE TABLE `actors` (
  `actor_id`      varchar(64) NOT NULL COMMENT 'ulid',
  `display_name`  varchar(64) NOT NULL COMMENT 'name',
  `registed_at`   bigint UNSIGNED NOT NULL COMMENT 'Unix time',
  `is_public`     boolean NOT NULL COMMENT 'user is public',
  `mail_address`  varchar(256) UNIQUE NOT NULL COMMENT 'mail address',
  PRIMARY KEY (`actor_id`)
)
DEFAULT CHARACTER SET utf8mb4
DEFAULT COLLATE utf8mb4_ja_0900_as_cs
;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `actors`;
-- +goose StatementEnd
