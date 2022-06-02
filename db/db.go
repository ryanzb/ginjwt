package db

import (
	"ginjwt/conf"
	"log"

	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var ProviderSet = wire.NewSet(NewGormDB)

func NewGormDB(cfg *conf.DB) *gorm.DB {
	db, err := gorm.Open(mysql.Open(cfg.DSN), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		log.Fatalf("failed opening connection to db: %v", err)
	}
	return db
}