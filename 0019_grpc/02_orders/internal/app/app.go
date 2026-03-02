package app

import (
	"context"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	grpcruntime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/igoroutine-courses/the_nature_of_microservices/orders/internal/controller"
	"github.com/igoroutine-courses/the_nature_of_microservices/orders/internal/repository"
	ordersv1 "github.com/igoroutine-courses/the_nature_of_microservices/orders/pkg/api/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

func Run(logger *zap.Logger) {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	repo := repository.NewInMemoryOrdersRepository()
	ctrl := controller.New(logger, repo)

	go runRest(ctx, logger)
	go runGrpc(logger, ctrl)

	<-ctx.Done()
	time.Sleep(time.Second * 3)
}

func runRest(ctx context.Context, logger *zap.Logger) {
	mux := grpcruntime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	address := "localhost:" + "50051"
	err := ordersv1.RegisterOrderServiceHandlerFromEndpoint(ctx, mux, address, opts)

	if err != nil {
		logger.Error("can not register grpc gateway", zap.Error(err))
		os.Exit(-1)
	}

	gatewayPort := ":" + "8080"
	logger.Info("gateway listening at port", zap.String("port", gatewayPort))

	if err = http.ListenAndServe(gatewayPort, mux); err != nil {
		logger.Error("gateway listen error", zap.Error(err))
	}
}

func runGrpc(logger *zap.Logger, orders ordersv1.OrderServiceServer) {
	port := ":" + "50051"
	lis, err := net.Listen("tcp", port)

	if err != nil {
		logger.Error("can not open tcp socket", zap.Error(err))
		os.Exit(-1)
	}

	s := grpc.NewServer()
	reflection.Register(s)

	ordersv1.RegisterOrderServiceServer(s, orders)

	logger.Info("grpc server listening at port", zap.String("port", port))

	if err = s.Serve(lis); err != nil {
		logger.Error("grpc server listen error", zap.Error(err))
	}
}
