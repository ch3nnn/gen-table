package dao

import (
	"context"
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
