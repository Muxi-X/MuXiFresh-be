package token

import (
	"errors"
	"fmt"
	"github.com/MuXiFresh-be/pkg/errno"
	"github.com/dgrijalva/jwt-go"
	"time"

	"github.com/MuXiFresh-be/log"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	jwtKey string

	ErrTokenInvalid = errors.New("the token is invalid")
	ErrTokenExpired = errors.New("the token is expired")
)

// getJwtKey get jwtKey.
func getJwtKey() string {
	if jwtKey == "" {
		jwtKey = viper.GetString("jwt_secret")
	}
	return jwtKey
}

// TokenPayload is a required payload when generates token.
type TokenPayload struct {
	Email   string        `json:"email"`
	Expired time.Duration `json:"expired"` // 有效时间（nanosecond）
}

// TokenResolve means returned payload when resolves token.
type TokenResolve struct {
	Email     string `json:"email"`
	ExpiresAt int64  `json:"expires_at"` // 过期时间（时间戳，10位）
}

// GenerateToken generates token.
func GenerateToken(email string, day time.Duration) (string, error) {
	claims := &TokenClaims{
		Email:     email,
		ExpiresAt: time.Now().Unix() + int64(day.Seconds()),
	}

	encodedString := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := encodedString.SignedString([]byte(getJwtKey()))
	if err != nil {
		return err.Error(), errno.ErrFormToken
	}

	return token, nil
}

// ResolveToken resolves token.
func ResolveToken(tokenStr string) (*TokenResolve, error) {
	claims := &TokenClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(getJwtKey()), nil
	})

	if err != nil {
		log.Error("Token parsing failed because of an internal error", zap.String("cause", err.Error()))
		return nil, err
	}

	if !token.Valid {
		log.Error("Token parsing failed; the token is invalid.")
		return nil, ErrTokenInvalid
	}

	t := &TokenResolve{
		Email:     claims.Email,
		ExpiresAt: claims.ExpiresAt,
	}
	return t, nil
}
