package models

import (
	"time"

	"github.com/teris-io/shortid"
	"gorm.io/gorm"
)

var (
	VehicleStatus = []string{
		"Checked In",    // vehicle received, job card created
		"In Service",    // under inspection / being worked on
		"On Hold",       // awaiting parts / customer approval
		"Quality Check", // work done, final check before delivery
		"Ready",         // handed over to customer
	}

	VehicleTypes = []string{
		// Scooters
		"Access 125",
		"Avenis 125",
		"Burgman Street 125",
		"E-Access", // electric

		// Commuter / Street
		"Gixxer",    // 150cc
		"Gixxer SF", // 150cc faired
		"Intruder 150",
		"Intruder 150 Fi",

		// Mid-range Sport
		"Gixxer 250",
		"Gixxer SF 250",
		"Gixxer SF 250 Flex Fuel",

		// Adventure
		"V-Strom SX", // 250cc
		"V-Strom 800DE",

		// Big Bikes
		"GSX-8S",
		"GSX-8R",
		"Hayabusa",
	}

	VehicleIssues = []string{
		"Mechanical",      // engine, gearbox, suspension, brakes
		"Electrical",      // battery, wiring, lights, indicators
		"Body & Frame",    // dents, scratches, panels, handles
		"Routine Service", // oil change, tyres, filters, chain
	}
)

type VehicleModel struct {
	DB *gorm.DB
}

type Vehicle struct {
	ID           string        `gorm:"primaryKey; size:14" json:"id"`
	Status       string        `gorm:"not null" json:"status"`
	CustomerName string        `gorm:"not null" json:"customerName"`
	Phone        string        `gorm:"not null" json:"phone"`
	Address      string        `gorm:"not null" json:"address"`
	Items        []VehicleItem `gorm:"foreignKey:ServiceID" json:"items"`
	CreatedAt    time.Time     `json:"createdAt"`
	// ChasisNumber string        `gorm:"not null" json:"chasisNumber"`
}

type VehicleItem struct {
	ID             string `gorm:"primaryKey; size:14" json:"id"`
	ServiceId      string `gorm:"index;size:14;not null" json:"serviceId"`
	Issue          string `gorm:"not null" json:"issue"`
	Vehicle        string `gorm:"not null" json:"vehicle"`
	IssueReproduce string `json:"IssueReproduce"`
}

// hook
func (v *Vehicle) BeforeCreate(ctx *gorm.DB) error {
	if v.ID == "" {
		v.ID = shortid.MustGenerate()
	}

	return nil
}

func (vi *VehicleItem) BeforeCreate(ctx *gorm.DB) error {
	if vi.ID == "" {
		vi.ID = shortid.MustGenerate()
	}

	return nil
}

// func (v *VehicleModel) CreateVehicle(vehicle *Vehicle) error {
// 	var existing Vehicle
// 	err := v.DB.Where("chasis_number = ? AND status != ?", vehicle.ChasisNumber, "Ready").First(&existing).Error
// 	if err == nil {
// 		return fmt.Errorf("Vehicle with chasis number %s already exists", vehicle.ChasisNumber)
// 	}
// 	if err == gorm.ErrRecordNotFound {
// 		return v.DB.Create(vehicle).Error
// 	}

// 	return err

// }

func (v *VehicleModel) CreateVehicle(vehicle *Vehicle) error {
	return v.DB.Create(vehicle).Error
}

func (v *VehicleModel) GetVehicle(id string) (*Vehicle, error) {
	var vehicle Vehicle
	err := v.DB.Preload("Items").First(&vehicle, "id = ?", id).Error // Preload issues
	return &vehicle, err
}

func (v *VehicleModel) GetAllVehicles() ([]Vehicle, error) {
	var vehicles []Vehicle
	err := v.DB.Preload("Items").Order("created_at desc").Find(&vehicles).Error
	return vehicles, err
}

func (v *VehicleModel) UpdateVehicleStatus(id string, status string) error {
	return v.DB.Model(&Vehicle{}).Where("id = ?", id).Update("status", status).Error
}

func (v *VehicleModel) DeleteVehicle(id string) error {
	return v.DB.Select("Items").Delete(&Vehicle{ID: id}).Error

}
