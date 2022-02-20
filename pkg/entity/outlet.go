package entity

import "github.com/majoo_test/soal_1/internal/pkg"

type Outlet struct {
	EntityID   `gorm:"embedded"`
	MerchantID int64  `json:"merchant_id" gorm:"type:bigint(20);foreignKey;not null"`
	OutletName string `json:"outlet_name" gorm:"type:varchar(40);not null"`
	BaseEntity `gorm:"embedded"`
	Merchant   Merchant `json:"-" gorm:"foreignKey:MerchantID"`
}

type OutletOmzet struct {
	MerchantName string  `json:"merchant_name"`
	OutletName   string  `json:"outlet_name"`
	Omzet        float64 `json:"omzet"`
	CreatedAt    string  `json:"created_at"`
}

type OutletService interface {
	FindOmzetReportNovember(int64, *Pagination) ([]OutletOmzet, *pkg.Errors)
}

type OutletRepository interface {
	FindAll(*Pagination) ([]Outlet, error)
	FindByID(int64) (Outlet, error)
	FindByMerchantID([]int64) ([]Outlet, error)
}

func (e *Outlet) TableName() string { return "Outlets" }
