package database

import "gorm.io/gorm"

//go:generate mockgen -package mock -source=db.go -destination=mock/db.go
type DB interface {
	Create(interface{}) error
	Updates(interface{}) (int64, error)
	Where(interface{}, ...interface{}) DB
	First(interface{}, ...interface{}) error
	Find(interface{}, ...interface{}) (int64, error)
	Count(interface{}) int64
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

func (s *db) Find(out interface{}, where ...interface{}) (int64, error) {
	tx := s.DB.Find(out, where...)
	return tx.RowsAffected, tx.Error
}

func (s *db) Count(model interface{}) int64 {
	var count int64
	s.DB.Model(model).Count(&count)
	return count
}

func (d *db) Updates(value interface{}) (int64, error) {
	tx := d.DB.Updates(value)
	return tx.RowsAffected, tx.Error
}
