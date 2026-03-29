package models

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DBModel struct {
	Vehicle VehicleModel
}

func InitDB(dataSourceName string) (*DBModel, error) {
	db, err := gorm.Open(sqlite.Open(dataSourceName), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("Failed to migrate database: %v", err)
	}
	err = db.AutoMigrate(&Vehicle{}, &VehicleItem{})
	if err != nil {
		return nil, fmt.Errorf("Failed to migrate database: %v", err)

	}

	dbModel := &DBModel{
		Vehicle: VehicleModel{DB: db},
	}

	return dbModel, nil
}
