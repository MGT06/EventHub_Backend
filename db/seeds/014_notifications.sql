INSERT INTO public.notifications VALUES
	(1, 5, 'Registrasi berhasil', 'Kamu terdaftar di Go Concurrency Workshop.', 'registration', NULL, '2026-09-26 13:41:33.601703+00', NULL),
	(2, 5, 'Event dimulai 2 hari lagi', 'Go Concurrency Workshop dimulai 09:00 di Bandung.', 'event_reminder', '2026-09-26 15:41:33.601703+00', '2026-09-26 14:41:33.601703+00', NULL),
	(3, 6, 'Balasan diskusi', 'Rizky membalas pertanyaanmu di Go Concurrency Workshop.', 'discussion_reply', NULL, '2026-09-26 15:41:33.601703+00', NULL),
	(4, 8, 'Registrasi berhasil', 'Kamu terdaftar di AI Product Design Summit.', 'registration', '2026-09-26 14:41:33.601703+00', '2026-09-26 13:41:33.601703+00', NULL),
	(5, 8, 'Anggota baru komunitas', 'Jakarta AI & ML Club baru saja mencapai 2.000 member.', 'community', NULL, '2026-09-22 16:41:33.601703+00', NULL),
	(6, 9, 'Update jadwal event', 'Jadwal Startup Pitch Night diperbarui.', 'event_updated', NULL, '2026-09-24 16:41:33.601703+00', NULL),
	(7, 10, 'Event minggu depan', 'Product Management Masterclass dimulai pukul 10:00.', 'event_reminder', NULL, '2026-09-21 16:41:33.601703+00', NULL);

SELECT pg_catalog.setval('public.notifications_id_seq', 7, true);
