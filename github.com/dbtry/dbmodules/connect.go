package dbmodules

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// GormConnect is DB connect setting
func GormConnect() *gorm.DB {
	dbms := "mysql"
	user := "root"
	pass := "root"
	host := "tcp(0.0.0.0:3306)"
	dbname := "kids_development"

	connect := user + ":" + pass + "@" + host + "/" + dbname
	db, err := gorm.Open(dbms, connect)

	if err != nil {
		panic(err.Error())
	}

	return db
}
