/*
 Navicat MySQL Data Transfer

 Source Server         : Toasobi的连接
 Source Server Type    : MySQL
 Source Server Version : 80029
 Source Host           : localhost:3306
 Source Schema         : lanshan

 Target Server Type    : MySQL
 Target Server Version : 80029
 File Encoding         : 65001

 Date: 26/07/2022 11:16:39
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for availabilities
-- ----------------------------
DROP TABLE IF EXISTS `availabilities`;
CREATE TABLE `availabilities`  (
  `Id_of_clothes` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '服装编码',
  `Id_of_providers` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '供应商编码',
  `Level` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '服装质量等级',
  PRIMARY KEY (`Id_of_clothes`, `Id_of_providers`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of availabilities
-- ----------------------------
INSERT INTO `availabilities` VALUES ('A1', 'P1', '合格');
INSERT INTO `availabilities` VALUES ('A2', 'P1', '不合格');

-- ----------------------------
-- Table structure for cloths
-- ----------------------------
DROP TABLE IF EXISTS `cloths`;
CREATE TABLE `cloths`  (
  `Id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '服装编码',
  `Size` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '服装尺码',
  `Price` int NULL DEFAULT NULL COMMENT '销售价格',
  `Variety` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '服装类型',
  PRIMARY KEY (`Id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '服装信息表\r\n' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of cloths
-- ----------------------------
INSERT INTO `cloths` VALUES ('A1', 'S', 50, 'A');
INSERT INTO `cloths` VALUES ('A2', 'S', 80, 'A');
INSERT INTO `cloths` VALUES ('B1', 'M', 100, 'B');

-- ----------------------------
-- Table structure for stores
-- ----------------------------
DROP TABLE IF EXISTS `stores`;
CREATE TABLE `stores`  (
  `Id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '仓库编码',
  `Capacity` int NULL DEFAULT NULL COMMENT '仓库容量',
  PRIMARY KEY (`Id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '仓库信息表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of stores
-- ----------------------------
INSERT INTO `stores` VALUES ('A', 1000);
INSERT INTO `stores` VALUES ('B', 10000);

-- ----------------------------
-- Table structure for suppliers
-- ----------------------------
DROP TABLE IF EXISTS `suppliers`;
CREATE TABLE `suppliers`  (
  `Id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '供应商编码',
  `Name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '供应商名称',
  PRIMARY KEY (`Id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of suppliers
-- ----------------------------
INSERT INTO `suppliers` VALUES ('P1', '腾讯');

SET FOREIGN_KEY_CHECKS = 1;
