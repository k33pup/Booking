package http

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	openapi "github.com/k33pup/Booking/internal/pkg/api/http/generated_api/generated_server/go"
	"github.com/k33pup/Booking/internal/pkg/config"
	"github.com/k33pup/Booking/internal/usecases"
	"net/http"
)

type Server struct {
	BookedRoomUseCase usecases.IBookedRoomUseCase
	Router            *gin.Engine
	OpenApiServer     *openapi.DefaultAPIService
	HttpServer        *http.Server
}

func NewServer(bookedRoomUseCase usecases.IBookedRoomUseCase) *Server {
	router := gin.Default()

	openApiServer := openapi.NewDefaultAPIService(bookedRoomUseCase)

	DefaultAPIController := openapi.NewDefaultAPIController(openApiServer)

	httpRouter := openapi.NewRouter(DefaultAPIController)

	serverConfig, _ := config.LoadServerConfig()

	serverAddress := fmt.Sprintf("%s:%s", serverConfig.Host, serverConfig.Port)

	httpServer := http.Server{
		Addr:    serverAddress,
		Handler: httpRouter,
	}

	return &Server{
		BookedRoomUseCase: bookedRoomUseCase,
		Router:            router,
		OpenApiServer:     openApiServer,
		HttpServer:        &httpServer,
	}
}

func (s *Server) Start() error {
	fmt.Println("Starting server...")
	err := s.HttpServer.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	fmt.Println("Stopping server...")
	err := s.HttpServer.Shutdown(ctx)
	if err != nil {
		return err
	}
	return nil
}
