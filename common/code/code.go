package code

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
	"strings"
)

const (
	Ok = iota
	Err
)

const (
	Forbidden = "禁止操作"
	SysErr    = "系统错误"
	ParamErr  = "参数错误"
	FreezeErr = "账号被冻结"
	UserErr   = "用户不存在"
)

var errMsg = map[int]string{
	Ok:  "success",
	Err: "error",
}

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data,omitempty"`
}

func Succ(c *gin.Context, data any) {
	c.JSON(Ok, Response{
		Code: Ok,
		Msg:  errMsg[Ok],
		Data: data,
	})
	c.Abort()
}

func Fail(c *gin.Context, err error) {
	msg := validateFail(err)
	if msg == "" {
		msg = errMsg[Err]
	}
	c.JSON(http.StatusOK, Response{
		Code: Err,
		Msg:  msg,
	})
	c.Abort()
}

func validateFail(err error) string {
	msg := err.Error()
	if len(msg) > 4 && msg[0:4] == "Key:" {
		s := strings.Split(msg, "\n")
		re := regexp.MustCompile(`'(\w+)'`)
		var sr []string
		for _, v := range s {
			match := re.FindStringSubmatch(v)
			if len(match) > 1 {
				sr = append(sr, strings.ToLower(match[1]))
			}
		}
		msg = "Param Error: " + strings.Join(sr, ",")
	}
	return msg
}

func UnLogin(c *gin.Context) {
	c.Status(http.StatusUnauthorized)
	c.Abort()
}

func UnAuth(c *gin.Context) {
	c.Status(http.StatusForbidden)
	c.Abort()
}
