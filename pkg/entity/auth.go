package entity

import (
	"github.com/lestrrat-go/jwx/jwt"
	"github.com/majoo_test/soal_1/internal/pkg"
)

type AuthService interface {
	UserLogin(*User) (string, *pkg.Errors)
	UserAuth(string) (jwt.Token, *pkg.Errors)
}
