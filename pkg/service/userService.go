package service

import (
	"fmt"

	"github.com/majoo_test/soal_1/internal/pkg"
	"github.com/majoo_test/soal_1/pkg/entity"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo entity.UserRepository
}

func NewUserService(repo entity.UserRepository) *UserService {
	return &UserService{repo}
}

func (srv *UserService) FindUsers(query map[string]interface{}, pagination *entity.Pagination) (results []entity.User, err *pkg.Errors) {
	results = make([]entity.User, 0)

	if v, ok := query["id"]; ok {
		if value, ok := v.(int64); ok && value != 0 {

			if data, e := srv.repo.FindByID(value); e != nil {
				err = pkg.NewError(
					fmt.Sprintf("Could not retrieve data : %s", e.Error()),
					500,
				)
			} else {
				results = append(results, data)
			}
			return
		}
	}

	var e error
	if results, e = srv.repo.FindAll(pagination); e != nil {
		err = pkg.NewError(
			fmt.Sprintf("Could not retrieve data : %s", e.Error()),
			500,
		)
	}
	return
}

func (srv *UserService) hashPassword(password string) (hashResult []byte, err error) {
	pwd := []byte(password)
	hashResult, err = bcrypt.GenerateFromPassword(pwd, 10)
	return
}
