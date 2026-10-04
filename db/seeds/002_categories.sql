INSERT INTO public.categories VALUES
	(1, 'Technology', '2026-09-26 16:41:33.601703+00'),
	(2, 'Programming', '2026-09-26 16:41:33.601703+00'),
	(3, 'Design', '2026-09-26 16:41:33.601703+00'),
	(4, 'Business', '2026-09-26 16:41:33.601703+00'),
	(5, 'Career', '2026-09-26 16:41:33.601703+00'),
	(6, 'Music', '2026-09-26 16:41:33.601703+00');


SELECT pg_catalog.setval('public.categories_id_seq', 6, true);
