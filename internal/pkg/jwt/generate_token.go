package jwt

import (
	"os"
	"time"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateAccessToken(userLoginResponse *web.AuthResponse) (string, error) {
	expireTime := time.Now().Add(time.Hour * 2).Unix()
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["id"] = userLoginResponse.ID
	claims["username"] = userLoginResponse.Username
	claims["email"] = userLoginResponse.Email
	claims["role_name"] = userLoginResponse.RoleName
	claims["exp"] = expireTime

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	validToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return validToken, nil
}
func GenerateRefreshToken(userLoginResponse *web.AuthResponse) (string, error) {
	expireTime := time.Now().Add(time.Hour * 24).Unix()
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["id"] = userLoginResponse.ID
	claims["username"] = userLoginResponse.Username
	claims["email"] = userLoginResponse.Email
	claims["role_name"] = userLoginResponse.RoleName
	claims["user_create"] = userLoginResponse.CreatedAt
	claims["exp"] = expireTime

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	validToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return validToken, nil
}
