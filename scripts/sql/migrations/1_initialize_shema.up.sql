-- ----------------------------
-- Table structure for links
-- ----------------------------
DROP TABLE IF EXISTS `qurls`;
Create TABLE `qurls`
(
    `id`         bigint(20)   NOT NULL AUTO_INCREMENT COMMENT '任务ID 主键',
    `url`        varchar(255) NOT NULL COMMENT '长网址',
    `hash`       char(16)     NOT NULL COMMENT 'md5 16 hash',
    `created_at` timestamp    NOT NULL NULL,
    `deleted_at` timestamp DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE,
    KEY qurls_hash_index (hash) USING BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  AUTO_INCREMENT = 10000
  COLLATE = utf8mb4_general_ci
    PARTITION BY RANGE (id) (
        PARTITION p0 VALUES LESS THAN (100000),
        PARTITION p1 VALUES LESS THAN (200000),
        PARTITION p2 VALUES LESS THAN (300000),
        PARTITION p3 VALUES LESS THAN (400000),
        PARTITION p4 VALUES LESS THAN (500000),
        PARTITION p5 VALUES LESS THAN (600000),
        PARTITION p6 VALUES LESS THAN (700000),
        PARTITION p7 VALUES LESS THAN (800000),
        PARTITION p8 VALUES LESS THAN (900000),
        PARTITION p9 VALUES LESS THAN (1000000)
        );
create index qurls_deleted_at_index
    on qurls (deleted_at);
