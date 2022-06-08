DROP TABLE IF EXISTS `short_url`;
CREATE TABLE `short_url`  (
    `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '短链Id',
    `app_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '所属应用标识',
    `slug` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '访问的短地址路径',
    `url` varchar(4096) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '原始URL',
    `remark` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注',
    `hints` int(11) NOT NULL DEFAULT 0 COMMENT '访问次数',
    `expire_at` bigint(20) NOT NULL DEFAULT 0 COMMENT '过期时间戳',
    `create_time` datetime NOT NULL COMMENT '创建时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB COMMENT = '短链接';

