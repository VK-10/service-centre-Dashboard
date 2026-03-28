package models

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DBModel struct {
	Vehicle VehicleModel
}

func Init(dataSourceName string) (*DBModel, error) {
	db, err := gorm.Open(sqlite.Open(dataSourceName), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("Failed to migrate database: %v", err)
	}
	err = db.AutoMigrate(&Vehicle{}, &VehicleIssueItem{})
	if err != nil {
		return nil, fmt.Errorf("Faield t o migrate databse %v", err)

	}

	dbModel = &DBModel{
		Vehicle: VehicleModel{DB: db},
	}

	return dbModel, nil
}
