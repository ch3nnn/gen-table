/**
 * @Author: chentong
 * @Date: 2024/07/21 下午9:39
 */

package dao

import (
	"context"
	"testing"

	"gen-table/example/dal/query"

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
	dao := NewSiteDao(context.Background())
	findOne, err := dao.FindOne([]func(dao gen.Dao) gen.Dao{
		dao.WhereByID(1),
		dao.WhereByTitle("ch3nnn Github 开源"),
		dao.WhereByURL("https://github.com/ch3nnn"),
	}...)

	assert.NoError(t, err)
	assert.Equal(t, findOne.ID, int64(1))
	assert.Equal(t, findOne.Title, "ch3nnn Github 开源")
}

func TestFindAll(t *testing.T) {
	dao := NewSiteDao(context.Background())
	sites, err := dao.FindAll(dao.WhereByID(1))
	if err != nil {
		return
	}

	assert.Len(t, sites, 1)

}

func TestFindPage(t *testing.T) {
	dao := NewSiteDao(context.Background())
	findPage, count, err := dao.FindPage(1, 10, nil, dao.WhereByID(1))
	if err != nil {
		return
	}

	assert.Len(t, findPage, 1)
	assert.Equal(t, count, int64(1))
}
