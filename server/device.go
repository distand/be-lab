package server

import (
	"be-lab/common/code"
	"be-lab/model/req"
	"github.com/gin-gonic/gin"
)

func (s *Server) DeviceList(c *gin.Context) {
	p := &req.ListReq{}
	if err := c.BindQuery(p); err != nil {
		code.Fail(c, err)
		return
	}
	list := s.Service.DeviceList(c, p)
	code.Succ(c, list)
}

func (s *Server) DeviceSave(c *gin.Context) {
	p := &req.DeviceSave{}
	if err := c.BindJSON(p); err != nil {
		code.Fail(c, err)
		return
	}
	err := s.Service.DeviceSave(c, p)
	if err != nil {
		code.Fail(c, err)
		return
	}
	code.Succ(c, nil)
}
