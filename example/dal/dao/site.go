package dao

import (
	"context"

	"gen-table/example/dal/query"
)

type (
	ISiteDao interface {
		iSiteDao
	}

	customSiteDao struct {
		siteDao
	}
)

func NewSiteDao(ctx context.Context) ISiteDao {
	return &customSiteDao{
		siteDao{
			siteDo: query.Site.WithContext(ctx),
		},
	}
}
