package dao

import (
	"context"

	"github.com/ch3nnn/gen-table/example/dal/model"
	"github.com/ch3nnn/gen-table/example/dal/query"
	"gorm.io/gen"
)

var _ iCustomGenSiteFunc = (*customSiteDao)(nil)

type (
	// not edit interface name
	iCustomGenSiteFunc interface {
		iGenSiteFunc

		// custom func ....
		// ...

		FirstOrInit(whereFunc ...func(gen.Dao) gen.Dao) (*model.Site, error)
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

// FindCount rewrite
func (d *customSiteDao) FindCount(_ ...func(dao gen.Dao) gen.Dao) (int64, error) {
	if _, err := d.siteDao.FindCount(); err != nil {
		return 0, err
	}

	return -99, nil

}

// FirstOrInit custom
func (d *customSiteDao) FirstOrInit(whereFunc ...func(gen.Dao) gen.Dao) (*model.Site, error) {
	return d.siteDo.Scopes(whereFunc...).FirstOrInit()
}
