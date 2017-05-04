create database bus_system;
use bus_system;

# ************************************************************
# Sequel Pro SQL dump
# Version 4541
#
# http://www.sequelpro.com/
# https://github.com/sequelpro/sequelpro
#
# Host: 127.0.0.1 (MySQL 5.7.17)
# Database: bus_system
# Generation Time: 2017-05-04 06:37:45 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table bus_basic
# ------------------------------------------------------------

DROP TABLE IF EXISTS `bus_basic`;

CREATE TABLE `bus_basic` (
  `id` char(36) COLLATE utf8_unicode_ci NOT NULL,
  `bus_license` varchar(45) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `custom_id` varchar(6) COLLATE utf8_unicode_ci DEFAULT NULL,
  `brand_id` char(36) COLLATE utf8_unicode_ci DEFAULT NULL,
  `register_date` date DEFAULT NULL,
  `VIN` varchar(30) COLLATE utf8_unicode_ci DEFAULT NULL,
  `engine_no` varchar(30) COLLATE utf8_unicode_ci DEFAULT NULL,
  `persons_capacity` int(11) DEFAULT '0',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `bus_license_UNIQUE` (`bus_license`),
  UNIQUE KEY `uc_engine_no` (`engine_no`),
  KEY `bus_brand_fk_idx` (`brand_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;



# Dump of table bus_brand
# ------------------------------------------------------------

DROP TABLE IF EXISTS `bus_brand`;

CREATE TABLE `bus_brand` (
  `id` char(36) NOT NULL DEFAULT '',
  `name` varchar(100) NOT NULL DEFAULT '',
  `model` varchar(200) DEFAULT NULL,
  `alias` varchar(50) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table charge_record
# ------------------------------------------------------------

DROP TABLE IF EXISTS `charge_record`;

CREATE TABLE `charge_record` (
  `id` char(36) NOT NULL DEFAULT '',
  `bus_id` char(36) NOT NULL DEFAULT '',
  `record_date` date NOT NULL,
  `record_staff_id` char(36) NOT NULL DEFAULT '',
  `mileage` int(11) DEFAULT NULL,
  `charged_TWH` decimal(10,0) DEFAULT NULL,
  `remain_percent` decimal(10,0) DEFAULT NULL,
  `final_percent` decimal(10,0) DEFAULT '100',
  `charged_time` int(11) NOT NULL DEFAULT '1',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `MULT_UNKEY` (`bus_id`,`record_date`,`charged_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table dictionary
# ------------------------------------------------------------

DROP TABLE IF EXISTS `dictionary`;

CREATE TABLE `dictionary` (
  `id` char(36) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `name` varchar(50) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `type` int(11) NOT NULL COMMENT '1:job type; 2:department; 3:driver type',
  `is_active` tinyint(4) NOT NULL DEFAULT '1',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;



# Dump of table income
# ------------------------------------------------------------

DROP TABLE IF EXISTS `income`;

CREATE TABLE `income` (
  `id` char(36) COLLATE utf8_unicode_ci NOT NULL,
  `income_type` varchar(45) COLLATE utf8_unicode_ci NOT NULL,
  `amount` decimal(10,2) NOT NULL,
  `record_date` date DEFAULT NULL,
  `time_unit` tinyint(2) DEFAULT NULL,
  `year` int(11) DEFAULT NULL,
  `month` int(11) DEFAULT NULL,
  `day` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;



# Dump of table line_fare_income
# ------------------------------------------------------------

DROP TABLE IF EXISTS `line_fare_income`;

CREATE TABLE `line_fare_income` (
  `id` char(36) NOT NULL,
  `line_no` int(11) DEFAULT NULL,
  `carrying_amount` decimal(10,2) NOT NULL,
  `actual_amount` decimal(10,2) DEFAULT NULL,
  `worn_coin_amount` decimal(10,2) DEFAULT NULL,
  `bus_numbers` int(11) DEFAULT NULL,
  `counting_date` date NOT NULL,
  `counting_staff1_id` char(36) NOT NULL DEFAULT '',
  `counting_staff2_id` char(36) DEFAULT NULL,
  `record_staff_id` char(36) DEFAULT NULL,
  `original_evidence` blob,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table mileage_record
# ------------------------------------------------------------

DROP TABLE IF EXISTS `mileage_record`;

CREATE TABLE `mileage_record` (
  `id` char(36) NOT NULL DEFAULT '',
  `bus_id` char(36) NOT NULL DEFAULT '',
  `record_date` date NOT NULL,
  `record_staff_id` char(36) NOT NULL DEFAULT '',
  `mileage` int(11) DEFAULT NULL,
  `daily_mileage` int(11) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `record_date` (`record_date`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table operation_data
# ------------------------------------------------------------

DROP TABLE IF EXISTS `operation_data`;

CREATE TABLE `operation_data` (
  `id` char(36) COLLATE utf8_unicode_ci NOT NULL,
  `operation_date` date NOT NULL,
  `bus_id` char(36) COLLATE utf8_unicode_ci NOT NULL,
  `driver` char(36) COLLATE utf8_unicode_ci NOT NULL,
  `recorder` char(36) COLLATE utf8_unicode_ci DEFAULT NULL,
  `total_km` int(11) DEFAULT NULL,
  `charge_quantity` double DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;



# Dump of table salaries
# ------------------------------------------------------------

DROP TABLE IF EXISTS `salaries`;

CREATE TABLE `salaries` (
  `id` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci COMMENT='工资表';



# Dump of table salary_element
# ------------------------------------------------------------

DROP TABLE IF EXISTS `salary_element`;

CREATE TABLE `salary_element` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `element_name` varchar(200) COLLATE utf8_unicode_ci NOT NULL,
  `base_amount` decimal(9,2) NOT NULL,
  `is_active` tinyint(4) NOT NULL,
  `job_type` bit(8) NOT NULL DEFAULT b'11111111' COMMENT '11111111-all , 00000001-司机, 00000010-维修, 00000100-技术, 00001000-保障, 10000000-管理',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci COMMENT='员工工资构成';



# Dump of table staff
# ------------------------------------------------------------

DROP TABLE IF EXISTS `staff`;

CREATE TABLE `staff` (
  `id` char(36) COLLATE utf8_unicode_ci NOT NULL,
  `name` varchar(100) COLLATE utf8_unicode_ci NOT NULL,
  `job_type_id` char(36) COLLATE utf8_unicode_ci NOT NULL DEFAULT '' COMMENT ' 00000001-司机, 00000010-维修, 00000100-技术, 00001000-保障, 10000000-管理',
  `onboard_time` date DEFAULT NULL,
  `personal_id` varchar(20) COLLATE utf8_unicode_ci DEFAULT NULL,
  `driver_type_id` char(36) COLLATE utf8_unicode_ci DEFAULT NULL,
  `is_internship` tinyint(4) DEFAULT '0',
  `is_multitime_hired` tinyint(4) DEFAULT '0',
  `first_onboard_time` date DEFAULT NULL,
  `phone` varchar(100) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT 'if there are multiple phone numbers, use comma slip them.',
  `department_id` char(36) COLLATE utf8_unicode_ci DEFAULT NULL,
  `emergency_contact` varchar(100) COLLATE utf8_unicode_ci DEFAULT NULL,
  `emergency_contact_phone` varchar(45) COLLATE utf8_unicode_ci DEFAULT NULL,
  `emergency_contact_relation` varchar(45) COLLATE utf8_unicode_ci DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `is_resign` tinyint(4) unsigned DEFAULT '0',
  `resign_date` date DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `personal_id_UNIQUE` (`personal_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci COMMENT='员工表';



# Dump of table staff_resign
# ------------------------------------------------------------

DROP TABLE IF EXISTS `staff_resign`;

CREATE TABLE `staff_resign` (
  `id` char(36) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `staff_id` char(36) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `resign_date` date NOT NULL,
  `resign_reason` varchar(1000) COLLATE utf8_unicode_ci DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;


