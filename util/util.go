package util

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/teris-io/shortid"
	"time"
)

func GenShortId() (string, error) {
	return shortid.Generate()
}

func GetReqID(c *gin.Context) string {
	v, ok := c.Get("X-Request-Id")
	if !ok {
		return ""
	}
	if requestID, ok := v.(string); ok {
		return requestID
	}
	return ""
}

// GetExpiredTime get token expired time from env or config file.
func GetExpiredTime() time.Duration {
	day := viper.GetInt("token.expired")
	return time.Hour * 24 * time.Duration(day)
}
