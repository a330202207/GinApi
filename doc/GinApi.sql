SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for api_admin
-- ----------------------------
DROP TABLE IF EXISTS `api_admin`;
CREATE TABLE `api_admin` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '管理员ID',
  `user_name` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '管理员名称',
  `password` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '密码',
  `phone` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '电话',
  `login_ip` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '登录时IP',
  `login_date` timestamp NULL DEFAULT NULL COMMENT '登录时间',
  `login_cnt` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '登录次数',
  `create_ip` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '创建时IP',
  `status` tinyint(2) unsigned NOT NULL DEFAULT '1' COMMENT '状态：1-正常，2-禁止，3-已删除',
  `created_at` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建日期',
  `updated_at` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '修改时间',
  `deleted_at` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `user_name` (`user_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='系统管理员表';

-- ----------------------------
-- Records of api_admin
-- ----------------------------
BEGIN;
INSERT INTO `api_admin` VALUES (1, 'admin', '$2a$10$zzjAmJrsR0hk8UBbL9P3OOTLBBNEjtME1G5s3Vl2./.TwHrroDwkm', '15202980611', '127.0.0.1', '2020-01-09 18:34:18', 1, '127.0.0.1', 1, 1547726169, 1568193871, 1578566058);
COMMIT;

-- ----------------------------
-- Table structure for api_admin_role
-- ----------------------------
DROP TABLE IF EXISTS `api_admin_role`;
CREATE TABLE `api_admin_role` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `admin_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '管理员ID',
  `role_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '角色ID',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='系统管理员所属角色中间表';

-- ----------------------------
-- Records of api_admin_role
-- ----------------------------
BEGIN;
INSERT INTO `api_admin_role` VALUES (3, 1, 1);
COMMIT;

-- ----------------------------
-- Table structure for api_menu
-- ----------------------------
DROP TABLE IF EXISTS `api_menu`;
CREATE TABLE `api_menu` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `parent_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '上级ID(1为顶级菜单)',
  `name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '资源名称',
  `menu_router` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '菜单路由',
  `order_by` tinyint(2) unsigned NOT NULL DEFAULT '0' COMMENT '排序',
  `created_at` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建日期',
  `updated_at` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='系统资源表';

-- ----------------------------
-- Records of api_menu
-- ----------------------------
BEGIN;
INSERT INTO `api_menu` VALUES (1, 0, '系统管理', '', 1, 1569112106, 1569112106);
INSERT INTO `api_menu` VALUES (2, 0, '管理员管理', '', 2, 1569112124, 1569112124);
INSERT INTO `api_menu` VALUES (3, 9, '角色列表', '/admin/role/role_index.html', 1, 1569212318, 1569212318);
INSERT INTO `api_menu` VALUES (4, 9, '菜单列表', '/admin/menu/menu_index.html', 2, 1569212379, 1569212379);
INSERT INTO `api_menu` VALUES (5, 10, '管理员列表', '/admin/admin/admin_index.html', 1, 1569212469, 1569212469);
INSERT INTO `api_menu` VALUES (6, 31, '添加角色页', '/admin/role/role_create.html', 1, 1569212509, 1569212509);
INSERT INTO `api_menu` VALUES (7, 31, '添加角色', '/admin/role/add', 2, 1569212531, 1569212531);
INSERT INTO `api_menu` VALUES (8, 31, '删除角色', '/admin/role/del', 3, 1569212620, 1569212620);
INSERT INTO `api_menu` VALUES (9, 31, '编辑角色页', '/admin/role/role_edit.html', 4, 1569212645, 1569212645);
INSERT INTO `api_menu` VALUES (10, 31, '保存角色', '/admin/role/save', 5, 1569212674, 1569212674);
INSERT INTO `api_menu` VALUES (11, 32, '添加菜单页', '/admin/menu/menu_create.html', 1, 1569212828, 1569212828);
INSERT INTO `api_menu` VALUES (12, 32, '添加菜单', '/admin/menu/add', 2, 1569212852, 1569212852);
INSERT INTO `api_menu` VALUES (13, 32, '添加下级菜单页', '/admin/menu/menu_add.html', 3, 1569212964, 1569212964);
INSERT INTO `api_menu` VALUES (14, 32, '删除菜单', '/admin/menu/del', 4, 1569212978, 1569212978);
INSERT INTO `api_menu` VALUES (15, 32, '编辑菜单页', '/admin/menu/menu_edit.html', 5, 1569213012, 1569213012);
INSERT INTO `api_menu` VALUES (16, 32, '保存菜单', '/admin/menu/save', 6, 1569213034, 1569213034);
INSERT INTO `api_menu` VALUES (17, 33, '添加管理员页', '/admin/admin/admin_create.html', 1, 1569213087, 1569213087);
INSERT INTO `api_menu` VALUES (18, 33, '添加管理员', '/admin/admin/add', 2, 1569213115, 1569213115);
INSERT INTO `api_menu` VALUES (19, 33, '删除管理员', '/admin/admin/del', 3, 1569213137, 1569213137);
INSERT INTO `api_menu` VALUES (20, 33, '编辑管理员页', '/admin/admin/admin_edit.html', 4, 1569213159, 1569213159);
INSERT INTO `api_menu` VALUES (21, 33, '保存管理员', '/admin/admin/save', 5, 1569213182, 1569213182);
INSERT INTO `api_menu` VALUES (22, 0, '后台首页', '/admin/backend_index.html', 15, 1575881087, 1575881087);
COMMIT;

-- ----------------------------
-- Table structure for api_role
-- ----------------------------
DROP TABLE IF EXISTS `api_role`;
CREATE TABLE `api_role` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '角色ID',
  `name` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '角色名',
  `status` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '状态：1-正常，2-已删除',
  `created_at` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建日期',
  `updated_at` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '修改时间',
  `deleted_at` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='系统角色表';

-- ----------------------------
-- Records of api_role
-- ----------------------------
BEGIN;
INSERT INTO `api_role` VALUES (1, '管理员', 1, 1568252204, 1568252204, 0);
COMMIT;

-- ----------------------------
-- Table structure for api_role_menu
-- ----------------------------
DROP TABLE IF EXISTS `api_role_menu`;
CREATE TABLE `api_role_menu` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `role_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '角色ID',
  `menu_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '菜单ID',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='系统角色对应菜单中间表';

-- ----------------------------
-- Records of api_role_menu
-- ----------------------------
BEGIN;
INSERT INTO `api_role_menu` VALUES (1, 1, 34);
INSERT INTO `api_role_menu` VALUES (2, 1, 35);
INSERT INTO `api_role_menu` VALUES (3, 1, 36);
INSERT INTO `api_role_menu` VALUES (4, 1, 38);
INSERT INTO `api_role_menu` VALUES (5, 1, 39);
INSERT INTO `api_role_menu` VALUES (6, 1, 40);
INSERT INTO `api_role_menu` VALUES (7, 1, 41);
INSERT INTO `api_role_menu` VALUES (8, 1, 42);
INSERT INTO `api_role_menu` VALUES (9, 1, 43);
INSERT INTO `api_role_menu` VALUES (10, 1, 44);
INSERT INTO `api_role_menu` VALUES (11, 1, 32);
INSERT INTO `api_role_menu` VALUES (12, 1, 37);
INSERT INTO `api_role_menu` VALUES (13, 1, 31);
INSERT INTO `api_role_menu` VALUES (14, 1, 9);
INSERT INTO `api_role_menu` VALUES (15, 1, 45);
INSERT INTO `api_role_menu` VALUES (16, 1, 46);
INSERT INTO `api_role_menu` VALUES (17, 1, 47);
INSERT INTO `api_role_menu` VALUES (18, 1, 48);
INSERT INTO `api_role_menu` VALUES (19, 1, 49);
INSERT INTO `api_role_menu` VALUES (20, 1, 33);
INSERT INTO `api_role_menu` VALUES (21, 1, 10);
INSERT INTO `api_role_menu` VALUES (22, 1, 53);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
