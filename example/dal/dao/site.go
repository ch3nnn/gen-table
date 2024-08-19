package dao

import (
	"context"
	"time"

	"github.com/ch3nnn/gen-table/example/dal/model"
	"github.com/ch3nnn/gen-table/example/dal/query"
	"gorm.io/gen"
)

var _ iCustomGenSiteFunc = (*customSiteDao)(nil)

type (
	// ISiteDao not edit interface name
	ISiteDao interface {
		iWhereSiteFunc
		WithContext(ctx context.Context) iCustomGenSiteFunc

		// TODO Custom WhereFunc ....
		// ...

		WhereBetweenByCreatedAt(left time.Time, right time.Time) func(dao gen.Dao) gen.Dao
	}

	// not edit interface name
	iCustomGenSiteFunc interface {
		iGenSiteFunc

		// TODO Custom DaoFunc ....
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

func (d *customSiteDao) WithContext(ctx context.Context) iCustomGenSiteFunc {
	d.siteDo = d.siteDo.WithContext(ctx)
	return d
}

func (d *customSiteDao) WhereBetweenByCreatedAt(left time.Time, right time.Time) func(dao gen.Dao) gen.Dao {
	return func(dao gen.Dao) gen.Dao {
		return dao.Where(query.Site.CreatedAt.Between(left, right))
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
