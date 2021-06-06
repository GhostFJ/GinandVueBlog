/*
 Navicat MySQL Data Transfer

 Source Server         : 本机
 Source Server Type    : MySQL
 Source Server Version : 50718
 Source Host           : localhost
 Source Database       : blog

 Target Server Type    : MySQL
 Target Server Version : 50718
 File Encoding         : utf-8

 Date: 01/03/2020 20:15:51 PM
*/

SET NAMES utf8;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
--  Table structure for `blog_article`
-- ----------------------------
DROP TABLE IF EXISTS `blog_article`;
CREATE TABLE `blog_article` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `tag_id` int unsigned DEFAULT '0' COMMENT '标签ID',
  `category_id` int unsigned DEFAULT '0' COMMENT '分类ID',
  `title` varchar(100) DEFAULT '' COMMENT '文章标题',
  `desc` varchar(255) DEFAULT '' COMMENT '简述',
  `content` text,
  `created_on` int DEFAULT NULL,
  `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
  `modified_on` int unsigned DEFAULT '0' COMMENT '修改时间',
  `modified_by` varchar(255) DEFAULT '' COMMENT '修改人',
  `deleted_on` int unsigned DEFAULT '0',
  `state` tinyint unsigned DEFAULT '1' COMMENT '状态 0为禁用1为启用',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COMMENT='文章管理';


-- ----------------------------
--  Records of `blog_article`
-- ----------------------------
BEGIN;
INSERT INTO `blog_article` VALUES ('1', '1', '2', '测试', '程序测试', '个人博客项目开发使用，API测试分析', '1575945565', 'test-created', '1575945829', 'test-created-edit', '0', '1');
COMMIT;



-- ----------------------------
--  Table structure for `blog_tag`
-- ----------------------------
DROP TABLE IF EXISTS `blog_tag`;
CREATE TABLE `blog_tag` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) DEFAULT '' COMMENT '标签名称',
  `created_on` int unsigned DEFAULT '0' COMMENT '创建时间',
  `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
  `modified_on` int unsigned DEFAULT '0' COMMENT '修改时间',
  `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
  `deleted_on` int unsigned DEFAULT '0',
  `state` tinyint unsigned DEFAULT '1' COMMENT '状态 0为禁用、1为启用',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8 COMMENT='文章标签管理';


-- ----------------------------
--  Records of `blog_tag`
-- ----------------------------

BEGIN;
INSERT INTO `blog_tag` VALUES 
('1', 'Golang', '1575937993', 'test', '0', '', '0', '1'), 
('2', 'Java', '1575940902', 'test', '0', '', '0', '0'), 
('3', 'Python', '1575940931', 'admin', '0', '', '0', '0'), 
('4', 'C/C++', '1575967899', 'admin', '0', '', '0', '1');
COMMIT;



-- ----------------------------
--  Table structure for `category`
-- ----------------------------
DROP TABLE IF EXISTS `blog_category`;
CREATE TABLE `blog_category` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) DEFAULT '' COMMENT '分类名称',
  `created_on` int unsigned DEFAULT '0' COMMENT '创建时间',
  `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
  `modified_on` int unsigned DEFAULT '0' COMMENT '修改时间',
  `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
  `deleted_on` int unsigned DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8 COMMENT='文章分类管理';


-- ----------------------------
--  Records of `blog_category`
-- ----------------------------

BEGIN;
INSERT INTO `blog_category` VALUES 
('1', '软件工程', '1575937993', 'test', '0', '', '0'), 
('2', '计算机网络', '1575940902', 'test', '0', '', '0'), 
('3', '嵌入式开发', '1575940931', 'test', '0', '', '0'), 
('4', '生活', '1575967899', 'admin', '0', '', '0');
COMMIT;


-- ----------------------------
--  Table structure for `blog_comments`
-- ----------------------------
DROP TABLE IF EXISTS `blog_comments`;
CREATE TABLE `blog_comments` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `article_id` int unsigned DEFAULT '0' COMMENT '文章ID',
  `user_id` int unsigned DEFAULT '0' COMMENT '用户ID',
  `content` text COMMENT '评论内容',
  `created_on` int unsigned DEFAULT '0' COMMENT '创建时间',
  `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
  `modified_on` int unsigned DEFAULT '0' COMMENT '修改时间',
  `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
  `deleted_on` int unsigned DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 COMMENT='评论管理';

-- ----------------------------
--  Records of `blog_comments`
-- ----------------------------
BEGIN;
INSERT INTO `blog_comments` VALUES 
('1', '1', '1', 'test', '0', '', '0', '', '0'),
('2', '1', '1', '写的不错', '0', '', '0', '', '0');
COMMIT;



-- ----------------------------
--  Table structure for `blog_user`
-- ----------------------------
DROP TABLE IF EXISTS `blog_user`;
CREATE TABLE `blog_user` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(50) DEFAULT '' COMMENT '账号',
  `password` varchar(50) DEFAULT '' COMMENT '密码',
  `secret` varchar(20) NOT NULL DEFAULT '' COMMENT 'jwt动态密钥 注销，修改密码时候 改变',
  `status` int DEFAULT '1' COMMENT '状态: 1 正常 0 软删 -1',
  `created_on` int unsigned DEFAULT NULL COMMENT '创建时间',
  `modified_on` int unsigned DEFAULT NULL COMMENT '更新时间',
  `deleted_on` int unsigned DEFAULT '0' COMMENT '删除时间戳',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COMMENT='用户管理';


-- ----------------------------
--  Records of `blog_user`
--  admin default password 123456 
-- ----------------------------
BEGIN;
INSERT INTO `blog_user` VALUES ('1', 'admin', 'e10adc3949ba59abbe56e057f20f883e', '', '1', '0', '0', '0');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
