package router

import (
	"github.com/didanslmn/movie-reservation-api/internal/user/handler"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(rg *gin.RouterGroup, h *handler.UserHandler, auth gin.HandlerFunc) {
	authGroup := rg.Group("/auth")
	{
		authGroup.POST("/register", h.Register)
		authGroup.POST("/login", h.Login)
	}

	userGroup := rg.Group("/user")
	{
		userGroup.GET("/profile", auth, h.Profile)
		userGroup.PUT("/profile", auth, h.UpdateProfile)
		userGroup.PUT("/change-password", auth, h.ChangePassword)
	}
}
