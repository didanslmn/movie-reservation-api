package router

import (
	genreHandler "github.com/didanslmn/movie-reservation-api/internal/genre/handler"
	genreRepository "github.com/didanslmn/movie-reservation-api/internal/genre/repository"
	genreRouter "github.com/didanslmn/movie-reservation-api/internal/genre/router" // Perbaikan di sini
	genreService "github.com/didanslmn/movie-reservation-api/internal/genre/service"
	"github.com/didanslmn/movie-reservation-api/internal/middleware"
	movieHandler "github.com/didanslmn/movie-reservation-api/internal/movie/handler"
	movieRepository "github.com/didanslmn/movie-reservation-api/internal/movie/repository"
	movieRouter "github.com/didanslmn/movie-reservation-api/internal/movie/router" // Perbaikan di sini
	movieService "github.com/didanslmn/movie-reservation-api/internal/movie/service"
	userHandler "github.com/didanslmn/movie-reservation-api/internal/user/handler"
	userRepository "github.com/didanslmn/movie-reservation-api/internal/user/repository"
	userRouter "github.com/didanslmn/movie-reservation-api/internal/user/router" // Perbaikan di sini
	userService "github.com/didanslmn/movie-reservation-api/internal/user/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB, jwtSecret string) *gin.Engine {
	r := gin.Default()

	// === User Setup ===
	userRepo := userRepository.NewUserRepository(db)
	userSvc := userService.NewUserService(userRepo, jwtSecret)
	userHdl := userHandler.NewUserHandler(userSvc)
	authMiddleware := middleware.JWTAuthMiddleware(jwtSecret)

	// === Genre Setup ===
	genreRepo := genreRepository.NewRepositoryGenre(db)
	genreSvc := genreService.NewGenreService(genreRepo)
	genreHdl := genreHandler.NewGenreHandler(genreSvc)

	// === Movie Setup ===
	movieRepo := movieRepository.NewRepositoryMovie(db)
	movieSvc := movieService.NewMovieService(movieRepo, genreRepo)
	movieHdl := movieHandler.NewMovieHandler(movieSvc)

	// Register routes
	api := r.Group("/api/v1")
	userRouter.RegisterUserRoutes(api, userHdl, authMiddleware)
	authGroup := api.Group("/")
	authGroup.Use(authMiddleware)
	genreRouter.RegisterGenreRoutes(authGroup, genreHdl, jwtSecret)
	movieRouter.RegisterMovieRoutes(authGroup, movieHdl, jwtSecret)

	return r
}
