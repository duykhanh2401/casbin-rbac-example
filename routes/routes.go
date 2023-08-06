package routes

import (
	"github.com/duykhanh2401/casbin-rbac-example/controller"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(app *gin.Engine) {
	app.POST("/add_role", controller.AddRole())
	app.POST("/add_role_for_user", controller.AddRoleForUser())
	app.POST("/delete_role_for_user", controller.DeleteRoleForUser())
	app.POST("/add_api_for_role", controller.AddAPIForRole())
	app.POST("/delete_api_for_role", controller.DeleteAPIForRole())
	app.POST("/add_role_for_api", controller.AddRoleForAPI())
	app.POST("/delete_role_for_api", controller.DeleteRoleForAPI())
	app.POST("/delete_role", controller.DeleteRole())
	app.GET("/admin", controller.CheckAuth(), controller.Notification())
}
