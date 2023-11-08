package jwt

import (
	"fmt"
	"os"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/golang-jwt/jwt/v5"
)

func ExtractToken(tokenString string) (*web.UserLoginResponse, error) {

	type CustomClaims struct {
		Email    string `json:"email"`
		Username string `json:"username"`
		Role     string `json:"role"`
		jwt.RegisteredClaims
	}

	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		extractedToken := &web.UserLoginResponse{
			Username: claims.Username,
			Email:    claims.Email,
			Role:     web.Role(claims.Role),
		}
		return extractedToken, nil
	}

	return nil, fmt.Errorf("Invalid token")

}
