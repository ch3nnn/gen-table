// Code generated. DO NOT EDIT.
// Code generated. DO NOT EDIT.
// Code generated. DO NOT EDIT.

package dao

import (
	"time"

	"gen-table/example/dal/model"
	"gen-table/example/dal/query"

	"github.com/duke-git/lancet/v2/structs"
	"gorm.io/gen"
)

var _ ISite = (*SiteDao)(nil)

type iSite interface {
	WhereByID(id *int64) func(dao gen.Dao) gen.Dao
	WhereByCategoryID(categoryId *int64) func(dao gen.Dao) gen.Dao
	WhereByTitle(title *string) func(dao gen.Dao) gen.Dao
	WhereByThumb(thumb *string) func(dao gen.Dao) gen.Dao
	WhereByDescription(description *string) func(dao gen.Dao) gen.Dao
	WhereByURL(url *string) func(dao gen.Dao) gen.Dao
	WhereByCreatedAt(createdAt *time.Time) func(dao gen.Dao) gen.Dao
	WhereByUpdatedAt(updatedAt *time.Time) func(dao gen.Dao) gen.Dao
	WhereByIsUsed(isUsed *bool) func(dao gen.Dao) gen.Dao
	WhereByType(type_ *string) func(dao gen.Dao) gen.Dao

	Create(m *model.Site) (*model.Site, error)
	Delete(whereFunc ...func(dao gen.Dao) gen.Dao) error
	DeletePhysical(whereFunc ...func(dao gen.Dao) gen.Dao) error
	Update(m interface{}, whereFunc ...func(dao gen.Dao) gen.Dao) (*model.Site, error)
	Select(whereFunc ...func(dao gen.Dao) gen.Dao) (*model.Site, error)
	SelectList(whereFunc ...func(dao gen.Dao) gen.Dao) ([]*model.Site, error)
	SelectPage(offset int, limit int, whereFunc ...func(dao gen.Dao) gen.Dao) ([]*model.Site, int64, error)
}

type siteDao struct {
	siteDo query.ISiteDo
}

func (s *siteDao) WhereByID(id *int64) func(dao gen.Dao) gen.Dao {
	return func(dao gen.Dao) gen.Dao {
		if id != nil {
			return dao.Where(query.Site.ID.Eq(*id))
		}
		return dao
	}
}

func (s *siteDao) WhereByCategoryID(categoryId *int64) func(dao gen.Dao) gen.Dao {
	return func(dao gen.Dao) gen.Dao {
		if categoryId != nil {
			return dao.Where(query.Site.CategoryID.Eq(*categoryId))
		}
		return dao
	}
}

func (s *siteDao) WhereByTitle(title *string) func(dao gen.Dao) gen.Dao {
	return func(dao gen.Dao) gen.Dao {
		if title != nil {
			return dao.Where(query.Site.Title.Eq(*title))
		}
		return dao
	}
}

func (s *siteDao) WhereByThumb(thumb *string) func(dao gen.Dao) gen.Dao {
	return func(dao gen.Dao) gen.Dao {
		if thumb != nil {
			return dao.Where(query.Site.Thumb.Eq(*thumb))
		}
		return dao
	}
}

func (s *siteDao) WhereByDescription(description *string) func(dao gen.Dao) gen.Dao {
	return func(dao gen.Dao) gen.Dao {
		if description != nil {
			return dao.Where(query.Site.Description.Eq(*description))
		}
		return dao
	}
}

func (s *siteDao) WhereByURL(url *string) func(dao gen.Dao) gen.Dao {
	return func(dao gen.Dao) gen.Dao {
		if url != nil {
			return dao.Where(query.Site.URL.Eq(*url))
		}
		return dao
	}
}

func (s *siteDao) WhereByCreatedAt(createdAt *time.Time) func(dao gen.Dao) gen.Dao {
	return func(dao gen.Dao) gen.Dao {
		if createdAt != nil {
			return dao.Where(query.Site.CreatedAt.Eq(*createdAt))
		}
		return dao
	}
}

func (s *siteDao) WhereByUpdatedAt(updatedAt *time.Time) func(dao gen.Dao) gen.Dao {
	return func(dao gen.Dao) gen.Dao {
		if updatedAt != nil {
			return dao.Where(query.Site.UpdatedAt.Eq(*updatedAt))
		}
		return dao
	}
}

func (s *siteDao) WhereByIsUsed(isUsed *bool) func(dao gen.Dao) gen.Dao {
	return func(dao gen.Dao) gen.Dao {
		if isUsed != nil {
			return dao.Where(query.Site.IsUsed.Is(*isUsed))
		}
		return dao
	}
}

func (s *siteDao) WhereByType(type_ *string) func(dao gen.Dao) gen.Dao {
	return func(dao gen.Dao) gen.Dao {
		if type_ != nil {
			return dao.Where(query.Site.Type.Eq(*type_))
		}
		return dao
	}
}

func (s *siteDao) Create(m *model.Site) (*model.Site, error) {
	if err := s.siteDo.Create(m); err != nil {
		return nil, err
	}
	return s.Select(s.WhereByID(&m.ID))
}

func (s *siteDao) Select(whereFunc ...func(dao gen.Dao) gen.Dao) (*model.Site, error) {
	return s.siteDo.Scopes(whereFunc...).First()
}

func (s *siteDao) SelectList(whereFunc ...func(dao gen.Dao) gen.Dao) ([]*model.Site, error) {
	return s.siteDo.Scopes(whereFunc...).Find()
}

func (s *siteDao) SelectPage(offset int, limit int, whereFunc ...func(dao gen.Dao) gen.Dao) ([]*model.Site, int64, error) {
	return s.siteDo.Scopes(whereFunc...).FindByPage(offset, limit)
}

func (s *siteDao) Update(m interface{}, whereFunc ...func(dao gen.Dao) gen.Dao) (*model.Site, error) {
	toMap, err := structs.ToMap(m)
	if err != nil {
		return nil, err
	}

	if _, err := s.siteDo.Scopes(whereFunc...).Updates(toMap); err != nil {
		return nil, err
	}

	return s.Select(whereFunc...)
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
