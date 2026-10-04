INSERT INTO public.location_event VALUES
	(1, 'Bandung', '2026-09-26 16:41:33.601703+00'),
	(2, 'Jakarta', '2026-09-26 16:41:33.601703+00'),
	(3, 'Surabaya', '2026-09-26 16:41:33.601703+00'),
	(4, 'Yogyakarta', '2026-09-26 16:41:33.601703+00'),
	(5, 'Online', '2026-09-26 16:41:33.601703+00');


SELECT pg_catalog.setval('public.location_event_id_seq', 5, true);
