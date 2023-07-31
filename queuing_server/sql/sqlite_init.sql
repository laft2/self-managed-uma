PRAGMA foreign_keys=OFF;
BEGIN TRANSACTION;
DROP TABLE IF EXISTS `authorization_requests`;
CREATE TABLE `authorization_requests` (
    `id` INT NOT NULL,
    `grant_type` varchar(255) NOT NULL,
    `ticket` varchar(255) NOT NULL,
    `claim_token` varchar(255),
    `claim_token_format` varchar(255),
    `pct` varchar(255),
    `rpt` varchar(255),
    `scopes` varchar(511),
    `status` varchar(255) NOT NULL,
    `created_at` timestamp default CURRENT_TIMESTAMP NOT NULL,
    PRIMARY KEY (`id`)
);
INSERT INTO `authorization_requests` (`id`, `grant_type`, `ticket`, `scopes`, `status`) 
VALUES (1, "urn:this_is_sample", "this_is_not_meaningful","read", "waiting") ;
COMMIT;