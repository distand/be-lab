package utils

import (
	"be-lab/common"
	"github.com/gin-gonic/gin"
	"math/rand"
	"os"
	"strings"
	"time"
)

const (
	letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	// 6 bits to represent a letter index
	letterIdBits = 6
	// All 1-bits as many as letterIdBits
	letterIdMask = 1<<letterIdBits - 1
	letterIdMax  = 63 / letterIdBits
)

func RandStr(n int) string {
	var src = rand.NewSource(time.Now().UnixNano())
	sb := strings.Builder{}
	sb.Grow(n)
	// A rand.Int63() generates 63 random bits, enough for letterIdMax letters!
	for i, cache, remain := n-1, src.Int63(), letterIdMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdMax
		}
		if idx := int(cache & letterIdMask); idx < len(letters) {
			sb.WriteByte(letters[idx])
			i--
		}
		cache >>= letterIdBits
		remain--
	}
	return sb.String()
}

func EnvDefault(key, def string) string {
	if v, ok := os.LookupEnv(key); ok {
		return v
	}
	return def
}

func Openid(c *gin.Context) string {
	value, exists := c.Get(common.KeyOpenid)
	if !exists {
		return ""
	}
	return value.(string)
}

func Uid(c *gin.Context) int32 {
	value, exists := c.Get(common.KeyUid)
	if !exists {
		return 0
	}
	return value.(int32)
}

func IsAdmin(c *gin.Context) bool {
	value, exists := c.Get(common.KeyRole)
	if !exists || value.(int32) < common.AuthAdmin {
		return false
	}
	return true
}
