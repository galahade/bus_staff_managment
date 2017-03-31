create database bus_system;
use bus_system;

DROP TABLE IF EXISTS `charge_record`;

#Create table charge_record
CREATE TABLE `charge_record` (
  `id` char(36) NOT NULL DEFAULT '',
  `bus_id` char(36) NOT NULL DEFAULT '',
  `record_date` date NOT NULL,
  `record_staff_id` char(36) NOT NULL DEFAULT '',
  `mileage` int(11) DEFAULT NULL,
  `charged_TWH` decimal(10,0) DEFAULT NULL,
  `remain_percent` decimal(10,0) DEFAULT NULL,
  `final_percent` decimal(10,0) DEFAULT '100',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


#Create table bus_basic
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
  KEY `bus_brand_fk_idx` (`brand_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

#Create table bus_brand
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
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

#Create table income
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

#Create table operation_data
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

#Create table salaries
DROP TABLE IF EXISTS `salaries`;
CREATE TABLE `salaries` (
  `id` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci COMMENT='工资表';


#Create table salary_element
DROP TABLE IF EXISTS `salary_element`;
CREATE TABLE `salary_element` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `element_name` varchar(200) COLLATE utf8_unicode_ci NOT NULL,
  `base_amount` decimal(9,2) NOT NULL,
  `is_active` tinyint(4) NOT NULL,
  `job_type` bit(8) NOT NULL DEFAULT b'11111111' COMMENT '11111111-all , 00000001-司机, 00000010-维修, 00000100-技术, 00001000-保障, 10000000-管理',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci COMMENT='员工工资构成';

#Create table staff
DROP TABLE IF EXISTS `staff`;
CREATE TABLE `staff` (
  `id` char(36) COLLATE utf8_unicode_ci NOT NULL,
  `name` varchar(100) COLLATE utf8_unicode_ci NOT NULL,
  `staff_id` varchar(10) COLLATE utf8_unicode_ci DEFAULT NULL,
  `job_type` bit(8) NOT NULL COMMENT ' 00000001-司机, 00000010-维修, 00000100-技术, 00001000-保障, 10000000-管理',
  `onboard_time` date DEFAULT NULL,
  `personal_id` varchar(20) COLLATE utf8_unicode_ci DEFAULT NULL,
  `driver_type` varchar(45) COLLATE utf8_unicode_ci DEFAULT NULL,
  `is_internship` tinyint(4) DEFAULT '0',
  `is_multitime_hired` tinyint(4) DEFAULT '0',
  `first_onboard_time` date DEFAULT NULL,
  `phone` varchar(100) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT 'if there are multiple phone numbers, use comma slip them.',
  `department` varchar(100) COLLATE utf8_unicode_ci DEFAULT NULL,
  `emergency_contact` varchar(100) COLLATE utf8_unicode_ci DEFAULT NULL,
  `emergency_contact_phone` varchar(45) COLLATE utf8_unicode_ci DEFAULT NULL,
  `emergency_contact_relation` varchar(45) COLLATE utf8_unicode_ci DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `personal_id_UNIQUE` (`personal_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci COMMENT='员工表';