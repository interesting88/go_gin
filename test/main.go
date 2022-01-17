package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_gin/config"
	"go_gin/router"
)

func main(){
	gin.SetMode(gin.ReleaseMode)
	engine:=gin.Default()
	router.InitRouter(engine)
	engine.Run(config.PORT)
}
func init()  {
	fmt.Println("hello word", "", res, c)

}
