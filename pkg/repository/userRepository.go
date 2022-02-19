package repository

import (
	"github.com/majoo_test/soal_1/pkg/entity"
	"gorm.io/gorm"
)

type UserRepository struct {
	baseRepository
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	tableName := (&entity.User{}).TableName()
	// db.AutoMigrate(&entity.User{})
	baseRepository := New(db, tableName)
	return &UserRepository{baseRepository}
}

func (repo *UserRepository) FindAll(options *entity.Pagination) (results []entity.User, err error) {
	results = make([]entity.User, 0)
	var tx *gorm.DB
	if tx, err = repo.pagination(options); err != nil {
		return
	}
	err = tx.Table(repo.tableName).Order("id").Find(&results).Error
	return
}

func (repo *UserRepository) FindByUsernameAndPassword(username, password string) (data entity.User, err error) {
	err = repo.db.Where("username = ? AND password = ?", username, password).Order("id").Find(&data).Error
	return
}

func (repo *UserRepository) FindByID(id int64) (data entity.User, err error) {
	err = repo.db.Where("id = ?", id).Order("id").Find(&data).Error
	return
}

func (repo *UserRepository) AddUser(user *entity.User) (err error) {
	err = repo.DBInsert(user)
	return
}

func (repo *UserRepository) UpdateUser(userId int64, user *entity.User) (err error) {
	err = repo.UpdateByID(userId, user)
	return
}

func (repo *UserRepository) DeleteUser(userId int64) (err error) {
	user := entity.User{}
	user.ID = userId
	err = repo.DeleteByID(&user)
	return
}
