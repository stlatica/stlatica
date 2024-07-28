-- +goose Up
-- +goose StatementBegin
CREATE TABLE `user_relations` (
  `follow_user_id`   char(26) NOT NULL COMMENT 'user_id of the followed side',
  `follower_user_id` char(26) NOT NULL COMMENT 'user_id of the follower side',
  `created_at`       bigint UNSIGNED NOT NULL COMMENT 'Unix time',
  `updated_at`       bigint UNSIGNED NOT NULL COMMENT 'Unix time',
  PRIMARY KEY (`follow_user_id`, `follower_user_id`),
  CONSTRAINT `fk_user_relations_follow_user_id` FOREIGN KEY (`follow_user_id`) REFERENCES `users` (`user_id`),
  CONSTRAINT `fk_user_relations_follower_user_id` FOREIGN KEY (`follower_user_id`) REFERENCES `users` (`user_id`)
)
DEFAULT CHARACTER SET utf8mb4
DEFAULT COLLATE utf8mb4_ja_0900_as_cs
;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `user_relations`;
-- +goose StatementEnd
