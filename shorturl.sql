CREATE TABLE `short_url` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `long_url` text,
  `short_url` text,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;
