package entity

type Transaction struct {
	EntityID   `gorm:"embedded"`
	MerchantID int64   `json:"merchant_id" gorm:"type:bigint(20);not null"`
	OutletID   int64   `json:"outlet_id" gorm:"type:bigint(20);not null"`
	BillTotal  float64 `json:"bill_total" gorm:"type:double;not null"`
	BaseEntity `gorm:"embedded"`
	Merchant   Merchant `json:"-" gorm:"foreignKey:MerchantID"`
	Outlet     Outlet   `json:"-" gorm:"foreignKey:OutletID"`
}

type TransactionService interface {
}

type TransactionRepository interface {
	FindByMerchant([]int64, string, string, *Pagination) ([]Transaction, error)
	FindByOutlet([]int64, string, string, *Pagination) ([]Transaction, error)
}

func (e *Transaction) TableName() string { return "Transactions" }
