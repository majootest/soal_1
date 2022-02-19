package repository

import (
	"github.com/majoo_test/soal_1/pkg/entity"
	"gorm.io/gorm"
)

type MerchantRepository struct {
	baseRepository
}

func NewMerchantRepository(db *gorm.DB) *MerchantRepository {
	tableName := (&entity.Merchant{}).TableName()
	baseRepository := New(db, tableName)
	return &MerchantRepository{baseRepository}
}

func (repo *MerchantRepository) FindAll(options *entity.Pagination) (results []entity.Merchant, err error) {
	results = make([]entity.Merchant, 0)
	var tx *gorm.DB
	if tx, err = repo.pagination(options); err != nil {
		return
	}
	err = tx.Table(repo.tableName).Order("id").Find(&results).Error
	return
}

func (repo *MerchantRepository) FindByID(id int64) (data entity.Merchant, err error) {
	err = repo.db.Where("id = ?", id).Order("id").Find(&data).Error
	return
}

func (repo *MerchantRepository) FindByUserID(userId int64) (data []entity.Merchant, err error) {
	err = repo.db.Where("user_id = ?", userId).Order("id").Find(&data).Error
	return
}
