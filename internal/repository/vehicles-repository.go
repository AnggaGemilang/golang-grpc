package repository

import (
	"errors"
	db "len/go-grpc-api/config"
	"len/go-grpc-api/internal/models"
)

type VehicleRepository struct{}

var vehicleRepository *VehicleRepository

func (r *VehicleRepository) Add(vehicle *models.Vehicle) error {
	err := Create(&vehicle)
	return err
}

func GetVehicleRepository() *VehicleRepository {
	if vehicleRepository == nil {
		vehicleRepository = &VehicleRepository{}
	}
	return vehicleRepository
}

func (r *VehicleRepository) All() (*[]models.Vehicle, error) {
	var vehicles []models.Vehicle
	err := Find(&models.Vehicle{}, &vehicles, []string{}, "id asc")
	return &vehicles, err
}

func (r *VehicleRepository) Get(id uint64) (*models.Vehicle, error) {
	var vehicle models.Vehicle
	where := models.Vehicle{ID: id}
	_, err := First(&where, &vehicle, []string{})
	if err != nil {
		return nil, err
	}
	return &vehicle, err
}

func (r *VehicleRepository) Query(q string) (*[]models.MostVehicle, error) {
	var mostVehicles []models.MostVehicle
	var query string
	if q == "hourly" {
		query = "SELECT vehicle_type, COUNT(*) AS count FROM vehicles WHERE created_at >= NOW() - INTERVAL 1 HOUR GROUP BY vehicle_type HAVING COUNT(*) = (SELECT COUNT(*) AS max_count FROM vehicles WHERE created_at >= NOW() - INTERVAL 1 HOUR GROUP BY vehicle_type ORDER BY max_count DESC LIMIT 1) ORDER BY count DESC"
	} else if q == "daily" {
		query = "SELECT vehicle_type, COUNT(*) AS count FROM vehicles WHERE created_at >= NOW() - INTERVAL 1 DAY GROUP BY vehicle_type HAVING COUNT(*) = (SELECT COUNT(*) AS max_count FROM vehicles WHERE created_at >= NOW() - INTERVAL 1 DAY GROUP BY vehicle_type ORDER BY max_count DESC LIMIT 1) ORDER BY count DESC"
	} else if q == "lifetime" {
		query = "SELECT vehicle_type, COUNT(*) AS count FROM vehicles GROUP BY vehicle_type HAVING COUNT(*) = (SELECT COUNT(*) AS max_count FROM vehicles GROUP BY vehicle_type ORDER BY max_count DESC LIMIT 1) ORDER BY count DESC"
	} else {
		return nil, errors.New("Kata kunci hanya tersedia hourly (satu jam terakhir), daily (satu hari terakhir), dan lifetime (seluruh waktu)")
	}
	err := Query(query, &mostVehicles, []string{})
	return &mostVehicles, err
}

func (r *VehicleRepository) Delete(vehicle *models.Vehicle) error {
	return db.GetDB().Unscoped().Delete(&vehicle).Error
}
