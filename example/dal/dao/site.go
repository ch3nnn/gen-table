package dao

import (
	"context"

	"gen-table/example/dal/query"
)

type (
	ISite interface {
		iSite
	}

	SiteDao struct {
		siteDao
	}
)

func NewSiteDao(ctx context.Context) SiteDao {
	return SiteDao{
		siteDao{
			siteDo: query.Site.WithContext(ctx),
		},
	}
}
