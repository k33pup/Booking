package http

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/k33pup/Booking/payment_svc/internal/pkg/config"
	openapi "github.com/k33pup/Booking/payment_svc/internal/pkg/transport/http/generated_api/generated_server/go"
	"github.com/k33pup/Booking/payment_svc/internal/usecases"
	"net/http"
)

type Server struct {
	PaymentUseCase usecases.IPaymentUseCase
	Router         *gin.Engine
	OpenApiServer  *openapi.DefaultAPIService
	HttpServer     *http.Server
}

func NewServer(paymentUseCase usecases.IPaymentUseCase) *Server {
	router := gin.Default()

	openApiServer := openapi.NewDefaultAPIService(paymentUseCase)

	DefaultAPIController := openapi.NewDefaultAPIController(openApiServer)

	httpRouter := openapi.NewRouter(DefaultAPIController)

	serverConfig, _ := config.LoadServerConfig()

	serverAddress := fmt.Sprintf("%s:%s", serverConfig.Host, serverConfig.Port)

	httpServer := http.Server{
		Addr:    serverAddress,
		Handler: httpRouter,
	}

	return &Server{
		PaymentUseCase: paymentUseCase,
		Router:         router,
		OpenApiServer:  openApiServer,
		HttpServer:     &httpServer,
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
