-- +migrate Up
INSERT INTO `permissions` (`id`, `title`) VALUES (1, 'user-delete');
INSERT INTO `permissions` (`id`, `title`) VALUES (2, 'user-list');

INSERT INTO `access_controls` (`actor_type`, `actor_id`,`permission_id`) VALUES ('role',2,1);
INSERT INTO `access_controls` (`actor_type`, `actor_id`,`permission_id`) VALUES ('role',2,2);

-- +migrate Down
DELETE FROM `permissions` WHERE `id` IN (1,2);
DELETE FROM `access_controls` WHERE `id` IN (1,2);
