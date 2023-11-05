-- MySQL dump 10.13  Distrib 5.6.27, for Linux (x86_64)
--
-- Host: localhost    Database: thsdb
-- ------------------------------------------------------
-- Server version	5.6.27

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
-- Table structure for table `services`
--
CREATE DATABASE thsdb;
USE thsdb;

DROP TABLE IF EXISTS `services`;
CREATE TABLE `services` (
  `service_id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `service_name` varchar(100) NOT NULL,
  `security_level` int(11) NOT NULL,
  `service_type` varchar(100) NOT NULL,
  PRIMARY KEY (`service_id`)
) ENGINE=InnoDB AUTO_INCREMENT=0005 DEFAULT CHARSET=latin1;

LOCK TABLES `services` WRITE;
/*!40000 ALTER TABLE `services` DISABLE KEYS */;
INSERT INTO `services` VALUES 
('0001', 'Proxmox', 'proxmox', 5, 'HyperVisor'),
('0002', 'Voyager Docker Host', 'voyager-docker-host', 5, 'ContainerHost'),
('0003', 'Enterprise Docker Host', 'enterprise-docker-host', 5, 'ContainerHost'),
('0004', 'UniFi Dream Machine Pro', 'unifi-dream-machine-pro', 5, 'Networking');


UNLOCK TABLES;
