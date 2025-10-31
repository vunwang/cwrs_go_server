/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 90200 (9.2.0)
 Source Host           : localhost:3306
 Source Schema         : cwrs

 Target Server Type    : MySQL
 Target Server Version : 90200 (9.2.0)
 File Encoding         : 65001

 Date: 30/10/2025 11:14:26
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for sys_data_scope
-- ----------------------------
DROP TABLE IF EXISTS `sys_data_scope`;
CREATE TABLE `sys_data_scope`  (
  `scope_id` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '主键',
  `scope_type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '数据权限控制方式(字典)',
  `menu_id` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '菜单id',
  `dept_id` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '授权组织',
  `role_id` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '角色id',
  `created_user_id` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '创建用户',
  `created_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_user_id` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '更新用户',
  `updated_time` datetime NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`scope_id`) USING BTREE,
  UNIQUE INDEX `only_purview_id`(`scope_id` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '自定义数据权限范围' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_data_scope
-- ----------------------------

-- ----------------------------
-- Table structure for sys_dept
-- ----------------------------
DROP TABLE IF EXISTS `sys_dept`;
CREATE TABLE `sys_dept`  (
  `dept_id` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '主键',
  `dept_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '组织名称',
  `parent_id` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '上级组织',
  `dept_level` varchar(768) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '组织层级',
  `dept_status` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '状态(字典)',
  `dept_sort` int NOT NULL COMMENT '排序',
  `created_user_id` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '创建用户',
  `created_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_user_id` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '更新用户',
  `updated_time` datetime NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`dept_id`) USING BTREE,
  UNIQUE INDEX `only_dept_id`(`dept_id` ASC) USING BTREE,
  INDEX `dept_level`(`dept_level` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '组织' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_dept
-- ----------------------------
INSERT INTO `sys_dept` VALUES ('128a4d4ca62b4a579094396a7f3ef00b', '研发二组', '8791cb9d85c2432ea258513ee561a92d', '0,12cb92fc10724b4ca762bb38bb624fd2,a2aa4cf0ba1f46049cbe5a5d966446e8,8791cb9d85c2432ea258513ee561a92d,128a4d4ca62b4a579094396a7f3ef00b', '1', 2, '953b77a21366497094303f293cc07424', '2025-08-12 10:35:58', NULL, NULL);
INSERT INTO `sys_dept` VALUES ('12cb92fc10724b4ca762bb38bb624fd2', 'XXXX科技有限公司', '0', '0,12cb92fc10724b4ca762bb38bb624fd2', '1', 1, '953b77a21366497094303f293cc07424', '2025-08-12 10:34:33', '953b77a21366497094303f293cc07424', '2025-10-29 18:15:42');
INSERT INTO `sys_dept` VALUES ('3ce5166aab7447728908e7c6f9608e8a', '研发一部', '8791cb9d85c2432ea258513ee561a92d', '0,12cb92fc10724b4ca762bb38bb624fd2,a2aa4cf0ba1f46049cbe5a5d966446e8,8791cb9d85c2432ea258513ee561a92d,3ce5166aab7447728908e7c6f9608e8a', '1', 1, '953b77a21366497094303f293cc07424', '2025-09-29 18:04:51', '953b77a21366497094303f293cc07424', '2025-09-29 18:05:01');
INSERT INTO `sys_dept` VALUES ('8791cb9d85c2432ea258513ee561a92d', '研发部', 'a2aa4cf0ba1f46049cbe5a5d966446e8', '0,12cb92fc10724b4ca762bb38bb624fd2,a2aa4cf0ba1f46049cbe5a5d966446e8,8791cb9d85c2432ea258513ee561a92d', '1', 1, '953b77a21366497094303f293cc07424', '2025-08-12 10:35:39', NULL, NULL);
INSERT INTO `sys_dept` VALUES ('a2aa4cf0ba1f46049cbe5a5d966446e8', '江苏总部', '12cb92fc10724b4ca762bb38bb624fd2', '0,12cb92fc10724b4ca762bb38bb624fd2,a2aa4cf0ba1f46049cbe5a5d966446e8', '1', 1, '953b77a21366497094303f293cc07424', '2025-08-12 10:34:44', '953b77a21366497094303f293cc07424', '2025-10-15 15:29:57');

-- ----------------------------
-- Table structure for sys_dict
-- ----------------------------
DROP TABLE IF EXISTS `sys_dict`;
CREATE TABLE `sys_dict`  (
  `dict_id` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '主键',
  `dict_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '名称',
  `dict_code` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '编码',
  `dict_status` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '状态(1正常 0禁用)',
  `dict_sort` int NOT NULL DEFAULT 0 COMMENT '排序',
  `desc` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '描述',
  `created_user_id` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '创建用户;所有表中都要有该字段，数据权限需要使用',
  `created_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_user_id` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '更新用户',
  `updated_time` datetime NULL DEFAULT NULL COMMENT '更新时间',
  UNIQUE INDEX `only_dict_code`(`dict_code` ASC) USING BTREE,
  UNIQUE INDEX `only_dict_name`(`dict_name` ASC) USING BTREE,
  UNIQUE INDEX `only_dict_id`(`dict_id` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '数据字典' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_dict
-- ----------------------------
INSERT INTO `sys_dict` VALUES ('8ab8c45424d542adbb2cd5085a875e0b', '数据权限', 'sys_data_purview', '1', 4, '数据权限all-全部 this_dept本部门 this_dept_and_below本部门及以下 only_me仅自己 customise自定义', '953b77a21366497094303f293cc07424', '2025-09-10 10:49:33', '953b77a21366497094303f293cc07424', '2025-09-10 10:55:49');
INSERT INTO `sys_dict` VALUES ('530b6fbf503942fea608a7117deace69', '数据权限控制方式', 'sys_data_purview_type', '0', 20, '数据权限控制方式1-菜单', '953b77a21366497094303f293cc07424', '2025-09-10 11:41:44', '953b77a21366497094303f293cc07424', '2025-09-10 18:26:57');
INSERT INTO `sys_dict` VALUES ('c489a70ee5e34226aeaf5521ca9a7425', '性别', 'sys_gender', '1', 2, '性别 1-男 2-女 3保密', '953b77a21366497094303f293cc07424', '2025-08-05 17:39:28', NULL, NULL);
INSERT INTO `sys_dict` VALUES ('3f3b1076e1c341fdbe965fb328945374', '菜单类型', 'sys_menu_type', '1', 3, '菜单类型 1-目录 2-菜单 3按钮', '953b77a21366497094303f293cc07424', '2025-07-28 17:09:03', '953b77a21366497094303f293cc07424', '2025-08-05 17:39:32');
INSERT INTO `sys_dict` VALUES ('80ab9f8cd75b4d309c73c8f42393e715', '状态', 'sys_status', '1', 1, '状态1-正常 2-禁用', '953b77a21366497094303f293cc07424', '2025-07-28 17:08:13', '953b77a21366497094303f293cc07424', '2025-10-15 15:30:20');
INSERT INTO `sys_dict` VALUES ('5cd8b511b05c4c2fa3260a119b314eb1', '定时任务cron表达式', 'sys_task_cron_expr', '1', 7, '定时任务执行时间cron表达式', '953b77a21366497094303f293cc07424', '2025-09-18 14:39:54', '953b77a21366497094303f293cc07424', '2025-09-18 15:12:51');
INSERT INTO `sys_dict` VALUES ('4e0b57eb18f3446299d8265f3b4adb1e', '任务状态', 'sys_task_status', '1', 6, '定时任务状态', '953b77a21366497094303f293cc07424', '2025-09-18 15:13:21', NULL, NULL);

-- ----------------------------
-- Table structure for sys_dict_item
-- ----------------------------
DROP TABLE IF EXISTS `sys_dict_item`;
CREATE TABLE `sys_dict_item`  (
  `dict_item_id` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '主键',
  `dict_code` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '字典编码',
  `item_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '名称',
  `item_value` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '字典值',
  `item_color` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '展示颜色',
  `item_status` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '状态(字典 1正常 0禁用)',
  `item_select` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '下拉菜单是否显示(1显示 0隐藏)',
  `item_sort` int NULL DEFAULT NULL COMMENT '排序',
  `created_user_id` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '创建用户',
  `created_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_user_id` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '更新用户',
  `updated_time` datetime NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`dict_item_id`) USING BTREE,
  UNIQUE INDEX `only_item_id`(`dict_item_id` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '字典项' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_dict_item
-- ----------------------------
INSERT INTO `sys_dict_item` VALUES ('1091ba541f344e0c89a1975e9cb63797', 'sys_menu_type', '菜单', '2', '#00B42A', '1', '1', 2, '953b77a21366497094303f293cc07424', '2025-07-28 14:48:29', '953b77a21366497094303f293cc07424', '2025-08-13 14:45:42');
INSERT INTO `sys_dict_item` VALUES ('27f78f25b3d44d5b9cbb5a43d57b0110', 'sys_task_cron_expr', '每月1号凌晨1点执行', '每月1号凌晨1点执行,0 0 1 1 * *', '#F1590E', '1', '1', 12, '953b77a21366497094303f293cc07424', '2025-09-19 14:35:57', NULL, NULL);
INSERT INTO `sys_dict_item` VALUES ('286221b88ef8489599a9129224588b74', 'sys_data_purview', '本组织及以下', 'this_dept_and_below', '#F1590E', '1', '1', 3, '953b77a21366497094303f293cc07424', '2025-09-10 10:57:42', '953b77a21366497094303f293cc07424', '2025-10-10 11:02:29');
INSERT INTO `sys_dict_item` VALUES ('2c6954594d8c4a72bd0af1eb5942ee3a', 'sys_task_cron_expr', '每天凌晨2点执行', '每天凌晨2点执行,0 0 2 * * *', '#F1590E', '1', '1', 7, '953b77a21366497094303f293cc07424', '2025-09-18 17:03:20', NULL, NULL);
INSERT INTO `sys_dict_item` VALUES ('2da398035fc74ee99656aac3ef0668af', 'sys_status', '禁用', '0', 'rgb(255, 140, 0)', '1', '1', 2, '953b77a21366497094303f293cc07424', '2025-07-24 15:19:12', '953b77a21366497094303f293cc07424', '2025-09-30 15:29:27');
INSERT INTO `sys_dict_item` VALUES ('338f7dd676754b27818e2c8c8e0db2e2', 'sys_task_status', '停止', '2', '#F1590E', '1', '1', 2, '953b77a21366497094303f293cc07424', '2025-09-18 15:13:55', NULL, NULL);
INSERT INTO `sys_dict_item` VALUES ('37bf6ed5c7644cbeb43326bfde419255', 'sys_task_cron_expr', '每年1月1日0点执行', '每年1月1日0点执行,0 0 0 1 1 *', '#F1590E', '1', '1', 13, '953b77a21366497094303f293cc07424', '2025-09-19 14:37:05', NULL, NULL);
INSERT INTO `sys_dict_item` VALUES ('3ec1457ba5104c7c815dd4edb805a2ca', 'sys_menu_type', '目录', '1', '#FF7D00', '1', '1', 1, '953b77a21366497094303f293cc07424', '2025-07-28 14:45:36', '953b77a21366497094303f293cc07424', '2025-08-13 14:45:45');
INSERT INTO `sys_dict_item` VALUES ('52611e0e9d7e418b8744bf4e76fdf3a1', 'sys_gender', '保密', '3', '#9FD4FD', '1', '1', 3, '953b77a21366497094303f293cc07424', '2025-08-05 17:40:30', '953b77a21366497094303f293cc07424', '2025-08-13 14:46:20');
INSERT INTO `sys_dict_item` VALUES ('58438420b04344089bf1df6fb34b4d1e', 'sys_data_purview', '仅自己', 'only_me', '#93BEFF', '1', '1', 4, '953b77a21366497094303f293cc07424', '2025-09-10 10:58:45', NULL, NULL);
INSERT INTO `sys_dict_item` VALUES ('5ab1ad43a3774f758a18a96f1ff3ec3f', 'sys_task_cron_expr', '每5分钟执行一次', '每5分钟执行一次,0 */5 * * * *', '#F1590E', '1', '1', 5, '953b77a21366497094303f293cc07424', '2025-09-18 16:59:41', '953b77a21366497094303f293cc07424', '2025-09-18 17:00:56');
INSERT INTO `sys_dict_item` VALUES ('5bf10b38e70a4ad79e551e4c19d6d7b5', 'sys_gender', '女', '2', '#FF7D00', '1', '1', 2, '953b77a21366497094303f293cc07424', '2025-08-05 17:40:07', '953b77a21366497094303f293cc07424', '2025-08-13 14:46:17');
INSERT INTO `sys_dict_item` VALUES ('5c9f3e8c500c48969cfd40f4bb6ecd4d', 'sys_data_purview', '全部', 'all', '#00B42A', '1', '1', 1, '953b77a21366497094303f293cc07424', '2025-09-10 10:56:09', '953b77a21366497094303f293cc07424', '2025-10-11 17:42:03');
INSERT INTO `sys_dict_item` VALUES ('5ea54337ccbb480cb062e994ca7e0eeb', 'sys_task_cron_expr', '每周一上午9点执行', '每周一上午9点执行,0 0 9 * * 1', '#F1590E', '1', '1', 11, '953b77a21366497094303f293cc07424', '2025-09-18 17:09:33', NULL, NULL);
INSERT INTO `sys_dict_item` VALUES ('6f67d63e5a3d4df8a6af75c6ce0ad7cc', 'sys_data_purview', '本组织', 'this_dept', '#3C7EFF', '1', '1', 2, '953b77a21366497094303f293cc07424', '2025-09-10 10:56:33', '953b77a21366497094303f293cc07424', '2025-10-10 11:02:21');
INSERT INTO `sys_dict_item` VALUES ('6f97244e39484993aa8cd7957aa67fdf', 'sys_gender', '男', '1', '#9FDB1D', '1', '1', 1, '953b77a21366497094303f293cc07424', '2025-08-05 17:39:54', '953b77a21366497094303f293cc07424', '2025-08-13 14:46:14');
INSERT INTO `sys_dict_item` VALUES ('7998755d3a3b40198d769cb8b07a122b', 'sys_task_cron_expr', '每天上午9点执行', '每天上午9点执行,0 0 9 * * *', '#F1590E', '1', '1', 9, '953b77a21366497094303f293cc07424', '2025-09-18 17:07:03', NULL, NULL);
INSERT INTO `sys_dict_item` VALUES ('79da4df4188d4a4783152d8c92e42bd9', 'sys_data_purview_type', '菜单', '1', '#F1590E', '1', '1', 1, '953b77a21366497094303f293cc07424', '2025-09-10 18:26:45', NULL, NULL);
INSERT INTO `sys_dict_item` VALUES ('7e0e2d2ca3114197a1f541e5ed89605b', 'sys_task_cron_expr', '每天中午12点执行', '每天中午12点执行,0 0 12 * * *', '#F1590E', '1', '1', 10, '953b77a21366497094303f293cc07424', '2025-09-18 17:05:25', '953b77a21366497094303f293cc07424', '2025-09-18 17:05:55');
INSERT INTO `sys_dict_item` VALUES ('9bce3292bf884f70a4aec3b6006cc6d7', 'sys_task_cron_expr', '每小时执行一次', '每小时执行一次,0 0 * * * *', '#F1590E', '1', '1', 6, '953b77a21366497094303f293cc07424', '2025-09-18 17:00:36', NULL, NULL);
INSERT INTO `sys_dict_item` VALUES ('a19b711890df41d7b842eba53bea18d2', 'sys_task_cron_expr', '每10秒执行一次', '每10秒执行一次,*/10 * * * * *', '#F1590E', '1', '1', 2, '953b77a21366497094303f293cc07424', '2025-09-18 16:56:54', '953b77a21366497094303f293cc07424', '2025-09-18 17:01:29');
INSERT INTO `sys_dict_item` VALUES ('a8fdf414073c48e482073312f484872e', 'sys_task_cron_expr', '每天上午8点执行', '每天上午8点执行,0 0 8 * * *', '#F1590E', '1', '1', 8, '953b77a21366497094303f293cc07424', '2025-09-18 17:04:31', '953b77a21366497094303f293cc07424', '2025-09-18 17:07:24');
INSERT INTO `sys_dict_item` VALUES ('ac8bf13e3279447abcfce1f018d649a3', 'sys_task_cron_expr', '每分钟执行一次', '每分钟执行一次,0 * * * * *', '#F1590E', '1', '1', 4, '953b77a21366497094303f293cc07424', '2025-09-18 16:58:05', '953b77a21366497094303f293cc07424', '2025-09-18 17:01:08');
INSERT INTO `sys_dict_item` VALUES ('b9a75b5802ac4cf1a34a03a1acf92a64', 'sys_menu_type', '按钮', '3', '#3C7EFF', '1', '1', 3, '953b77a21366497094303f293cc07424', '2025-07-28 14:49:02', '953b77a21366497094303f293cc07424', '2025-08-13 15:07:45');
INSERT INTO `sys_dict_item` VALUES ('ba0521595c1e440d8bad2061236e7343', 'sys_task_status', '运行中', '1', '#00B42A', '1', '1', 1, '953b77a21366497094303f293cc07424', '2025-09-18 15:13:44', NULL, NULL);
INSERT INTO `sys_dict_item` VALUES ('be1cf78e8f894846b7f6c64dba0d4430', 'sys_status', '正常', '1', '#00B42A', '1', '1', 1, '953b77a21366497094303f293cc07424', '2025-07-24 15:12:42', '953b77a21366497094303f293cc07424', '2025-10-15 15:30:24');
INSERT INTO `sys_dict_item` VALUES ('c7bc3ebea1b842da8527a4c8ed826c94', 'sys_task_cron_expr', '每秒执行一次', '每秒执行一次,* * * * * *', '#F1590E', '1', '1', 1, '953b77a21366497094303f293cc07424', '2025-09-18 16:54:52', '953b77a21366497094303f293cc07424', '2025-09-18 17:01:41');
INSERT INTO `sys_dict_item` VALUES ('f6c54cecd6dc48898759878730aac7cc', 'sys_data_purview', '自定义', 'custom', '#9FDB1D', '1', '1', 5, '953b77a21366497094303f293cc07424', '2025-09-10 10:59:33', NULL, NULL);
INSERT INTO `sys_dict_item` VALUES ('f874836d57624efda2fc7756020a7015', 'sys_task_cron_expr', '每30秒执行一次', '每30秒执行一次,*/30* * * * *', '#F1590E', '1', '1', 3, '953b77a21366497094303f293cc07424', '2025-09-18 16:57:22', '953b77a21366497094303f293cc07424', '2025-09-18 17:01:21');

-- ----------------------------
-- Table structure for sys_login_log
-- ----------------------------
DROP TABLE IF EXISTS `sys_login_log`;
CREATE TABLE `sys_login_log`  (
  `info_id` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '访问ID',
  `login_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '登录账号',
  `ipaddr` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '登录IP地址',
  `login_location` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '登录地点',
  `browser` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '浏览器类型',
  `os` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '操作系统',
  `status` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '0' COMMENT '登录状态（0成功 1失败）',
  `msg` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '提示消息',
  `login_time` datetime NULL DEFAULT NULL COMMENT '登录时间',
  `module` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '登录模块',
  PRIMARY KEY (`info_id`) USING BTREE,
  UNIQUE INDEX `info_id`(`info_id` ASC) USING BTREE COMMENT 'id唯一索引'
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '系统访问记录' ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of sys_login_log
-- ----------------------------

-- ----------------------------
-- Table structure for sys_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_menu`;
CREATE TABLE `sys_menu`  (
  `menu_id` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '主键',
  `parent_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '0' COMMENT '上级菜单',
  `path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '路由路径',
  `android_path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'android路由',
  `ios_path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'ios路由',
  `component` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '组件路径',
  `android_component` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'android组件路径',
  `ios_component` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'ios组件路径',
  `redirect` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '重定向',
  `type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '类型(1目录2菜单3按钮)',
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '菜单标题',
  `svg_icon` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '自定义图标(优先显示)',
  `icon` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '菜单图标',
  `hidden` tinyint NULL DEFAULT NULL COMMENT '是否隐藏(1-true 0-false)',
  `keep_alive` tinyint NULL DEFAULT NULL COMMENT '是否缓存(1-true 0-false)',
  `breadcrumb` tinyint NULL DEFAULT NULL COMMENT '面包屑(1-true 0-false)',
  `always_show` tinyint NULL DEFAULT NULL COMMENT '目录-总是显示(1-true 0-false)',
  `showIn_tabs` tinyint NULL DEFAULT NULL COMMENT '菜单-页签显示(1-true 0-false)',
  `affix` tinyint NULL DEFAULT NULL COMMENT '菜单-是否固定在标签页(1-true 0-false)',
  `status` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '状态(1-启用 0-禁用)',
  `active_menu` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '激活的菜单项',
  `permission` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '按钮-权限标识',
  `client_type` int NULL DEFAULT NULL COMMENT '客户端类型(1-平台 2-APP)',
  `sort` int NULL DEFAULT NULL COMMENT '菜单排序',
  `created_user_id` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '创建用户',
  `created_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_user_id` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '更新用户',
  `updated_time` datetime NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`menu_id`) USING BTREE,
  UNIQUE INDEX `only_menu_id`(`menu_id` ASC) USING BTREE,
  UNIQUE INDEX `only_permission`(`permission` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '菜单' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_menu
-- ----------------------------
INSERT INTO `sys_menu` VALUES ('03581a90c87d4fd18bdbc77a7a8b7d42', 'c28884e311364d4abd7a13ab3fd89bdb', NULL, NULL, NULL, NULL, NULL, NULL, NULL, '3', '字典编辑', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, '1', NULL, 'sys:dict:edit', 1, 2, '953b77a21366497094303f293cc07424', '2025-07-30 10:26:27', NULL, NULL);
INSERT INTO `sys_menu` VALUES ('09051c22d95047c38aa196ea9e9b6e73', '9b316608d8494955807dddb160c42857', NULL, '2222', '22', NULL, '22', '22', NULL, '2', '222', NULL, 'https://plcsh.oss-cn-shanghai.aliyuncs.com/images/202509/15162542aJzHvH.png', NULL, NULL, NULL, NULL, NULL, NULL, '1', NULL, NULL, 2, 1, '953b77a21366497094303f293cc07424', '2025-09-15 16:25:51', '953b77a21366497094303f293cc07424', '2025-09-25 17:05:34');
INSERT INTO `sys_menu` VALUES ('0a16a6b6abd24d35afe3d7d28cc41171', 'd92867e8d4b443c3ba07024cf1f83def', NULL, NULL, NULL, NULL, NULL, NULL, NULL, '3', '列表是否缓存调整', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, '1', NULL, 'sys:menu:keepAlive', 1, 5, '953b77a21366497094303f293cc07424', '2025-08-13 16:00:25', NULL, NULL);
INSERT INTO `sys_menu` VALUES ('0b087f935c0c44139b0cc18fe1862dda', '90695bd4b3af406a94c54170515c3b5c', NULL, NULL, NULL, NULL, NULL, NULL, NULL, '3', '新增', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, '1', NULL, 'sys:timer:add', 1, 1, '953b77a21366497094303f293cc07424', '2025-09-29 16:52:12', NULL, NULL);
INSERT INTO `sys_menu` VALUES ('0c2ec443c2fa43eba2521ca4bdd5ffaa', '7a39fb17ab654a1da60b93d73b422789', NULL, NULL, NULL, NULL, NULL, NULL, NULL, '3', '新增', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, '1', NULL, 'sys:user:add', 1, 1, '953b77a21366497094303f293cc07424', '2025-10-13 17:40:12', NULL, NULL);
INSERT INTO `sys_menu` VALUES ('0c36b3a959b343eda48e696e0fb7dfaf', 'fc1c718052ef4490b8dfaf83f2185dc2', '/system/post', NULL, NULL, 'system/post/index', NULL, NULL, '', '2', '岗位管理', '', 'IconUserGroup', 0, 1, 1, NULL, 1, 0, '1', NULL, NULL, 1, 6, '953b77a21366497094303f293cc07424', '2025-07-30 17:54:33', '953b77a21366497094303f293cc07424', '2025-07-30 17:54:47');
INSERT INTO `sys_menu` VALUES ('1672341312774125a340b04f9caabbac', '2785b90df06d4c299225234165a17298', NULL, NULL, NULL, NULL, NULL, NULL, NULL, '3', '添加', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, '1', NULL, 'sys:btn:add', 2, 1, '953b77a21366497094303f293cc07424', '2025-09-09 10:39:29', '953b77a21366497094303f293cc07424', '2025-09-09 10:40:10');
INSERT INTO `sys_menu` VALUES ('19ababda3a604d308677eed597a11add', '7bb71cc033034b9589a571ded86e2dc8', NULL, NULL, NULL, NULL, NULL, NULL, NULL, '3', '新增', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, '1', NULL, 'sys:param:add', 1, 1, '953b77a21366497094303f293cc07424', '2025-09-29 16:41:59', '953b77a21366497094303f293cc07424', '2025-09-29 16:42:13');
INSERT INTO `sys_menu` VALUES ('1bab070bc46648639ed916bd1ce5e807', 'c28884e311364d4abd7a13ab3fd89bdb', NULL, NULL, NULL, NULL, NULL, NULL, NULL, '3', '字典新增', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, '1', NULL, 'sys:dict:add', 1, 1, '953b77a21366497094303f293cc07424', '2025-07-30 10:25:45', NULL, NULL);
INSERT INTO `sys_menu` VALUES ('1d91eb70b45f4d5ea8f279323c8caafa', 'c28884e311364d4abd7a13ab3fd89bdb', NULL, NULL, NULL, NULL, NULL, NULL, NULL, '3', '列表字典项状态调整', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, '1', NULL, 'sys:dictItem:itemStatus', 1, 7, '953b77a21366497094303f293cc07424', '2025-08-13 16:07:07', NULL, NULL);
INSERT INTO `sys_menu` VALUES ('1ea410dae6e34503b539123eacf7e33b', '7bb71cc033034b9589a571ded86e2dc8', NULL, NULL, NULL, NULL, NULL, NULL, NULL, '3', '全局参数', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, '1', NULL, 'sys:param:tab', 1, 4, '953b77a21366497094303f293cc07424', '2025-10-14 18:17:43', NULL, NULL);
INSERT INTO `sys_menu` VALUES ('2304e56278124d2c852fed8ba5219ce5', 'fc1c718052ef4490b8dfaf83f2185dc2', '/system/account', NULL, NULL, 'system/account/index', NULL, NULL, '', '2', '用户账户', '', 'IconTool', 1, 0, 1, NULL, 0, 0, '1', NULL, NULL, 1, 11, '953b77a21366497094303f293cc07424', '2025-07-29 14:21:14', '953b77a21366497094303f293cc07424', '2025-10-30 09:38:42');
INSERT INTO `sys_menu` VALUES ('2785b90df06d4c299225234165a17298', '41c8df9c14964a8db858a113e0cd9c78', NULL, '/lib', '/bin', NULL, '123', '456', NULL, '2', '测试菜单lss', NULL, 'https://plcsh.oss-cn-shanghai.aliyuncs.com/images/202509/15162610vQcNZf.jpg', NULL, NULL, NULL, NULL, NULL, NULL, '1', NULL, NULL, 2, 74, '953b77a21366497094303f293cc07424', '2025-09-03 14:10:21', '953b77a21366497094303f293cc07424', '2025-09-25 17:06:36');
INSERT INTO `sys_menu` VALUES ('30ab7879300841e29f725a934fd82ba8', '7bb71cc033034b9589a571ded86e2dc8', NULL, NULL, NULL, NULL, NULL, NULL, NULL, '3', '编辑', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, '1', NULL, 'sys:param:edit', 1, 2, '953b77a21366497094303f293cc07424', '2025-09-29 16:42:33', NULL, NULL);
INSERT INTO `sys_menu` VALUES ('328b482b66b24db2864f0e7e6e0469c8', 'c28884e311364d4abd7a13ab3fd89bdb', NULL, NULL, NULL, NULL, NULL, NULL, NULL, '3', '字典项新增', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, '1', NULL, 'sys:dictItem:add', 1, 4, '953b77a21366497094303f293cc07424', '2025-07-30 10:30:07', '953b77a21366497094303f293cc07424', '2025-07-30 10:36:13');
INSERT INTO `sys_menu` VALUES ('41c8df9c14964a8db858a113e0cd9c78', '0', NULL, '/9/9/9', '777', NULL, '999', '888', NULL, '1', 'tp001', NULL, 'https://plcsh.oss-cn-shanghai.aliyuncs.com/images/202509/15161331TqrMjA.jpg', NULL, NULL, NULL, NULL, NULL, NULL, '1', NULL, NULL, 2, 1, '953b77a21366497094303f293cc07424', '2025-09-12 09:42:16', '953b77a21366497094303f293cc07424', '2025-09-25 17:06:25');
INSERT INTO `sys_menu` VALUES ('45af68a141a74f298fb4c00a9e992c70', '90695bd4b3af406a94c54170515c3b5c', NULL, NULL, NULL, NULL, NULL, NULL, NULL, '3', '删除', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, '1', NULL, 'sys:timer:del', 1, 3, '953b77a21366497094303f293cc07424', '2025-09-29 16:52:43', NULL, NULL);
INSERT INTO `sys_menu` VALUES ('477f429a7f4e486c9204fd5f8f058a61', '7a39fb17ab654a1da60b93d73b422789', NULL, NULL, NULL, NULL, NULL, NULL, NULL, '3', '编辑', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, '1', NULL, 'sys:user:edit', 1, 2, '953b77a21366497094303f293cc07424', '2025-10-13 17:40:40', NULL, NULL);
INSERT INTO `sys_menu` VALUES ('491ea08df95f47488680bce1d68f433a', '90695bd4b3af406a94c54170515c3b5c', NULL, NULL, NULL, NULL, NULL, NULL, NULL, '3', '修改', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, '1', NULL, 'sys:timer:edit', 1, 2, '953b77a21366497094303f293cc07424', '2025-09-29 16:52:32', NULL, NULL);
INSERT INTO `sys_menu` VALUES ('514a8e2cf3694e108c13c9ebd45282e8', 'c28884e311364d4abd7a13ab3fd89bdb', NULL, NULL, NULL, NULL, NULL, NULL, NULL, '3', '列表字典项下拉菜单状态调整', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, '1', NULL, 'sys:dictItem:itemSelect', 1, 8, '953b77a21366497094303f293cc07424', '2025-08-13 16:07:37', NULL, NULL);
INSERT INTO `sys_menu` VALUES ('6475177fa06145fbb0b2d10af62f4baf', '0c36b3a959b343eda48e696e0fb7dfaf', NULL, NULL, NULL, NULL, NULL, NULL, NULL, '3', '新增', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, '1', NULL, 'sys:post:add', 1, 1, '953b77a21366497094303f293cc07424', '2025-07-30 18:29:46', NULL, NULL);
INSERT INTO `sys_menu` VALUES ('65591853590e4550984ab20f777c87a0', 'd92867e8d4b443c3ba07024cf1f83def', NULL, NULL, NULL, NULL, NULL, NULL, NULL, '3', '编辑', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, '1', NULL, 'sys:menu:edit', 1, 2, '953b77a21366497094303f293cc07424', '2025-07-30 09:32:30', '953b77a21366497094303f293cc07424', '2025-07-30 09:46:54');
INSERT INTO `sys_menu` VALUES ('674e522160944a53ad53652f7cb63f5b', 'fc1c718052ef4490b8dfaf83f2185dc2', '/log/operation', NULL, NULL, 'log/operation/index', NULL, NULL, '', '2', '操作日志', '', 'IconBookmark', 0, 1, 1, NULL, 1, 0, '1', NULL, NULL, 1, 10, '953b77a21366497094303f293cc07424', '2025-09-25 10:30:48', '953b77a21366497094303f293cc07424', '2025-10-30 09:38:47');
INSERT INTO `sys_menu` VALUES ('746c7853a14a4fd2b7ac2886b82e304b', 'c28884e311364d4abd7a13ab3fd89bdb', NULL, NULL, NULL, NULL, NULL, NULL, NULL, '3', '字典项编辑', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, '1', NULL, 'sys:dictItem:edit', 1, 5, '953b77a21366497094303f293cc07424', '2025-07-30 10:30:34', '953b77a21366497094303f293cc07424', '2025-07-30 10:36:18');
INSERT INTO `sys_menu` VALUES ('7a39fb17ab654a1da60b93d73b422789', 'fc1c718052ef4490b8dfaf83f2185dc2', '/system/user', NULL, NULL, 'system/user/index', NULL, NULL, '', '2', '用户管理', '', 'icon-user', 0, 1, 1, NULL, 1, 0, '1', NULL, NULL, 1, 1, '953b77a21366497094303f293cc07424', '2025-07-15 11:45:55', '949eb85f8ca446828b5826f745e2c824', '2025-10-10 17:41:16');
INSERT INTO `sys_menu` VALUES ('7bb71cc033034b9589a571ded86e2dc8', 'fc1c718052ef4490b8dfaf83f2185dc2', '/system/parameter', NULL, NULL, 'system/parameter/index', NULL, NULL, '', '2', '系统参数', '', 'IconEnglishFill', 0, 1, 1, NULL, 1, 0, '1', NULL, NULL, 1, 7, '953b77a21366497094303f293cc07424', '2025-09-25 10:55:33', '953b77a21366497094303f293cc07424', '2025-10-17 11:47:22');
INSERT INTO `sys_menu` VALUES ('808dabc1a96540cf95fcfab1ac53ed98', 'd92867e8d4b443c3ba07024cf1f83def', NULL, NULL, NULL, NULL, NULL, NULL, NULL, '3', '列表状态调整', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, '1', NULL, 'sys:menu:status', 1, 4, '953b77a21366497094303f293cc07424', '2025-07-30 10:24:07', '953b77a21366497094303f293cc07424', '2025-07-30 10:36:03');
INSERT INTO `sys_menu` VALUES ('8cc5876c17e249f0907bdfd7846214ba', '0c36b3a959b343eda48e696e0fb7dfaf', NULL, NULL, NULL, NULL, NULL, NULL, NULL, '3', '删除', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, '1', NULL, 'sys:post:del', 1, 3, '953b77a21366497094303f293cc07424', '2025-07-30 18:30:41', NULL, NULL);
INSERT INTO `sys_menu` VALUES ('8f81d3e08b944a0ea883f53e01598501', '0c36b3a959b343eda48e696e0fb7dfaf', NULL, NULL, NULL, NULL, NULL, NULL, NULL, '3', '编辑', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, '1', NULL, 'sys:post:edit', 1, 2, '953b77a21366497094303f293cc07424', '2025-07-30 18:30:17', '953b77a21366497094303f293cc07424', '2025-07-30 18:30:46');
INSERT INTO `sys_menu` VALUES ('90695bd4b3af406a94c54170515c3b5c', 'fc1c718052ef4490b8dfaf83f2185dc2', '/system/timer', NULL, NULL, 'system/timer/index', NULL, NULL, '', '2', '定时任务', '', 'IconClockCircle', 0, 1, 1, NULL, 1, 0, '1', NULL, NULL, 1, 9, '953b77a21366497094303f293cc07424', '2025-09-29 13:55:07', '953b77a21366497094303f293cc07424', '2025-10-09 18:13:56');
INSERT INTO `sys_menu` VALUES ('99f7a123b0af45c3a16f0430eb9c29c0', 'f6c078279883409ea44989bb80935a8f', NULL, NULL, NULL, NULL, NULL, NULL, NULL, '3', '编辑', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, '1', NULL, 'sys:dept:edit', 1, 2, '953b77a21366497094303f293cc07424', '2025-07-30 17:46:26', NULL, NULL);
INSERT INTO `sys_menu` VALUES ('9b316608d8494955807dddb160c42857', '0', NULL, '/lib', '/bin', NULL, '18', 'non', NULL, '1', '测试APP目录', NULL, 'https://plcsh.oss-cn-shanghai.aliyuncs.com/images/202509/15162231yBQtkX.png', NULL, NULL, NULL, NULL, NULL, NULL, '1', NULL, NULL, 2, 1, '953b77a21366497094303f293cc07424', '2025-09-03 13:59:19', '953b77a21366497094303f293cc07424', '2025-09-18 09:23:15');
INSERT INTO `sys_menu` VALUES ('9c4610cb8cd14f7c90893d877ac9e364', '7a39fb17ab654a1da60b93d73b422789', NULL, NULL, NULL, NULL, NULL, NULL, NULL, '3', '删除', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, '1', NULL, 'sys:user:del', 1, 3, '953b77a21366497094303f293cc07424', '2025-10-13 17:41:08', NULL, NULL);
INSERT INTO `sys_menu` VALUES ('9ecb0ab22ee0479c8a8f228f7896dcd2', 'a52a6f56b2294289a731ff4d8cff42a9', NULL, NULL, NULL, NULL, NULL, NULL, NULL, '3', '编辑', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, '1', NULL, 'sys:role:edit', 1, 2, '953b77a21366497094303f293cc07424', '2025-10-13 17:48:25', NULL, NULL);
INSERT INTO `sys_menu` VALUES ('a3bacdad9dec428e948e472f4b744825', 'a52a6f56b2294289a731ff4d8cff42a9', NULL, NULL, NULL, NULL, NULL, NULL, NULL, '3', '删除', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, '1', NULL, 'sys:role:del', 1, 3, '953b77a21366497094303f293cc07424', '2025-10-13 17:48:38', NULL, NULL);
INSERT INTO `sys_menu` VALUES ('a52a6f56b2294289a731ff4d8cff42a9', 'fc1c718052ef4490b8dfaf83f2185dc2', '/system/role', NULL, NULL, 'system/role/index', NULL, NULL, '', '2', '角色管理', '', 'IconCommon', 0, 1, 1, NULL, 1, 0, '1', NULL, NULL, 1, 2, '953b77a21366497094303f293cc07424', '2025-07-29 14:17:07', NULL, NULL);
INSERT INTO `sys_menu` VALUES ('a794c6385b964cd2a1f274d17a86afed', '7bb71cc033034b9589a571ded86e2dc8', NULL, NULL, NULL, NULL, NULL, NULL, NULL, '3', '删除', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, '1', NULL, 'sys:param:del', 1, 3, '953b77a21366497094303f293cc07424', '2025-09-29 16:42:47', NULL, NULL);
INSERT INTO `sys_menu` VALUES ('a97f98a231794561a74820ed612b7e04', 'c28884e311364d4abd7a13ab3fd89bdb', NULL, NULL, NULL, NULL, NULL, NULL, NULL, '3', '字典项删除', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, '1', NULL, 'sys:dictItem:del', 1, 6, '953b77a21366497094303f293cc07424', '2025-07-30 10:30:58', '953b77a21366497094303f293cc07424', '2025-07-30 10:36:23');
INSERT INTO `sys_menu` VALUES ('b1914145854b42ec99ca8c7561d6eaaa', 'd92867e8d4b443c3ba07024cf1f83def', NULL, NULL, NULL, NULL, NULL, NULL, NULL, '3', '新增', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, '1', NULL, 'sys:menu:add', 1, 1, '953b77a21366497094303f293cc07424', '2025-07-18 15:19:40', '953b77a21366497094303f293cc07424', '2025-07-30 09:46:47');
INSERT INTO `sys_menu` VALUES ('b6ef85e05bbf462788ee3b5de782d686', 'd92867e8d4b443c3ba07024cf1f83def', NULL, NULL, NULL, NULL, NULL, NULL, NULL, '3', '删除', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, '1', NULL, 'sys:menu:del', 1, 3, '953b77a21366497094303f293cc07424', '2025-07-30 09:32:50', '953b77a21366497094303f293cc07424', '2025-07-30 09:47:01');
INSERT INTO `sys_menu` VALUES ('c28884e311364d4abd7a13ab3fd89bdb', 'fc1c718052ef4490b8dfaf83f2185dc2', '/system/dict', NULL, NULL, 'system/dict/index', NULL, NULL, '', '2', '字典管理', '', 'icon-bookmark', 0, 1, 1, NULL, 1, 0, '1', NULL, NULL, 1, 5, '953b77a21366497094303f293cc07424', '2025-07-18 10:53:10', '953b77a21366497094303f293cc07424', '2025-08-06 15:01:11');
INSERT INTO `sys_menu` VALUES ('ccd84ed949bb4a3ba74f0ead6eb4829f', 'f6c078279883409ea44989bb80935a8f', NULL, NULL, NULL, NULL, NULL, NULL, NULL, '3', '新增', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, '1', NULL, 'sys:dept:add', 1, 1, '953b77a21366497094303f293cc07424', '2025-07-30 17:45:58', '953b77a21366497094303f293cc07424', '2025-09-09 14:03:45');
INSERT INTO `sys_menu` VALUES ('d92867e8d4b443c3ba07024cf1f83def', 'fc1c718052ef4490b8dfaf83f2185dc2', '/system/menu', NULL, NULL, 'system/menu/index', NULL, NULL, '', '2', '菜单管理', '', 'icon-menu', 0, 1, 1, NULL, 1, 0, '1', NULL, NULL, 1, 4, '953b77a21366497094303f293cc07424', '2025-07-18 10:52:04', '953b77a21366497094303f293cc07424', '2025-08-06 15:04:02');
INSERT INTO `sys_menu` VALUES ('e3cdc038d53c4278a95d2c05d8486ed1', '9b316608d8494955807dddb160c42857', NULL, '/8888', '1111', NULL, '', '11', NULL, '2', '111', NULL, 'https://plcsh.oss-cn-shanghai.aliyuncs.com/images/202509/15161529JMJdyp.png', NULL, NULL, NULL, NULL, NULL, NULL, '1', NULL, NULL, 2, 1, '953b77a21366497094303f293cc07424', '2025-09-15 16:15:48', NULL, NULL);
INSERT INTO `sys_menu` VALUES ('e7443eea983e4abbbdfcd4d824c29f8d', '7a39fb17ab654a1da60b93d73b422789', NULL, NULL, NULL, NULL, NULL, NULL, NULL, '3', '重置密码', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, '1', NULL, 'sys:user:reset', 1, 4, '953b77a21366497094303f293cc07424', '2025-10-13 17:44:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES ('f15324e4e60a4839bbb4dd1a5c755545', 'd92867e8d4b443c3ba07024cf1f83def', NULL, NULL, NULL, NULL, NULL, NULL, NULL, '3', '列表是否隐藏调整', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, '1', NULL, 'sys:menu:hidden', 1, 6, '953b77a21366497094303f293cc07424', '2025-08-13 16:01:03', '953b77a21366497094303f293cc07424', '2025-08-13 16:01:28');
INSERT INTO `sys_menu` VALUES ('f636bbd19db744bf8f0e1b4511072b54', 'f6c078279883409ea44989bb80935a8f', NULL, NULL, NULL, NULL, NULL, NULL, NULL, '3', '删除', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, '1', NULL, 'sys:dept:del', 1, 3, '953b77a21366497094303f293cc07424', '2025-07-30 17:46:46', NULL, NULL);
INSERT INTO `sys_menu` VALUES ('f6c078279883409ea44989bb80935a8f', 'fc1c718052ef4490b8dfaf83f2185dc2', '/system/dept', NULL, NULL, 'system/dept/index', NULL, NULL, '', '2', '组织管理', '', 'IconMindMapping', 0, 1, 1, NULL, 1, 0, '1', NULL, NULL, 1, 3, '953b77a21366497094303f293cc07424', '2025-07-29 14:18:37', '953b77a21366497094303f293cc07424', '2025-08-06 15:04:07');
INSERT INTO `sys_menu` VALUES ('fba3ca1ccf23460f8454bd1b2d4a6f1a', 'a52a6f56b2294289a731ff4d8cff42a9', NULL, NULL, NULL, NULL, NULL, NULL, NULL, '3', '新增', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, '1', NULL, 'sys:role:add', 1, 1, '953b77a21366497094303f293cc07424', '2025-10-13 17:48:04', NULL, NULL);
INSERT INTO `sys_menu` VALUES ('fc1c718052ef4490b8dfaf83f2185dc2', '0', '/system', NULL, NULL, 'Layout', NULL, NULL, '', '1', '系统管理', 'menu-system', '', 0, 0, 1, 0, NULL, 0, '1', NULL, NULL, 1, 1, '953b77a21366497094303f293cc07424', '2025-07-15 11:43:18', '953b77a21366497094303f293cc07424', '2025-10-15 15:30:14');
INSERT INTO `sys_menu` VALUES ('fe14c59a39724673ae431b8269a9f013', 'c28884e311364d4abd7a13ab3fd89bdb', NULL, NULL, NULL, NULL, NULL, NULL, NULL, '3', '字典删除', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, '1', NULL, 'sys:dict:del', 1, 3, '953b77a21366497094303f293cc07424', '2025-07-30 10:26:53', NULL, NULL);

-- ----------------------------
-- Table structure for sys_oper_log
-- ----------------------------
DROP TABLE IF EXISTS `sys_oper_log`;
CREATE TABLE `sys_oper_log`  (
  `oper_id` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '日志主键',
  `oper_user_id` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '操作用户id',
  `method` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '请求方法',
  `path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '请求路径',
  `ip` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '0' COMMENT '客户端IP',
  `status` int NULL DEFAULT NULL COMMENT '响应状态码',
  `req_body` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '请求体',
  `res_body` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '响应体',
  `latency` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '耗时',
  `oper_time` datetime NOT NULL COMMENT '操作时间',
  PRIMARY KEY (`oper_id`, `oper_time`) USING BTREE,
  UNIQUE INDEX `uniq_oper_id`(`oper_id` ASC, `oper_time` ASC) USING BTREE COMMENT 'id唯一索引',
  INDEX `idx_oper_time`(`oper_time` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '操作日志' ROW_FORMAT = COMPACT PARTITION BY RANGE (to_days(`oper_time`))
PARTITIONS 4
(PARTITION `p20251030` VALUES LESS THAN (739920) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20251031` VALUES LESS THAN (739921) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20251101` VALUES LESS THAN (739922) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p_future` VALUES LESS THAN (MAXVALUE) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 )
;

-- ----------------------------
-- Table structure for sys_param
-- ----------------------------
DROP TABLE IF EXISTS `sys_param`;
CREATE TABLE `sys_param`  (
  `param_id` char(32) CHARACTER SET latin1 COLLATE latin1_general_ci NOT NULL DEFAULT '' COMMENT '主键',
  `sys_logo` varchar(255) CHARACTER SET latin1 COLLATE latin1_general_ci NOT NULL DEFAULT '' COMMENT 'logo(上传)',
  `sys_title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '平台名称',
  `large_screen_title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '大屏标题',
  `map_display_area` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '地图默认显示区域',
  `annual_inspection` int NOT NULL COMMENT '年检提醒(默认30天)',
  `contract` int NOT NULL COMMENT '合同到期提醒(默认30天)',
  `insurance` int NOT NULL COMMENT '保险到期预警(默认30天)',
  `dept_id` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '所属组织(all)',
  `created_user_id` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '创建用户',
  `created_time` datetime NOT NULL COMMENT '创建时间',
  `updated_user_id` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '更新用户',
  `updated_time` datetime NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`param_id`) USING BTREE,
  UNIQUE INDEX `only_param_id`(`param_id` ASC) USING BTREE,
  UNIQUE INDEX `only_dept_id`(`dept_id` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '参数管理表' ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of sys_param
-- ----------------------------
INSERT INTO `sys_param` VALUES ('b306c13201c545c68c84854cc08586ed', 'http://localhost:9091/api/images/CWRS_LOGO.png', 'CWRS', 'CWRS', '江苏徐州', 30, 30, 30, 'a2aa4cf0ba1f46049cbe5a5d966446e8', '953b77a21366497094303f293cc07424', '2025-10-30 09:37:45', NULL, NULL);
INSERT INTO `sys_param` VALUES ('ff6aff0f7e914a2ca006556baa8f345c', 'http://localhost:9091/api/images/CWRS_LOGO.png', 'CWRS管理系统', 'CWRS管理系统', '江苏徐州', 30, 30, 30, 'all', '953b77a21366497094303f293cc07424', '2025-09-18 08:49:49', '953b77a21366497094303f293cc07424', '2025-10-29 18:23:30');

-- ----------------------------
-- Table structure for sys_post
-- ----------------------------
DROP TABLE IF EXISTS `sys_post`;
CREATE TABLE `sys_post`  (
  `post_id` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '主键',
  `post_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '岗位名称',
  `post_code` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '岗位编码',
  `post_status` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '状态(字典)',
  `post_sort` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '排序',
  `dept_id` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '所属组织',
  `desc` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '描述',
  `created_user_id` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '创建用户;所有表中都要有该字段，数据权限需要使用',
  `created_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_user_id` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '更新用户',
  `updated_time` datetime NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`post_id`) USING BTREE,
  UNIQUE INDEX `only_post_id`(`post_id` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '岗位' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_post
-- ----------------------------
INSERT INTO `sys_post` VALUES ('4f8fa7d0726c4080a15a4cdd4a51f72f', '总经理', 'CEO', '1', '2', 'a2aa4cf0ba1f46049cbe5a5d966446e8', '江苏总部总经理', '953b77a21366497094303f293cc07424', '2025-08-04 11:40:55', '953b77a21366497094303f293cc07424', '2025-09-09 09:15:44');
INSERT INTO `sys_post` VALUES ('a90cee574f364bf891d9c2a21643eeba', '董事长', 'Chairman', '1', '1', '12cb92fc10724b4ca762bb38bb624fd2', '', '953b77a21366497094303f293cc07424', '2025-08-06 11:30:51', '953b77a21366497094303f293cc07424', '2025-10-21 08:53:15');

-- ----------------------------
-- Table structure for sys_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_role`;
CREATE TABLE `sys_role`  (
  `role_id` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '主键',
  `role_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '角色名称',
  `role_code` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '角色编码',
  `role_status` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '状态(字典)',
  `role_sort` int NOT NULL COMMENT '排序',
  `is_builtin` int NOT NULL DEFAULT 0 COMMENT '是否内置1是 0否',
  `desc` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '描述',
  `dept_id` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '所属组织',
  `parent_id` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '父级id',
  `role_level` varchar(768) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '角色层级',
  `created_user_id` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '创建用户',
  `created_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_user_id` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '更新人',
  `updated_time` datetime NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`role_id`) USING BTREE,
  UNIQUE INDEX `only_role_code`(`role_code` ASC) USING BTREE,
  UNIQUE INDEX `only_role_id`(`role_id` ASC) USING BTREE,
  UNIQUE INDEX `only_dept_role`(`role_name` ASC, `dept_id` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '角色' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_role
-- ----------------------------
INSERT INTO `sys_role` VALUES ('5c9e854262ce42918a1b9d12500ca9ff', '超级管理员', 'sys_admin', '1', 1, 1, '系统内置', 'all', '', '', '953b77a21366497094303f293cc07424', '2025-07-15 11:49:42', '953b77a21366497094303f293cc07424', '2025-09-29 16:54:10');
INSERT INTO `sys_role` VALUES ('6afbfef6f5ba4d0eb3a2efc505a519c8', '系统用户', 'sys_user', '1', 2, 0, NULL, 'a2aa4cf0ba1f46049cbe5a5d966446e8', '', '', '953b77a21366497094303f293cc07424', '2025-10-30 09:29:09', NULL, NULL);

-- ----------------------------
-- Table structure for sys_role_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_menu`;
CREATE TABLE `sys_role_menu`  (
  `role_menu_id` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '主键',
  `dept_id` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '组织id',
  `role_id` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '角色id',
  `menu_id` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '菜单id',
  `menu_type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '菜单类型',
  `data_purview` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '按菜单-数据权限(字典)',
  `created_user_id` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '创建用户',
  `created_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_user_id` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '更新用户',
  `updated_time` datetime NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`role_menu_id`) USING BTREE,
  UNIQUE INDEX `only_role_menu_id`(`role_menu_id` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '角色菜单关联关系' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_role_menu
-- ----------------------------
INSERT INTO `sys_role_menu` VALUES ('055aca2030a24ca0aad6a7179f177bff', '12cb92fc10724b4ca762bb38bb624fd2', 'e34b22d6f29b4334acaac995988f700f', 'ccd84ed949bb4a3ba74f0ead6eb4829f', '3', '', '953b77a21366497094303f293cc07424', '2025-10-23 09:11:39', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('06b92d5bbf9a484da4de02d2a67fe588', 'a2aa4cf0ba1f46049cbe5a5d966446e8', '65ecc19d88e64e39b718c5f96f326f6f', 'a52a6f56b2294289a731ff4d8cff42a9', '2', 'this_dept', '953b77a21366497094303f293cc07424', '2025-10-15 15:29:47', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('070056a902024d3c9ea6873413bbbc0f', 'a2aa4cf0ba1f46049cbe5a5d966446e8', '65ecc19d88e64e39b718c5f96f326f6f', '8cc5876c17e249f0907bdfd7846214ba', '3', '', '953b77a21366497094303f293cc07424', '2025-10-15 15:29:47', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('0701eae120104b71a5ba27fd8a080750', '12cb92fc10724b4ca762bb38bb624fd2', 'e34b22d6f29b4334acaac995988f700f', '80a528c9ffe24cdfac795813f61a0c89', '2', 'this_dept_and_below', '953b77a21366497094303f293cc07424', '2025-10-23 09:11:39', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('070ecd3e55434d4f8023674ebeffa7f9', 'a2aa4cf0ba1f46049cbe5a5d966446e8', '805b5bac807043c9aedde416c0be9343', 'fba3ca1ccf23460f8454bd1b2d4a6f1a', '3', '', '953b77a21366497094303f293cc07424', '2025-10-23 09:11:48', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('0886aed8f01e433d8638665503048bef', 'a2aa4cf0ba1f46049cbe5a5d966446e8', '65ecc19d88e64e39b718c5f96f326f6f', 'a794c6385b964cd2a1f274d17a86afed', '3', '', '953b77a21366497094303f293cc07424', '2025-10-15 15:29:47', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('0b851914916b4ca8b934d937eac3fda6', '12cb92fc10724b4ca762bb38bb624fd2', 'e34b22d6f29b4334acaac995988f700f', 'f636bbd19db744bf8f0e1b4511072b54', '3', '', '953b77a21366497094303f293cc07424', '2025-10-23 09:11:39', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('0c34b965801e4bb2b41b36ef848425c7', '12cb92fc10724b4ca762bb38bb624fd2', 'e34b22d6f29b4334acaac995988f700f', '30d7c3bfc9d646309c52766b5041422e', '3', '', '953b77a21366497094303f293cc07424', '2025-10-23 09:11:39', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('0c754e0ffebf4ab481d46ce1d6d6a545', 'a2aa4cf0ba1f46049cbe5a5d966446e8', '805b5bac807043c9aedde416c0be9343', 'e7443eea983e4abbbdfcd4d824c29f8d', '3', '', '953b77a21366497094303f293cc07424', '2025-10-23 09:11:48', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('11e7e52d9c91432ba3feadd9a39c80cd', 'a2aa4cf0ba1f46049cbe5a5d966446e8', '805b5bac807043c9aedde416c0be9343', '230fe955336445189f93986a9807a17e', '3', '', '953b77a21366497094303f293cc07424', '2025-10-23 09:11:48', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('14a906d5c0614100bae6b0543a3ba1da', '12cb92fc10724b4ca762bb38bb624fd2', 'e34b22d6f29b4334acaac995988f700f', '0b087f935c0c44139b0cc18fe1862dda', '3', '', '953b77a21366497094303f293cc07424', '2025-10-23 09:11:39', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('16154b326d2d467e9a1b2c08db73a4b2', '12cb92fc10724b4ca762bb38bb624fd2', 'e34b22d6f29b4334acaac995988f700f', '9c4610cb8cd14f7c90893d877ac9e364', '3', '', '953b77a21366497094303f293cc07424', '2025-10-23 09:11:39', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('166563f45c134f2c9787a195e290d3f0', 'a2aa4cf0ba1f46049cbe5a5d966446e8', '805b5bac807043c9aedde416c0be9343', '6a184a4fc46d4aaf864644689cfb7c51', '3', '', '953b77a21366497094303f293cc07424', '2025-10-23 09:11:48', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('16d7b885140b4ea7a4c71c174c9f7c48', 'a2aa4cf0ba1f46049cbe5a5d966446e8', '65ecc19d88e64e39b718c5f96f326f6f', '328b482b66b24db2864f0e7e6e0469c8', '3', '', '953b77a21366497094303f293cc07424', '2025-10-15 15:29:47', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('1944e80da668405283413355970880f4', 'a2aa4cf0ba1f46049cbe5a5d966446e8', '805b5bac807043c9aedde416c0be9343', 'f636bbd19db744bf8f0e1b4511072b54', '3', '', '953b77a21366497094303f293cc07424', '2025-10-23 09:11:48', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('1a411c1a73584385b327feb2248eb763', 'a2aa4cf0ba1f46049cbe5a5d966446e8', '65ecc19d88e64e39b718c5f96f326f6f', '90695bd4b3af406a94c54170515c3b5c', '2', 'this_dept', '953b77a21366497094303f293cc07424', '2025-10-15 15:29:47', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('2076d86f4d3347c18d284c3e8a4d900c', '12cb92fc10724b4ca762bb38bb624fd2', 'e34b22d6f29b4334acaac995988f700f', '99f7a123b0af45c3a16f0430eb9c29c0', '3', '', '953b77a21366497094303f293cc07424', '2025-10-23 09:11:39', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('2094a00a101f413f9d53878e76953933', '12cb92fc10724b4ca762bb38bb624fd2', 'e34b22d6f29b4334acaac995988f700f', 'a3bacdad9dec428e948e472f4b744825', '3', '', '953b77a21366497094303f293cc07424', '2025-10-23 09:11:39', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('2115488ea7a34a37b7f338c5dd3306f7', 'a2aa4cf0ba1f46049cbe5a5d966446e8', '65ecc19d88e64e39b718c5f96f326f6f', 'fe14c59a39724673ae431b8269a9f013', '3', '', '953b77a21366497094303f293cc07424', '2025-10-15 15:29:47', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('212a3528d40c4421a416e9fb0f279a0b', '12cb92fc10724b4ca762bb38bb624fd2', 'e34b22d6f29b4334acaac995988f700f', '8cc5876c17e249f0907bdfd7846214ba', '3', '', '953b77a21366497094303f293cc07424', '2025-10-23 09:11:39', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('2301f030368b44fab48e10df6495f5ba', 'a2aa4cf0ba1f46049cbe5a5d966446e8', '6afbfef6f5ba4d0eb3a2efc505a519c8', 'fba3ca1ccf23460f8454bd1b2d4a6f1a', '3', '', '953b77a21366497094303f293cc07424', '2025-10-30 09:29:09', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('24caa086a2244f49ac8319b782c33fe5', 'a2aa4cf0ba1f46049cbe5a5d966446e8', '65ecc19d88e64e39b718c5f96f326f6f', '7bb71cc033034b9589a571ded86e2dc8', '2', 'this_dept', '953b77a21366497094303f293cc07424', '2025-10-15 15:29:47', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('260cbb2cb380400287fe8c6dc3c7901d', 'a2aa4cf0ba1f46049cbe5a5d966446e8', '65ecc19d88e64e39b718c5f96f326f6f', 'b1914145854b42ec99ca8c7561d6eaaa', '3', '', '953b77a21366497094303f293cc07424', '2025-10-15 15:29:47', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('2757e2f4de934d3c908d826b96798e37', '12cb92fc10724b4ca762bb38bb624fd2', 'e34b22d6f29b4334acaac995988f700f', 'b1914145854b42ec99ca8c7561d6eaaa', '3', '', '953b77a21366497094303f293cc07424', '2025-10-23 09:11:39', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('278afbd18da746909644234aa9a26375', 'a2aa4cf0ba1f46049cbe5a5d966446e8', '65ecc19d88e64e39b718c5f96f326f6f', '8f81d3e08b944a0ea883f53e01598501', '3', '', '953b77a21366497094303f293cc07424', '2025-10-15 15:29:47', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('297c5ec6b3734e7fbbf564e13e663bc4', '12cb92fc10724b4ca762bb38bb624fd2', 'e34b22d6f29b4334acaac995988f700f', 'a52a6f56b2294289a731ff4d8cff42a9', '2', 'this_dept_and_below', '953b77a21366497094303f293cc07424', '2025-10-23 09:11:39', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('2a01e95a925f4abbb1307ba35a242c49', 'a2aa4cf0ba1f46049cbe5a5d966446e8', '65ecc19d88e64e39b718c5f96f326f6f', '746c7853a14a4fd2b7ac2886b82e304b', '3', '', '953b77a21366497094303f293cc07424', '2025-10-15 15:29:47', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('2c4235128a5245058d01a01084979741', 'a2aa4cf0ba1f46049cbe5a5d966446e8', '65ecc19d88e64e39b718c5f96f326f6f', 'a3bacdad9dec428e948e472f4b744825', '3', '', '953b77a21366497094303f293cc07424', '2025-10-15 15:29:47', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('2d527c92868f49d98ad207cd3e57a814', 'a2aa4cf0ba1f46049cbe5a5d966446e8', '805b5bac807043c9aedde416c0be9343', '04772007211d4b889e060ef87eb1bf65', '1', '', '953b77a21366497094303f293cc07424', '2025-10-23 09:11:48', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('2f0881a31d964e348905b94ed1e6096a', 'a2aa4cf0ba1f46049cbe5a5d966446e8', '6afbfef6f5ba4d0eb3a2efc505a519c8', '9c4610cb8cd14f7c90893d877ac9e364', '3', '', '953b77a21366497094303f293cc07424', '2025-10-30 09:29:09', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('2fa0336c4567419a9f12ec7e97c02d55', '12cb92fc10724b4ca762bb38bb624fd2', 'e34b22d6f29b4334acaac995988f700f', 'b2b60dce2ae74c708e88428dcb1add28', '3', '', '953b77a21366497094303f293cc07424', '2025-10-23 09:11:39', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('3150bb160d6c46c68d9e5f791042d7b5', 'a2aa4cf0ba1f46049cbe5a5d966446e8', '805b5bac807043c9aedde416c0be9343', '0562128730274a47909e0ddd9959d116', '3', '', '953b77a21366497094303f293cc07424', '2025-10-23 09:11:48', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('3456eeea63a4435ba32f92d4aa628ec2', 'a2aa4cf0ba1f46049cbe5a5d966446e8', '805b5bac807043c9aedde416c0be9343', '7d9d00ccb13e4249baa22332f83ef8fc', '2', '', '953b77a21366497094303f293cc07424', '2025-10-23 09:11:48', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('356a8be6094c44a6b7f90652fae45ba9', 'a2aa4cf0ba1f46049cbe5a5d966446e8', '6afbfef6f5ba4d0eb3a2efc505a519c8', 'e7443eea983e4abbbdfcd4d824c29f8d', '3', '', '953b77a21366497094303f293cc07424', '2025-10-30 09:29:09', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('38cbd5192f384a949b4aeea3b6560c2b', '12cb92fc10724b4ca762bb38bb624fd2', 'e34b22d6f29b4334acaac995988f700f', 'e7443eea983e4abbbdfcd4d824c29f8d', '3', '', '953b77a21366497094303f293cc07424', '2025-10-23 09:11:39', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('38e5b11493b144848e6f86a6f47372f8', '12cb92fc10724b4ca762bb38bb624fd2', 'e34b22d6f29b4334acaac995988f700f', '19ababda3a604d308677eed597a11add', '3', '', '953b77a21366497094303f293cc07424', '2025-10-23 09:11:39', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('3c1a410cfcba447c99af95b81ec274fa', '12cb92fc10724b4ca762bb38bb624fd2', 'e34b22d6f29b4334acaac995988f700f', '6475177fa06145fbb0b2d10af62f4baf', '3', '', '953b77a21366497094303f293cc07424', '2025-10-23 09:11:39', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('3ea1be2884ee4682a430bf64f85f3b94', 'a2aa4cf0ba1f46049cbe5a5d966446e8', '65ecc19d88e64e39b718c5f96f326f6f', '0c2ec443c2fa43eba2521ca4bdd5ffaa', '3', '', '953b77a21366497094303f293cc07424', '2025-10-15 15:29:47', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('4198dbf182114c91ac3d0b82b4d7c9b6', 'a2aa4cf0ba1f46049cbe5a5d966446e8', '805b5bac807043c9aedde416c0be9343', 'ccd84ed949bb4a3ba74f0ead6eb4829f', '3', '', '953b77a21366497094303f293cc07424', '2025-10-23 09:11:48', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('4797d80712ff48f384a542cc5fdd9097', '12cb92fc10724b4ca762bb38bb624fd2', 'e34b22d6f29b4334acaac995988f700f', 'cd2c5bb506954b89a1bb48c35d6f747f', '3', '', '953b77a21366497094303f293cc07424', '2025-10-23 09:11:39', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('482c379e260d4e1293eff22fd1040527', 'a2aa4cf0ba1f46049cbe5a5d966446e8', '805b5bac807043c9aedde416c0be9343', '91a86b43c4294973941343627c95dab4', '3', '', '953b77a21366497094303f293cc07424', '2025-10-23 09:11:48', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('4832bc03f4604114925095f4c9ab6bd3', 'a2aa4cf0ba1f46049cbe5a5d966446e8', '805b5bac807043c9aedde416c0be9343', 'a52a6f56b2294289a731ff4d8cff42a9', '2', 'this_dept', '953b77a21366497094303f293cc07424', '2025-10-23 09:11:48', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('49fab817235f4243a4a1ecaca560266a', '12cb92fc10724b4ca762bb38bb624fd2', 'e34b22d6f29b4334acaac995988f700f', '9b316608d8494955807dddb160c42857', '1', '', '953b77a21366497094303f293cc07424', '2025-10-23 09:11:39', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('4f10399f291144bc9b695bd49874957f', '12cb92fc10724b4ca762bb38bb624fd2', 'e34b22d6f29b4334acaac995988f700f', 'd29c6af4cd864a14ae1afa1ddfdd42f6', '3', '', '953b77a21366497094303f293cc07424', '2025-10-23 09:11:39', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('4f1197e9425547f2ba43d027261b49cb', 'a2aa4cf0ba1f46049cbe5a5d966446e8', '805b5bac807043c9aedde416c0be9343', '575dec0ef6964c11a429c13c03b8b964', '3', '', '953b77a21366497094303f293cc07424', '2025-10-23 09:11:48', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('50a931181ad44c51a778207b2b3681c9', 'a2aa4cf0ba1f46049cbe5a5d966446e8', '6afbfef6f5ba4d0eb3a2efc505a519c8', 'a3bacdad9dec428e948e472f4b744825', '3', '', '953b77a21366497094303f293cc07424', '2025-10-30 09:29:09', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('53e723ad27eb45d08562f7ce5425ee4b', '12cb92fc10724b4ca762bb38bb624fd2', 'e34b22d6f29b4334acaac995988f700f', 'e3cdc038d53c4278a95d2c05d8486ed1', '2', '', '953b77a21366497094303f293cc07424', '2025-10-23 09:11:39', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('5570b98c11a644fe84388f4316405cd9', '12cb92fc10724b4ca762bb38bb624fd2', 'e34b22d6f29b4334acaac995988f700f', '7f5ef7885be54be782a61461000eb130', '3', '', '953b77a21366497094303f293cc07424', '2025-10-23 09:11:39', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('5646ecca01ca4aeaac1c6137c52062d5', 'a2aa4cf0ba1f46049cbe5a5d966446e8', '65ecc19d88e64e39b718c5f96f326f6f', 'b6ef85e05bbf462788ee3b5de782d686', '3', '', '953b77a21366497094303f293cc07424', '2025-10-15 15:29:47', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('5bc9cd1ad32d454ab62a63c6e8ae3e16', 'a2aa4cf0ba1f46049cbe5a5d966446e8', '65ecc19d88e64e39b718c5f96f326f6f', '491ea08df95f47488680bce1d68f433a', '3', '', '953b77a21366497094303f293cc07424', '2025-10-15 15:29:47', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('5ca1077f27954ff1896a393e1576f45d', 'a2aa4cf0ba1f46049cbe5a5d966446e8', '805b5bac807043c9aedde416c0be9343', '9ecb0ab22ee0479c8a8f228f7896dcd2', '3', '', '953b77a21366497094303f293cc07424', '2025-10-23 09:11:48', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('5e877efa1af349c2ab7b4e74879d83c7', 'a2aa4cf0ba1f46049cbe5a5d966446e8', '6afbfef6f5ba4d0eb3a2efc505a519c8', 'f636bbd19db744bf8f0e1b4511072b54', '3', '', '953b77a21366497094303f293cc07424', '2025-10-30 09:29:09', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('5fa3404761d340e7b535cc3f620d3540', 'a2aa4cf0ba1f46049cbe5a5d966446e8', '65ecc19d88e64e39b718c5f96f326f6f', 'e7443eea983e4abbbdfcd4d824c29f8d', '3', '', '953b77a21366497094303f293cc07424', '2025-10-15 15:29:47', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('60e8fd9d0d5c485da108f0931c8d1527', '12cb92fc10724b4ca762bb38bb624fd2', 'e34b22d6f29b4334acaac995988f700f', 'a794c6385b964cd2a1f274d17a86afed', '3', '', '953b77a21366497094303f293cc07424', '2025-10-23 09:11:39', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('62cca7fe3a7c4e4e9035edf2637ed003', 'a2aa4cf0ba1f46049cbe5a5d966446e8', '805b5bac807043c9aedde416c0be9343', 'f6c078279883409ea44989bb80935a8f', '2', 'this_dept', '953b77a21366497094303f293cc07424', '2025-10-23 09:11:48', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('6482abe1e4d6447caa7aed324f4090c0', '12cb92fc10724b4ca762bb38bb624fd2', 'e34b22d6f29b4334acaac995988f700f', '8f81d3e08b944a0ea883f53e01598501', '3', '', '953b77a21366497094303f293cc07424', '2025-10-23 09:11:39', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('667113aae29e401e8ebd2d6c832ddfc9', '12cb92fc10724b4ca762bb38bb624fd2', 'e34b22d6f29b4334acaac995988f700f', '7d9d00ccb13e4249baa22332f83ef8fc', '2', 'this_dept_and_below', '953b77a21366497094303f293cc07424', '2025-10-23 09:11:39', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('674fbe66556748debd7aa07325842e5c', 'a2aa4cf0ba1f46049cbe5a5d966446e8', '6afbfef6f5ba4d0eb3a2efc505a519c8', 'f6c078279883409ea44989bb80935a8f', '2', 'this_dept_and_below', '953b77a21366497094303f293cc07424', '2025-10-30 09:29:09', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('694fb636d44a46a899a71f824a494917', '12cb92fc10724b4ca762bb38bb624fd2', 'e34b22d6f29b4334acaac995988f700f', '24303c4b66a34032b7ef943ea903ec42', '2', '', '953b77a21366497094303f293cc07424', '2025-10-23 09:11:39', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('696ee7e2df86441ea9f3c2a0603ef208', 'a2aa4cf0ba1f46049cbe5a5d966446e8', '65ecc19d88e64e39b718c5f96f326f6f', '99f7a123b0af45c3a16f0430eb9c29c0', '3', '', '953b77a21366497094303f293cc07424', '2025-10-15 15:29:47', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('6ac4c8312d5f42dfb2857b1dda7fd60a', 'a2aa4cf0ba1f46049cbe5a5d966446e8', '65ecc19d88e64e39b718c5f96f326f6f', '19ababda3a604d308677eed597a11add', '3', '', '953b77a21366497094303f293cc07424', '2025-10-15 15:29:47', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('6b0f7b73ea9840329066ed2bfdd41f1a', '12cb92fc10724b4ca762bb38bb624fd2', 'e34b22d6f29b4334acaac995988f700f', '9e1c556a8f014a41a65595d65a93c1d1', '2', '', '953b77a21366497094303f293cc07424', '2025-10-23 09:11:39', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('6bc00e8c5e1e423aa1c6b84f65802713', 'a2aa4cf0ba1f46049cbe5a5d966446e8', '65ecc19d88e64e39b718c5f96f326f6f', '80a528c9ffe24cdfac795813f61a0c89', '2', 'this_dept', '953b77a21366497094303f293cc07424', '2025-10-15 15:29:47', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('6c1b97d691ef4fcd9013f0c80ae4cdbf', '12cb92fc10724b4ca762bb38bb624fd2', 'e34b22d6f29b4334acaac995988f700f', '1ea410dae6e34503b539123eacf7e33b', '3', '', '953b77a21366497094303f293cc07424', '2025-10-23 09:11:39', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('6e4bbc22cac2468895e6a894cf71c559', 'a2aa4cf0ba1f46049cbe5a5d966446e8', '6afbfef6f5ba4d0eb3a2efc505a519c8', '9ecb0ab22ee0479c8a8f228f7896dcd2', '3', '', '953b77a21366497094303f293cc07424', '2025-10-30 09:29:09', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('6e65a70a48b54e07ac4f25a2da71e9d0', 'a2aa4cf0ba1f46049cbe5a5d966446e8', '65ecc19d88e64e39b718c5f96f326f6f', 'ccd84ed949bb4a3ba74f0ead6eb4829f', '3', '', '953b77a21366497094303f293cc07424', '2025-10-15 15:29:47', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('6e70b38666904aa08fb33ae60004ba9e', 'a2aa4cf0ba1f46049cbe5a5d966446e8', '65ecc19d88e64e39b718c5f96f326f6f', '65591853590e4550984ab20f777c87a0', '3', '', '953b77a21366497094303f293cc07424', '2025-10-15 15:29:47', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('6fd09d3cacf848da813496e89fb28682', '12cb92fc10724b4ca762bb38bb624fd2', 'e34b22d6f29b4334acaac995988f700f', 'fba3ca1ccf23460f8454bd1b2d4a6f1a', '3', '', '953b77a21366497094303f293cc07424', '2025-10-23 09:11:39', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('7180ab6eeb2a47e18effbdebe91b46ca', 'a2aa4cf0ba1f46049cbe5a5d966446e8', '65ecc19d88e64e39b718c5f96f326f6f', '0c36b3a959b343eda48e696e0fb7dfaf', '2', 'this_dept', '953b77a21366497094303f293cc07424', '2025-10-15 15:29:47', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('72c3369935514fc890722254e47e2f92', 'a2aa4cf0ba1f46049cbe5a5d966446e8', '6afbfef6f5ba4d0eb3a2efc505a519c8', '6475177fa06145fbb0b2d10af62f4baf', '3', '', '953b77a21366497094303f293cc07424', '2025-10-30 09:29:09', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('72fc289063c94f07994a385520728628', 'a2aa4cf0ba1f46049cbe5a5d966446e8', '65ecc19d88e64e39b718c5f96f326f6f', '1bab070bc46648639ed916bd1ce5e807', '3', '', '953b77a21366497094303f293cc07424', '2025-10-15 15:29:47', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('74506a6c5de84ec78495681ccc552a60', 'a2aa4cf0ba1f46049cbe5a5d966446e8', '6afbfef6f5ba4d0eb3a2efc505a519c8', '477f429a7f4e486c9204fd5f8f058a61', '3', '', '953b77a21366497094303f293cc07424', '2025-10-30 09:29:09', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('7696eefa67984c24bbb447d8be5e18ab', '12cb92fc10724b4ca762bb38bb624fd2', 'e34b22d6f29b4334acaac995988f700f', 'd92867e8d4b443c3ba07024cf1f83def', '2', 'this_dept_and_below', '953b77a21366497094303f293cc07424', '2025-10-23 09:11:39', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('7a229ff29e864ea890326ad9242b409e', '12cb92fc10724b4ca762bb38bb624fd2', 'e34b22d6f29b4334acaac995988f700f', '2304e56278124d2c852fed8ba5219ce5', '2', 'this_dept_and_below', '953b77a21366497094303f293cc07424', '2025-10-23 09:11:39', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('7a6b1a0808d44a10be2ef90b84d45e93', '12cb92fc10724b4ca762bb38bb624fd2', 'e34b22d6f29b4334acaac995988f700f', '997ceabcb58c487a9d61ac2672390529', '3', '', '953b77a21366497094303f293cc07424', '2025-10-23 09:11:39', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('7b53342ecbfe411a881847a844649c3b', 'a2aa4cf0ba1f46049cbe5a5d966446e8', '6afbfef6f5ba4d0eb3a2efc505a519c8', '7a39fb17ab654a1da60b93d73b422789', '2', 'this_dept_and_below', '953b77a21366497094303f293cc07424', '2025-10-30 09:29:09', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('7fdbad3c323243d8a94ce51a1384821f', 'a2aa4cf0ba1f46049cbe5a5d966446e8', '805b5bac807043c9aedde416c0be9343', 'd29c6af4cd864a14ae1afa1ddfdd42f6', '3', '', '953b77a21366497094303f293cc07424', '2025-10-23 09:11:48', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('81508baf173545ffaf42dda615dd884b', 'a2aa4cf0ba1f46049cbe5a5d966446e8', '805b5bac807043c9aedde416c0be9343', '99f7a123b0af45c3a16f0430eb9c29c0', '3', '', '953b77a21366497094303f293cc07424', '2025-10-23 09:11:48', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('832caf48397f48218559b5ea3afdf9cc', '12cb92fc10724b4ca762bb38bb624fd2', 'e34b22d6f29b4334acaac995988f700f', '04772007211d4b889e060ef87eb1bf65', '1', 'this_dept_and_below', '953b77a21366497094303f293cc07424', '2025-10-23 09:11:39', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('834584a16ad74f0bafa78c94fa23097e', 'a2aa4cf0ba1f46049cbe5a5d966446e8', '65ecc19d88e64e39b718c5f96f326f6f', '0b087f935c0c44139b0cc18fe1862dda', '3', '', '953b77a21366497094303f293cc07424', '2025-10-15 15:29:47', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('846d23d4ae004defa8424f2fd036b580', 'a2aa4cf0ba1f46049cbe5a5d966446e8', '6afbfef6f5ba4d0eb3a2efc505a519c8', '0c2ec443c2fa43eba2521ca4bdd5ffaa', '3', '', '953b77a21366497094303f293cc07424', '2025-10-30 09:29:09', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('84b839289f414668820db1d2a49da4b9', '12cb92fc10724b4ca762bb38bb624fd2', 'e34b22d6f29b4334acaac995988f700f', '0c36b3a959b343eda48e696e0fb7dfaf', '2', 'this_dept_and_below', '953b77a21366497094303f293cc07424', '2025-10-23 09:11:39', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('84eaf6eefee34532b122de73136afd0b', 'a2aa4cf0ba1f46049cbe5a5d966446e8', '805b5bac807043c9aedde416c0be9343', '7c2fbb95bba34637a8f6d07ba48ca546', '3', '', '953b77a21366497094303f293cc07424', '2025-10-23 09:11:48', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('856883e05b6e4adcb61d11e344ff5ff0', '12cb92fc10724b4ca762bb38bb624fd2', 'e34b22d6f29b4334acaac995988f700f', '9ecb0ab22ee0479c8a8f228f7896dcd2', '3', '', '953b77a21366497094303f293cc07424', '2025-10-23 09:11:39', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('8b900759d45e489e826252299c4348a1', 'a2aa4cf0ba1f46049cbe5a5d966446e8', '6afbfef6f5ba4d0eb3a2efc505a519c8', '0c36b3a959b343eda48e696e0fb7dfaf', '2', 'this_dept_and_below', '953b77a21366497094303f293cc07424', '2025-10-30 09:29:09', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('8cc3e4a2d1934cc09a0cd7de0a2088b1', '12cb92fc10724b4ca762bb38bb624fd2', 'e34b22d6f29b4334acaac995988f700f', '674e522160944a53ad53652f7cb63f5b', '2', 'this_dept_and_below', '953b77a21366497094303f293cc07424', '2025-10-23 09:11:39', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('8ea0264f188442ffa797d0de08777a6f', '12cb92fc10724b4ca762bb38bb624fd2', 'e34b22d6f29b4334acaac995988f700f', '7bb71cc033034b9589a571ded86e2dc8', '2', 'this_dept_and_below', '953b77a21366497094303f293cc07424', '2025-10-23 09:11:39', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('8f8556d705644018b362b1004597a17d', '12cb92fc10724b4ca762bb38bb624fd2', 'e34b22d6f29b4334acaac995988f700f', '575dec0ef6964c11a429c13c03b8b964', '3', '', '953b77a21366497094303f293cc07424', '2025-10-23 09:11:39', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('8f8a40d84a394c0781d8cebfa535c6c5', 'a2aa4cf0ba1f46049cbe5a5d966446e8', '65ecc19d88e64e39b718c5f96f326f6f', 'd92867e8d4b443c3ba07024cf1f83def', '2', 'this_dept', '953b77a21366497094303f293cc07424', '2025-10-15 15:29:47', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('902efa61f8b44f8d8329e9ef2b03252f', 'a2aa4cf0ba1f46049cbe5a5d966446e8', '805b5bac807043c9aedde416c0be9343', '674e522160944a53ad53652f7cb63f5b', '2', '', '953b77a21366497094303f293cc07424', '2025-10-23 09:11:48', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('94a80e867b044617b4e53cea5cc923fe', 'a2aa4cf0ba1f46049cbe5a5d966446e8', '65ecc19d88e64e39b718c5f96f326f6f', '808dabc1a96540cf95fcfab1ac53ed98', '3', '', '953b77a21366497094303f293cc07424', '2025-10-15 15:29:47', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('962053da50f043e2ac98bd615d77b096', '12cb92fc10724b4ca762bb38bb624fd2', 'e34b22d6f29b4334acaac995988f700f', 'f6c078279883409ea44989bb80935a8f', '2', 'this_dept_and_below', '953b77a21366497094303f293cc07424', '2025-10-23 09:11:39', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('97b9ac4936554dc8bd775aed2f88154c', 'a2aa4cf0ba1f46049cbe5a5d966446e8', '65ecc19d88e64e39b718c5f96f326f6f', '6475177fa06145fbb0b2d10af62f4baf', '3', '', '953b77a21366497094303f293cc07424', '2025-10-15 15:29:47', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('99bd0f9855d6474db9e97408bcdd04d3', 'a2aa4cf0ba1f46049cbe5a5d966446e8', '805b5bac807043c9aedde416c0be9343', 'b2b60dce2ae74c708e88428dcb1add28', '3', '', '953b77a21366497094303f293cc07424', '2025-10-23 09:11:48', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('9cc2e95bdeca4343915bdfee4ad5733e', '12cb92fc10724b4ca762bb38bb624fd2', 'e34b22d6f29b4334acaac995988f700f', '6a184a4fc46d4aaf864644689cfb7c51', '3', '', '953b77a21366497094303f293cc07424', '2025-10-23 09:11:39', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('9d5c43215c4c4770944552a8645631b6', '12cb92fc10724b4ca762bb38bb624fd2', 'e34b22d6f29b4334acaac995988f700f', '3487c80fded34204849859598f014c8a', '3', '', '953b77a21366497094303f293cc07424', '2025-10-23 09:11:39', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('9feaf48668de48fbb3bba7cb6da12ca4', '12cb92fc10724b4ca762bb38bb624fd2', 'e34b22d6f29b4334acaac995988f700f', '230fe955336445189f93986a9807a17e', '3', '', '953b77a21366497094303f293cc07424', '2025-10-23 09:11:39', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('a0a0661786924af3891e4703388f424c', '12cb92fc10724b4ca762bb38bb624fd2', 'e34b22d6f29b4334acaac995988f700f', 'fc1c718052ef4490b8dfaf83f2185dc2', '1', 'this_dept_and_below', '953b77a21366497094303f293cc07424', '2025-10-23 09:11:39', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('a8b060a6a846437a8e237c40c4d0cf23', 'a2aa4cf0ba1f46049cbe5a5d966446e8', '65ecc19d88e64e39b718c5f96f326f6f', '7f5ef7885be54be782a61461000eb130', '3', '', '953b77a21366497094303f293cc07424', '2025-10-15 15:29:47', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('aa8cc987a0e24f29943f108d52a9f148', 'a2aa4cf0ba1f46049cbe5a5d966446e8', '65ecc19d88e64e39b718c5f96f326f6f', '2304e56278124d2c852fed8ba5219ce5', '2', 'this_dept', '953b77a21366497094303f293cc07424', '2025-10-15 15:29:47', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('ab66b4b58d034f5398b40002ed3c31d3', '12cb92fc10724b4ca762bb38bb624fd2', 'e34b22d6f29b4334acaac995988f700f', '90695bd4b3af406a94c54170515c3b5c', '2', 'this_dept_and_below', '953b77a21366497094303f293cc07424', '2025-10-23 09:11:39', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('ad4b1b318a354bff909f41e55705e46b', 'a2aa4cf0ba1f46049cbe5a5d966446e8', '805b5bac807043c9aedde416c0be9343', '7a39fb17ab654a1da60b93d73b422789', '2', 'this_dept', '953b77a21366497094303f293cc07424', '2025-10-23 09:11:48', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('af30895b3f234a099af62f5252374906', '12cb92fc10724b4ca762bb38bb624fd2', 'e34b22d6f29b4334acaac995988f700f', '30ab7879300841e29f725a934fd82ba8', '3', '', '953b77a21366497094303f293cc07424', '2025-10-23 09:11:39', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('afd419da5163466e9284fb2573c02db4', 'a2aa4cf0ba1f46049cbe5a5d966446e8', '65ecc19d88e64e39b718c5f96f326f6f', 'fc1c718052ef4490b8dfaf83f2185dc2', '1', 'this_dept', '953b77a21366497094303f293cc07424', '2025-10-15 15:29:47', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('b00c772563c14deb9c9cbce495be0370', '12cb92fc10724b4ca762bb38bb624fd2', 'e34b22d6f29b4334acaac995988f700f', '0562128730274a47909e0ddd9959d116', '3', '', '953b77a21366497094303f293cc07424', '2025-10-23 09:11:39', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('b0ed78ccd4de4ed990510f36c4914e46', 'a2aa4cf0ba1f46049cbe5a5d966446e8', '6afbfef6f5ba4d0eb3a2efc505a519c8', '99f7a123b0af45c3a16f0430eb9c29c0', '3', '', '953b77a21366497094303f293cc07424', '2025-10-30 09:29:09', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('b433d09d4e6346bb8c0a1158c296a9e8', 'a2aa4cf0ba1f46049cbe5a5d966446e8', '65ecc19d88e64e39b718c5f96f326f6f', '30ab7879300841e29f725a934fd82ba8', '3', '', '953b77a21366497094303f293cc07424', '2025-10-15 15:29:47', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('b53b78667c0f49de9472c88901abd335', 'a2aa4cf0ba1f46049cbe5a5d966446e8', '805b5bac807043c9aedde416c0be9343', '477f429a7f4e486c9204fd5f8f058a61', '3', '', '953b77a21366497094303f293cc07424', '2025-10-23 09:11:48', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('b87e1ed5d1674aebacf6bcd27c8734c7', 'a2aa4cf0ba1f46049cbe5a5d966446e8', '65ecc19d88e64e39b718c5f96f326f6f', 'a97f98a231794561a74820ed612b7e04', '3', '', '953b77a21366497094303f293cc07424', '2025-10-15 15:29:47', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('badbcf0fbded433b8fb6c668be165b56', 'a2aa4cf0ba1f46049cbe5a5d966446e8', '805b5bac807043c9aedde416c0be9343', '2a91b59dc12946ac80d0710f76bdd6f2', '1', '', '953b77a21366497094303f293cc07424', '2025-10-23 09:11:48', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('bd1f57e29a704459be77046193768840', 'a2aa4cf0ba1f46049cbe5a5d966446e8', '6afbfef6f5ba4d0eb3a2efc505a519c8', 'ccd84ed949bb4a3ba74f0ead6eb4829f', '3', '', '953b77a21366497094303f293cc07424', '2025-10-30 09:29:09', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('bfd3f712e51d428f8b6d1a210122e287', 'a2aa4cf0ba1f46049cbe5a5d966446e8', '805b5bac807043c9aedde416c0be9343', '9e1c556a8f014a41a65595d65a93c1d1', '2', '', '953b77a21366497094303f293cc07424', '2025-10-23 09:11:48', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('c0b967e84cfa4248b3ac7ad29d1617c8', 'a2aa4cf0ba1f46049cbe5a5d966446e8', '65ecc19d88e64e39b718c5f96f326f6f', '45af68a141a74f298fb4c00a9e992c70', '3', '', '953b77a21366497094303f293cc07424', '2025-10-15 15:29:47', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('c1316fad600d44538077c1b0d32b9d82', 'a2aa4cf0ba1f46049cbe5a5d966446e8', '65ecc19d88e64e39b718c5f96f326f6f', 'f636bbd19db744bf8f0e1b4511072b54', '3', '', '953b77a21366497094303f293cc07424', '2025-10-15 15:29:47', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('c34989edf9434fe78d3be7c4364d45b9', '12cb92fc10724b4ca762bb38bb624fd2', 'e34b22d6f29b4334acaac995988f700f', '0a16a6b6abd24d35afe3d7d28cc41171', '3', '', '953b77a21366497094303f293cc07424', '2025-10-23 09:11:39', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('c481a3f10d124916b581bf7299ab8c3c', 'a2aa4cf0ba1f46049cbe5a5d966446e8', '6afbfef6f5ba4d0eb3a2efc505a519c8', 'a52a6f56b2294289a731ff4d8cff42a9', '2', 'this_dept_and_below', '953b77a21366497094303f293cc07424', '2025-10-30 09:29:09', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('c48e2a2dc6474411964ec218e345779a', 'a2aa4cf0ba1f46049cbe5a5d966446e8', '65ecc19d88e64e39b718c5f96f326f6f', '9ecb0ab22ee0479c8a8f228f7896dcd2', '3', '', '953b77a21366497094303f293cc07424', '2025-10-15 15:29:47', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('c9ea5f8a59d24647bfc1810dd930f0f0', 'a2aa4cf0ba1f46049cbe5a5d966446e8', '6afbfef6f5ba4d0eb3a2efc505a519c8', '8f81d3e08b944a0ea883f53e01598501', '3', '', '953b77a21366497094303f293cc07424', '2025-10-30 09:29:09', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('cafe51cdc8ab497cb3a28022b82e150e', '12cb92fc10724b4ca762bb38bb624fd2', 'e34b22d6f29b4334acaac995988f700f', '7c2fbb95bba34637a8f6d07ba48ca546', '3', '', '953b77a21366497094303f293cc07424', '2025-10-23 09:11:39', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('cd6cb756f3024460b6451d4f1d252c6e', 'a2aa4cf0ba1f46049cbe5a5d966446e8', '65ecc19d88e64e39b718c5f96f326f6f', '0a16a6b6abd24d35afe3d7d28cc41171', '3', '', '953b77a21366497094303f293cc07424', '2025-10-15 15:29:47', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('cdfb5a28e7944682b29588fc037d308e', '12cb92fc10724b4ca762bb38bb624fd2', 'e34b22d6f29b4334acaac995988f700f', '7a39fb17ab654a1da60b93d73b422789', '2', 'this_dept_and_below', '953b77a21366497094303f293cc07424', '2025-10-23 09:11:39', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('d101397a1ad94048ad8cacd402bc4600', 'a2aa4cf0ba1f46049cbe5a5d966446e8', '65ecc19d88e64e39b718c5f96f326f6f', 'f15324e4e60a4839bbb4dd1a5c755545', '3', '', '953b77a21366497094303f293cc07424', '2025-10-15 15:29:47', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('d22bb9d5335048559ecb78fffbdd1f51', 'a2aa4cf0ba1f46049cbe5a5d966446e8', '65ecc19d88e64e39b718c5f96f326f6f', '03581a90c87d4fd18bdbc77a7a8b7d42', '3', '', '953b77a21366497094303f293cc07424', '2025-10-15 15:29:47', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('d2e12f4bc2dc40228030d816da1a4975', 'a2aa4cf0ba1f46049cbe5a5d966446e8', '65ecc19d88e64e39b718c5f96f326f6f', '997ceabcb58c487a9d61ac2672390529', '3', '', '953b77a21366497094303f293cc07424', '2025-10-15 15:29:47', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('d3a8a59b3d8a42b69a4e8da0387c9615', 'a2aa4cf0ba1f46049cbe5a5d966446e8', '6afbfef6f5ba4d0eb3a2efc505a519c8', '2304e56278124d2c852fed8ba5219ce5', '2', 'this_dept_and_below', '953b77a21366497094303f293cc07424', '2025-10-30 09:29:09', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('d554321766f743de94884931dccd78fc', 'a2aa4cf0ba1f46049cbe5a5d966446e8', '805b5bac807043c9aedde416c0be9343', '9c4610cb8cd14f7c90893d877ac9e364', '3', '', '953b77a21366497094303f293cc07424', '2025-10-23 09:11:48', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('d6d76585d524428f98ebbd0d14e8b886', 'a2aa4cf0ba1f46049cbe5a5d966446e8', '805b5bac807043c9aedde416c0be9343', '30d7c3bfc9d646309c52766b5041422e', '3', '', '953b77a21366497094303f293cc07424', '2025-10-23 09:11:48', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('d6f53f3c67cb4b89bba1965c4380731c', 'a2aa4cf0ba1f46049cbe5a5d966446e8', '805b5bac807043c9aedde416c0be9343', '520f5be4b4414369a7a679e256db8d81', '3', '', '953b77a21366497094303f293cc07424', '2025-10-23 09:11:48', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('d8804cbed024456cb2c6f79160533a9a', 'a2aa4cf0ba1f46049cbe5a5d966446e8', '65ecc19d88e64e39b718c5f96f326f6f', 'c28884e311364d4abd7a13ab3fd89bdb', '2', 'this_dept', '953b77a21366497094303f293cc07424', '2025-10-15 15:29:47', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('d8f828f945144c68b2a293a003ae3c83', 'a2aa4cf0ba1f46049cbe5a5d966446e8', '805b5bac807043c9aedde416c0be9343', 'fc1c718052ef4490b8dfaf83f2185dc2', '1', 'this_dept', '953b77a21366497094303f293cc07424', '2025-10-23 09:11:48', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('dc5486f102a54136823fb590e2669673', '12cb92fc10724b4ca762bb38bb624fd2', 'e34b22d6f29b4334acaac995988f700f', '91a86b43c4294973941343627c95dab4', '3', '', '953b77a21366497094303f293cc07424', '2025-10-23 09:11:39', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('dee199880e6440cc9e997cb7f0295e29', '12cb92fc10724b4ca762bb38bb624fd2', 'e34b22d6f29b4334acaac995988f700f', '520f5be4b4414369a7a679e256db8d81', '3', '', '953b77a21366497094303f293cc07424', '2025-10-23 09:11:39', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('e140a32654f34c4990dbebd2bb3ae60a', 'a2aa4cf0ba1f46049cbe5a5d966446e8', '65ecc19d88e64e39b718c5f96f326f6f', 'fba3ca1ccf23460f8454bd1b2d4a6f1a', '3', '', '953b77a21366497094303f293cc07424', '2025-10-15 15:29:47', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('e159727d21c64457b3bd9911a8cc99f3', 'a2aa4cf0ba1f46049cbe5a5d966446e8', '65ecc19d88e64e39b718c5f96f326f6f', '1d91eb70b45f4d5ea8f279323c8caafa', '3', '', '953b77a21366497094303f293cc07424', '2025-10-15 15:29:47', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('e41de6068fa942698a636f7e59518cde', 'a2aa4cf0ba1f46049cbe5a5d966446e8', '65ecc19d88e64e39b718c5f96f326f6f', '514a8e2cf3694e108c13c9ebd45282e8', '3', '', '953b77a21366497094303f293cc07424', '2025-10-15 15:29:47', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('e5af9cd0632b436788d8bca69308f02b', 'a2aa4cf0ba1f46049cbe5a5d966446e8', '6afbfef6f5ba4d0eb3a2efc505a519c8', 'fc1c718052ef4490b8dfaf83f2185dc2', '1', 'this_dept_and_below', '953b77a21366497094303f293cc07424', '2025-10-30 09:29:09', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('e870ad809c494ef88397d4e07c527be5', '12cb92fc10724b4ca762bb38bb624fd2', 'e34b22d6f29b4334acaac995988f700f', '0c2ec443c2fa43eba2521ca4bdd5ffaa', '3', '', '953b77a21366497094303f293cc07424', '2025-10-23 09:11:39', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('eafdcda3c11e463d85f8f62bcb9e6528', 'a2aa4cf0ba1f46049cbe5a5d966446e8', '65ecc19d88e64e39b718c5f96f326f6f', '477f429a7f4e486c9204fd5f8f058a61', '3', '', '953b77a21366497094303f293cc07424', '2025-10-15 15:29:47', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('ef5afb7517d54322b6fb553dfcaaf36c', 'a2aa4cf0ba1f46049cbe5a5d966446e8', '65ecc19d88e64e39b718c5f96f326f6f', 'cd2c5bb506954b89a1bb48c35d6f747f', '3', '', '953b77a21366497094303f293cc07424', '2025-10-15 15:29:47', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('f1103e8b84e24fef897fb59b7105d92f', 'a2aa4cf0ba1f46049cbe5a5d966446e8', '805b5bac807043c9aedde416c0be9343', '24303c4b66a34032b7ef943ea903ec42', '2', '', '953b77a21366497094303f293cc07424', '2025-10-23 09:11:48', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('f28739b701a347558afe36c0ed69d616', 'a2aa4cf0ba1f46049cbe5a5d966446e8', '805b5bac807043c9aedde416c0be9343', 'a3bacdad9dec428e948e472f4b744825', '3', '', '953b77a21366497094303f293cc07424', '2025-10-23 09:11:48', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('f3ebe7b308004bbab14659c9b63faa75', 'a2aa4cf0ba1f46049cbe5a5d966446e8', '805b5bac807043c9aedde416c0be9343', '0c2ec443c2fa43eba2521ca4bdd5ffaa', '3', '', '953b77a21366497094303f293cc07424', '2025-10-23 09:11:48', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('f46d703804bc48adb62c0bc24f786cea', 'a2aa4cf0ba1f46049cbe5a5d966446e8', '65ecc19d88e64e39b718c5f96f326f6f', '9c4610cb8cd14f7c90893d877ac9e364', '3', '', '953b77a21366497094303f293cc07424', '2025-10-15 15:29:47', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('f48ccb565f4747679b7afd955528fb09', '12cb92fc10724b4ca762bb38bb624fd2', 'e34b22d6f29b4334acaac995988f700f', '477f429a7f4e486c9204fd5f8f058a61', '3', '', '953b77a21366497094303f293cc07424', '2025-10-23 09:11:39', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('f8c2b17c536647b1b159da5b917cce78', 'a2aa4cf0ba1f46049cbe5a5d966446e8', '805b5bac807043c9aedde416c0be9343', '3487c80fded34204849859598f014c8a', '3', '', '953b77a21366497094303f293cc07424', '2025-10-23 09:11:48', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('f9fa03cfbff84385b7c04cbefef3de76', 'a2aa4cf0ba1f46049cbe5a5d966446e8', '6afbfef6f5ba4d0eb3a2efc505a519c8', '8cc5876c17e249f0907bdfd7846214ba', '3', '', '953b77a21366497094303f293cc07424', '2025-10-30 09:29:09', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('fa219d1db01344c893fb9dda8efb991b', '12cb92fc10724b4ca762bb38bb624fd2', 'e34b22d6f29b4334acaac995988f700f', '6de8d2de77874f56b58d5b8d3b794cbc', '2', 'this_dept_and_below', '953b77a21366497094303f293cc07424', '2025-10-23 09:11:39', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('fb6e9c82c4db4ac383f79e29e9f9189a', '12cb92fc10724b4ca762bb38bb624fd2', 'e34b22d6f29b4334acaac995988f700f', '2a91b59dc12946ac80d0710f76bdd6f2', '1', '', '953b77a21366497094303f293cc07424', '2025-10-23 09:11:39', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES ('fc6e4143498a4929ba55c69362d87e90', 'a2aa4cf0ba1f46049cbe5a5d966446e8', '65ecc19d88e64e39b718c5f96f326f6f', '7a39fb17ab654a1da60b93d73b422789', '2', 'this_dept', '953b77a21366497094303f293cc07424', '2025-10-15 15:29:47', NULL, NULL);

-- ----------------------------
-- Table structure for sys_setting
-- ----------------------------
DROP TABLE IF EXISTS `sys_setting`;
CREATE TABLE `sys_setting`  (
  `setting_id` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '主键',
  `created_user_id` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '创建用户',
  `created_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_user_id` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '更新用户',
  `updated_time` datetime NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`setting_id`) USING BTREE,
  UNIQUE INDEX `only_setting_id`(`setting_id` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '配置信息' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_setting
-- ----------------------------

-- ----------------------------
-- Table structure for sys_task
-- ----------------------------
DROP TABLE IF EXISTS `sys_task`;
CREATE TABLE `sys_task`  (
  `task_id` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '任务ID',
  `task_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '任务名称',
  `cron_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '执行时间(字典-sys_task_cron_expr)',
  `cron_expr` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT 'Cron表达式-字典值',
  `func_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '要执行的函数名',
  `task_params` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL COMMENT 'JSON格式参数',
  `task_status` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '任务状态(字典1-运行中 2-停止)',
  `last_run_time` datetime NULL DEFAULT NULL COMMENT '最近一次开启/关闭时间',
  `next_run_time` datetime NULL DEFAULT NULL COMMENT '下次执行时间-预留',
  `desc` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '描述',
  `created_user_id` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '创建用户',
  `created_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_user_id` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '更新用户',
  `updated_time` datetime NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`task_id`) USING BTREE,
  UNIQUE INDEX `uniq_task_id`(`task_id` ASC) USING BTREE,
  UNIQUE INDEX `uniq_task_name`(`task_name` ASC) USING BTREE,
  INDEX `idx_task_status`(`task_status` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '定时任务表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_task
-- ----------------------------
INSERT INTO `sys_task` VALUES ('2e10f86cb8b0407aaf26c3031c7ab5d1', '测试任务', '每秒执行一次', '* * * * * *', 'ceshi1', '{}', '2', NULL, NULL, '测试定时任务', '953b77a21366497094303f293cc07424', '2025-10-30 10:20:13', NULL, NULL);

-- ----------------------------
-- Table structure for sys_user
-- ----------------------------
DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user`  (
  `user_id` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '主键',
  `user_phone` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '手机号',
  `user_name` varchar(90) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户名',
  `nick_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '昵称',
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '密码',
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '头像',
  `email` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '邮箱',
  `gender` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '性别(字典)',
  `birth` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '出生日期',
  `user_status` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '状态(字典)',
  `signature` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '个性签名',
  `desc` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '描述',
  `created_user_id` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '创建用户',
  `created_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_user_id` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '更新用户',
  `updated_time` datetime NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`user_id`) USING BTREE,
  UNIQUE INDEX `only_user_phone`(`user_phone` ASC) USING BTREE,
  UNIQUE INDEX `only_user_name`(`user_name` ASC) USING BTREE,
  UNIQUE INDEX `only_user_id`(`user_id` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '管理用户' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_user
-- ----------------------------
INSERT INTO `sys_user` VALUES ('332314d172134598a250e4d05e6ea2d9', '15250854548', 'user', 'user', '$2a$14$.tz4JgCKtyoCXK5uaXBwpuEK7OOKxSheyxmy/bdElkZBW/.suDA.C', 'http://localhost:9091/api/images/nv1.jpg', '', '1', '2025-10-17', '1', '工程师 | 值班中', '', '953b77a21366497094303f293cc07424', '2025-10-17 11:23:11', '953b77a21366497094303f293cc07424', '2025-10-30 09:29:20');
INSERT INTO `sys_user` VALUES ('953b77a21366497094303f293cc07424', '19951500648', 'admin', 'admin', '$2a$14$o9NIns3vyhgoy9cSdfl2IOSPnM6chE0akwy.sVV5siSHp9MhgGVV2', 'http://localhost:9091/api/images/nv3.jpg', '', '1', '1995-04-05', '1', '工程师 | 值班中', '超级管理员', '953b77a21366497094303f293cc07424', '2025-07-14 09:15:07', '953b77a21366497094303f293cc07424', '2025-10-17 10:24:58');

-- ----------------------------
-- Table structure for sys_user_dept_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_dept_role`;
CREATE TABLE `sys_user_dept_role`  (
  `user_dept_role_id` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '主键',
  `user_id` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户id',
  `dept_id` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '所属组织id(为all时为系统用户)',
  `role_id` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '角色id',
  `post_id` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '所属岗位',
  `data_purview` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '用户-数据权限(字典)-预留暂时未用',
  `created_user_id` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '创建用户',
  PRIMARY KEY (`user_dept_role_id`) USING BTREE,
  UNIQUE INDEX `uniq_user_dept_role_id`(`user_dept_role_id` ASC) USING BTREE,
  UNIQUE INDEX `uniq_user_dept`(`user_id` ASC, `dept_id` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户组织角色关联关系' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_user_dept_role
-- ----------------------------
INSERT INTO `sys_user_dept_role` VALUES ('47881145ae834c5b965468a62a0ea6eb', '332314d172134598a250e4d05e6ea2d9', 'a2aa4cf0ba1f46049cbe5a5d966446e8', '6afbfef6f5ba4d0eb3a2efc505a519c8', '', '', '953b77a21366497094303f293cc07424');
INSERT INTO `sys_user_dept_role` VALUES ('6470780e1d394ecbb17f74bfa66c36bf', '953b77a21366497094303f293cc07424', 'all', '5c9e854262ce42918a1b9d12500ca9ff', 'a90cee574f364bf891d9c2a21643eeba', '', '');

SET FOREIGN_KEY_CHECKS = 1;
