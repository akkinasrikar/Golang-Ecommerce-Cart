package database

import "gorm.io/gorm"

type DB interface {
	Create(interface{}) error
	Where(interface{}, ...interface{}) DB
	First(interface{}, ...interface{}) error
}

type db struct {
	*gorm.DB
}

func NewDb(gormDB *gorm.DB) *db {
	return &db{DB: gormDB}
}

func (d *db) Create(value interface{}) error {
	return d.DB.Create(value).Error
}

func (d *db) Where(query interface{}, args ...interface{}) DB {
	return &db{DB: d.DB.Where(query, args...)}
}

func (d *db) First(dest interface{}, args ...interface{}) error {
	return d.DB.First(dest, args...).Error
}
