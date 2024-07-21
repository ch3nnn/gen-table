/**
 * @Author: chentong
 * @Date: 2024/07/21 下午9:39
 */

package dao

import (
	"context"
	"fmt"
	"testing"

	"gen-table/example/dal/query"

	"github.com/duke-git/lancet/v2/pointer"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestNewSiteDao(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("../../test.db?_busy_timeout=5000"))
	if err != nil {
		return
	}

	query.SetDefault(db)

	dao := NewSiteDao(context.Background())
	sites, err := dao.SelectList(dao.WhereByID(pointer.Of(int64(1))))
	if err != nil {
		return
	}

	fmt.Println(sites[0])

}
