package router

import (
	"github.com/didanslmn/movie-reservation-api/internal/genre/handler"
	"github.com/didanslmn/movie-reservation-api/internal/genre/repository"
	"github.com/didanslmn/movie-reservation-api/internal/genre/service"
	mHandler "github.com/didanslmn/movie-reservation-api/internal/movie/handler"
	movieRepository "github.com/didanslmn/movie-reservation-api/internal/movie/repository"
	movieService "github.com/didanslmn/movie-reservation-api/internal/movie/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// === Genre setup ===
	genreRepo := repository.NewRepositoryGenre(db)
	genreService := service.NewGenreService(genreRepo)
	genreHandler := handler.NewGenreHandler(genreService)

	// === Movie setup ===
	movieRepo := movieRepository.NewRepositoryMovie(db)
	movieSvc := movieService.NewMovieService(movieRepo, genreRepo)
	movieHandler := mHandler.NewMovieHandler(movieSvc)

	// Genre routes
	genreRoutes := r.Group("/genres")
	{
		genreRoutes.POST("", genreHandler.CreateGenre)
		genreRoutes.GET("", genreHandler.GetAllGenres)
		genreRoutes.GET("/:id", genreHandler.GetGenre)
		genreRoutes.PUT("/:id", genreHandler.UpdateGenre)
		genreRoutes.DELETE("/:id", genreHandler.DeleteGenre)
	}

	// Movie routes
	movieRoutes := r.Group("/movies")
	{
		movieRoutes.POST("", movieHandler.CreateMovie)
		movieRoutes.GET("", movieHandler.GetAllMovies)
		movieRoutes.GET("/:id", movieHandler.GetMovie)
		movieRoutes.PUT("/:id", movieHandler.UpdateMovie)
		movieRoutes.DELETE("/:id", movieHandler.DeleteMovie)
		movieRoutes.GET("/genre/:id", movieHandler.GetMoviesByGenre)
	}

	return r
}
