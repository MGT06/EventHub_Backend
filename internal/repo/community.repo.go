package repo

import (
	"context"

	"github.com/MGT06/EventHub_Backend.git/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CommunityRepo struct {
	db *pgxpool.Pool
}

func NewCommunityRepo(db *pgxpool.Pool) *CommunityRepo {
	return &CommunityRepo{
		db: db,
	}
}

func (c *CommunityRepo) GetAllCommunity(ctx context.Context) ([]model.CommunityDetail, error) {
	query := `SELECT c.id, c.community_name, c.description, c.image_community_url, STRING_AGG(cg.category_name, ', ') AS "category" 
FROM communities c 
JOIN community_categories cc ON c.id = cc.community_id
JOIN categories cg ON cg.id = cc.category_id
GROUP BY  c.id, c.community_name, c.description, c.image_community_url`

	res, err := c.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	var communities []model.CommunityDetail

	for res.Next() {
		var community model.CommunityDetail
		if err := res.Scan(&community.Id_community, &community.Community_name, &community.Description, &community.Image_community_url, &community.Category_name); err != nil {
			return nil, err
		}
		communities = append(communities, community)
	}

	if res.Err() != nil {
		return nil, res.Err()
	}

	return communities, nil
}
