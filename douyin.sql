/*
 Navicat Premium Data Transfer

 Source Server         : Mysql
 Source Server Type    : MySQL
 Source Server Version : 50712
 Source Host           : localhost:3306
 Source Schema         : douyin

 Target Server Type    : MySQL
 Target Server Version : 50712
 File Encoding         : 65001

 Date: 22/08/2023 15:33:33
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for chats
-- ----------------------------
DROP TABLE IF EXISTS `chats`;
CREATE TABLE `chats`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `sender_id` bigint(20) NOT NULL,
  `receiver_id` bigint(20) NOT NULL,
  `content` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `publish_time` bigint(20) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 13 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of chats
-- ----------------------------
INSERT INTO `chats` VALUES (1, 1, 2, '你好1', 1692182935);
INSERT INTO `chats` VALUES (2, 2, 1, '你好2', 1692243369);
INSERT INTO `chats` VALUES (3, 1, 3, '你好3', 1692243406);
INSERT INTO `chats` VALUES (4, 3, 1, '你好4', 1692243472);
INSERT INTO `chats` VALUES (5, 1, 2, '你好 123', 1692597296);
INSERT INTO `chats` VALUES (6, 1, 3, '11111111111111111111', 1692610039);
INSERT INTO `chats` VALUES (7, 1, 3, '11111111111111111111', 1692617340);
INSERT INTO `chats` VALUES (8, 1, 3, '11111111111111111111', 1692620563);
INSERT INTO `chats` VALUES (9, 1, 2, '你好好呀', 1692673340);
INSERT INTO `chats` VALUES (10, 1, 3, '大', 1692673439);
INSERT INTO `chats` VALUES (11, 3, 1, 'rwr ', 1692674678);
INSERT INTO `chats` VALUES (12, 3, 4, 'hello\n', 1692675849);

-- ----------------------------
-- Table structure for comments
-- ----------------------------
DROP TABLE IF EXISTS `comments`;
CREATE TABLE `comments`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) NOT NULL,
  `video_id` bigint(20) NOT NULL,
  `comment` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `publish_time` bigint(20) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 11 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of comments
-- ----------------------------
INSERT INTO `comments` VALUES (1, 1, 1, '好看1', 1692243369);
INSERT INTO `comments` VALUES (2, 3, 2, '好看2', 1692243370);
INSERT INTO `comments` VALUES (3, 1, 1, '好看4', 1692530164);
INSERT INTO `comments` VALUES (4, 1, 2, '好看好看', 1692612449);
INSERT INTO `comments` VALUES (5, 1, 3, '123', 1692612469);
INSERT INTO `comments` VALUES (6, 1, 4, '111', 1692687536);
INSERT INTO `comments` VALUES (7, 1, 1, '333', 1692687549);
INSERT INTO `comments` VALUES (8, 1, 1, '231', 1692688024);
INSERT INTO `comments` VALUES (9, 1, 3, '123', 1692688345);
INSERT INTO `comments` VALUES (10, 1, 4, '1321', 1692688353);

-- ----------------------------
-- Table structure for favorites
-- ----------------------------
DROP TABLE IF EXISTS `favorites`;
CREATE TABLE `favorites`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) NOT NULL,
  `video_id` bigint(20) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 16 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of favorites
-- ----------------------------
INSERT INTO `favorites` VALUES (1, 2, 1);
INSERT INTO `favorites` VALUES (5, 1, 1);
INSERT INTO `favorites` VALUES (6, 1, 2);
INSERT INTO `favorites` VALUES (8, 1, 3);
INSERT INTO `favorites` VALUES (10, 3, 3);
INSERT INTO `favorites` VALUES (12, 3, 4);
INSERT INTO `favorites` VALUES (13, 3, 1);
INSERT INTO `favorites` VALUES (14, 3, 2);

-- ----------------------------
-- Table structure for relations
-- ----------------------------
DROP TABLE IF EXISTS `relations`;
CREATE TABLE `relations`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `host_id` bigint(20) NOT NULL,
  `guest_id` bigint(20) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 38 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of relations
-- ----------------------------
INSERT INTO `relations` VALUES (1, 4, 1);
INSERT INTO `relations` VALUES (18, 3, 2);
INSERT INTO `relations` VALUES (20, 3, 1);
INSERT INTO `relations` VALUES (21, 2, 3);
INSERT INTO `relations` VALUES (23, 2, 1);
INSERT INTO `relations` VALUES (34, 1, 4);
INSERT INTO `relations` VALUES (35, 1, 3);
INSERT INTO `relations` VALUES (36, 3, 4);
INSERT INTO `relations` VALUES (37, 4, 3);

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `password` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `follow_count` bigint(20) NULL DEFAULT 0,
  `follower_count` bigint(20) NULL DEFAULT 0,
  `avatar` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `background_image` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `signature` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `total_favorited` bigint(20) NULL DEFAULT 0,
  `work_count` bigint(20) NULL DEFAULT 0,
  `favorite_count` bigint(20) NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 13 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `users` VALUES (1, 'test111', '$2a$04$ltwpPhqRZB5x9fNbhXsEaOfBAYpXR5MqeESjr4BHFnsdICHGr7bBu', 2, 3, 'https://avatar-1316481827.cos.ap-https://avatar-1316481827.cos.ap-shanghai.myqcloud.com/1_avatar.jpeg', 'https://avatar-1316481827.cos.ap-shanghai.myqcloud.com/1_background.jpeg', '救赎之道，就在其中', 6, 4, 3);
INSERT INTO `users` VALUES (2, '222', '$2a$04$EJp2n/53X7xEZQi9rC6u3ejluOSFGVBqocUseiBmCxYGtCe7yZR.O', 2, 1, 'https://avatar-1316481827.cos.ap-https://avatar-1316481827.cos.ap-shanghai.myqcloud.com/1_avatar.jpeg', 'https://avatar-1316481827.cos.ap-shanghai.myqcloud.com/1_background.jpeg', '222', 0, 0, 1);
INSERT INTO `users` VALUES (3, '3333', '$2a$04$IWxZIHrYz1CkFfOobStgjOugZZB5/XqKMCfhBZhHzX5XmhBGKXJqW', 3, 3, 'https://avatar-1316481827.cos.ap-https://avatar-1316481827.cos.ap-shanghai.myqcloud.com/1_avatar.jpeg', 'https://avatar-1316481827.cos.ap-shanghai.myqcloud.com/1_background.jpeg', '333', 0, 0, 4);
INSERT INTO `users` VALUES (4, '4444', '$2a$04$ogd4nKu3Do6rT7LZq7r0aeNDqiFFJwTT/Ficedw2M5jyjLmWvQhmu', 2, 2, 'https://avatar-1316481827.cos.ap-https://avatar-1316481827.cos.ap-shanghai.myqcloud.com/1_avatar.jpeg', 'https://avatar-1316481827.cos.ap-shanghai.myqcloud.com/1_background.jpeg', '444', 0, 0, 0);

-- ----------------------------
-- Table structure for videos
-- ----------------------------
DROP TABLE IF EXISTS `videos`;
CREATE TABLE `videos`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `author_id` bigint(20) NOT NULL,
  `play_url` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `favorite_count` bigint(20) NULL DEFAULT 0,
  `comment_count` bigint(20) NULL DEFAULT 0,
  `title` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `publish_time` bigint(20) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of videos
-- ----------------------------
INSERT INTO `videos` VALUES (1, 1, 'https://video-1316481827.cos.ap-shanghai.myqcloud.com/1_瑞秋好好看.mp4', 'https://cover-1316481827.cos.ap-shanghai.myqcloud.com/1_瑞秋好好看.jpeg', 1, 3, '瑞秋好好看', 1692182935);
INSERT INTO `videos` VALUES (2, 1, 'https://video-1316481827.cos.ap-shanghai.myqcloud.com/1_迈阿密赛后梅西ins现状.mp4', 'https://cover-1316481827.cos.ap-shanghai.myqcloud.com/1_迈阿密赛后梅西ins现状.jpeg', 2, 1, '迈阿密赛后梅西ins现状', 1692243369);
INSERT INTO `videos` VALUES (3, 1, 'https://video-1316481827.cos.ap-shanghai.myqcloud.com/1_迈阿密赛后虎扑评分.mp4', 'https://cover-1316481827.cos.ap-shanghai.myqcloud.com/1_迈阿密赛后虎扑评分.jpeg', 2, 2, '迈阿密赛后虎扑评分', 1692243406);
INSERT INTO `videos` VALUES (4, 1, 'https://video-1316481827.cos.ap-shanghai.myqcloud.com/1_大型纪录片-上海交通大学传奇.mp4', 'https://cover-1316481827.cos.ap-shanghai.myqcloud.com/1_大型纪录片-上海交通大学传奇.jpeg', 1, 2, '大型纪录片-上海交通大学传奇', 1692243472);

SET FOREIGN_KEY_CHECKS = 1;
