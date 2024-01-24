/*
SQLyog Ultimate v11.3 (64 bit)
MySQL - 5.7.32-log : Database - yiyuanguahao
*********************************************************************
*/


/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`yiyuanguahao` /*!40100 DEFAULT CHARACTER SET utf8 */;

USE `yiyuanguahao`;

/*Table structure for table `chat` */

DROP TABLE IF EXISTS `chat`;

CREATE TABLE `chat` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `yonghu_id` int(11) DEFAULT NULL COMMENT '提问用户',
  `chat_issue` varchar(200) DEFAULT NULL COMMENT '问题',
  `issue_time` timestamp NULL DEFAULT NULL COMMENT '问题时间 Search111',
  `chat_reply` varchar(200) DEFAULT NULL COMMENT '回复',
  `reply_time` timestamp NULL DEFAULT NULL COMMENT '回复时间 Search111',
  `zhuangtai_types` int(255) DEFAULT NULL COMMENT '状态',
  `chat_types` int(11) DEFAULT NULL COMMENT '数据类型',
  `insert_time` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8 COMMENT='在线咨询';

/*Data for the table `chat` */

insert  into `chat`(`id`,`yonghu_id`,`chat_issue`,`issue_time`,`chat_reply`,`reply_time`,`zhuangtai_types`,`chat_types`,`insert_time`) values (1,1,'询问一些信息','2022-03-11 21:10:30',NULL,NULL,2,1,'2022-03-11 21:10:31'),(2,1,'管理在后台可回复','2022-03-11 21:10:41',NULL,NULL,2,1,'2022-03-11 21:10:41'),(3,1,NULL,NULL,'管理回复用户的信息','2022-03-11 21:13:19',NULL,2,'2022-03-11 21:13:20'),(4,1,NULL,NULL,'已回复的信息将不能回复','2022-03-11 21:13:38',NULL,2,'2022-03-11 21:13:39');

/*Table structure for table `config` */

DROP TABLE IF EXISTS `config`;

CREATE TABLE `config` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `name` varchar(100) NOT NULL COMMENT '配置参数名称',
  `value` varchar(100) DEFAULT NULL COMMENT '配置参数值',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8 COMMENT='配置文件';

/*Data for the table `config` */

insert  into `config`(`id`,`name`,`value`) values (4,'轮播图1','http://localhost:8080/yiyuanguahao/upload/1647000993989.jpg'),(5,'轮播图2','http://localhost:8080/yiyuanguahao/upload/1647001010142.jpg'),(6,'轮播图3','http://localhost:8080/yiyuanguahao/upload/1647001038348.jpg');

/*Table structure for table `dictionary` */

DROP TABLE IF EXISTS `dictionary`;

CREATE TABLE `dictionary` (
  `id` int(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `dic_code` varchar(200) DEFAULT NULL COMMENT '字段',
  `dic_name` varchar(200) DEFAULT NULL COMMENT '字段名',
  `code_index` int(11) DEFAULT NULL COMMENT '编码',
  `index_name` varchar(200) DEFAULT NULL COMMENT '编码名字  Search111 ',
  `super_id` int(11) DEFAULT NULL COMMENT '父字段id',
  `beizhu` varchar(200) DEFAULT NULL COMMENT '备注',
  `create_time` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=38 DEFAULT CHARSET=utf8 COMMENT='字典表';

/*Data for the table `dictionary` */

insert  into `dictionary`(`id`,`dic_code`,`dic_name`,`code_index`,`index_name`,`super_id`,`beizhu`,`create_time`) values (14,'guahao_types','时间类型',1,'上午',NULL,NULL,'2022-03-11 17:54:04'),(15,'guahao_types','时间类型',2,'下午',NULL,NULL,'2022-03-11 17:54:04'),(16,'guahao_status_types','挂号状态',1,'未就诊',NULL,NULL,'2022-03-11 17:54:04'),(17,'guahao_status_types','挂号状态',2,'已就诊',NULL,NULL,'2022-03-11 17:54:04'),(18,'sex_types','性别',1,'男',NULL,NULL,'2022-03-11 17:54:05'),(19,'sex_types','性别',2,'女',NULL,NULL,'2022-03-11 17:54:05'),(20,'jiankangjiaoyu_types','健康教育类型',1,'健康教育类型1',NULL,NULL,'2022-03-11 17:54:05'),(21,'jiankangjiaoyu_types','健康教育类型',2,'健康教育类型2',NULL,NULL,'2022-03-11 17:54:05'),(22,'jiankangjiaoyu_types','健康教育类型',3,'健康教育类型3',NULL,NULL,'2022-03-11 17:54:05'),(23,'guahao_yesno_types','挂号审核',1,'审核中',NULL,NULL,'2022-03-11 17:54:05'),(24,'guahao_yesno_types','挂号审核',2,'通过',NULL,NULL,'2022-03-11 17:54:05'),(25,'guahao_yesno_types','挂号审核',3,'拒绝',NULL,NULL,'2022-03-11 17:54:05'),(26,'news_types','公告类型',1,'公告类型1',NULL,NULL,'2022-03-11 17:54:05'),(27,'news_types','公告类型',2,'公告类型2',NULL,NULL,'2022-03-11 17:54:05'),(28,'yisheng_types','科室',1,'科室1',NULL,NULL,'2022-03-11 17:54:05'),(29,'yisheng_types','科室',2,'科室2',NULL,NULL,'2022-03-11 17:54:05'),(30,'yisheng_types','科室',3,'科室3',NULL,NULL,'2022-03-11 17:54:05'),(31,'zhiwei_types','职位',1,'职位1',NULL,NULL,'2022-03-11 17:54:05'),(32,'zhiwei_types','职位',2,'职位2',NULL,NULL,'2022-03-11 17:54:05'),(33,'zhiwei_types','职位',3,'职位3',NULL,NULL,'2022-03-11 17:54:05'),(34,'chat_types','数据类型',1,'问题',NULL,NULL,'2022-03-11 17:54:05'),(35,'chat_types','数据类型',2,'回复',NULL,NULL,'2022-03-11 17:54:05'),(36,'zhuangtai_types','状态',1,'未回复',NULL,NULL,'2022-03-11 17:54:05'),(37,'zhuangtai_types','状态',2,'已回复',NULL,NULL,'2022-03-11 17:54:05');

/*Table structure for table `guahao` */

DROP TABLE IF EXISTS `guahao`;

CREATE TABLE `guahao` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键 ',
  `yisheng_id` int(11) DEFAULT NULL COMMENT '医生',
  `yonghu_id` int(11) DEFAULT NULL COMMENT '用户',
  `guahao_uuin_number` varchar(200) DEFAULT NULL COMMENT '就诊识别码',
  `guahao_time` date DEFAULT NULL COMMENT '挂号时间 Search111',
  `guahao_types` int(11) DEFAULT NULL COMMENT '时间类型 Search111',
  `guahao_status_types` int(11) DEFAULT NULL COMMENT '挂号状态',
  `guahao_yesno_types` int(11) DEFAULT NULL COMMENT '挂号审核',
  `guahao_yesno_text` text COMMENT '审核结果',
  `create_time` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8 COMMENT='挂号';

/*Data for the table `guahao` */

insert  into `guahao`(`id`,`yisheng_id`,`yonghu_id`,`guahao_uuin_number`,`guahao_time`,`guahao_types`,`guahao_status_types`,`guahao_yesno_types`,`guahao_yesno_text`,`create_time`) values (3,2,3,'339','2022-03-13',2,2,2,'通过审核','2022-03-11 17:54:11'),(4,3,1,'374','2022-03-13',1,NULL,1,NULL,'2022-03-11 17:54:11'),(15,1,1,'1647002713041','2022-03-12',2,NULL,1,NULL,'2022-03-11 20:45:17'),(16,3,1,'1647004198061','2022-03-13',2,NULL,1,NULL,'2022-03-11 21:10:04');

/*Table structure for table `jiankangjiaoyu` */

DROP TABLE IF EXISTS `jiankangjiaoyu`;

CREATE TABLE `jiankangjiaoyu` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `jiankangjiaoyu_name` varchar(200) DEFAULT NULL COMMENT '健康教育标题  Search111 ',
  `jiankangjiaoyu_types` int(11) DEFAULT NULL COMMENT '健康教育类型  Search111 ',
  `jiankangjiaoyu_photo` varchar(200) DEFAULT NULL COMMENT '健康教育图片',
  `insert_time` timestamp NULL DEFAULT NULL COMMENT '健康教育时间',
  `jiankangjiaoyu_content` text COMMENT '健康教育详情',
  `jiankangjiaoyu_delete` int(11) DEFAULT '1' COMMENT '假删',
  `create_time` timestamp NULL DEFAULT NULL COMMENT '创建时间 show1 show2 nameShow',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8 COMMENT='健康教育';

/*Data for the table `jiankangjiaoyu` */

insert  into `jiankangjiaoyu`(`id`,`jiankangjiaoyu_name`,`jiankangjiaoyu_types`,`jiankangjiaoyu_photo`,`insert_time`,`jiankangjiaoyu_content`,`jiankangjiaoyu_delete`,`create_time`) values (1,'健康教育标题1',3,'http://localhost:8080/yiyuanguahao/upload/1646998131255.jpeg','2022-03-11 17:54:11','<p>健康教育详情1</p>',1,'2022-03-11 17:54:11'),(2,'健康教育标题2',2,'http://localhost:8080/yiyuanguahao/upload/1646998124526.jpeg','2022-03-11 17:54:11','<p>健康教育详情2</p>',1,'2022-03-11 17:54:11'),(3,'健康教育标题3',2,'http://localhost:8080/yiyuanguahao/upload/1646998117251.jpeg','2022-03-11 17:54:11','<p>健康教育详情3</p>',1,'2022-03-11 17:54:11'),(4,'健康教育标题4',1,'http://localhost:8080/yiyuanguahao/upload/1646998108197.jpeg','2022-03-11 17:54:11','<p>健康教育详情4</p>',1,'2022-03-11 17:54:11'),(5,'健康教育标题5',1,'http://localhost:8080/yiyuanguahao/upload/1646998101142.jpeg','2022-03-11 17:54:11','<p>健康教育详情5</p>',1,'2022-03-11 17:54:11');

/*Table structure for table `news` */

DROP TABLE IF EXISTS `news`;

CREATE TABLE `news` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键 ',
  `news_name` varchar(200) DEFAULT NULL COMMENT '公告名称 Search111  ',
  `news_photo` varchar(200) DEFAULT NULL COMMENT '公告图片 ',
  `news_types` int(11) NOT NULL COMMENT '公告类型 ',
  `insert_time` timestamp NULL DEFAULT NULL COMMENT '公告发布时间 ',
  `news_content` text COMMENT '公告详情 ',
  `create_time` timestamp NULL DEFAULT NULL COMMENT '创建时间 show1 show2 nameShow',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8 COMMENT='公告信息';

/*Data for the table `news` */

insert  into `news`(`id`,`news_name`,`news_photo`,`news_types`,`insert_time`,`news_content`,`create_time`) values (1,'公告名称1','http://localhost:8080/yiyuanguahao/upload/news1.jpg',2,'2022-03-11 17:54:11','公告详情1','2022-03-11 17:54:11'),(2,'公告名称2','http://localhost:8080/yiyuanguahao/upload/news2.jpg',2,'2022-03-11 17:54:11','公告详情2','2022-03-11 17:54:11'),(3,'公告名称3','http://localhost:8080/yiyuanguahao/upload/news3.jpg',1,'2022-03-11 17:54:11','公告详情3','2022-03-11 17:54:11'),(4,'公告名称4','http://localhost:8080/yiyuanguahao/upload/news4.jpg',1,'2022-03-11 17:54:11','公告详情4','2022-03-11 17:54:11'),(5,'公告名称5','http://localhost:8080/yiyuanguahao/upload/news5.jpg',2,'2022-03-11 17:54:11','公告详情5','2022-03-11 17:54:11');

/*Table structure for table `token` */

DROP TABLE IF EXISTS `token`;

CREATE TABLE `token` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `userid` bigint(20) NOT NULL COMMENT '用户id',
  `username` varchar(100) NOT NULL COMMENT '用户名',
  `tablename` varchar(100) DEFAULT NULL COMMENT '表名',
  `role` varchar(100) DEFAULT NULL COMMENT '角色',
  `token` varchar(200) NOT NULL COMMENT '密码',
  `addtime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '新增时间',
  `expiratedtime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '过期时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8 COMMENT='token表';

/*Data for the table `token` */

insert  into `token`(`id`,`userid`,`username`,`tablename`,`role`,`token`,`addtime`,`expiratedtime`) values (1,1,'admin','users','管理员','fcs5cpaf4gh0qrubdjk9qzdih4ve6h1j','2022-03-11 19:09:39','2022-03-11 22:12:37'),(2,1,'a1','yonghu','用户','u99h1cduapxa6otyxn6husty2wumhhwa','2022-03-11 19:16:39','2022-03-11 22:10:52'),(3,2,'a2','yisheng','医生','r5sitkdkdkxh4wir9r3lnuk2yc98n0d4','2022-03-11 21:11:13','2022-03-11 22:11:13');

/*Table structure for table `users` */

DROP TABLE IF EXISTS `users`;

CREATE TABLE `users` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `username` varchar(100) NOT NULL COMMENT '用户名',
  `password` varchar(100) NOT NULL COMMENT '密码',
  `role` varchar(100) DEFAULT '管理员' COMMENT '角色',
  `addtime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '新增时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COMMENT='用户表';

/*Data for the table `users` */

insert  into `users`(`id`,`username`,`password`,`role`,`addtime`) values (1,'admin','admin','管理员','2022-05-01 00:00:00');

/*Table structure for table `yisheng` */

DROP TABLE IF EXISTS `yisheng`;

CREATE TABLE `yisheng` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键 ',
  `yisheng_uuid_number` varchar(200) DEFAULT NULL COMMENT '医生工号',
  `username` varchar(200) DEFAULT NULL COMMENT '账户',
  `password` varchar(200) DEFAULT NULL COMMENT '密码',
  `yisheng_name` varchar(200) DEFAULT NULL COMMENT '医生名称 Search111',
  `yisheng_types` int(11) DEFAULT NULL COMMENT '科室 Search111',
  `zhiwei_types` int(11) DEFAULT NULL COMMENT '职位 Search111',
  `yisheng_zhichneg` varchar(200) DEFAULT NULL COMMENT '职称',
  `yisheng_photo` varchar(200) DEFAULT NULL COMMENT '医生头像',
  `yisheng_phone` varchar(200) DEFAULT NULL COMMENT '联系方式',
  `yisheng_guahao` varchar(200) DEFAULT NULL COMMENT '挂号须知',
  `yisheng_email` varchar(200) DEFAULT NULL COMMENT '邮箱',
  `yisheng_new_money` decimal(10,2) DEFAULT NULL COMMENT '挂号价格',
  `yisheng_content` text COMMENT '履历介绍',
  `create_time` timestamp NULL DEFAULT NULL COMMENT '创建时间 show1 show2 photoShow',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8 COMMENT='医生';

/*Data for the table `yisheng` */

insert  into `yisheng`(`id`,`yisheng_uuid_number`,`username`,`password`,`yisheng_name`,`yisheng_types`,`zhiwei_types`,`yisheng_zhichneg`,`yisheng_photo`,`yisheng_phone`,`yisheng_guahao`,`yisheng_email`,`yisheng_new_money`,`yisheng_content`,`create_time`) values (1,'164699245161612','a1','123456','医生名称1',1,3,'职称1','http://localhost:8080/yiyuanguahao/upload/yisheng1.jpg','17703786901','挂号须知1','1@qq.com','217.68','履历介绍1','2022-03-11 17:54:11'),(2,'16469924516166','a2','123456','医生名称2',3,3,'职称2','http://localhost:8080/yiyuanguahao/upload/yisheng2.jpg','17703786902','挂号须知2','2@qq.com','339.54','履历介绍2','2022-03-11 17:54:11'),(3,'16469924516172','a3','123456','医生名称3',2,3,'职称3','http://localhost:8080/yiyuanguahao/upload/yisheng3.jpg','17703786903','挂号须知3','3@qq.com','347.75','履历介绍3','2022-03-11 17:54:11');

/*Table structure for table `yonghu` */

DROP TABLE IF EXISTS `yonghu`;

CREATE TABLE `yonghu` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `username` varchar(200) DEFAULT NULL COMMENT '账户',
  `password` varchar(200) DEFAULT NULL COMMENT '密码',
  `yonghu_name` varchar(200) DEFAULT NULL COMMENT '用户姓名 Search111 ',
  `yonghu_photo` varchar(255) DEFAULT NULL COMMENT '头像',
  `yonghu_phone` varchar(200) DEFAULT NULL COMMENT '用户手机号',
  `yonghu_id_number` varchar(200) DEFAULT NULL COMMENT '用户身份证号 ',
  `yonghu_email` varchar(200) DEFAULT NULL COMMENT '邮箱',
  `sex_types` int(11) DEFAULT NULL COMMENT '性别 Search111 ',
  `new_money` decimal(10,2) DEFAULT NULL COMMENT '余额 ',
  `yonghu_delete` int(11) DEFAULT '1' COMMENT '假删',
  `create_time` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8 COMMENT='用户';

/*Data for the table `yonghu` */

insert  into `yonghu`(`id`,`username`,`password`,`yonghu_name`,`yonghu_photo`,`yonghu_phone`,`yonghu_id_number`,`yonghu_email`,`sex_types`,`new_money`,`yonghu_delete`,`create_time`) values (1,'a1','123456','用户姓名1','http://localhost:8080/yiyuanguahao/upload/yonghu1.jpg','17703786901','410224199610232001','1@qq.com',1,'9746.59',1,'2022-03-11 17:54:11'),(2,'a2','123456','用户姓名2','http://localhost:8080/yiyuanguahao/upload/yonghu2.jpg','17703786902','410224199610232002','2@qq.com',1,'755.85',1,'2022-03-11 17:54:11'),(3,'a3','123456','用户姓名3','http://localhost:8080/yiyuanguahao/upload/yonghu3.jpg','17703786903','410224199610232003','3@qq.com',1,'825.49',1,'2022-03-11 17:54:11');

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
