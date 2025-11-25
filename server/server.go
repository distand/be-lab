package server

import (
	"be-lab/common/code"
	"be-lab/service"
	"github.com/gin-gonic/gin"
)

type Server struct {
	Service *service.Service
}

func NewServer(s *service.Service) *Server {
	return &Server{
		Service: s,
	}
}

func (s *Server) Index(c *gin.Context) {
	code.Succ(c, s.Service.Index(c))
}
