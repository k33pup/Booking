package app

import (
	"context"
	"errors"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/k33pup/Booking/hotel_svc/internal/pkg/config"
	"github.com/k33pup/Booking/hotel_svc/internal/pkg/repository/postgres"
	impl "github.com/k33pup/Booking/hotel_svc/internal/pkg/transport/grpc"
	apiv1pb "github.com/k33pup/Booking/hotel_svc/internal/pkg/transport/grpc/generated"
	"github.com/k33pup/Booking/hotel_svc/internal/usecases"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log/slog"
	"net"
	"net/http"
	"os"
	"sync"
)

type HotelService struct {
	grpcServer *grpc.Server
	grpcListen *net.Listener
	httpServer *http.Server
	waitGroup  *sync.WaitGroup

	cfg *config.Config
	log *slog.Logger
}

func NewHotelService(cfg *config.Config) *HotelService {
	logFile, err := os.OpenFile(cfg.GetLogPath(), os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	logH := slog.NewJSONHandler(logFile, &slog.HandlerOptions{})
	return &HotelService{
		cfg: cfg,
		log: slog.New(logH),
	}
}

func (h *HotelService) Init(ctx context.Context) error {
	// Initialization of grpc, postgres, kafka and other dependencies
	dsn := h.cfg.GetDSN()
	repo, err := postgres.NewHotelRepository(dsn)
	useCase := usecases.NewHotelUseCase(repo)
	server := impl.NewServer(useCase, h.log)

	h.waitGroup = &sync.WaitGroup{}

	grpcListen, err := net.Listen("tcp", h.cfg.GetGRCPAddress())
	h.grpcListen = &grpcListen

	if err != nil {
		h.log.InfoContext(ctx, "failed to create gRPC listener: %w", err)
		return fmt.Errorf("failed to create gRPC listener: %w", err)
	}

	h.grpcServer = grpc.NewServer()
	apiv1pb.RegisterHotelServiceServer(h.grpcServer, server)

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	err = apiv1pb.RegisterHotelServiceHandlerFromEndpoint(ctx, mux, h.cfg.GetGRCPAddress(), opts)
	if err != nil {
		h.log.InfoContext(ctx, "failed to register gRPC-Gateway: %w", err)
		return fmt.Errorf("failed to register gRPC-Gateway: %w", err)
	}

	h.httpServer = &http.Server{
		Addr:    h.cfg.GetHTTPAddress(),
		Handler: mux,
	}
	h.log.InfoContext(ctx, "internal service has been initialized successfully")
	return nil
}

func (h *HotelService) Start(ctx context.Context) error {
	cfg := config.NewConfig()
	h.waitGroup.Add(1)
	go func() {
		defer h.waitGroup.Done()
		if err := h.grpcServer.Serve(*h.grpcListen); err != nil && !errors.Is(err, grpc.ErrServerStopped) {
			h.log.InfoContext(ctx, "gRPC server error: ", err.Error())
		}
	}()
	h.log.InfoContext(ctx, fmt.Sprintf("gRPC server started on address %s", cfg.GetGRCPAddress()))

	h.waitGroup.Add(1)
	go func() {
		defer h.waitGroup.Done()
		if err := h.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			h.log.InfoContext(ctx, "HTTP server error: ", err.Error())
		}
	}()
	h.log.InfoContext(ctx, fmt.Sprintf("HTTP server started on address %s", cfg.GetHTTPAddress()))
	return nil
}

func (h *HotelService) Stop(ctx context.Context) error {
	if h.httpServer != nil {
		h.log.InfoContext(ctx, "Shutting down HTTP server...")
		if err := h.httpServer.Shutdown(ctx); err != nil {
			h.log.InfoContext(ctx, "Error shutting down HTTP server: ", err.Error())
		}
	}

	if h.grpcServer != nil {
		h.log.InfoContext(ctx, "Shutting down gRPC server...")
		h.grpcServer.GracefulStop()
	}
	return nil
}

func (h *HotelService) Wait(ctx context.Context) {
	h.waitGroup.Wait()
	h.log.InfoContext(ctx, "internal service has stopped successfully")
}
