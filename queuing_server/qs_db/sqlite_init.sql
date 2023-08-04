PRAGMA foreign_keys=OFF;
BEGIN TRANSACTION;
DROP TABLE IF EXISTS `authorization_requests`;
DROP TABLE IF EXISTS `users`;
DROP TABLE IF EXISTS `qs_ids`;
DROP TABLE IF EXISTS `tickets`;
CREATE TABLE `authorization_requests` (
    `request_id` INTEGER PRIMARY KEY AUTOINCREMENT,
    `ticket` varchar(256) NOT NULL,
    `requested_scopes` varchar(8192) NOT NULL,
    `client_request` varchar(8192) NOT NULL default "",
    `user_id` INTEGER NOT NULL,
    `status` varchar(15) NOT NULL,
    `created_at` timestamp default CURRENT_TIMESTAMP NOT NULL
);
INSERT INTO `authorization_requests` (`user_id`, `ticket`, `status`, `requested_scopes`) 
VALUES (1, "this_is_not_meaningful", "waiting", "this.is.sample");


CREATE TABLE `users` (
    `user_id` INTEGER PRIMARY KEY,
    `username` varchar(256) NOT NULL
    -- TODO: `passphrase` varchar(256) NOT NULL,
);
INSERT INTO `users` VALUES (1, "a");
CREATE TABLE `qs_ids` (
    `qs_id`  varchar(256) PRIMARY KEY,
    `user_id` INTEGER NOT NULL
);
INSERT INTO `qs_ids` VALUES ("sample_qs_id", 1);
CREATE TABLE `tickets` (
    `ticket_id` varchar(256) PRIMARY KEY,
    `qs_id` varchar(256) NOT NULL,
    `user_id` INTEGER NOT NULL
);

COMMIT;