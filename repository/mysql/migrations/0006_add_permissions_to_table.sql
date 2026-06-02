-- +migrate Up
INSERT INTO `permissions` (`id`, `title`) VALUES (1, 'user-delete');
INSERT INTO `permissions` (`id`, `title`) VALUES (2, 'user-list');

-- +migrate Down
DELETE FROM `permissions` WHERE `id` IN (1,2);
