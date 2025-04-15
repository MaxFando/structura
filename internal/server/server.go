package server

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"

	structurav1 "github.com/MaxFando/structura/api/grpc/gen/go/structura/v1"
	"github.com/MaxFando/structura/internal/server/interceptor"
	v1 "github.com/MaxFando/structura/internal/server/structura/v1"
	"github.com/MaxFando/structura/pkg/log"
)

const (
	defaultGRPCAddress = "localhost"
	defaultGRPCPort    = "50051"
	defaultHTTPPort    = "8080"

	defaultMaxRecvMsgSize = 1024 * 1024 * 50 // 50 MB
	defaultMaxSendMsgSize = 1024 * 1024 * 50 // 50 MB
)

type Server struct {
	grpcServer *grpc.Server
	httpServer *http.Server
	logger     log.Logger

	grpcPort string
	httpPort string

	errors chan error
}

func NewServer(logger log.Logger, structuraServer *v1.StructuraServer) (*Server, error) {
	var err error

	srv := new(Server)

	srv.grpcPort = defaultGRPCPort
	srv.httpPort = defaultHTTPPort
	srv.errors = make(chan error, 1)
	srv.logger = logger

	srv.grpcServer = initGRPCServer(logger, structuraServer)
	srv.httpServer, err = initHTTPServer(srv.grpcPort, srv.httpPort)
	if err != nil {
		return nil, fmt.Errorf("ошибка при инициализации HTTP сервера: %w", err)
	}

	return srv, nil
}

func (s *Server) Serve(ctx context.Context) {
	grpcDone := make(chan struct{})
	httpDone := make(chan struct{})

	go func() {
		defer close(grpcDone)
		s.serveGRPC(ctx)
	}()

	go func() {
		defer close(httpDone)
		s.serveHTTP(ctx)
	}()

	go func() {
		select {
		case <-ctx.Done():
			s.Shutdown(ctx)
		case <-grpcDone:
		case <-httpDone:
		}
	}()

	<-grpcDone
	<-httpDone
}

func (s *Server) Shutdown(ctx context.Context) {
	s.logger.Info(ctx, "Завершение работы сервера")

	s.grpcServer.GracefulStop()

	if s.httpServer != nil {
		if err := s.httpServer.Shutdown(ctx); err != nil {
			s.logger.Error(ctx, "Ошибка при завершении работы HTTP сервера", "error", err)
		}
	}

	close(s.errors)
}

func (s *Server) Notify() <-chan error {
	errorsCh := make(chan error)

	go func() {
		for err := range s.errors {
			errorsCh <- err
		}
		close(errorsCh)
	}()

	return errorsCh
}

func (s *Server) serveHTTP(ctx context.Context) {
	mux := runtime.NewServeMux(
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.HTTPBodyMarshaler{
			Marshaler: &runtime.JSONPb{
				MarshalOptions: protojson.MarshalOptions{
					UseProtoNames:     true,
					EmitUnpopulated:   true,
					EmitDefaultValues: true,
				},
			},
		}),
	)
	//nolint:staticcheck
	opts := []grpc.DialOption{grpc.WithInsecure()}

	if err := structurav1.RegisterStructuraServiceHandlerFromEndpoint(ctx, mux, defaultGRPCAddress+":"+s.grpcPort, opts); err != nil {
		s.sendError(ctx, err)
		return
	}

	httpServer := &http.Server{
		Addr:    ":" + s.httpPort,
		Handler: mux,
	}
	s.httpServer = httpServer

	s.logger.Info(ctx, "Запуск HTTP сервера на порту "+s.httpPort)
	if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		s.sendError(ctx, err)
	}
}

func (s *Server) serveGRPC(ctx context.Context) {
	grpcListener, err := net.Listen("tcp", ":"+s.grpcPort)
	if err != nil {
		s.sendError(ctx, err)
		return
	}

	s.logger.Info(ctx, "Запуск gRPC сервера на порту "+s.grpcPort)
	if err := s.grpcServer.Serve(grpcListener); err != nil {
		s.sendError(ctx, err)
	}
}

func initGRPCServer(logger log.Logger, structuraServer *v1.StructuraServer) *grpc.Server {
	server := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			interceptor.PanicRecoveryUnaryInterceptor(logger),
		),
		grpc.MaxRecvMsgSize(defaultMaxRecvMsgSize),
		grpc.MaxSendMsgSize(defaultMaxSendMsgSize),
		grpc.KeepaliveParams(keepalive.ServerParameters{
			Time:    30 * time.Second,
			Timeout: 20 * time.Second,
		}),
	)

	structurav1.RegisterStructuraServiceServer(server, structuraServer)
	reflection.Register(server)

	return server
}

func initHTTPServer(grpcPort, httpPort string) (*http.Server, error) {
	ctx := context.Background()

	mux := runtime.NewServeMux(
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.HTTPBodyMarshaler{
			Marshaler: &runtime.JSONPb{
				MarshalOptions: protojson.MarshalOptions{
					UseProtoNames:     true,
					EmitUnpopulated:   true,
					EmitDefaultValues: true,
				},
			},
		}),
	)
	//nolint:staticcheck
	opts := []grpc.DialOption{grpc.WithInsecure()}

	if err := structurav1.RegisterStructuraServiceHandlerFromEndpoint(ctx, mux, defaultGRPCAddress+":"+grpcPort, opts); err != nil {
		return nil, fmt.Errorf("ошибка при регистрации HTTP сервера: %w", err)
	}

	httpServer := &http.Server{
		Addr:    ":" + httpPort,
		Handler: mux,
	}

	return httpServer, nil
}

func (s *Server) sendError(ctx context.Context, err error) {
	select {
	case s.errors <- err:
		s.logger.Error(ctx, "Error sent to channel", "error", err)
	default:
		s.logger.Error(ctx, "Error channel is full, error dropped", "error", err)
	}
}
