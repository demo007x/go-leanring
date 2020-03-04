package model

import "time"

const (
	SexWomen   = "W"
	SexMan     = "M"
	SexUnKnown = "U"
)

type User struct {
	Id       int64     `xorm:"pk autoincr bigint(64)" from:"id" json:"id"`
	Mobile   string    `xorm:"varchar(20)" from:"mobile" json:"mobile"`
	Passwd   string    `xorm:"varchar(40)" from:"passwd" json:"_"`
	Avatar   string    `xorm:"varchar(150)" from:"avatar" json:"avatar"`
	Sex      string    `xorm:"varchar(2)" from:"sex" json:"sex"`
	Nickname string    `xorm:"varchar(20)" from:"nickname" json:"nickname"`
	Salt     string    `xorm:"varchar(10)" from:"salt" json:"_"`
	Online   int       `xorm:"int(10)" from:"online" json:"online"`
	Token    string    `xorm:"varchar(40)" from:"token" json:"token"`
	Memo     string    `xorm:"varchar(140)" from:"memo" json:"memo"`
	Createat time.Time `xorm:"datetime" from:"createat" json:"createat"`
}
