-- SQL SchemaをExportするための空のテーブルを作成
CREATE DATABASE IF NOT EXISTS `empty`;

GRANT ALL ON empty.* TO `docker` @`%`;
