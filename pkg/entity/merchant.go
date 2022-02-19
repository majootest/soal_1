package entity

import "github.com/majoo_test/soal_1/internal/pkg"

type Merchant struct {
	EntityID     `gorm:"embedded"`
	UserID       int64  `json:"user_id" gorm:"type:bigint(20);not null"`
	MerchantName string `json:"merchant_name" gorm:"type:varchar(40);not null"`
	BaseEntity   `gorm:"embedded"`
	User         User `json:"-" gorm:"foreignKey:UserID"`
}

type MerchantOmzet struct {
	MerchantName string  `json:"merchant_name"`
	Omzet        float64 `json:"omzet"`
}

type MerchantService interface {
	FindOmzetReportNovember(int64, *Pagination) ([]MerchantOmzet, *pkg.Errors)
}

type MerchantRepository interface {
	FindAll(*Pagination) ([]Merchant, error)
	FindByID(int64) (Merchant, error)
	FindByUserID(int64) ([]Merchant, error)
}

func (e *Merchant) TableName() string { return "Merchants" }
