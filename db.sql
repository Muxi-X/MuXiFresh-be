CREATE DATABASE `muxifresh`;
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
)ENGINE = InnoDB;
  -- DEFAULT CHARSET = utf8mb4
  -- COLLATE = utf8mb4_unicode_ci
  -- ROW_FORMAT = DYNAMIC;

DROP TABLE IF EXISTS `forms`;
CREATE TABLE `forms`
(
    `id`              int(11)     NOT NULL AUTO_INCREMENT,
    `email`           varchar(35) NOT NULL,
    `name`            varchar(20) NOT NULL, -- 姓名
    `avatar`          varchar(100) DEFAULT NULL, -- 头像 DEFAULT NULL 表示默认为空 
    `student_id`      char(10)  DEFAULT NULL, -- 学号
    `college`         varchar(20) DEFAULT NULL, -- 学院
    `major`           varchar(20) DEFAULT NULL, -- 专业
    `grade`           varchar(20) DEFAULT NULL, -- 年级
    `gender`          varchar(20) DEFAULT NULL, -- 性别
    `contact_way`     varchar(20) DEFAULT NULL, -- 联系方式
    `contact_number`  varchar(20) DEFAULT NULL, -- 联系号码
    `group`           varchar(20) DEFAULT NULL, -- 心动组别
    `reason`          varchar(100) DEFAULT NULL, -- 心动理由
    `understand`      varchar(100) DEFAULT NULL, -- 对组别的了解
    `self_introduction` varchar(100) DEFAULT NULL, -- 自我介绍
    `if_other_organization` int(3) DEFAULT NULL COMMENT '是否有加入/正在加入一些其他组织或担任学生工作 0-没有 1-有',
    PRIMARY KEY(`id`),
    UNIQUE KEY `email` (`email`)
)ENGINE = InnoDB;
  -- DEFAULT CHARSET = utf8mb4
  -- COLLATE = utf8mb4_unicode_ci
  -- ROW_FORMAT = DYNAMIC;

DROP TABLE IF EXISTS `schedules`;
CREATE TABLE `schedules`
(
    `id`              int(11)     NOT NULL AUTO_INCREMENT,
    `email`           varchar(35) NOT NULL,
    `name`            varchar(20) NOT NULL, -- 姓名
    `avatar`          varchar(100) DEFAULT NULL, -- 头像 DEFAULT NULL 表示默认为空 
    `student_id`      char(10)  DEFAULT NULL, -- 学号
    `college`         varchar(20) DEFAULT NULL, -- 学院
    `major`           varchar(20) DEFAULT NULL, -- 专业
    `group`           varchar(20) DEFAULT NULL, -- 报名组别
    `form_status`     int(3) DEFAULT NULL COMMENT '报名表状态 0-未提交 1-已提交', 
    `work_status`     int(3) DEFAULT NULL COMMENT '作业提交状态 0-未提交 1-已提交',
    `admission_status`int(3) DEFAULT NULL COMMENT '录取状态 0-未录取 1-已录取',
    PRIMARY KEY(`id`),
    UNIQUE KEY `email` (`email`)
)ENGINE = InnoDB;
  -- DEFAULT CHARSET = utf8mb4
  -- COLLATE = utf8mb4_unicode_ci
  -- ROW_FORMAT = DYNAMIC;