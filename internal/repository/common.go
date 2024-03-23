package repository

import (
	db "len/go-grpc-api/config"

	"github.com/jinzhu/gorm"
)

// Create
func Create(value interface{}) error {
	return db.GetDB().Create(value).Error
}

// Delete
func DeleteByID(model interface{}, id uint64) (count int64, err error) {
	db := db.GetDB().Where("id=?", id).Delete(model)
	err = db.Error
	if err != nil {
		return
	}
	count = db.RowsAffected
	return
}

// Query
func Query(where string, out interface{}, associations []string) error {
	db := db.GetDB().Raw(where)
	for _, a := range associations {
		db = db.Preload(a)
	}
	return db.Find(out).Error
}

// Find
func Find(where interface{}, out interface{}, associations []string, orders ...string) error {
	db := db.GetDB()
	for _, a := range associations {
		db = db.Preload(a)
	}
	db = db.Where(where)
	if len(orders) > 0 {
		for _, order := range orders {
			db = db.Order(order)
		}
	}
	return db.Find(out).Error
}

// First
func First(where interface{}, out interface{}, associations []string) (notFound bool, err error) {
	db := db.GetDB()
	for _, a := range associations {
		db = db.Preload(a)
	}
	err = db.Where(where).First(out).Error
	if err != nil {
		notFound = gorm.IsRecordNotFoundError(err)
	}
	return
}
