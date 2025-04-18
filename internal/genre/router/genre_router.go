package router

import (
	"github.com/didanslmn/movie-reservation-api/internal/genre/handler"
	"github.com/didanslmn/movie-reservation-api/internal/middleware"
	"github.com/didanslmn/movie-reservation-api/internal/user/model"
	"github.com/gin-gonic/gin"
)

func RegisterGenreRoutes(rg *gin.RouterGroup, h *handler.GenreHandler, jwtSecret string) {
	genre := rg.Group("/genres")

	// Admin routes
	adminRoutes := genre.Group("/")
	adminRoutes.Use(middleware.JWTAuthMiddleware(jwtSecret))
	adminRoutes.Use(middleware.RoleBasedAccess(model.RoleAdmin))
	{
		adminRoutes.POST("/", h.CreateGenre)
		adminRoutes.PUT("/:id", h.UpdateGenre)
		adminRoutes.DELETE("/:id", h.DeleteGenre)
	}

	// user , admin routes
	publicRoutes := genre.Group("/")
	publicRoutes.Use(middleware.JWTAuthMiddleware(jwtSecret))
	publicRoutes.Use(middleware.RoleBasedAccess(model.RoleUser, model.RoleAdmin))
	{
		publicRoutes.GET("/:id", h.GetGenre)
		publicRoutes.GET("/", h.GetAllGenres)
	}
}
