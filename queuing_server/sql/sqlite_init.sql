PRAGMA foreign_keys=OFF;
BEGIN TRANSACTION;
CREATE TABLE `authorization_requests` (
    `id` INT NOT NULL,
    `client_domain` varchar(255) NOT NULL,
    `rs_domain` varchar(255) NOT NULL,
    `scopes` varchar(511) NOT NULL,
    `state` varchar(255) NOT NULL,
    PRIMARY KEY (`id`)
);
COMMIT;