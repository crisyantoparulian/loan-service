package driver

import (
	"github.com/crisyantoparulian/loansvc/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewGormDatabase(cfg config.Database) (db *gorm.DB) {
	dsn := cfg.PostgreDSN
	dialector := postgres.Open(dsn)

	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic("db connection failed")
	}
	sqlDB.SetMaxIdleConns(cfg.MaxIddleConn)
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConn)

	if cfg.DebugMode {
		db = db.Debug()
	}

	return
}
