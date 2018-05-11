# magnet-go
magnet go vetsion
 
Magnet url can be  download by special software ,exp thunder uTorrent

more pople downloading ,more time you will save

let's get some picture about our project .



#database
table search_filelist
```sql
CREATE TABLE `search_filelist` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `info_hash` varchar(40) CHARACTER SET utf8 NOT NULL,
  `file_list` longtext CHARACTER SET utf8 NOT NULL,
  PRIMARY KEY (`info_hash`,`id`),
  UNIQUE KEY `id` (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=1331735 DEFAULT CHARSET=ascii;

```

table search_filelist
```sql
CREATE TABLE `search_hash` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `info_hash` varchar(40) NOT NULL,
  `category` varchar(20) NOT NULL,
  `data_hash` varchar(32) NOT NULL,
  `name` varchar(255) NOT NULL,
  `extension` varchar(20) NOT NULL,
  `classified` tinyint(1) NOT NULL,
  `source_ip` varchar(20) DEFAULT NULL,
  `tagged` tinyint(1) NOT NULL,
  `length` bigint(20) NOT NULL,
  `create_time` datetime NOT NULL,
  `last_seen` datetime NOT NULL,
  `requests` int(10) unsigned NOT NULL,
  `comment` varchar(255) DEFAULT NULL,
  `creator` varchar(20) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `info_hash` (`info_hash`),
  KEY `search_hash_tagged_50480647a28d03e1_uniq` (`tagged`)
) ENGINE=MyISAM AUTO_INCREMENT=1390673 DEFAULT CHARSET=utf8;

```

## Sphinx(coreseek) config
```
#
# Minimal Sphinx configuration sample (clean, simple, functional)
#

source main
{
        type                    = mysql

        sql_host                = 127.0.0.1
        sql_user                = root
        sql_pass                = #root
        sql_db                  = ssbc
        sql_port                = 3306  # optional, default is 3306

        sql_query_pre       = SET NAMES utf8
        sql_query = SELECT search_filelist.id AS id,search_hash.id AS hash_id, search_hash.name as name, CRC32(search_hash.category) AS category,\
                search_hash.length, UNIX_TIMESTAMP(search_hash.create_time) AS create_time, \
                UNIX_TIMESTAMP(last_seen) AS last_seen FROM search_hash,search_filelist where search_hash.info_hash=search_filelist.info_hash

        sql_attr_bigint         = length
#       sql_attr_timestamp      = info_hash
        sql_attr_timestamp      = create_time
        sql_attr_timestamp      = last_seen
        sql_attr_uint   = category
        sql_attr_uint   = hash_id
}


index main
{
    source                      = main
    path                        = /data/bt/index/db/main
    charset_type        = zh_cn.utf-8
    charset_dictpath = /usr/local/mmseg3/etc/

    docinfo            = extern
    mlock            = 0
    morphology        = none
    min_word_len        = 1
    html_strip                = 0
}


#index rt_main
#{
#       type                    = rt
#       #rt_mem_limit           = 512M

#       path                    = /data/bt/index/db/rt_main

#       rt_field                = name
#       rt_attr_bigint          = length
#       rt_attr_timestamp       = create_time
#       rt_attr_timestamp       = last_seen
#       rt_attr_uint    = category

#   ngram_len = 1
#    ngram_chars = U+3000..U+2FA1F
#}


indexer
{
        mem_limit               = 150M
}


searchd
{
        listen                  = 9312
#       listen                  = 9306:mysql41
        log                     = /data/bt/index/searchd.log
        query_log               = /data/bt/index/query.log
        read_timeout            = 5
        max_children            = 30
        max_matches            = 1000
        pid_file                = /data/bt/index/searchd.pid
        seamless_rotate         = 1
        preopen_indexes         = 1
        unlink_old              = 1
        #workers                        = threads # for RT to work
#       binlog_path             = /data/bt/index/binlog/
```

### home page
![](https://raw.githubusercontent.com/rhettli/magnet-php/master/pic/1.png)

### search page
![](https://raw.githubusercontent.com/rhettli/magnet-php/master/pic/2.png)

### search page
![](https://raw.githubusercontent.com/rhettli/magnet-php/master/pic/3.png)

### zbout page
![](https://raw.githubusercontent.com/rhettli/magnet-php/master/pic/4.png)

### donate
![](https://raw.githubusercontent.com/rhettli/magnet-php/master/pic/5.png)

### detail
![](https://raw.githubusercontent.com/rhettli/magnet-php/master/pic/6.png)

 
