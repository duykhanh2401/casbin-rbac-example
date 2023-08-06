package controller

import (
	"net/http"

	"github.com/duykhanh2401/casbin-rbac-example/models"
	"github.com/duykhanh2401/casbin-rbac-example/rbac"
	"github.com/gin-gonic/gin"
)

func AddRole() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var data models.RoleData

		if err := ctx.ShouldBindJSON(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, "can not get data")
			return
		}

		if data.API == nil {
			data.API = append(data.API, "http://localhost:3000")
		}
		// thêm các policy rule
		for _, api := range data.API {
			for _, method := range data.Method {
				_, err := rbac.RBAC.AddPolicy(data.Role, api, method)
				if err != nil {
					ctx.JSON(http.StatusBadRequest, "ERROR")
					return
				}
			}
		}

		ctx.JSON(http.StatusOK, "success added role: "+data.Role)
	}
}

func AddRoleForUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var data models.UserRole
		if err := ctx.ShouldBindJSON(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, "can not get data")
			return
		}

		_, err := rbac.RBAC.AddGroupingPolicy(data.UserID, data.Role)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, "cans not add role for this user")
			return
		}

		ctx.JSON(http.StatusOK, "success")
	}
}

func DeleteRoleForUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var data models.UserRole
		if err := ctx.ShouldBindJSON(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, "can not get data")
			return
		}

		_, err := rbac.RBAC.RemoveGroupingPolicy(data.UserID, data.Role)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, "cans not delete role for this user")
			return
		}
	}
}

func AddAPIForRole() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var data models.RoleData

		if err := ctx.ShouldBindJSON(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, "can not get data")
			return
		}

		if data.Method == nil {
			data.Method = append(data.Method, "GET")
		}

		for _, api := range data.API {
			for _, method := range data.Method {
				_, err := rbac.RBAC.AddPolicy(data.Role, api, method)
				if err != nil {
					ctx.JSON(http.StatusBadRequest, "ERROR")
					return
				}
			}
		}
		ctx.JSON(http.StatusOK, "success")
	}
}

func DeleteAPIForRole() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var data models.RoleAPI

		if err := ctx.ShouldBindJSON(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, "can not get data")
			return
		}

		allAction := rbac.RBAC.GetAllActions()
		for _, action := range allAction {
			rbac.RBAC.RemovePolicy(data.Role, data.API, action)
		}

		filteredPolicy := rbac.RBAC.GetFilteredPolicy(0, data.Role)
		if (len(filteredPolicy)) == 0 {
			_, err := rbac.RBAC.AddPolicy(data.Role, "http://localhost:8080", "GET")
			if err != nil {
				ctx.JSON(http.StatusBadRequest, "ERROR")
				return
			}
		}
		ctx.JSON(http.StatusOK, "success")

	}
}

func AddRoleForAPI() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var data models.APIData

		if err := ctx.ShouldBindJSON(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, "can not get data")
			return
		}
		for _, role := range data.Role {
			for _, method := range data.Method {
				_, err := rbac.RBAC.AddPolicy(role, data.API, method)
				if err != nil {
					ctx.JSON(http.StatusBadRequest, "ERROR")
					return
				}
			}
		}

		ctx.JSON(http.StatusOK, "success")

	}
}

func DeleteRoleForAPI() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var data models.APIRole

		if err := ctx.ShouldBindJSON(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, "can not get data")
			return
		}

		allAction := rbac.RBAC.GetAllActions()
		for _, role := range data.Role {
			for _, action := range allAction {
				_, err := rbac.RBAC.RemovePolicy(role, data.API, action)
				if err != nil {
					ctx.JSON(http.StatusBadRequest, "ERROR")

					return
				}
			}
		}

		ctx.JSON(http.StatusOK, "delete success")

	}
}

func DeleteRole() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var data models.Role

		if err := ctx.ShouldBindJSON(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, "can not get data")
			return
		}

		ok, err := rbac.RBAC.DeleteRole(data.Role)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, "ERROR")

			return
		}
		// nếu không có role thì in ra
		if !ok {
			ctx.JSON(http.StatusBadRequest, "do not have role: "+data.Role)
			return
		}
	}
}
