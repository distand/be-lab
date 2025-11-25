package server

import (
	"be-lab/common"
	"be-lab/common/code"
	"be-lab/common/utils"
	"be-lab/model/req"
	"errors"
	"github.com/gin-gonic/gin"
)

func (s *Server) Login(c *gin.Context) {
	openid := c.Query("openid")
	if len(openid) < 16 {
		code.Fail(c, errors.New(code.ParamErr))
		return
	}
	u, err := s.Service.Login(c, openid)
	if err != nil {
		code.Fail(c, err)
		return
	}
	code.Succ(c, u)
}

func (s *Server) UserList(c *gin.Context) {
	p := &req.ListReq{}
	if err := c.BindQuery(p); err != nil {
		code.Fail(c, err)
		return
	}
	list := s.Service.UserList(c, p)
	code.Succ(c, list)
}

func (s *Server) UserInfo(c *gin.Context) {
	u, _ := s.Service.UserInfo(c)
	code.Succ(c, u)
}

func (s *Server) UserSave(c *gin.Context) {
	p := &req.UserSave{}
	if err := c.BindJSON(p); err != nil {
		code.Fail(c, err)
		return
	}
	err := s.Service.UserSave(c, p)
	if err != nil {
		code.Fail(c, err)
		return
	}
	code.Succ(c, nil)
}

func (s *Server) CheckLogin(c *gin.Context) {
	c.Set(common.KeyOpenid, c.Param(common.KeyOpenid))
	user, err := s.Service.UserInfo(c)
	if err != nil || user == nil {
		code.UnLogin(c)
		return
	}
	c.Set(common.KeyUid, user.ID)
	c.Set(common.KeyRole, user.Role)
	c.Next()
}

func (s *Server) CheckAdmin(c *gin.Context) {
	if !utils.IsAdmin(c) {
		code.UnAuth(c)
		return
	}
	c.Next()
}
