/*
Navicat MySQL Data Transfer

Source Server         : 203.3.230.50
Source Server Version : 50720
Source Host           : 203.3.230.50:3306
Source Database       : allinone

Target Server Type    : MYSQL
Target Server Version : 50720
File Encoding         : 65001

Date: 2022-07-08 15:44:09
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for file_info
-- ----------------------------
DROP TABLE IF EXISTS `file_info`;
CREATE TABLE `file_info` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `sys_batch_id` int(11) DEFAULT NULL,
  `file_list_id` int(11) DEFAULT NULL,
  `filesize` int(11) DEFAULT NULL,
  `filemodtime` int(11) DEFAULT NULL,
  `status` int(1) DEFAULT NULL COMMENT '1,成功获取到.2,失败',
  `statusmsg` varchar(255) DEFAULT NULL,
  `filename` varchar(255) DEFAULT NULL,
  `filemd5` varchar(255) DEFAULT NULL,
  `ipaddr` varchar(255) DEFAULT NULL,
  `user` varchar(255) DEFAULT NULL,
  `creat_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `index_id` (`id`),
  KEY `index_ipaddr` (`ipaddr`),
  KEY `index_create` (`creat_at`)
) ENGINE=InnoDB AUTO_INCREMENT=62 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of file_info
-- ----------------------------
INSERT INTO `file_info` VALUES ('47', '11', '8', '0', '0', '2', 'Process exited with status 1', '/home/monitor/aaa.txt', '', '203.3.230.50', 'monitor', '2022-07-08 12:47:42');
INSERT INTO `file_info` VALUES ('48', '11', '6', '8621108', '1605606500', '1', '', '/opt/chaosblade/blade', '72d784aca3ecbf6ff5edaabba779596c', '203.3.254.197', 'root', '2022-07-08 12:47:44');
INSERT INTO `file_info` VALUES ('49', '11', '7', '0', '0', '2', 'ssh: handshake failed: ssh: unable to authenticate, attempted methods [none password], no supported methods remain', '/tmp/test/text.txt', '', '203.3.250.146', 'czbank', '2022-07-08 12:47:45');
INSERT INTO `file_info` VALUES ('50', '11', '5', '2476542', '1644203108', '1', '', '/opt/redis-6.2.6.tar.gz', 'f69ca96b39ca93001add922d558f9841', '203.3.254.197', 'root', '2022-07-08 12:47:45');
INSERT INTO `file_info` VALUES ('51', '11', '4', '180180204', '1646927836', '1', '', '/root/performance-web-console-0.0.1.jar', 'a8d32578f19d7cf0ac779091d9a739db', '203.3.254.197', 'root', '2022-07-08 12:47:46');
INSERT INTO `file_info` VALUES ('52', '11', '1', '7151742', '1649736860', '1', '', '/home/monitor/moco-runner-1.1.0-standalone.jar', 'b3d6b191ba2f3e4c02cf5c11aa7e22fc', '203.3.230.50', 'monitor', '2022-07-08 12:47:47');
INSERT INTO `file_info` VALUES ('53', '12', '8', '0', '0', '2', 'Process exited with status 1', '/home/monitor/aaa.txt', '', '203.3.230.50', 'monitor', '2022-07-08 12:48:29');
INSERT INTO `file_info` VALUES ('54', '12', '9', '0', '0', '2', 'Process exited with status 1', '/home/test/aasd.jj', 'test', '203.3.250.146', 'root', '2022-07-08 12:48:29');
INSERT INTO `file_info` VALUES ('55', '12', '6', '8621108', '1605606500', '1', '', '/opt/chaosblade/blade', '72d784aca3ecbf6ff5edaabba779596c', '203.3.254.197', 'root', '2022-07-08 12:48:30');
INSERT INTO `file_info` VALUES ('56', '12', '5', '2476542', '1644203108', '1', '', '/opt/redis-6.2.6.tar.gz', 'f69ca96b39ca93001add922d558f9842', '203.3.254.197', 'root', '2022-07-08 12:48:31');
INSERT INTO `file_info` VALUES ('57', '12', '7', '0', '0', '2', 'ssh: handshake failed: ssh: unable to authenticate, attempted methods [none password], no supported methods remain', '/tmp/test/text.txt', '', '203.3.250.146', 'czbank', '2022-07-08 12:48:31');
INSERT INTO `file_info` VALUES ('58', '12', '4', '180180204', '1646927836', '1', '', '/root/performance-web-console-0.0.1.jar', 'a8d32578f19d7cf0ac779091d9a739db', '203.3.254.197', 'root', '2022-07-08 12:48:31');
INSERT INTO `file_info` VALUES ('59', '12', '1', '7151742', '1649736860', '1', '', '/home/monitor/moco-runner-1.1.0-standalone.jar', 'b3d6b191ba2f3e4c02cf5c11aa7e22fc', '203.3.230.50', 'monitor', '2022-07-08 12:48:32');
INSERT INTO `file_info` VALUES ('60', '12', '10', '1', '1', '1', 'asd', '/tese', 'ets2', '203.3.test', 'asd', '2022-07-08 12:57:08');
INSERT INTO `file_info` VALUES ('61', '11', '11', '2', '2', '2', 'asdas', '/11', '12', '192.168.', 'asd', '2022-07-08 12:59:06');

-- ----------------------------
-- Table structure for file_list
-- ----------------------------
DROP TABLE IF EXISTS `file_list`;
CREATE TABLE `file_list` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `retype` int(1) DEFAULT NULL COMMENT '1 模糊匹配',
  `status` int(1) DEFAULT NULL,
  `sys_ip_id` int(11) DEFAULT NULL,
  `sys_userpasswd_id` int(11) DEFAULT NULL,
  `filepath` varchar(255) DEFAULT NULL,
  `filename` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `index_id` (`id`),
  KEY `index_unit` (`sys_ip_id`,`sys_userpasswd_id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of file_list
-- ----------------------------
INSERT INTO `file_list` VALUES ('1', '2', '1', '1', '1', '/home/monitor/', 'moco-runner-*standalone.jar');
INSERT INTO `file_list` VALUES ('4', '1', '1', '2', '2', '/root', 'performance-web-console-0.0.1.jar');
INSERT INTO `file_list` VALUES ('5', '2', '1', '2', '2', '/opt/', 'redis-6.2.6.tar.gz');
INSERT INTO `file_list` VALUES ('6', '2', '1', '2', '2', '/opt/chaosblade', 'blade');
INSERT INTO `file_list` VALUES ('7', '2', '1', '3', '3', '/tmp/test', 'text.txt');
INSERT INTO `file_list` VALUES ('8', '1', '1', '1', '1', '/home/monitor', 'aaa.txt');
INSERT INTO `file_list` VALUES ('9', '1', '1', '3', '2', '/home/test', 'aasd.jj');

-- ----------------------------
-- Table structure for file_result
-- ----------------------------
DROP TABLE IF EXISTS `file_result`;
CREATE TABLE `file_result` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `sys_batch_id_a` int(11) DEFAULT NULL,
  `sys_batch_id_b` int(11) DEFAULT NULL,
  `result_uuid` varchar(255) DEFAULT NULL,
  `filename` varchar(255) DEFAULT NULL,
  `ipaddr` varchar(255) DEFAULT NULL,
  `filemd5_a` varchar(255) DEFAULT NULL,
  `filemd5_b` varchar(255) DEFAULT NULL,
  `creat_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `index_id` (`id`),
  KEY `index_result_id` (`result_uuid`)
) ENGINE=InnoDB AUTO_INCREMENT=23 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of file_result
-- ----------------------------
INSERT INTO `file_result` VALUES ('17', '11', '12', 'a2994140-ca8f-46e0-b499-3b56953d9328', '/opt/redis-6.2.6.tar.gz', '203.3.254.197', 'f69ca96b39ca93001add922d558f9841', 'f69ca96b39ca93001add922d558f9842', '2022-07-08 14:28:20');
INSERT INTO `file_result` VALUES ('18', '11', '0', 'a2994140-ca8f-46e0-b499-3b56953d9328', '/11', '192.168.', '12', '', '2022-07-08 14:28:20');
INSERT INTO `file_result` VALUES ('19', '0', '12', 'a2994140-ca8f-46e0-b499-3b56953d9328', '/home/test/aasd.jj', '203.3.250.146', '', 'test', '2022-07-08 14:28:20');
INSERT INTO `file_result` VALUES ('20', '11', '12', '02902cb8-38b3-42a9-aca9-3a6626af65de', '/opt/redis-6.2.6.tar.gz', '203.3.254.197', 'f69ca96b39ca93001add922d558f9841', 'f69ca96b39ca93001add922d558f9842', '2022-07-08 14:28:35');
INSERT INTO `file_result` VALUES ('21', '11', '0', '02902cb8-38b3-42a9-aca9-3a6626af65de', '/11', '192.168.', '12', '', '2022-07-08 14:28:35');
INSERT INTO `file_result` VALUES ('22', '0', '12', '02902cb8-38b3-42a9-aca9-3a6626af65de', '/tese', '203.3.test', '', 'ets2', '2022-07-08 14:28:35');

-- ----------------------------
-- Table structure for sys_batch
-- ----------------------------
DROP TABLE IF EXISTS `sys_batch`;
CREATE TABLE `sys_batch` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `creat_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_batch
-- ----------------------------
INSERT INTO `sys_batch` VALUES ('2', '2022-07-07 12:38:01');
INSERT INTO `sys_batch` VALUES ('4', '2022-07-07 12:38:21');
INSERT INTO `sys_batch` VALUES ('5', '2022-07-07 15:25:11');
INSERT INTO `sys_batch` VALUES ('6', '2022-07-07 15:28:18');
INSERT INTO `sys_batch` VALUES ('7', '2022-07-07 15:29:49');
INSERT INTO `sys_batch` VALUES ('8', '2022-07-08 08:48:35');
INSERT INTO `sys_batch` VALUES ('9', '2022-07-08 08:49:05');
INSERT INTO `sys_batch` VALUES ('10', '2022-07-08 08:52:36');
INSERT INTO `sys_batch` VALUES ('11', '2022-07-08 12:47:42');
INSERT INTO `sys_batch` VALUES ('12', '2022-07-08 12:48:28');

-- ----------------------------
-- Table structure for sys_env
-- ----------------------------
DROP TABLE IF EXISTS `sys_env`;
CREATE TABLE `sys_env` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `envname` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `index_id` (`id`,`envname`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_env
-- ----------------------------
INSERT INTO `sys_env` VALUES ('1', 'DEV');
INSERT INTO `sys_env` VALUES ('2', 'FAT');
INSERT INTO `sys_env` VALUES ('3', 'UAT');
INSERT INTO `sys_env` VALUES ('4', 'LPT');
INSERT INTO `sys_env` VALUES ('5', 'PRO');
INSERT INTO `sys_env` VALUES ('8', 'asdsdf');

-- ----------------------------
-- Table structure for sys_ip
-- ----------------------------
DROP TABLE IF EXISTS `sys_ip`;
CREATE TABLE `sys_ip` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `ipaddr` varchar(255) DEFAULT NULL,
  `sys_env_id` int(11) DEFAULT NULL,
  `sys_name_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `index_ipaddr` (`id`,`ipaddr`) USING BTREE,
  KEY `index_id` (`id`,`ipaddr`,`sys_name_id`,`sys_env_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_ip
-- ----------------------------
INSERT INTO `sys_ip` VALUES ('1', '203.3.230.50', '2', '7');
INSERT INTO `sys_ip` VALUES ('2', '203.3.254.197', '1', '5');
INSERT INTO `sys_ip` VALUES ('3', '203.3.250.146', '8', '2');

-- ----------------------------
-- Table structure for sys_name
-- ----------------------------
DROP TABLE IF EXISTS `sys_name`;
CREATE TABLE `sys_name` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `index_id` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_name
-- ----------------------------
INSERT INTO `sys_name` VALUES ('1', '测试系统');
INSERT INTO `sys_name` VALUES ('2', 'mygirl');
INSERT INTO `sys_name` VALUES ('3', 'allinone');
INSERT INTO `sys_name` VALUES ('5', '青锋利剑');
INSERT INTO `sys_name` VALUES ('6', '紫薇软剑');
INSERT INTO `sys_name` VALUES ('7', '玄铁重剑');
INSERT INTO `sys_name` VALUES ('8', '腐朽木剑');

-- ----------------------------
-- Table structure for sys_userpasswd
-- ----------------------------
DROP TABLE IF EXISTS `sys_userpasswd`;
CREATE TABLE `sys_userpasswd` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user` varchar(255) DEFAULT NULL,
  `passwd` varchar(255) DEFAULT NULL,
  `sys_ip_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `index_id` (`id`) USING BTREE,
  KEY `index_u_p` (`user`,`passwd`,`sys_ip_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_userpasswd
-- ----------------------------
INSERT INTO `sys_userpasswd` VALUES ('3', 'czbank', 'root', '3');
INSERT INTO `sys_userpasswd` VALUES ('1', 'monitor', 'Kydl@2020', '1');
INSERT INTO `sys_userpasswd` VALUES ('2', 'root', 'Kydl@2020', '2');
