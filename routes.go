package main

import (
	"github.com/1tsuki/gl-market/interfaces/api/server/handler"
	"github.com/labstack/echo/middleware"
	"github.com/1tsuki/gl-market/domain/service"
)

func init() {
	config := middleware.JWTConfig{
		Claims:     &service.JwtCustomClaims{},
		SigningKey: []byte("secret"), // TODO: replace with some secure key
	}

	auth := e.Group("/auth")
	{
		auth.POST("/signUp", handler.SignUp)
		auth.POST("/signIn", handler.SignIn)
	}

	api := e.Group("/api")
	{
		// APIs under /api must be authorized
		api.Use(middleware.JWTWithConfig(config))
		v1 := api.Group("/v1")
		{
			users := v1.Group("/users")
			{
				users.GET("/:userID", nil)
				users.PUT("/:userID", nil)
			}
		}
	}
}
