package models

import (
	"database/sql"
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	gorm.Model
	Birthday   time.Time
	Age        int
	Name       string     `gorm:"size:255"`
	Num        int        `gorm:"AUTO_INCREMENT"`
	CreditCard CreditCard // 一对一,creditCard表的userId作外键
	Emails     []Email    // one-to-mony email 表中的userId作外键

	BillingAddress   Address // one-to-one 本表中BillingAddressID作外键
	BillingAddressID sql.NullInt64

	ShippingAddress   Address // One-To-One (属于 - 本表的ShippingAddressID作外键)
	ShippingAddressID int

	IgnoreMe int        `gorm:"-"`                          // 忽略这个字段
	Language []Language `gorm:"many2many:uuser_languages;"` // Many-To-Many , 'user_languages'是连接表
}

type Email struct {
	ID         int
	UserID     int    `gorm:"index"`                          // 外键, tag 'index' 表示为该列创建index索引
	Email      string `gorm:"type:varchar(100);unique_index"` // 设置类型, 病添加unique_index 索引
	Subscribed bool
}

type Address struct {
	ID       int
	Address1 string         `gorm:"not null; unique"`         // 设置字段不能为空,并唯一
	Address2 string         `gorm:"type:varchar(100);unique"` //
	Post     sql.NullString `gorm:"not null"`
}

type Language struct {
	ID   int
	Name string `gorm:"index:idx_name_code"`
	Code string `gorm:"index:index_name_code"`
}

type CreditCard struct {
	gorm.Model
	UserID uint
	Number string
}

type Animal struct {
	gorm.Model
	Name string `gorm:default:'galeone'`
	Age  int64
}
