package entity

type Entity interface {
	TableName() string
}

type EntityID struct {
	ID int64 `json:"id" gorm:"autoIncrement;primaryKey;not null"`
}

type BaseEntity struct {
	CreatedAt string `json:"created_at" gorm:"type:timestamp;not null;default:current_timestamp"`
	CreatedBy int64  `json:"created_by" gorm:"type:bigint(20);not null"`
	UpdatedAt string `json:"updated_at" gorm:"type:timestamp;not null;default:current_timestamp"`
	UpdatedBy int64  `json:"updated_by" gorm:"type:bigint(20);not null"`
}

type Pagination struct {
	Limit  int
	PageNo int
}

func (entity *BaseEntity) SetCreated() {

}
