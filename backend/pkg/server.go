package pkg

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"recipinder/pkg/controller"
)

var (
	Server *gin.Engine
)

func init() {

	Server = gin.Default()
	Server.Use(cors.Default())

	//ユーザ関連
	Server.POST("/sign/up", controller.SignUp())

}
