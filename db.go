package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func setupDatabase(cfg *Config) *gorm.DB {
	connectionString := fmt.Sprintf("%s:%s@(%s:%s)/%s?parseTime=True", cfg.Database.Username, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.DatabaseName)
	fmt.Print(connectionString)
	db, err := gorm.Open("mysql", connectionString)
	if err != nil {
		panic("failed to connect database")
	}

	db.SingularTable(true)
	return db
}
