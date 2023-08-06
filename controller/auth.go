package controller

import (
	"fmt"
	"net/http"

	"github.com/duykhanh2401/casbin-rbac-example/models"
	"github.com/duykhanh2401/casbin-rbac-example/rbac"
	"github.com/gin-gonic/gin"
)

func CheckAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var data models.User

		if err := ctx.ShouldBindJSON(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, "can not get data")
			ctx.Abort()
			return
		}

		path := "localhost:3000" + ctx.Request.URL.Path
		fmt.Println(path)
		ok, err := rbac.RBAC.Enforce(data.UserID, path, "GET")
		if !ok || err != nil {
			ctx.JSON(http.StatusBadRequest, "ERROR")
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}

func Notification() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Hello")
	}
}
