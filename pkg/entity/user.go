package entity

import "github.com/majoo_test/soal_1/internal/pkg"

type User struct {
	EntityID   `gorm:"embedded"`
	Name       string `json:"name" gorm:"type:varchar(45);default:null"`
	UserName   string `json:"user_name" gorm:"type:varchar(45);default:null"`
	Password   string `json:"password" gorm:"type:varchar(225):default:null"`
	BaseEntity `gorm:"embedded"`
}

type UserService interface {
	FindUsers(map[string]interface{}, *Pagination) ([]User, *pkg.Errors)
}

type UserRepository interface {
	FindAll(*Pagination) ([]User, error)
	FindByID(int64) (User, error)
	FindByUsernameAndPassword(string, string) (User, error)
	AddUser(*User) error
	UpdateUser(int64, *User) error
	DeleteUser(int64) error
}

func (e *User) TableName() string {
	return "Users"
}
