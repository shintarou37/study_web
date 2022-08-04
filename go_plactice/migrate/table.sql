CREATE TABLE `data1` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `created_at` datetime(3) DEFAULT NULL,
    `updated_at` datetime(3) DEFAULT NULL,
    `deleted_at` datetime(3) DEFAULT NULL,
    `title` longtext,
    `content` longtext,
    PRIMARY KEY (`id`),
    KEY `idx_data1_deleted_at` (`deleted_at`)
) ENGINE = InnoDB DEFAULT CHARSET = latin1
-- 構造体
-- type Data1 struct {
--     gorm.Model
-- 	Title    string
-- 	Content  string
-- }