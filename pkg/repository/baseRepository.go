package repository

import (
	"github.com/majoo_test/soal_1/pkg/entity"
	"gorm.io/gorm"
)

type baseRepository struct {
	db        *gorm.DB
	tableName string
}

func New(db *gorm.DB, tableName string) baseRepository {
	return baseRepository{db, tableName}
}

func (repo *baseRepository) DBInsert(data entity.Entity) (err error) {
	err = repo.db.Create(data).Error
	return
}

func (repo *baseRepository) UpdateByID(id int64, data entity.Entity) (err error) {
	err = repo.db.Table(repo.tableName).Where("id = ?", id).Updates(data).Error
	return
}

func (repo *baseRepository) DeleteByID(data entity.Entity) (err error) {
	err = repo.db.Delete(data).Error
	return
}

func (repo *baseRepository) pagination(options *entity.Pagination) (db *gorm.DB, err error) {
	limit := options.Limit
	if limit < 1 {
		limit = 10
	}
	db = repo.db.Limit(limit)
	if err = db.Error; err != nil {
		return
	}
	pageNo := options.PageNo
	if pageNo < 1 {
		pageNo = 1
	}
	skip := (pageNo - 1) * limit
	db = db.Offset(skip)
	err = db.Error
	return
}
