package server

import (
	"be-lab/common/code"
	"be-lab/model/req"
	"github.com/gin-gonic/gin"
	"time"
)

func (s *Server) BookingList(c *gin.Context) {
	p := &req.ListReq{}
	if err := c.BindQuery(p); err != nil {
		code.Fail(c, err)
		return
	}
	list := s.Service.BookingList(c, p)
	code.Succ(c, list)
}

func (s *Server) BookingSave(c *gin.Context) {
	p := &req.BookingSave{}
	if err := c.BindJSON(p); err != nil {
		code.Fail(c, err)
		return
	}
	err := s.Service.BookingSave(c, p)
	if err != nil {
		code.Fail(c, err)
		return
	}
	code.Succ(c, nil)
}

func (s *Server) BookingJob() {
	go func() {
		for {
			s.Service.BookingJob()
			time.Sleep(time.Second * 5)
		}
	}()
}
