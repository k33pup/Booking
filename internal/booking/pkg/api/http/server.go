package http

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	openapi "github.com/k33pup/Booking.git/internal/booking/api/generated/go"
	"github.com/k33pup/Booking.git/internal/booking/config"
	"github.com/k33pup/Booking.git/internal/booking/usecases"
	"log"
	"net/http"
	"sync"
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

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		log.Printf("Server started on %s\n", serverAddress)
		if err = http.ListenAndServe(serverAddress, router); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server: %v\n", err)
		}
	}()

	wg.Wait()

	// TODO graceful shut down

	return nil
}
