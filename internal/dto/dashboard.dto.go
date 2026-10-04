package dto

type DashboardOrganizer struct {
	TotalEvent    int	`json:"totalEvent"`
	TotalAttendes int	`json:"totalAttendes"`
	AVGFillRate   int	`json:"AVGFillRate"`
}

type DashboardAdmin struct {
	TotalUsers       int	`json:"totalUsers"`
	TotalEvents      int	`json:"totalEvents"`
	TotalCommunities int	`json:"totalCommunities"`
}