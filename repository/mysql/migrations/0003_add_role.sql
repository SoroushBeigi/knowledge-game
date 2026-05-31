-- +migrate Up
ALTER TABLE `users` ADD COLUMN `role` ENUM('user','admin') DEFAULT 'user';

-- +migrate Down
ALTER TABLE `users` DROP COLUMN `role`;
