package router

import (
	"github.com/didanslmn/movie-reservation-api/internal/middleware"
	"github.com/didanslmn/movie-reservation-api/internal/movie/handler"
	"github.com/didanslmn/movie-reservation-api/internal/user/model"
	"github.com/gin-gonic/gin"
)

func RegisterMovieRoutes(rg *gin.RouterGroup, h *handler.MovieHandler, jwtSecret string) {
	movie := rg.Group("/movies")

	// User & Admin routes
	publicRoutes := movie.Group("/")
	publicRoutes.Use(middleware.JWTAuthMiddleware(jwtSecret)) // pastikan middleware JWT dijalankan dulu
	publicRoutes.Use(middleware.RoleBasedAccess(model.RoleUser, model.RoleAdmin))
	{
		publicRoutes.GET("/:id", h.GetMovie)
		publicRoutes.GET("/", h.GetAllMovies)
		publicRoutes.GET("/genre/:genre_id", h.GetMoviesByGenre)
	}

	// Admin-only routes
	adminRoutes := movie.Group("/")
	adminRoutes.Use(middleware.JWTAuthMiddleware(jwtSecret))
	adminRoutes.Use(middleware.RoleBasedAccess(model.RoleAdmin))
	{
		adminRoutes.POST("/", h.CreateMovie)
		adminRoutes.PUT("/:id", h.UpdateMovie)
		adminRoutes.DELETE("/:id", h.DeleteMovie)
	}
}
