package usecase

import (
	"fmt"
	"os"
	"time"

	"github.com/Uwasaru/Sayka/domain/repository"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type JwtUsecase struct {
	ar repository.AuthRepository
}

type IJwtUsecase interface {
	GetUserIdFromJwtToken(ctx *gin.Context) (string, error)
}

func NewJwtUsecase(ar repository.AuthRepository) IJwtUsecase {
	return &JwtUsecase{ar: ar}
}

func (ju *JwtUsecase) GetUserIdFromJwtToken(ctx *gin.Context) (string, error) {
	tokenString := ctx.Request.Header.Get("jwt")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("rawPrivKey")), nil
	})
	if err != nil {
		return "", err
	}
	session := ""
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		session, ok = claims["session"].(string)
		if !ok {
			return "", fmt.Errorf("kata asa-sion")
		}
	} else {
		return "", fmt.Errorf("token empty")
	}
	expiry, err := ju.ar.GetExpiryFromSession(ctx, session)
	if err != nil || expiry.Before(time.Now()) {
		return "", fmt.Errorf("GetExpiryFromSession: %w", err)
	}
	userId, err := ju.ar.GetUserIdFromSession(ctx, session)
	if err != nil {
		return "", fmt.Errorf("GetUserIdFromSession: %w", err)
	}
	return userId, nil
}