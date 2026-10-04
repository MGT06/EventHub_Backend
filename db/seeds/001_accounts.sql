INSERT INTO public.accounts VALUES
	(1, 'Admin Utama', 'admin@eventhub.id', 'Mengelola akun, event, dan community di EventHub.', 'Jakarta', 'Platform Admin', '$2a$hash_admin', 'admin', 'https://i.pravatar.cc/150?img=1', 'active', true, '2026-09-26 16:41:33.601703+00', NULL),
	(2, 'Rizky Pratama', 'rizky@eventhub.id', 'Engineering Manager, suka ngoprek Go.', 'Bandung', 'Engineering Manager', '$2a$hash_rizky', 'organizer', 'https://i.pravatar.cc/150?img=11', 'active', true, '2026-09-26 16:41:33.601703+00', NULL),
	(3, 'Kevin Santoso', 'kevin@eventhub.id', 'AI Engineer, aktif di komunitas AI Jakarta.', 'Jakarta', 'AI Engineer', '$2a$hash_kevin', 'organizer', 'https://i.pravatar.cc/150?img=68', 'active', true, '2026-09-26 16:41:33.601703+00', NULL),
	(4, 'Raisa Nurdiana', 'raisa@eventhub.id', 'Product manager, suka bikin workshop online.', 'Jakarta', 'Product Manager', '$2a$hash_raisa', 'organizer', 'https://i.pravatar.cc/150?img=44', 'suspended', true, '2026-09-26 16:41:33.601703+00', NULL),
	(5, 'Dian Purnama', 'dian@example.com', 'Mahasiswa, belajar backend Go.', 'Bandung', NULL, '$2a$hash_dian', 'attendee', 'https://i.pravatar.cc/150?img=25', 'active', true, '2026-09-26 16:41:33.601703+00', NULL),
	(6, 'Ahmad Fauzan', 'ahmad@example.com', 'Suka ikut workshop teknis akhir pekan.', 'Bandung', NULL, '$2a$hash_ahmad', 'attendee', 'https://i.pravatar.cc/150?img=13', 'active', true, '2026-09-26 16:41:33.601703+00', NULL),
	(7, 'Siti Rahayu', 'siti@example.com', NULL, 'Jakarta', NULL, '$2a$hash_siti', 'attendee', 'https://i.pravatar.cc/150?img=32', 'active', true, '2026-09-26 16:41:33.601703+00', NULL),
	(8, 'Indira Kusuma', 'indira@example.com', 'Data scientist, tertarik AI product design.', 'Jakarta', NULL, '$2a$hash_indira', 'attendee', 'https://i.pravatar.cc/150?img=45', 'active', true, '2026-09-26 16:41:33.601703+00', NULL),
	(9, 'Fikri Ramadhan', 'fikri@example.com', NULL, 'Surabaya', NULL, '$2a$hash_fikri', 'attendee', 'https://i.pravatar.cc/150?img=60', 'inactive', true, '2026-09-26 16:41:33.601703+00', NULL),
	(10, 'Bella Anjani', 'bella@example.com', 'Musisi dan penggemar produksi musik.', 'Surabaya', NULL, '$2a$hash_bella', 'attendee', 'https://i.pravatar.cc/150?img=36', 'active', true, '2026-09-26 16:41:33.601703+00', NULL),
	(12, 'admin', 'admin@gmail.com', NULL, NULL, NULL, '$argon2id$v=19$m=65536,t=2,p=2$vXL2bYUUt+WdvnT6crHNpQ$hFGTWP/n/1rvZ7RYF7krlKsApFtca+F4z6CTaYc+xho', 'admin', NULL, 'active', NULL, '2026-09-27 21:40:00.409891+00', NULL),
	(13, 'organizer', 'organizer@gmail.com', NULL, NULL, NULL, '$argon2id$v=19$m=65536,t=2,p=2$XXMS+efQd4G/2Qzin12AMA$boK+ktWhsAzEU3T87+a/HNARxAKr6XU9fzPi6TmehwE', 'organizer', NULL, 'active', NULL, '2026-09-27 21:45:38.664232+00', NULL),
	(11, 'givta', 'givta123@gmail.com', 'lotnok', 'Bekasi', 'CEO', '$argon2id$v=19$m=65536,t=2,p=2$pZpMpJs6TEIDvigGEGbMBA$9G8GeQv19VjYRRjhVgP5gSVsrpEwbOpOr0OgfVb6QZQ', 'attendee', NULL, 'active', NULL, '2026-09-26 17:01:19.926874+00', NULL),
	(15, '123alfan', 'alfan@gmail.com', 'pecinta wasawho', 'purwo', 'CEO', '$argon2id$v=19$m=65536,t=2,p=2$odGXFBPs2eqOdEjJmdGTLA$rprAvhHTE9kpgZ7emIsTCugv8WI1dm+337NhcwVS4CY', 'attendee', 'public/img/persons/1791101385088716400_alfan123.png', 'active', NULL, '2026-09-28 03:55:42.219979+00', NULL);

SELECT pg_catalog.setval('public.accounts_id_seq', 15, true);
