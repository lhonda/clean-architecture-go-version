-- MySQL dump 10.13  Distrib 5.7.35, for Linux (x86_64)
--
-- Host: 172.17.0.2    Database: clean_arch
-- ------------------------------------------------------
-- Server version	8.0.26

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
-- Table structure for table `orders`
--


use mysql;
INSERT IGNORE INTO user (host, user, select_priv, insert_priv,update_priv, delete_priv) VALUES ('%', 'clean_architecture_go_version','Y','Y','Y','Y');
commit;
SET PASSWORD for 'clean_architecture_go_version'@'%' =PASSWORD('clean_architecture_go_version');
flush privileges;

create database if not exists clean_architecture_go_version;
use clean_architecture_go_version;

DROP TABLE IF EXISTS `orders`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `orders`
(
    `id`         varchar(32)  DEFAULT NULL,
    `owner`      varchar(255) DEFAULT NULL,
    `created_at` timestamp NULL DEFAULT NULL,
    UNIQUE KEY `order_id_uindex` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `orders`
--

LOCK
TABLES `orders` WRITE;
/*!40000 ALTER TABLE `orders` DISABLE KEYS */;
/*!40000 ALTER TABLE `orders` ENABLE KEYS */;
UNLOCK
TABLES;

--
-- Table structure for table `pizza`
--

DROP TABLE IF EXISTS `pizza`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `pizza`
(
    `id`          varchar(32)  DEFAULT NULL,
    `name`        varchar(255) DEFAULT NULL,
    `ingredients` varchar(255) DEFAULT NULL,
    `order_id`    varchar(32)  DEFAULT NULL,
    `created_at`  timestamp NULL DEFAULT NULL,
    UNIQUE KEY `pizza_id_uindex` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `pizza`
--

LOCK
TABLES `pizza` WRITE;
/*!40000 ALTER TABLE `pizza` DISABLE KEYS */;
/*!40000 ALTER TABLE `pizza` ENABLE KEYS */;
UNLOCK
TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

insert into pizza(id,name, ingredients, order_id, created_at) values ("4c7966aa33dd4e5a803846bd34d21e9f","queijo","queijo;molho","4c7966aa33dd4e5a803846bd34d21e9d",now());
insert into pizza(id,name, ingredients, order_id, created_at) values ("4c7966aa33dd3e5a803846bd34d21e9f","presunto","presunto;molho","4c7966aa33dd4e5a803846bd34d21e92",now());
GRANT ALL PRIVILEGES ON *.* TO '*'@'%' WITH GRANT OPTION;
-- Dump completed on 2021-08-11 16:44:00
