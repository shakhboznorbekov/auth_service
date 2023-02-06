package api

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/auth_service/api/docs"
	"github.com/auth_service/api/handler"
	"github.com/auth_service/config"
	"github.com/auth_service/pkg/helper"
	"github.com/auth_service/storage"
)

func SetUpApi(cfg *config.Config, r *gin.Engine, storage storage.StorageI) {

	handlerV1 := handler.NewHandlerV1(cfg, storage)

	r.Use(customCORSMiddleware())


	v2 := r.Group("/v2")

	r.POST("/login", handlerV1.Login)
	r.POST("/register", handlerV1.Register)

	
	v2.Use(checkTokenClient())

	r.POST("/user", handlerV1.CreateUser)
	r.GET("/user/:id", handlerV1.GetUserById)
	v2.GET("/user", handlerV1.GetUserList)
	// r.PUT("/user/:id", handlerV1.UpdateUser)
	// r.DELETE("/user/:id", handlerV1.DeleteUser)

	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}

func checkTokenClient() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if _, ok := ctx.Request.Header["Authorization"]; ok {
			_, err := helper.ExtractClaims(ctx.Request.Header["Authorization"][0], config.Load().AuthSecretKey)
			if err != nil {
				ctx.AbortWithError(http.StatusForbidden, errors.New("not found password"))
				return
			} else {
				ctx.Next()
			}
		}
	}
}

func customCORSMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE, HEAD")
		c.Header("Access-Control-Allow-Headers", "Platform-Id, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Max-Age", "3600")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
