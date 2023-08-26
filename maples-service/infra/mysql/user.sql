CREATE TABLE `user` (
                                      `id` int(11) NOT NULL AUTO_INCREMENT,
                                      `name` varchar(255) NOT NULL,
                                      `sex` int(11) NOT NULL  DEFAULT 1,
                                      `phone` varchar(20) NOT NULL  DEFAULT '00000000000',
                                      `birthday` varchar(30) NOT NULL default '0000-00-00',
                                      `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                      `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                      PRIMARY KEY (`id`),
                                      KEY `index_create_time` (`create_time`),
                                      KEY `index_update_time` (`update_time`)
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;