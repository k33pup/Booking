package http

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/k33pup/Booking.git/internal/usecases"
	"github.com/k33pup/Booking.git/pkg/config"
	openapi "github.com/k33pup/Booking.git/pkg/generated_api/generated/go"
	"net/http"
)

type Server struct {
	BookedRoomUseCase usecases.IBookedRoomUseCase
	Router            *gin.Engine
	OpenApiServer     *openapi.DefaultAPIService
}

func NewServer(bookedRoomUseCase usecases.IBookedRoomUseCase) *Server {
	router := gin.Default()

	openApiServer := openapi.NewDefaultAPIService(bookedRoomUseCase)

	return &Server{
		BookedRoomUseCase: bookedRoomUseCase,
		Router:            router,
		OpenApiServer:     openApiServer,
	}
}

func (s *Server) Start(ctx context.Context) error {
	DefaultAPIController := openapi.NewDefaultAPIController(s.OpenApiServer)

	router := openapi.NewRouter(DefaultAPIController)

	serverConfig, err := config.LoadServerConfig()
	if err != nil {
		return err
	}

	serverAddress := fmt.Sprintf("%s:%s", serverConfig.Host, serverConfig.Port)

	fmt.Println("Server started on %s\n", serverAddress)
	if err = http.ListenAndServe(serverAddress, router); err != nil && err != http.ErrServerClosed {
		return err
	}

	// TODO graceful shut down

	return nil
}
