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

func NewSiteDao() ISiteDao {
	return &customSiteDao{
		siteDao{
			siteDo: query.Site.WithContext(context.Background()),
		},
	}
}
