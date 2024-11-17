# timelines

## Description

<details>
<summary><strong>Table Definition</strong></summary>

```sql
CREATE TABLE `timelines` (
  `timeline_id` char(26) COLLATE utf8mb4_ja_0900_as_cs NOT NULL COMMENT 'ulid',
  `user_id` char(26) COLLATE utf8mb4_ja_0900_as_cs NOT NULL COMMENT 'user_id',
  `name` varchar(64) COLLATE utf8mb4_ja_0900_as_cs NOT NULL COMMENT 'name',
  `description` varchar(255) COLLATE utf8mb4_ja_0900_as_cs NOT NULL COMMENT 'description',
  `created_at` bigint unsigned NOT NULL COMMENT 'Unix time',
  `updated_at` bigint unsigned NOT NULL COMMENT 'Unix time',
  PRIMARY KEY (`timeline_id`),
  KEY `fk_timelines_users` (`user_id`),
  CONSTRAINT `fk_timelines_users` FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_ja_0900_as_cs
```

</details>

## Columns

| Name | Type | Default | Nullable | Children | Parents | Comment |
| ---- | ---- | ------- | -------- | -------- | ------- | ------- |
| timeline_id | char(26) |  | false | [timeline_users](timeline_users.md) |  | ulid |
| user_id | char(26) |  | false |  | [users](users.md) | user_id |
| name | varchar(64) |  | false |  |  | name |
| description | varchar(255) |  | false |  |  | description |
| created_at | bigint unsigned |  | false |  |  | Unix time |
| updated_at | bigint unsigned |  | false |  |  | Unix time |

## Constraints

| Name | Type | Definition |
| ---- | ---- | ---------- |
| fk_timelines_users | FOREIGN KEY | FOREIGN KEY (user_id) REFERENCES users (user_id) |
| PRIMARY | PRIMARY KEY | PRIMARY KEY (timeline_id) |

## Indexes

| Name | Definition |
| ---- | ---------- |
| fk_timelines_users | KEY fk_timelines_users (user_id) USING BTREE |
| PRIMARY | PRIMARY KEY (timeline_id) USING BTREE |

## Relations

![er](timelines.svg)

---

> Generated by [tbls](https://github.com/k1LoW/tbls)