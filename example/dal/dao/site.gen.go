// Code generated by gen-table. DO NOT EDIT.
// Code generated by gen-table. DO NOT EDIT.
// Code generated by gen-table. DO NOT EDIT.

package dao

import (
	"time"

	"github.com/ch3nnn/gen-table/example/dal/model"
	"github.com/ch3nnn/gen-table/example/dal/query"

	"gorm.io/gen"
	"gorm.io/gen/field"
)

var _ iSiteDao = (*siteDao)(nil)

// ------------------------------------
// Site  ColumnName
// ------------------------------------
type iWhereSiteFunc interface {
	WhereByID(id int) func(dao gen.Dao) gen.Dao
	WhereByCategoryID(categoryId int) func(dao gen.Dao) gen.Dao
	WhereByTitle(title string) func(dao gen.Dao) gen.Dao
	WhereByThumb(thumb string) func(dao gen.Dao) gen.Dao
	WhereByDescription(description string) func(dao gen.Dao) gen.Dao
	WhereByURL(url string) func(dao gen.Dao) gen.Dao
	WhereByCreatedAt(createdAt time.Time) func(dao gen.Dao) gen.Dao
	WhereByUpdatedAt(updatedAt time.Time) func(dao gen.Dao) gen.Dao
	WhereByIsUsed(isUsed bool) func(dao gen.Dao) gen.Dao
	WhereByType(type_ string) func(dao gen.Dao) gen.Dao
}

// ------------------------------------
// Site  Generate Function
// ------------------------------------
type iGenSiteFunc interface {
	Create(m *model.Site) (*model.Site, error)
	Delete(whereFunc ...func(dao gen.Dao) gen.Dao) error
	DeletePhysical(whereFunc ...func(dao gen.Dao) gen.Dao) error
	Update(v interface{}, whereFunc ...func(dao gen.Dao) gen.Dao) (rowsAffected int64, err error)
	FindCount(whereFunc ...func(dao gen.Dao) gen.Dao) (int64, error)
	FindOne(whereFunc ...func(dao gen.Dao) gen.Dao) (*model.Site, error)
	FindAll(whereFunc ...func(dao gen.Dao) gen.Dao) ([]*model.Site, error)
	FindPage(page int, pageSize int, orderColumns []field.Expr, whereFunc ...func(dao gen.Dao) gen.Dao) ([]*model.Site, int64, error)
	Scan(result interface{}, whereFunc ...func(dao gen.Dao) gen.Dao) (err error)
	ScanPage(page int, pageSize int, orderColumns []field.Expr, result interface{}, whereFunc ...func(dao gen.Dao) gen.Dao) (count int64, err error)
}

type iSiteDao interface {
	iWhereSiteFunc
	iGenSiteFunc
}

type siteDao struct {
	siteDo query.ISiteDo
}

func (s *siteDao) WhereByID(id int) func(dao gen.Dao) gen.Dao {
	return func(dao gen.Dao) gen.Dao {
		return dao.Where(query.Site.ID.Eq(id))
	}
}

func (s *siteDao) WhereByCategoryID(categoryId int) func(dao gen.Dao) gen.Dao {
	return func(dao gen.Dao) gen.Dao {
		return dao.Where(query.Site.CategoryID.Eq(categoryId))
	}
}

func (s *siteDao) WhereByTitle(title string) func(dao gen.Dao) gen.Dao {
	return func(dao gen.Dao) gen.Dao {
		return dao.Where(query.Site.Title.Eq(title))
	}
}

func (s *siteDao) WhereByThumb(thumb string) func(dao gen.Dao) gen.Dao {
	return func(dao gen.Dao) gen.Dao {
		return dao.Where(query.Site.Thumb.Eq(thumb))
	}
}

func (s *siteDao) WhereByDescription(description string) func(dao gen.Dao) gen.Dao {
	return func(dao gen.Dao) gen.Dao {
		return dao.Where(query.Site.Description.Eq(description))
	}
}

func (s *siteDao) WhereByURL(url string) func(dao gen.Dao) gen.Dao {
	return func(dao gen.Dao) gen.Dao {
		return dao.Where(query.Site.URL.Eq(url))
	}
}

func (s *siteDao) WhereByCreatedAt(createdAt time.Time) func(dao gen.Dao) gen.Dao {
	return func(dao gen.Dao) gen.Dao {
		return dao.Where(query.Site.CreatedAt.Eq(createdAt))
	}
}

func (s *siteDao) WhereByUpdatedAt(updatedAt time.Time) func(dao gen.Dao) gen.Dao {
	return func(dao gen.Dao) gen.Dao {
		return dao.Where(query.Site.UpdatedAt.Eq(updatedAt))
	}
}

func (s *siteDao) WhereByIsUsed(isUsed bool) func(dao gen.Dao) gen.Dao {
	return func(dao gen.Dao) gen.Dao {
		return dao.Where(query.Site.IsUsed.Is(isUsed))
	}
}

func (s *siteDao) WhereByType(type_ string) func(dao gen.Dao) gen.Dao {
	return func(dao gen.Dao) gen.Dao {
		return dao.Where(query.Site.Type.Eq(type_))
	}
}

func (s *siteDao) Create(m *model.Site) (*model.Site, error) {
	if err := s.siteDo.Create(m); err != nil {
		return nil, err
	}
	return s.FindOne(s.WhereByID(m.ID))
}

func (s *siteDao) FindCount(whereFunc ...func(dao gen.Dao) gen.Dao) (int64, error) {
	return s.siteDo.Scopes(whereFunc...).Count()
}

func (s *siteDao) FindOne(whereFunc ...func(dao gen.Dao) gen.Dao) (*model.Site, error) {
	return s.siteDo.Scopes(whereFunc...).First()
}

func (s *siteDao) FindAll(whereFunc ...func(dao gen.Dao) gen.Dao) ([]*model.Site, error) {
	return s.siteDo.Scopes(whereFunc...).Find()
}

func (s *siteDao) FindPage(page int, pageSize int, orderColumns []field.Expr, whereFunc ...func(dao gen.Dao) gen.Dao) ([]*model.Site, int64, error) {
	return s.siteDo.Scopes(whereFunc...).Order(orderColumns...).FindByPage((page-1)*pageSize, pageSize)
}

// 注意 当通过 struct 更新时，GORM 只会更新非零字段
// 如果想确保指定字段被更新，使用 map 来完成更新操作
func (s *siteDao) Update(v interface{}, whereFunc ...func(dao gen.Dao) gen.Dao) (rowsAffected int64, err error) {
	info, err := s.siteDo.Scopes(whereFunc...).Updates(v)
	if err != nil {
		return rowsAffected, err
	}

	return info.RowsAffected, nil
}

func (s *siteDao) Delete(whereFunc ...func(dao gen.Dao) gen.Dao) error {
	if _, err := s.siteDo.Scopes(whereFunc...).Delete(); err != nil {
		return err
	}
	return nil
}

func (s *siteDao) DeletePhysical(whereFunc ...func(dao gen.Dao) gen.Dao) error {
	if _, err := s.siteDo.Unscoped().Scopes(whereFunc...).Delete(); err != nil {
		return err
	}
	return nil
}

func (s *siteDao) Scan(result interface{}, whereFunc ...func(dao gen.Dao) gen.Dao) (err error) {
	return s.siteDo.Scopes(whereFunc...).Scan(result)
}

func (s *siteDao) ScanPage(page int, pageSize int, orderColumns []field.Expr, result interface{}, whereFunc ...func(dao gen.Dao) gen.Dao) (count int64, err error) {
	return s.siteDo.Scopes(whereFunc...).Order(orderColumns...).ScanByPage(result, (page-1)*pageSize, pageSize)
}
