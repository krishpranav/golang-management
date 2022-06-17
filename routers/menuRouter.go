package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/krishpranav/golang-management/controllers"
)

func MenuRouters(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/menu", controllers.GetMenus())
	incomingRoutes.GET("/menu/:menus_id", controllers.GetMenu())
	incomingRoutes.POST("/menus", controllers.CreateMenu())
	incomingRoutes.PATCH("/menus/:menu_id", controllers.UpdateMenu())
}
