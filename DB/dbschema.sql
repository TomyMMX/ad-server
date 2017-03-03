CREATE DATABASE `addb`;

CREATE TABLE `folder` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `parentid` int(11) NOT NULL DEFAULT '0',
  `name` varchar(255) NOT NULL,
  `lastmodified` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniqueNameInFolder` (`parentid`,`name`),
  KEY `parentIdx` (`parentid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `ad` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `folderid` int(11) NOT NULL,
  `name` varchar(255) NOT NULL,
  `url` varchar(512) CHARACTER SET ascii NOT NULL,
  `lastmodified` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniqueNameInFolder` (`folderid`,`name`),
  KEY `folderId_idx` (`folderid`),
  CONSTRAINT `folderId` FOREIGN KEY (`folderid`) REFERENCES `folder` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


