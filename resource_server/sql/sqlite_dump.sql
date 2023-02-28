PRAGMA foreign_keys=OFF;
BEGIN TRANSACTION;
CREATE TABLE `users` (
    `id` varchar(255) NOT NULL,
    `email` varchar(255) NOT NULL DEFAULT '',
    `salt` varchar(255) NOT NULL,
    `password` varchar(255) NOT NULL,
    PRIMARY KEY (`id`)
);
/* sample user: userid=user, password=a */
INSERT INTO users VALUES('user','user@example.com','eaghea','3128028ac86a592d3b48451cebc44d3c7227cd6b4160496bb84ad9d1437202d1');
COMMIT;
