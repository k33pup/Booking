package http

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	openapi "github.com/k33pup/Booking.git/internal/pkg/api/http/generated_api/generated_server/go"
	"github.com/k33pup/Booking.git/internal/pkg/config"
	"github.com/k33pup/Booking.git/internal/usecases"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os/signal"
	"syscall"
	"time"
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

func (s *Server) Start() error {
	DefaultAPIController := openapi.NewDefaultAPIController(s.OpenApiServer)

	router := openapi.NewRouter(DefaultAPIController)

	serverConfig, err := config.LoadServerConfig()
	if err != nil {
		return err
	}

	serverAddress := fmt.Sprintf("%s:%s", serverConfig.Host, serverConfig.Port)

	httpServer := http.Server{
		Addr:    serverAddress,
		Handler: router,
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	group, ctx := errgroup.WithContext(ctx)
	group.Go(func() error {
		fmt.Println("Server started on %s\n", serverAddress)
		err := httpServer.ListenAndServe()
		if err != nil {
			return err
		}
		return nil
	})

	group.Go(func() error {
		<-ctx.Done()
		fmt.Println("Start to shut down server")
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
		defer cancel()

		err := httpServer.Shutdown(shutdownCtx)
		if err != nil {
			return err
		}
		fmt.Println("Server shutted down")
		return nil
	})

	err = group.Wait()
	if err != nil {
		return err
	}

	return nil
}
