package orm

import "time"

type DBConfig struct {
	Name string  // 数据库链接名称
	IsMater bool // 主数据库
	Driver string
	Dsn string
	MaxLifetime time.Duration
	MaxOpenConns int //最大链接数量
}

type KConfig struct {
	DBConfig []DBConfig
}