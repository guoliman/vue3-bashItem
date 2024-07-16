-- MySQL dump 10.13  Distrib 5.7.41, for osx10.18 (x86_64)
--
-- Host: 127.0.0.1    Database: youlai_boot
-- ------------------------------------------------------
-- Server version	5.7.36

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `sys_dept`
--

DROP TABLE IF EXISTS `sys_dept`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_dept` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `name` varchar(64) NOT NULL DEFAULT '' COMMENT '部门名称',
  `parent_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '父节点id',
  `sort` bigint(20) DEFAULT '0' COMMENT '显示顺序',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '1' COMMENT '状态(1:正常',
  `deleted` tinyint(3) unsigned DEFAULT '0' COMMENT '逻辑删除标识(1:已删除',
  `create_time` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime(3) DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=20 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='部门表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_dept`
--

LOCK TABLES `sys_dept` WRITE;
/*!40000 ALTER TABLE `sys_dept` DISABLE KEYS */;
INSERT INTO `sys_dept` VALUES (1,'技术部',0,1,1,0,'0001-01-01 00:00:00.000','2024-07-01 22:54:34.402'),(2,'研发组',1,2,1,0,'0001-01-01 00:00:00.000','2024-07-01 22:54:53.099'),(3,'测试组',1,1,1,0,'0001-01-01 00:00:00.000','2024-07-01 22:54:45.346'),(12,'aa',0,1,1,0,'2023-06-16 17:09:53.000','2023-06-19 14:55:10.520'),(18,'a3',12,1,1,0,'2023-06-29 19:34:49.913','2023-06-29 19:34:49.913'),(19,'a2',12,1,1,0,'2023-09-04 12:48:36.928','2023-09-04 12:48:36.928');
/*!40000 ALTER TABLE `sys_dept` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_dict`
--

DROP TABLE IF EXISTS `sys_dict`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_dict` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `type_code` varchar(64) DEFAULT NULL COMMENT '字典类型编码',
  `name` varchar(50) DEFAULT NULL COMMENT '字典项-名称',
  `value` varchar(50) DEFAULT NULL COMMENT '字典项-值',
  `sort` bigint(20) DEFAULT NULL COMMENT '排序',
  `status` tinyint(3) unsigned DEFAULT NULL COMMENT '状态(1:正常',
  `defaulted` tinyint(3) unsigned DEFAULT '0' COMMENT '是否默认(1:是',
  `remark` varchar(255) DEFAULT NULL COMMENT '备注',
  `create_time` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime(3) DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=34 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='字典数据表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_dict`
--

LOCK TABLES `sys_dict` WRITE;
/*!40000 ALTER TABLE `sys_dict` DISABLE KEYS */;
INSERT INTO `sys_dict` VALUES (1,'gender','男','1',1,1,0,NULL,'2019-05-05 13:07:52.000','2022-06-12 23:20:39.000'),(2,'gender','女','2',2,1,0,NULL,'2019-04-19 11:33:00.000','2019-07-02 14:23:05.000'),(3,'gender','未知','0',1,1,0,NULL,'2020-10-17 08:09:31.000','2020-10-17 08:09:31.000'),(6,'b0','b2','b2',1,1,0,'b2','2023-06-16 00:02:47.000','2023-07-04 17:41:06.856'),(10,'b0','b4','b4',1,1,0,'b4','2023-06-16 00:06:42.000','2023-06-16 00:06:42.000'),(11,'b0','a9-1','a9-1',1,1,0,'a9-123232','2023-06-16 00:16:50.000','2023-07-04 17:29:28.438'),(14,'a8','a8-1','a8-1',1,1,0,'a8-1','2023-06-16 00:18:12.000','2023-06-16 00:18:12.000'),(15,'a8','a7-1','a7-1',1,1,0,'a7-1','2023-06-16 00:18:31.000','2023-06-16 00:18:31.000'),(18,'a8','11188','11188',1,1,0,'11188','2023-06-16 00:37:21.000','2023-06-16 01:31:44.000'),(22,'a13','a1','aa',1,1,0,'','2023-06-16 00:41:49.000','2023-06-16 00:41:49.000'),(29,'a8','aaab','bbb',1,1,0,'','2023-06-16 01:39:34.000','2023-06-16 01:39:34.000'),(30,'b0','a7-1','a7-1',1,1,0,'','2023-07-04 17:19:48.855','2023-07-04 17:19:48.855'),(31,'b0','b2','b2',1,1,0,'','2023-07-04 17:40:28.493','2023-07-04 17:40:28.493'),(32,'b1','b0-1','b0-2',1,1,0,'','2023-07-04 17:53:41.000','2023-07-04 17:56:07.417'),(33,'b1','b0-1','b0-1',1,1,0,'','2023-07-04 17:54:02.000','2023-07-04 17:55:34.719');
/*!40000 ALTER TABLE `sys_dict` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_dict_type`
--

DROP TABLE IF EXISTS `sys_dict_type`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_dict_type` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键 ',
  `name` varchar(50) DEFAULT NULL COMMENT '类型名称',
  `code` varchar(50) DEFAULT NULL COMMENT '类型编码 唯一',
  `status` tinyint(3) unsigned DEFAULT NULL COMMENT '状态(0:正常',
  `remark` varchar(255) DEFAULT NULL COMMENT '备注',
  `create_time` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime(3) DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `type_code` (`code`) USING BTREE,
  UNIQUE KEY `code` (`code`)
) ENGINE=InnoDB AUTO_INCREMENT=23 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='字典类型表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_dict_type`
--

LOCK TABLES `sys_dict_type` WRITE;
/*!40000 ALTER TABLE `sys_dict_type` DISABLE KEYS */;
INSERT INTO `sys_dict_type` VALUES (1,'性别','gender',1,'cccd','2019-12-06 19:03:32.000','2023-06-15 19:10:44.000'),(2,'岗位','quarters',1,'岗位细分','2023-06-15 19:54:37.000','2023-06-15 19:55:01.000'),(11,'a13','a13',1,'','2023-06-15 21:04:30.000','2023-07-04 17:37:22.101'),(12,'a4','a4',1,'','2023-06-15 21:04:45.000','2023-06-15 21:04:45.000'),(15,'a7','a7',1,'','2023-06-15 21:05:05.000','2023-06-15 21:05:05.000'),(16,'a8','a8',1,'','2023-06-15 21:05:13.000','2023-07-04 17:37:14.383'),(18,'b0','b0',1,'','2023-06-15 21:05:27.000','2023-07-04 17:37:17.815'),(22,'b0','b1',1,'','2023-07-04 17:42:23.000','2023-09-04 12:48:45.156');
/*!40000 ALTER TABLE `sys_dict_type` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_menu`
--

DROP TABLE IF EXISTS `sys_menu`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_menu` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `parent_id` bigint(20) NOT NULL COMMENT '父节点id',
  `name` varchar(64) NOT NULL DEFAULT '' COMMENT '菜单名称',
  `type` tinyint(4) NOT NULL COMMENT '菜单类型(1:菜单；2:目录；3:外链；4:按钮)',
  `path` varchar(128) DEFAULT '' COMMENT '路由路径(浏览器地址栏路径)',
  `component` varchar(128) DEFAULT NULL COMMENT '组件路径(vue页面完整路径)',
  `perm` varchar(191) DEFAULT '' COMMENT '权限标识',
  `visible` tinyint(3) unsigned NOT NULL DEFAULT '1' COMMENT '显示状态(1-显示',
  `sort` bigint(20) DEFAULT '0' COMMENT '排序',
  `icon` varchar(64) DEFAULT NULL COMMENT '菜单图标',
  `redirect` varchar(128) DEFAULT NULL COMMENT '跳转路径',
  `create_time` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime(3) DEFAULT NULL COMMENT '更新时间',
  `api_path` varchar(255) DEFAULT '' COMMENT '后端请求地址',
  `api_type` varchar(64) DEFAULT '' COMMENT '后端接口类型',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=147 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='菜单管理';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_menu`
--

LOCK TABLES `sys_menu` WRITE;
/*!40000 ALTER TABLE `sys_menu` DISABLE KEYS */;
INSERT INTO `sys_menu` VALUES (1,0,'系统管理',2,'/system','Layout','',1,0,'system','/system/user','2021-08-28 09:12:21.000','2023-07-06 16:11:11.967','',''),(2,1,'用户管理',1,'user','system/user/index','',1,3,'user','','2021-08-28 09:12:21.000','2023-06-29 15:02:46.730','',''),(3,1,'角色管理',1,'role','system/role/index',NULL,1,2,'role',NULL,'2021-08-28 09:12:21.000','2021-08-28 09:12:21.000','',''),(4,1,'菜单管理',1,'menu','system/menu/index','',1,1,'menu','','2021-08-28 09:12:21.000','2023-06-29 15:02:38.098','',''),(5,1,'部门管理',1,'dept','system/dept/index',NULL,1,4,'tree',NULL,'2021-08-28 09:12:21.000','2021-08-28 09:12:21.000','',''),(6,1,'字典管理',1,'dict','system/dict/index',NULL,1,5,'dict',NULL,'2021-08-28 09:12:21.000','2021-08-28 09:12:21.000','',''),(26,0,'外部链接',2,'/external-link','Layout',NULL,1,8,'link','noredirect','2022-02-17 22:51:20.000','2022-02-17 22:51:20.000','',''),(30,26,'百度外链',3,'https://www.baidu.com','','',1,1,'document','','2022-02-18 00:01:40.000','2023-07-06 16:10:49.078','',''),(31,2,'用户新增',4,'','','sys:user:add',1,2,'','','2022-10-23 11:04:08.000','2023-06-29 11:04:17.903','/api/authAdmin/users/userOperate','POST'),(32,2,'用户编辑',4,'','','sys:user:edit',1,2,'','','2022-10-23 11:04:08.000','2023-06-29 11:04:34.639','/api/authAdmin/users/userOperate','PUT'),(33,2,'用户删除',4,'','','sys:user:delete',1,3,'','','2022-10-23 11:04:08.000','2023-06-29 11:04:42.512','/api/authAdmin/users/userOperate','DELETE'),(36,0,'组件封装',2,'/component','Layout',NULL,1,10,'menu','','2022-10-31 09:18:44.000','2022-10-31 09:18:47.000','',''),(37,36,'富文本编辑器',1,'wang-editor','demo/wang-editor',NULL,1,1,'','',NULL,NULL,'',''),(38,36,'图片上传',1,'upload','demo/upload',NULL,1,2,'','','2022-11-20 23:16:30.000','2022-11-20 23:16:32.000','',''),(39,36,'图标选择器',1,'icon-selector','demo/icon-selector',NULL,1,3,'','','2022-11-20 23:16:30.000','2022-11-20 23:16:32.000','',''),(40,0,'接口',2,'/api','Layout',NULL,1,7,'api','','2022-02-17 22:51:20.000','2022-02-17 22:51:20.000','',''),(41,40,'接口文档',1,'apidoc','demo/api-doc',NULL,1,1,'api','','2022-02-17 22:51:20.000','2022-02-17 22:51:20.000','',''),(70,3,'角色新增',4,'','','sys:role:add',1,2,'','','2023-05-20 23:39:09.000','2023-06-29 14:27:31.520','/api/authAdmin/roles/roleOperate','POST'),(71,3,'角色编辑',4,'','','sys:role:edit',1,2,'','','2023-05-20 23:40:31.000','2023-06-29 14:27:39.058','/api/authAdmin/roles/roleOperate','PUT'),(72,3,'角色删除',4,'','','sys:role:delete',1,2,'','','2023-05-20 23:41:08.000','2023-06-29 14:28:09.031','/api/authAdmin/roles/roleOperate','DELETE'),(73,4,'菜单新增',4,'','','sys:menu:add',1,3,'','','2023-05-20 23:41:35.000','2023-06-29 14:45:45.913','/api/authAdmin/menus/menuOperate','POST'),(74,4,'菜单编辑',4,'','','sys:menu:edit',1,3,'','','2023-05-20 23:41:58.000','2023-06-29 14:45:32.740','/api/authAdmin/menus/menuOperate','PUT'),(75,4,'菜单删除',4,'','','sys:menu:delete',1,3,'','','2023-05-20 23:44:18.000','2023-06-29 14:45:39.475','/api/authAdmin/menus/menuOperate','DELETE'),(76,5,'部门新增',4,'','','sys:dept:add',1,3,'','','2023-05-20 23:45:00.000','2023-06-29 14:57:18.462','/api/authAdmin/dept/deptTree','POST'),(77,5,'部门编辑',4,'','','sys:dept:edit',1,3,'','','2023-05-20 23:46:16.000','2023-06-29 14:57:22.412','/api/authAdmin/dept/deptTree','PUT'),(78,5,'部门删除',4,'','','sys:dept:delete',1,3,'','','2023-05-20 23:46:36.000','2023-06-29 14:57:26.246','/api/authAdmin/dept/deptTree','DELETE'),(79,6,'字典类型新增',4,'','','sys:dict_type:add',1,2,'','','2023-05-21 00:16:06.000','2023-06-29 15:05:04.949','/api/authAdmin/dict/dictTypes','POST'),(81,6,'字典类型编辑',4,'','','sys:dict_type:edit',1,2,'','','2023-05-21 00:27:37.000','2023-06-29 15:04:01.630','/api/authAdmin/dict/dictTypes','PUT'),(84,6,'字典类型删除',4,'','','sys:dict_type:delete',1,2,'','','2023-05-21 00:29:39.000','2023-06-29 15:05:11.551','/api/authAdmin/dict/dictTypes','DELETE'),(85,6,'字典数据新增',4,'','','sys:dict:add',1,4,'','','2023-05-21 00:46:56.000','2023-06-29 15:05:38.783','/api/authAdmin/dict/dictSon','POST'),(86,6,'字典数据编辑',4,'','','sys:dict:edit',1,4,'','','2023-05-21 00:47:36.000','2023-06-29 15:05:16.303','/api/authAdmin/dict/dictSon','PUT'),(87,6,'字典数据删除',4,'','','sys:dict:delete',1,4,'','','2023-05-21 00:48:10.000','2023-06-29 15:05:26.550','/api/authAdmin/dict/dictSon','DELETE'),(88,2,'重置密码',4,'','','sys:user:reset_pwd',1,4,'','','2023-05-21 00:49:18.000','2023-06-29 17:42:17.425','/api/authAdmin/users/userPassChange','PATCH'),(90,94,'a1',2,'/systemaaa','Layout','',1,0,'api','aba','2023-06-19 18:04:20.000','2023-07-04 23:30:28.634','',''),(91,90,'a222',1,'user','system/user/indexaaa','',0,2,'advert','','2023-06-19 18:05:05.000','2023-06-29 17:56:49.743','',''),(92,91,'a3w',4,'','','a222:a3w:a31',1,1,'','','2023-06-19 18:05:31.000','2023-06-29 10:54:01.498','/afa','GET'),(94,90,'bmenu',1,'bmenu','bmenu/index','',1,1,'drag','','2023-06-28 13:53:50.000','2023-06-29 10:52:27.531','',''),(96,0,'c2',2,'/c2','Layout','',0,1,'api','/system/user','2023-06-28 15:54:13.000','2023-06-30 13:43:32.043','',''),(98,94,'bm1',4,'',NULL,'a1:bmenu:bm1',1,1,NULL,NULL,'2023-06-29 10:53:11.908','2023-06-29 10:53:11.908','/abaa','POST'),(99,2,'获取用户',4,'','','sys:user:get',1,1,'','','2023-06-29 11:03:45.000','2023-06-29 14:27:02.551','/api/authAdmin/users/userOperate','GET'),(100,2,'用户状态更改',4,'',NULL,'sys:user:userStatusChange',1,4,NULL,NULL,'2023-06-29 11:06:37.974','2023-06-29 11:06:37.974','/api/authAdmin/users/userStatusChange','PATCH'),(101,2,'获取all部门和all角色',4,'','','sys:user:getDeptRole',1,4,'','','2023-06-29 11:07:41.000','2023-06-29 20:17:20.181','/api/authAdmin/users/getDeptRole','GET'),(102,3,'获取角色',4,'',NULL,'sys:role:get',1,1,'lab',NULL,'2023-06-29 14:26:39.268','2023-06-29 14:26:39.268','/api/authAdmin/roles/roleOperate','GET'),(103,3,'获取角色的菜单权限',4,'','','sys:role:PostSonMenu',1,3,'','','2023-06-29 14:29:51.000','2023-06-29 14:31:59.460','/api/authAdmin/roles/roleSonMenu','POST'),(104,3,'更新角色的菜单权限',4,'','','sys:role:PutSonMenu',1,3,'','','2023-06-29 14:31:00.000','2023-06-29 14:31:11.029','/api/authAdmin/roles/roleSonMenu','PUT'),(105,3,'获取角色下拉列表',4,'','','sys:role:roleOptions',1,3,'','','2023-06-29 14:33:28.000','2023-06-29 14:33:50.677','/api/authAdmin/roles/roleOptions','GET'),(106,4,'菜单获取',4,'','','sys:menu:get',1,2,'','','2023-06-29 14:43:29.000','2023-06-29 14:44:55.826','/api/authAdmin/menus/menuOperate','GET'),(107,4,'角色分配权限时获取下拉菜单',4,'','','sys:menu:menuSelect',1,4,'','','2023-06-29 14:49:28.000','2023-07-05 09:47:46.510','/api/authAdmin/menus/menuSelect','GET'),(109,5,'获取部门',4,'','','sys:dept:get',1,2,'','','2023-06-29 14:54:43.000','2023-06-29 14:57:02.295','/api/authAdmin/dept/deptTree','GET'),(110,5,'获取部门下拉菜单',4,'','','sys:dept:getSelect',1,4,'','','2023-06-29 14:56:48.000','2023-06-29 14:57:47.639','/api/authAdmin/dept/deptOptions','GET'),(111,6,'字典类型获取',4,'','','sys:dict_type:get',1,1,'','','2023-06-29 15:06:44.000','2023-06-29 18:01:11.164','/api/authAdmin/dict/dictTypes','GET'),(112,6,'字典数据获取',4,'',NULL,'sys:dict:get',1,3,NULL,NULL,'2023-06-29 15:07:39.382','2023-06-29 15:07:39.382','/api/authAdmin/dict/dictSon','GET'),(113,0,'SLO',2,'/gdir','Layout','',1,1,'系统设置','','2023-06-29 18:27:56.000','2024-07-01 22:58:43.764','',''),(114,122,'gmenu1',1,'gmenu1','gmenu1','',1,1,'advert','','2023-06-29 18:28:24.000','2023-07-04 23:40:29.404','',''),(117,113,'gmenu2',1,'gmenu2','gmenu2','',1,1,'rabbitmq','','2023-06-29 18:31:42.000','2023-09-15 15:19:59.493','',''),(118,113,'baidu',3,'https://www.baidu.com','','',1,1,'api','','2023-06-29 18:33:49.000','2023-06-29 18:36:37.608','',''),(120,90,'c3',1,'c3','c3','',1,1,'','','2023-06-30 08:40:59.000','2023-06-30 08:46:07.978','',''),(121,90,'cc',1,'cc','cc','',1,1,'','','2023-06-30 08:50:51.000','2023-06-30 08:58:31.296','',''),(122,113,'aa11',1,'aa11','aa11','',1,1,'api',NULL,'2023-07-04 23:39:20.952','2023-07-04 23:39:20.952','',''),(123,4,'获取目录和菜单select',4,'','','sys:dirMenuS:get',1,4,'','','2023-07-05 09:49:06.000','2023-07-05 09:55:09.775','/api/authAdmin/menus/dirMenuS','GET'),(133,0,'多级菜单',2,'/demoDir','Layout','',1,3,'multi_level','','2023-07-06 15:29:40.000','2023-07-06 16:08:04.568','',''),(134,133,'第二层目录-1',1,'twoDir-1','demoDir/towDir-aa-1/index','',1,1,'nested',NULL,'2023-07-06 15:31:09.007','2023-07-06 15:31:09.007','',''),(135,134,'a菜单-1',1,'aaMenu-1','demoDir/towDir-aa-1/aaMenu-1','',1,1,NULL,NULL,'2023-07-06 15:32:18.876','2023-07-06 15:32:18.876','',''),(136,134,'a菜单-2',1,'aaMenu-2','demoDir/towDir-aa-1/aaMenu-2','',1,1,NULL,NULL,'2023-07-06 15:33:11.314','2023-07-06 15:33:11.314','',''),(137,133,'第二层目录-2',1,'twoDir-bb-2','demoDir/towDir-bb-2/index','',1,1,'shopping',NULL,'2023-07-06 15:34:28.427','2023-07-06 15:34:28.427','',''),(138,137,'bb菜单-1',1,'bbMenu-1','demoDir/towDir-bb-2/bbMenu-1','',1,1,NULL,NULL,'2023-07-06 15:35:41.982','2023-07-06 15:35:41.982','',''),(139,137,'bb菜单-2',1,'bbMenu-2','demoDir/towDir-bb-2/bbMenu-2','',1,1,'redis',NULL,'2023-07-06 15:36:02.280','2023-07-06 15:36:02.280','',''),(140,134,'第三层目录-1',1,'threeDir-cc-1','demoDir/towDir-aa-1/threeDir-cc-1/index','',1,3,'order','','2023-07-06 15:37:33.000','2023-07-06 16:09:00.631','',''),(141,140,'cc菜单-1.1',1,'ccMenu1-One','demoDir/towDir-aa-1/threeDir-cc-1/ccMenu1-One','',1,1,'number','','2023-07-06 15:38:18.000','2023-07-06 15:40:03.926','',''),(142,140,'cc菜单1-2',1,'ccMenu1-Two','demoDir/towDir-aa-1/threeDir-cc-1/ccMenu1-Two','',1,1,'monitor','','2023-07-06 15:39:22.000','2023-07-06 15:39:33.654','',''),(143,134,'第三层目录-2',1,'threeDir-cc-2','demoDir/towDir-aa-1/threeDir-cc-2/index','',1,3,'refresh','','2023-07-06 15:41:07.000','2023-07-06 16:09:08.589','',''),(144,143,'cc菜单2.1',1,'ccMenu2-One','demoDir/towDir-aa-1/threeDir-cc-2/ccMenu2-One','',1,1,'theme',NULL,'2023-07-06 15:42:09.038','2023-07-06 15:42:09.038','',''),(145,143,'cc菜单2.2',1,'ccMenu2-Two','demoDir/towDir-aa-1/threeDir-cc-2/ccMenu2-Two','',1,1,'link',NULL,'2023-07-06 15:42:50.184','2023-07-06 15:42:50.184','',''),(146,134,'百度外链',3,'https://www.baidu.com',NULL,'',1,2,NULL,NULL,'2023-07-06 16:08:50.385','2023-07-06 16:08:50.385','','');
/*!40000 ALTER TABLE `sys_menu` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_role`
--

DROP TABLE IF EXISTS `sys_role`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_role` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(64) NOT NULL DEFAULT '' COMMENT '角色名称',
  `code` varchar(32) DEFAULT NULL COMMENT '角色编码',
  `sort` bigint(20) DEFAULT '0' COMMENT '排序',
  `status` tinyint(1) DEFAULT '1' COMMENT '角色状态(1-正常；0-停用)',
  `data_scope` tinyint(3) unsigned DEFAULT '1' COMMENT '数据权限(0-所有数据；1-部门及子部门数据；2-本部门数据；3-本人数据)',
  `deleted` tinyint(1) NOT NULL DEFAULT '0' COMMENT '逻辑删除标识(0-未删除；1-已删除)',
  `create_time` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime(3) DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `name` (`name`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='角色表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_role`
--

LOCK TABLES `sys_role` WRITE;
/*!40000 ALTER TABLE `sys_role` DISABLE KEYS */;
INSERT INTO `sys_role` VALUES (1,'超级管理员','ROOT',1,1,0,0,'2021-05-21 14:56:51.000','2018-12-23 16:00:00.000'),(2,'系统管理员','ADMIN',2,1,1,0,'2021-03-25 12:39:54.000',NULL),(3,'访问游客','GUEST',3,1,2,0,'2021-05-26 15:49:05.000','2019-05-05 16:00:00.000'),(5,'系统管理员2','ADMIN1',2,1,1,0,'2021-03-25 12:39:54.000',NULL),(12,'系统管理员9','ADMIN1',2,1,1,0,'2021-03-25 12:39:54.000',NULL);
/*!40000 ALTER TABLE `sys_role` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_role_menu`
--

DROP TABLE IF EXISTS `sys_role_menu`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_role_menu` (
  `role_id` bigint(20) NOT NULL COMMENT '角色ID',
  `menu_id` bigint(20) NOT NULL COMMENT '菜单ID'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='角色和菜单关联表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_role_menu`
--

LOCK TABLES `sys_role_menu` WRITE;
/*!40000 ALTER TABLE `sys_role_menu` DISABLE KEYS */;
INSERT INTO `sys_role_menu` VALUES (2,1),(2,2),(2,3),(2,4),(2,5),(2,6),(2,11),(2,12),(2,19),(2,18),(2,17),(2,13),(2,14),(2,15),(2,16),(2,9),(2,10),(2,37),(2,32),(2,33),(2,39),(2,34),(2,26),(2,30),(2,31),(2,36),(2,38),(2,39),(2,40),(2,41),(2,1),(2,2),(2,3),(2,4),(2,5),(2,6),(2,26),(2,30),(2,31),(2,32),(2,33),(2,36),(2,37),(2,38),(2,39),(2,40),(2,41),(2,70),(2,71),(2,72),(2,73),(2,74),(2,75),(2,76),(2,77),(2,78),(2,79),(2,81),(2,84),(2,85),(2,86),(2,87),(2,88);
/*!40000 ALTER TABLE `sys_role_menu` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_user`
--

DROP TABLE IF EXISTS `sys_user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(64) DEFAULT NULL COMMENT '用户名',
  `nickname` varchar(64) DEFAULT NULL COMMENT '昵称',
  `gender` tinyint(3) unsigned DEFAULT '1' COMMENT '性别((1:男',
  `password` varchar(100) DEFAULT NULL COMMENT '密码',
  `dept_id` bigint(20) DEFAULT NULL COMMENT '部门ID',
  `avatar` varchar(191) DEFAULT '' COMMENT '用户头像',
  `mobile` varchar(20) DEFAULT NULL COMMENT '联系方式',
  `status` tinyint(3) unsigned DEFAULT '1' COMMENT '用户状态((1:正常',
  `email` varchar(128) DEFAULT NULL COMMENT '用户邮箱',
  `deleted` tinyint(3) unsigned DEFAULT '0' COMMENT '逻辑删除标识(0:未删除',
  `create_time` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime(3) DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `login_name` (`username`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=297 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='用户信息表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_user`
--

LOCK TABLES `sys_user` WRITE;
/*!40000 ALTER TABLE `sys_user` DISABLE KEYS */;
INSERT INTO `sys_user` VALUES (1,'root','gavin公司',0,'$2a$10$xVWsNOhHrCxh5UbpCE7/HuJ.PAOKcYAqRxD2CO2nVnJS.IAXkr5aq',0,'https://oss.youlai.tech/youlai-boot/2023/05/16/811270ef31f548af9cffc026dfc3777b.gif','17621590365',1,'youlaitech@163.com',0,'0001-01-01 00:00:00.000','2024-07-01 22:53:48.951'),(2,'admin','系统管理员',1,'uwpz7oPaCwxWC52/i2aPYA==',1,'https://oss.youlai.tech/youlai-boot/2023/05/16/811270ef31f548af9cffc026dfc3777b.gif','17621210366',1,'',0,'2019-10-10 13:41:22.000','2023-06-29 17:45:24.743'),(3,'test','测试小用户',1,'$2a$10$xVWsNOhHrCxh5UbpCE7/HuJ.PAOKcYAqRxD2CO2nVnJS.IAXkr5aq',3,'https://oss.youlai.tech/youlai-boot/2023/05/16/811270ef31f548af9cffc026dfc3777b.gif','17621210366',1,'youlaitech@163.com',0,'2021-06-05 01:31:29.000','2021-06-05 01:31:29.000'),(292,'ff','王五',1,'uwpz7oPaCwxWC52/i2aPYA==',18,'https://oss.youlai.tech/youlai-boot/2023/05/16/811270ef31f548af9cffc026dfc3777b.gif','13344560601',0,'aa@qq.com',0,'2023-06-23 17:28:33.000','2024-07-01 23:05:00.943'),(293,'aa','李四',1,'8MsTa9dZLI2sqdYcX6Wy7A==',2,'https://oss.youlai.tech/youlai-boot/2023/05/16/811270ef31f548af9cffc026dfc3777b.gif','',1,'',0,'2023-06-28 11:28:11.000','2024-07-01 23:04:51.550'),(296,'a2','张三',1,NULL,12,'https://oss.youlai.tech/youlai-boot/2023/05/16/811270ef31f548af9cffc026dfc3777b.gif','',1,'',0,'2023-09-04 12:48:23.000','2024-07-01 23:04:32.806');
/*!40000 ALTER TABLE `sys_user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_user_role`
--

DROP TABLE IF EXISTS `sys_user_role`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_user_role` (
  `user_id` bigint(20) NOT NULL COMMENT '用户ID',
  `role_id` bigint(20) NOT NULL COMMENT '角色ID',
  PRIMARY KEY (`user_id`,`role_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='用户和角色关联表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_user_role`
--

LOCK TABLES `sys_user_role` WRITE;
/*!40000 ALTER TABLE `sys_user_role` DISABLE KEYS */;
INSERT INTO `sys_user_role` VALUES (0,3),(0,4),(0,5),(1,1),(2,1),(2,2),(3,3),(292,3),(293,3),(296,1);
/*!40000 ALTER TABLE `sys_user_role` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2024-07-16 23:59:58
