package seeder

import (
	"motorcycle-rent-api/app/constant"
	"motorcycle-rent-api/app/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func SeedMotorcycles(db *gorm.DB) {
	motorcycles := []model.Motorcycle{
		{
			PlateNumber: "DK1234ABC",
			Brand:       "Yamaha Mio",
			Type:        constant.MotorcycleTypeMatic,
			Year:        2020,
			Status:      constant.MotorcycleStatusAvailable,
			PricePerDay: 200000,
		},
		{
			PlateNumber: "DK5678DEF",
			Brand:       "Honda CBR",
			Type:        constant.MotorcycleTypeManual,
			Year:        2019,
			Status:      constant.MotorcycleStatusAvailable,
			PricePerDay: 300000,
		},
		{
			PlateNumber: "DK9012GHI",
			Brand:       "Suzuki Satria",
			Type:        constant.MotorcycleTypeManual,
			Year:        2021,
			Status:      constant.MotorcycleStatusMaintenance,
			PricePerDay: 250000,
		},
		{
			PlateNumber: "DK3456JKO",
			Brand:       "Kawasaki Ninja",
			Type:        constant.MotorcycleTypeManual,
			Year:        2022,
			Status:      constant.MotorcycleStatusAvailable,
			PricePerDay: 400000,
		},
		{
			PlateNumber: "DK7890MNO",
			Brand:       "Ducati Panigale",
			Type:        constant.MotorcycleTypeManual,
			Year:        2022,
			Status:      constant.MotorcycleStatusAvailable,
			PricePerDay: 400000,
		},
		{
			PlateNumber: "DK3345PQZ",
			Brand:       "Yamaha R1",
			Type:        constant.MotorcycleTypeMatic,
			Year:        2020,
			Status:      constant.MotorcycleStatusAvailable,
			PricePerDay: 300000,
		},
		{
			PlateNumber: "DK6789STU",
			Brand:       "Honda Vario",
			Type:        constant.MotorcycleTypeMatic,
			Year:        2019,
			Status:      constant.MotorcycleStatusAvailable,
			PricePerDay: 210000,
		},
		{
			PlateNumber: "DK0123VWX",
			Brand:       "Suzuki Burgman",
			Type:        constant.MotorcycleTypeMatic,
			Year:        2021,
			Status:      constant.MotorcycleStatusMaintenance,
			PricePerDay: 250000,
		},
		{
			PlateNumber: "DK3456JKL",
			Brand:       "Kawasaki Ninja",
			Type:        constant.MotorcycleTypeMatic,
			Year:        2022,
			Status:      constant.MotorcycleStatusAvailable,
			PricePerDay: 400000,
		},
		{
			PlateNumber: "DK7890MNC",
			Brand:       "Ducati Panigale",
			Type:        constant.MotorcycleTypeMatic,
			Year:        2022,
			Status:      constant.MotorcycleStatusAvailable,
			PricePerDay: 200000,
		},
		{
			PlateNumber: "DK2345PQR",
			Brand:       "Yamaha R1",
			Type:        constant.MotorcycleTypeMatic,
			Year:        2020,
			Status:      constant.MotorcycleStatusAvailable,
			PricePerDay: 250000,
		},
		{
			PlateNumber: "DK6789G",
			Brand:       "Honda Vario",
			Type:        constant.MotorcycleTypeMatic,
			Year:        2019,
			Status:      constant.MotorcycleStatusAvailable,
			PricePerDay: 210000,
		},
	}

	if len(motorcycles) > 0 {
		err := db.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "plate_number"}},
			DoNothing: true,
		}).Create(&motorcycles).Error
		if err != nil {
			panic("Failed to seed motorcycles: " + err.Error())
		}
	}
}
