package postgresql

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/spf13/viper"
)

const (
	host     = "postgres"
	port     = 5432
	user     = "user"
	password = "password"
	dbname   = "user"
)

// Connection gets connection of postgresql database
func Connection() (db *gorm.DB) {

	host := viper.Get("PGHOST")
	user := viper.Get("PGUSER")
	pass := viper.Get("PGPASSWORD")
	dsn := fmt.Sprintf("host=%v user=%v password=%v", host, user, pass)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	return db
}
