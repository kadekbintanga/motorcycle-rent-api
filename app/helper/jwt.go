package helper

import (
	"errors"
	"motorcycle-rent-api/app/model"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type AdminAuthClaims struct {
	ID      uint   `json:"id"`
	UUID    string `json:"uuid"`
	Expired int64  `json:"exp"`
	Email   string `json:"email"`
	Name    string `json:"name"`
}

func GenerateJWTAdmin(admin *model.Admin, secretKey string, expiredDuration time.Duration) (string, error) {
	tokenClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":   time.Now().Add(expiredDuration).Unix(),
		"id":    admin.ID,
		"uuid":  admin.UUID.String(),
		"email": admin.Email,
		"name":  admin.Name,
	})

	signedToken, err := tokenClaim.SignedString([]byte(secretKey))
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

func ValidateJWTAdmin(tokenString, secret string, db *gorm.DB) (*AdminAuthClaims, error) {
	jwtSecret := []byte(secret)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token")
		}
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	if exp, ok := claims["exp"].(float64); ok {
		if time.Now().Unix() > int64(exp) {
			return nil, errors.New("token expired")
		}
	}

	var adminFound *model.Admin
	err = db.Preload("Role").Preload("Role.Features").Where("email = ?", strings.ToLower(claims["email"].(string))).First(&adminFound).Error
	if err != nil || adminFound == nil {
		return nil, errors.New("unable to get user data")
	}

	formattedClaims := AdminAuthClaims{
		ID:      uint(claims["id"].(float64)),
		UUID:    claims["uuid"].(string),
		Expired: int64(claims["exp"].(float64)),
		Email:   claims["email"].(string),
		Name:    claims["name"].(string),
	}

	return &formattedClaims, nil
}
