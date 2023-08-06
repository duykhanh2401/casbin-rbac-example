package main

import (
	"github.com/duykhanh2401/casbin-rbac-example/rbac"
	"github.com/duykhanh2401/casbin-rbac-example/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	ginMode := gin.DebugMode
	gin.SetMode(ginMode)

	rbac.Setup()
	engine := gin.New()
	routes.RegisterRoutes(engine)
	engine.Run(":3000")

}
