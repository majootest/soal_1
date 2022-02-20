package repository

import (
	"github.com/majoo_test/soal_1/pkg/entity"
	"gorm.io/gorm"
)

type TransactionRepository struct {
	baseRepository
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	tableName := (&entity.Transaction{}).TableName()
	// db.AutoMigrate(&entity.Transaction{})
	baseRepository := New(db, tableName)
	return &TransactionRepository{baseRepository}
}

func (repo *TransactionRepository) FindByMerchant(merchantIds []int64, dateFrom string, dateTo string, options *entity.Pagination) (results []entity.Transaction, err error) {
	results = make([]entity.Transaction, 0)
	var tx *gorm.DB
	if tx, err = repo.pagination(options); err != nil {
		return
	}
	tx = tx.Table(repo.tableName)

	err = tx.Select("merchant_id, DATE_FORMAT(created_at, '%Y-%m-%d') as created_at, SUM(bill_total) as bill_total").Where("merchant_id IN ? AND created_at >= ? AND created_at <= ?", merchantIds, dateFrom, dateTo).Group("merchant_id,DATE_FORMAT(created_at, '%Y-%m-%d')").Find(&results).Error
	return
}

func (repo *TransactionRepository) FindByOutlet(outletIds []int64, dateFrom string, dateTo string, options *entity.Pagination) (results []entity.Transaction, err error) {
	results = make([]entity.Transaction, 0)
	var tx *gorm.DB
	if tx, err = repo.pagination(options); err != nil {
		return
	}
	tx = tx.Table(repo.tableName)

	err = tx.Select("merchant_id, outlet_id, DATE_FORMAT(created_at, '%Y-%m-%d') as created_at, SUM(bill_total) as bill_total").Where("outlet_id IN ? AND created_at >= ? AND created_at <= ?", outletIds, dateFrom, dateTo).Group("merchant_id,outlet_id,DATE_FORMAT(created_at, '%Y-%m-%d')").Find(&results).Error
	return
}
