package dao

import (
	"context"

	"gen-table/example/dal/query"
)

var _ iCustomGenSiteFunc = (*customSiteDao)(nil)

type (
	// not edit interface name
	iCustomGenSiteFunc interface {
		iGenSiteFunc

		// custom func ....
		// ...
	}
	// not edit interface name
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
