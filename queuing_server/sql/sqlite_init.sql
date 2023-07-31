PRAGMA foreign_keys=OFF;
BEGIN TRANSACTION;
DROP TABLE IF EXISTS `authorization_requests`;
CREATE TABLE `authorization_requests` (
    `id` INT NOT NULL,
    `client_domain` varchar(255) NOT NULL,
    `rs_domain` varchar(255) NOT NULL,
    `scopes` varchar(511) NOT NULL,
    `status` varchar(255) NOT NULL,
    `added_at` timestamp default CURRENT_TIMESTAMP NOT NULL,
    PRIMARY KEY (`id`)
);
INSERT INTO `authorization_requests` (`id`, `client_domain`, `rs_domain`, `scopes`, `status`) 
VALUES (1, "www.example.com", "rs.example.com", "read", "waiting") ;
COMMIT;