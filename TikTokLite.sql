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

 Date: 12/05/2022 11:01:22
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for Comment
-- ----------------------------
DROP TABLE IF EXISTS `Comment`;
CREATE TABLE `Comment`  (
  `commentId` bigint NOT NULL,
  `userId` bigint NOT NULL,
  `VideoId` bigint NOT NULL,
  `comment` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `time` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`commentId`) USING BTREE,
  INDEX `commentUser`(`userId`) USING BTREE,
  INDEX `commentVideo`(`VideoId`) USING BTREE,
  CONSTRAINT `commentUser` FOREIGN KEY (`userId`) REFERENCES `User` (`userId`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `commentVideo` FOREIGN KEY (`VideoId`) REFERENCES `Video` (`videoId`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for Favorite
-- ----------------------------
DROP TABLE IF EXISTS `Favorite`;
CREATE TABLE `Favorite`  (
  `favoriteId` bigint NOT NULL,
  `userId` bigint NOT NULL,
  `videoId` bigint NOT NULL,
  PRIMARY KEY (`favoriteId`) USING BTREE,
  INDEX `favoriteUser`(`userId`) USING BTREE,
  INDEX `favoriteVideo`(`videoId`) USING BTREE,
  CONSTRAINT `favoriteUser` FOREIGN KEY (`userId`) REFERENCES `User` (`userId`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `favoriteVideo` FOREIGN KEY (`videoId`) REFERENCES `Video` (`videoId`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for Relation
-- ----------------------------
DROP TABLE IF EXISTS `Relation`;
CREATE TABLE `Relation`  (
  `relationId` bigint NOT NULL,
  `followId` bigint NOT NULL,
  `followerId` bigint NOT NULL,
  PRIMARY KEY (`relationId`) USING BTREE,
  INDEX `FollowId`(`followId`) USING BTREE,
  INDEX `FollowerId`(`followerId`) USING BTREE,
  CONSTRAINT `FollowerId` FOREIGN KEY (`followerId`) REFERENCES `User` (`userId`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `FollowId` FOREIGN KEY (`followId`) REFERENCES `User` (`userId`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for User
-- ----------------------------
DROP TABLE IF EXISTS `User`;
CREATE TABLE `User`  (
  `userId` bigint NOT NULL,
  `userName` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `followCount` bigint NULL DEFAULT NULL,
  `followerCount` bigint NULL DEFAULT NULL,
  PRIMARY KEY (`userId`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for Video
-- ----------------------------
DROP TABLE IF EXISTS `Video`;
CREATE TABLE `Video`  (
  `videoId` bigint NOT NULL,
  `authorId` bigint NOT NULL,
  `playUrl` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `coverUrl` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `favoriteCount` bigint NULL DEFAULT NULL,
  `commentCount` bigint NULL DEFAULT NULL,
  `publishTime` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`videoId`) USING BTREE,
  INDEX `user`(`authorId`) USING BTREE,
  CONSTRAINT `user` FOREIGN KEY (`authorId`) REFERENCES `User` (`userId`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
