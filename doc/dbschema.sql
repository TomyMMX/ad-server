CREATE DATABASE `addb` /*!40100 DEFAULT CHARACTER SET utf8 */;

CREATE TABLE `folder` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `parrentid` int(11) NOT NULL DEFAULT '0',
  `name` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


CREATE TABLE `ad` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `folderid` int(11) NOT NULL,
  `name` varchar(255) NOT NULL,
  `url` varchar(512) CHARACTER SET ascii NOT NULL,
  `lastmodified` datetime NOT NULL,
  PRIMARY KEY (`id`),
  KEY `folderId_idx` (`folderid`),
  CONSTRAINT `folderId` FOREIGN KEY (`folderid`) REFERENCES `folder` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
