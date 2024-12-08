package middleware

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"katalisRobo/component-store/data/current"
	"katalisRobo/component-store/helper"
	"net/http"
	"strings"
)

func WithAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			unauthorizedResponse(ctx)
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			unauthorizedResponse(ctx)
			return
		}

		auths := strings.Split(authHeader, " ")
		if len(auths) != 2 {
			unauthorizedResponse(ctx)
			return
		}

		data, err := helper.DecryptJWT(auths[1])
		if err != nil {
			unauthorizedResponse(ctx)
			return
		}

		authUser := current.AuthUser{
			UserEmail: fmt.Sprintf("%v", data["user_email"]),
			Role:      fmt.Sprintf("%v", data["role"]),
		}

		ctxCustomerEmail := context.WithValue(ctx.Request.Context(), "authUser", authUser)
		ctx.Request = ctx.Request.WithContext(ctxCustomerEmail)
		ctx.Next()
	}
}

func unauthorizedResponse(ctx *gin.Context) {
	ctx.JSON(http.StatusUnauthorized, gin.H{
		"message": "unauthorized",
	})
	ctx.Abort()
	return
}
