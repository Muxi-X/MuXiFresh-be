package auth

import (
	"errors"
	"github.com/MuXiFresh-be/log"
	"github.com/MuXiFresh-be/model"
	userModel "github.com/MuXiFresh-be/model/user"
	"github.com/MuXiFresh-be/pkg/token"
	"github.com/MuXiFresh-be/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	// ErrMissingHeader means the `Authorization` header was empty.
	ErrMissingHeader = errors.New("the length of the `Authorization` header is zero")
	// ErrTokenInvalid means the token is invalid.
	ErrTokenInvalid = errors.New("the token is invalid")
)

// Context is the context of the JSON web token.
type Context struct {
	Email     string
	Role      uint32
	Id        uint32
	ExpiresAt int64 // 过期时间（时间戳，10位）
}

// Parse parses the token, and returns the context if the token is valid.
func Parse(tokenString string) (*Context, error) {
	t, err := token.ResolveToken(tokenString)
	if err != nil {
		return nil, err
	}

	var db *model.Database
	var user userModel.UserModel
	db.Init()
	defer db.Close()
	model.DB.Self.Where("email = ?", t.Email).First(&user)
	return &Context{
		Email:     t.Email,
		ExpiresAt: t.ExpiresAt,
		Role:      user.Role,
		Id:        user.Id,
	}, nil
}

// ParseRequest gets the token from the header and
// pass it to the Parse function to parse the token.
func ParseRequest(c *gin.Context) (*Context, error) {
	header := c.Request.Header.Get("Authorization")
	if len(header) == 0 {
		return nil, ErrMissingHeader
	}

	// check if in blacklist
	exist, err := service.CheckInBlacklist(header)
	if err != nil {
		log.Error("CheckInBlacklist error", zap.String("cause", err.Error()))
		return nil, err
	} else if exist {
		return nil, ErrTokenInvalid
	}

	return Parse(header)
}
