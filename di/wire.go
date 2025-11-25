//go:build wireinject
// +build wireinject

package di

import (
	"be-lab/dal"
	"be-lab/server"
	"be-lab/service"
	"github.com/google/wire"
)

func InitServer() *server.Server {
	wire.Build(server.NewServer, service.NewService, dal.NewDal)
	return &server.Server{}
}
