package repository

import (
	"github.com/majoo_test/soal_1/pkg/entity"
	"gorm.io/gorm"
)

type OutletRepository struct {
	baseRepository
}

func NewOutletRepository(db *gorm.DB) *OutletRepository {
	tableName := (&entity.Outlet{}).TableName()
	baseRepository := New(db, tableName)
	return &OutletRepository{baseRepository}
}

func (repo *OutletRepository) FindAll(options *entity.Pagination) (results []entity.Outlet, err error) {
	results = make([]entity.Outlet, 0)
	var tx *gorm.DB
	if tx, err = repo.pagination(options); err != nil {
		return
	}
	err = tx.Table(repo.tableName).Order("id").Find(&results).Error
	return
}

func (repo *OutletRepository) FindByID(id int64) (data entity.Outlet, err error) {
	err = repo.db.Where("id = ?", id).Order("id").Find(&data).Error
	return
}

func (repo *OutletRepository) FindByMerchantID(merchantIds []int64) (data []entity.Outlet, err error) {
	err = repo.db.Where("user_id IN ?", merchantIds).Order("id").Find(&data).Error
	return
}
