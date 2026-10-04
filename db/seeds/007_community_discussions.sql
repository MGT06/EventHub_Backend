INSERT INTO public.community_discussions VALUES
	(1, 1, 2, NULL, 'Welcome semuanya! Senang banyak member baru bulan ini.', '2026-09-26 16:41:33.601703+00', NULL),
	(2, 1, 6, NULL, 'Ada yang udah nyoba fitur generics di Go 1.22? Mau bahas di meetup berikutnya.', '2026-09-26 16:41:33.601703+00', NULL),
	(3, 1, 5, 2, 'Belum coba, tapi penasaran juga!', '2026-09-26 16:41:33.601703+00', NULL),
	(4, 2, 3, NULL, 'Lagi cari pembicara buat meetup AI product strategy bulan depan, DM kalau tertarik.', '2026-09-26 16:41:33.601703+00', NULL);

SELECT pg_catalog.setval('public.community_discussions_id_seq', 4, true);
