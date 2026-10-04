INSERT INTO public.event_discussions VALUES
	(1, 1, 5, NULL, 'Super excited buat ini — bakal ada live coding session enggak?', '2026-09-26 16:41:33.601703+00', NULL),
	(2, 1, 6, 1, 'Ada! Sesi sore full hands-on. Bawa laptop dengan Go 1.22+ ya.', '2026-09-26 16:41:33.601703+00', NULL),
	(3, 1, 7, NULL, 'Ada parkir di dekat lokasi enggak ya? Saya dari luar kota.', '2026-09-26 16:41:33.601703+00', NULL),
	(4, 1, 2, 3, 'Ada, parkiran gedung sebelah cukup luas.', '2026-09-26 16:41:33.601703+00', NULL),
	(5, 2, 8, NULL, 'Slide-nya bakal dibagikan setelah acara enggak?', '2026-09-26 16:41:33.601703+00', NULL);

SELECT pg_catalog.setval('public.event_discussions_id_seq', 5, true);
