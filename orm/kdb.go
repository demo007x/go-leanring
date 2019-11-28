package orm

import "database/sql"

func RegisterDataBase(config KConfig)  {
	for _,dbConf := range config.DBConfig {
		db, error := sql.Open(dbConf.Driver, dbConf.Dsn)
		if error != nil {
			panic(error)
		}

		if dbConf.MaxLifetime > 0 {
			db.SetConnMaxLifetime(dbConf.MaxLifetime)
		}

		if dbConf.MaxOpenConns > 0 {
			db.SetMaxOpenConns(dbConf.MaxOpenConns)
		}

		if dbConf.Name == "" {
			dbConf.Name = defaultGroupName
		}
		m.addDB(dbConf.Name, dbConf.IsMater, db)
	}
}
