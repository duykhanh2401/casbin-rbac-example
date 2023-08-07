package controller

import (
	"net/http"

	"github.com/duykhanh2401/casbin-rbac-example/models"
	"github.com/duykhanh2401/casbin-rbac-example/rbac"
	"github.com/gin-gonic/gin"
)

func AddRoleCategory() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var data models.RoleCategory

		if err := ctx.ShouldBindJSON(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, "can not get data")
			ctx.Abort()
			return
		}

		path := "http://localhost:3000/category/:slug"
		for _, role := range data.Role {
			for _, method := range data.Method {
				_, err := rbac.RBAC.AddNamedPolicy("p2", role, path, method)
				if err != nil {
					ctx.JSON(http.StatusBadRequest, "ERROR")
					return
				}
			}
		}
	}
}
