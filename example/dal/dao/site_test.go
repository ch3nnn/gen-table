/**
 * @Author: chentong
 * @Date: 2024/07/21 下午9:39
 */

package dao

import (
	"context"
	"testing"

	"github.com/ch3nnn/gen-table/example/dal/model"
	"github.com/ch3nnn/gen-table/example/dal/query"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func TestMain(m *testing.M) {
	db, err := gorm.Open(sqlite.Open("../../test.db?_busy_timeout=5000"))
	if err != nil {
		return
	}

	query.SetDefault(db)
	m.Run()

}

func TestFindOne(t *testing.T) {
	siteDao := NewSiteDao()
	findOne, err := siteDao.WithContext(context.Background()).FindOne([]func(dao gen.Dao) gen.Dao{
		siteDao.WhereByID(1),
		siteDao.WhereByTitle("ch3nnn Github 开源"),
		siteDao.WhereByURL("https://github.com/ch3nnn"),
	}...)

	assert.NoError(t, err)
	assert.Equal(t, findOne.ID, 1)
	assert.Equal(t, findOne.Title, "ch3nnn Github 开源")
}

func TestFindAll(t *testing.T) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "findAll", "findAll")

	siteDao := NewSiteDao()
	sites, err := siteDao.WithContext(ctx).FindAll(siteDao.WhereByID(1))
	if err != nil {
		return
	}

	assert.Len(t, sites, 1)

}

func TestFindPage(t *testing.T) {
	siteDao := NewSiteDao()
	findPage, count, err := siteDao.WithContext(context.Background()).FindPage(1, 10, nil, siteDao.WhereByID(1))
	if err != nil {
		return
	}

	assert.Len(t, findPage, 1)
	assert.Equal(t, count, int64(1))
}

func TestRewrite_FindCount(t *testing.T) {
	dao := NewSiteDao()
	count, err := dao.WithContext(context.Background()).FindCount()

	assert.NoError(t, err)
	assert.Equal(t, int64(-99), count)
}

func TestCustom_FirstOrInit(t *testing.T) {
	site := &model.Site{
		ID:          2,
		Title:       "hello",
		Description: "doc",
	}

	dao := NewSiteDao()
	first, err := dao.WithContext(context.Background()).
		FirstOrInit(
			dao.WhereByID(2),
			dao.WhereByTitle("hello"),
			dao.WhereByDescription("doc"),
		)
	assert.NoError(t, err)
	assert.Equal(t, site, first)
}
