package service

import (
	"crypto/md5"
	"fmt"
	"time"

	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/lestrrat-go/jwx/jws"
	"github.com/lestrrat-go/jwx/jwt"
	"github.com/majoo_test/soal_1/configs"
	"github.com/majoo_test/soal_1/internal/pkg"
	"github.com/majoo_test/soal_1/pkg/entity"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo entity.UserRepository
}

func NewAuthService(userRepo entity.UserRepository) *AuthService {
	return &AuthService{userRepo}
}

func (srv *AuthService) UserLogin(user *entity.User) (jwtToken string, err *pkg.Errors) {

	if user.UserName == "" || user.Password == "" {
		err = pkg.NewError("username and password field is required", 400)
		return
	}

	pwd := md5.Sum([]byte(user.Password))
	getuser, e := srv.userRepo.FindByUsernameAndPassword(user.UserName, string(pwd[:]))
	if e != nil {
		err = pkg.NewError(
			fmt.Sprintf("Could not retrieve user data : %s", e.Error()),
			500,
		)
		return
	}
	if getuser.ID == 0 {
		err = pkg.NewError("wrong user_name or password", 404)
		return
	}

	jwtToken, e = srv.generateToken(
		map[string]interface{}{
			"iss":      getuser.UserName,
			"sub":      getuser.UserName,
			"exp":      int64(36000),
			"iat":      time.Now().Format("2006-01-02 15:04:05"),
			"user_id":  getuser.ID,
			"iat_unix": time.Now().Unix(),
		},
	)
	if e != nil {
		err = pkg.NewError(
			fmt.Sprintf("Generate token failed : %s", e.Error()),
			500,
		)
	}
	return
}

func (srv *AuthService) UserAuth(jwtToken string) (token jwt.Token, err *pkg.Errors) {

	if jwtToken == "" {
		err = pkg.NewError("Access token is required", 400)
		return
	}

	token, e := srv.validateToken(jwtToken)
	if e != nil {
		err = pkg.NewError("Invalid JWT Token", 400)
		return
	}

	iat, _ := token.Get("iat")
	issuedAt, e := time.Parse("2006-01-02 15:04:05", iat.(string))
	if e != nil {
		err = pkg.NewError(
			fmt.Sprintf("Parse time failed : %s", e.Error()),
			500,
		)
		return
	}

	timeNow := time.Now().Unix()

	exp, _ := token.Get("exp")
	expiryTime := issuedAt.Unix() + exp.(int64)
	if timeNow > expiryTime {
		err = pkg.NewError("Unauthorized, Access token expired", 401)
	}
	return
}

func (srv *AuthService) validateToken(payload string) (token jwt.Token, err error) {

	key := []byte(configs.HASHKEY)

	jwkKey, err := jwk.New(key)
	if err != nil {
		return
	}

	token, err = jwt.Parse(
		[]byte(payload),
		jwt.WithVerify(jwa.HS256, jwkKey),
	)
	return
}

func (srv *AuthService) generateToken(payload map[string]interface{}) (jwtToken string, err error) {
	jwtHeader := jws.NewHeaders()
	jwtHeader.Set("typ", "JWT")
	jwtHeader.Set("alg", jwa.HS256)

	key := []byte(configs.HASHKEY)

	jwkKey, err := jwk.New(key)
	if err != nil {
		err = pkg.NewError(err.Error(), 403)
		return
	}
	token := jwt.New()
	for k, v := range payload {
		token.Set(k, v)
	}
	tokenPayload, err := jwt.Sign(token, jwa.HS256, jwkKey, jwt.WithHeaders(jwtHeader))
	if err != nil {
		err = pkg.NewError(err.Error(), 403)
		return
	}
	jwtToken = string(tokenPayload)
	return
}

func (srv *AuthService) validatePassword(hashedPassword, password string) (err error) {
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return
}
