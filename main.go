package main

import (
	"be-lab/common/infra"
	"be-lab/common/utils"
	"be-lab/di"
	"be-lab/router"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func main() {
	f, _ := os.OpenFile(utils.EnvDefault("GO_RUN_LOG", "access.log"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	infra.Init()
	s := di.InitServer()
	s.BookingJob()
	r := router.SetupRouter(s)
	err := r.Run(fmt.Sprintf("%s:%s", infra.Cfg.Server.Host, infra.Cfg.Server.Port))
	if err != nil {
		panic(err)
	}
}
