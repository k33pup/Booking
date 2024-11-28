package http

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	openapi "github.com/k33pup/Booking.git/internal/booking/api/generated/go"
	"github.com/k33pup/Booking.git/internal/booking/config"
	"github.com/k33pup/Booking.git/internal/booking/usecases"
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

	serverAddress := fmt.Sprintf("%s:%d", serverConfig.Host, serverConfig.Port)

	go func() {
		http.ListenAndServe(serverAddress, router)
	}()

	return nil
}