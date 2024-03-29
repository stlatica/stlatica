-- +goose Up
-- +goose StatementBegin
CREATE TABLE `user_authentication` (
  `user_id`     char(64) NOT NULL COMMENT 'UserID',
  `user_password`   varchar(64) NOT NULL COMMENT 'UserPassword',
  PRIMARY KEY(`user_id`),
  FOREIGN KEY (`user_id`) REFERENCES `actors` (`actor_id`)
)
DEFAULT CHARACTER SET utf8mb4
DEFAULT COLLATE utf8mb4_ja_0900_as_cs
;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `user_password`;
-- +goose StatementEnd
