package model

type DashboardOrganizer struct {
	TotalEvent    int `db:"total_event_created"`
	TotalAttendes int `db:"total_attendees_joined"`
	AVGFillRate   int `db:"avg_fill_rate"`
}

type DashboardAdmin struct {
	TotalUsers       int `db:"total_users"`
	TotalEvents      int `db:"total_events"`
	TotalCommunities int `db:"total_communities"`
}
