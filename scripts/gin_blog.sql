drop database if exists `gin_blog`;
create database if not exists `gin_blog` default character set utf8mb4 default collate utf8mb4_0900_ai_ci;
use `gin_blog`;

CREATE TABLE `blog_tag` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `name` varchar(100) DEFAULT '' COMMENT '標籤名稱',
    `created_on` int(10) unsigned DEFAULT '0' COMMENT '創建時間',
    `created_by` varchar(100) DEFAULT '' COMMENT '創建人',
    `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改時間',
    `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
    `deleted_on` int(10) unsigned DEFAULT '0' COMMENT '刪除時間',
    `is_del` tinyint(3) unsigned DEFAULT '0' COMMENT '是否刪除 0為未刪除、1為已刪除',
    `state` tinyint(3) unsigned DEFAULT '1' COMMENT '狀態 0為禁用、1為啟用',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = '標籤管理';

CREATE TABLE `blog_article` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `title` varchar(100) DEFAULT '' COMMENT '文章標題',
    `desc` varchar(255) DEFAULT '' COMMENT '文章簡述',
    `cover_image_url` varchar(255) DEFAULT '' COMMENT '封面圖片地址',
    `content` longtext COMMENT '文章內容',
    `created_on` int(10) unsigned DEFAULT '0' COMMENT '新建時間',
    `created_by` varchar(100) DEFAULT '' COMMENT '創建人',
    `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改時間',
    `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
    `deleted_on` int(10) unsigned DEFAULT '0' COMMENT '刪除時間',
    `is_del` tinyint(3) unsigned DEFAULT '0' COMMENT '是否刪除 0為未刪除、1為已刪除',
    `state` tinyint(3) unsigned DEFAULT '1' COMMENT '狀態 0為禁用、1為啟用',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = '文章管理';

CREATE TABLE `blog_article_tag` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `article_id` int(10) unsigned NOT NULL COMMENT '文章ID',
    `tag_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '標籤ID',
    `created_on` int(10) unsigned DEFAULT '0' COMMENT '創建時間',
    `created_by` varchar(100) DEFAULT '' COMMENT '創建人',
    `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改時間',
    `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
    `deleted_on` int(10) unsigned DEFAULT '0' COMMENT '刪除時間',
    `is_del` tinyint(3) unsigned DEFAULT '0' COMMENT '是否刪除 0為未刪除、1為已刪除',
    PRIMARY KEY (`id`, `article_id`, `tag_id`),
    FOREIGN KEY (`article_id`) REFERENCES `blog_article` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION,
    FOREIGN KEY (`tag_id`) REFERENCES `blog_tag` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = '文章標籤關聯';