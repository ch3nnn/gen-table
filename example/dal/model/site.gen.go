// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameSite = "site"

// Site mapped from table <site>
type Site struct {
	ID          int64      `gorm:"column:id;type:INTEGER" json:"id"`
	CategoryID  int64      `gorm:"column:category_id;type:int(11)" json:"category_id"`
	Title       string     `gorm:"column:title;type:varchar(50)" json:"title"`
	Thumb       string     `gorm:"column:thumb;type:varchar(200)" json:"thumb"`
	Description string     `gorm:"column:description;type:varchar(300)" json:"description"`
	URL         string     `gorm:"column:url;type:varchar(200)" json:"url"`
	CreatedAt   *time.Time `gorm:"column:created_at;type:datetime;not null;default:CURRENT_TIMESTAMP not null" json:"created_at"`
	UpdatedAt   *time.Time `gorm:"column:updated_at;type:datetime" json:"updated_at"`
	IsUsed      bool       `gorm:"column:is_used;type:bool;default:false" json:"is_used"`
	Type        string     `gorm:"column:type;type:varchar(50)" json:"type"`
}

// TableName Site's table name
func (*Site) TableName() string {
	return TableNameSite
}
