package routes

import (
	"tps/controllers"

	"github.com/gin-gonic/gin"
)


func AppRoutes(router *gin.Engine){
	router.POST("/tokens",controllers.Token)

}