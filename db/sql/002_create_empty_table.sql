-- SQL SchemaをExportするための空のテーブルを作成
CREATE DATABASE IF NOT EXISTS `em`;

GRANT ALL ON em.* TO `docker` @`%`;
