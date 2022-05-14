/*
 Navicat Premium Data Transfer

 Source Server         : aliyun
 Source Server Type    : MySQL
 Source Server Version : 80029
 Source Host           : 112.74.109.70:3306
 Source Schema         : TikTokLite

 Target Server Type    : MySQL
 Target Server Version : 80029
 File Encoding         : 65001

 Date: 14/05/2022 16:33:43
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for Comment
-- ----------------------------
DROP TABLE IF EXISTS `Comment`;
CREATE TABLE `Comment`  (
  `comment_id` bigint NOT NULL AUTO_INCREMENT,
  `user_id` bigint NOT NULL,
  `video_id` bigint NOT NULL,
  `comment` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `time` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`comment_id`) USING BTREE,
  INDEX `commentUser`(`user_id`) USING BTREE,
  INDEX `commentVideo`(`video_id`) USING BTREE,
  CONSTRAINT `commentuser` FOREIGN KEY (`user_id`) REFERENCES `User` (`user_id`) ON DELETE CASCADE ON UPDATE RESTRICT,
  CONSTRAINT `commentvideo` FOREIGN KEY (`video_id`) REFERENCES `Video` (`video_id`) ON DELETE CASCADE ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for Favorite
-- ----------------------------
DROP TABLE IF EXISTS `Favorite`;
CREATE TABLE `Favorite`  (
  `favorite_id` bigint NOT NULL AUTO_INCREMENT,
  `user_id` bigint NOT NULL,
  `video_id` bigint NOT NULL,
  PRIMARY KEY (`favorite_id`) USING BTREE,
  INDEX `favoriteUser`(`user_id`) USING BTREE,
  INDEX `favoriteVideo`(`video_id`) USING BTREE,
  CONSTRAINT `favoriteuser` FOREIGN KEY (`user_id`) REFERENCES `User` (`user_id`) ON DELETE CASCADE ON UPDATE RESTRICT,
  CONSTRAINT `favoritevideo` FOREIGN KEY (`video_id`) REFERENCES `Video` (`video_id`) ON DELETE CASCADE ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for Relation
-- ----------------------------
DROP TABLE IF EXISTS `Relation`;
CREATE TABLE `Relation`  (
  `relation_id` bigint NOT NULL AUTO_INCREMENT,
  `follow_id` bigint NOT NULL,
  `follower_id` bigint NOT NULL,
  PRIMARY KEY (`relation_id`) USING BTREE,
  INDEX `FollowId`(`follow_id`) USING BTREE,
  INDEX `FollowerId`(`follower_id`) USING BTREE,
  CONSTRAINT `followerid` FOREIGN KEY (`follower_id`) REFERENCES `User` (`user_id`) ON DELETE CASCADE ON UPDATE RESTRICT,
  CONSTRAINT `followid` FOREIGN KEY (`follow_id`) REFERENCES `User` (`user_id`) ON DELETE CASCADE ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for User
-- ----------------------------
DROP TABLE IF EXISTS `User`;
CREATE TABLE `User`  (
  `user_id` bigint NOT NULL AUTO_INCREMENT,
  `user_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `follow_count` bigint NULL DEFAULT NULL,
  `follower_count` bigint NULL DEFAULT NULL,
  `token` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  PRIMARY KEY (`user_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for Video
-- ----------------------------
DROP TABLE IF EXISTS `Video`;
CREATE TABLE `Video`  (
  `video_id` bigint NOT NULL AUTO_INCREMENT,
  `author_id` bigint NOT NULL,
  `play_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `cover_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `favorite_count` bigint NULL DEFAULT NULL,
  `comment_count` bigint NULL DEFAULT NULL,
  `publish_time` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`video_id`) USING BTREE,
  INDEX `user`(`author_id`) USING BTREE,
  CONSTRAINT `authorid` FOREIGN KEY (`author_id`) REFERENCES `User` (`user_id`) ON DELETE CASCADE ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

SET FOREIGN_KEY_CHECKS = 1;
