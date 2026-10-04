INSERT INTO public.communities VALUES
	(1, 'Bandung Go Community', 'Komunitas Go di Bandung — meetup mingguan dan workshop.', 'https://images.unsplash.com/photo-1607799279861-4dd421887fb3', true, '2026-09-26 16:41:33.601703+00', NULL),
	(2, 'Jakarta AI & ML Club', 'Peneliti dan praktisi ML/LLM di Jakarta.', 'https://images.unsplash.com/photo-1526374965328-7f61d4dc18c5', true, '2026-09-26 16:41:33.601703+00', NULL),
	(3, 'Indonesia Frontend Devs', 'Komunitas frontend terbesar di Indonesia.', 'https://images.unsplash.com/photo-1555066931-4365d14bab8c', true, '2026-09-26 16:41:33.601703+00', NULL),
	(4, 'Music Tech Community', 'Produser dan sound engineer yang suka teknologi musik.', 'https://images.unsplash.com/photo-1531651008558-ed1740375b39', true, '2026-09-26 16:41:33.601703+00', NULL);


SELECT pg_catalog.setval('public.communities_id_seq', 4, true);
