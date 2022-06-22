-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`
(
    `id`            int(11)     NOT NULL AUTO_INCREMENT,
    `name`          varchar(20) NOT NULL,
    `email`         varchar(35) NOT NULL,
    `avatar`        varchar(100) DEFAULT NULL,
    `student_id`    char(10)     DEFAULT NULL,
    `hash_password` varchar(100) NOT NULL,
    `role`          int(11)     NOT NULL COMMENT '权限 0-无权限用户 1-普通学生用户 2（闲置，可能学生管理员） 3-团队成员 4-团队管理员',
    `message`       varchar(200) DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `email` (`email`),
    KEY `role` (`role`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci
  ROW_FORMAT = DYNAMIC;
