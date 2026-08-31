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
	"github.com/igoroutine-courses/microservices.ecommerce.loms/internal/controller"
	lomsctrl "github.com/igoroutine-courses/microservices.ecommerce.loms/internal/controller/loms"
	productctrl "github.com/igoroutine-courses/microservices.ecommerce.loms/internal/controller/product"
	stocksctrl "github.com/igoroutine-courses/microservices.ecommerce.loms/internal/controller/stocks"
	stocksrepo "github.com/igoroutine-courses/microservices.ecommerce.loms/internal/repository/stocks"
	db "github.com/igoroutine-courses/microservices.ecommerce.loms/migrations"
	lomssrv "github.com/igoroutine-courses/microservices.ecommerce.pkg/generated/loms/api/loms/v1"
	productsrv "github.com/igoroutine-courses/microservices.ecommerce.pkg/generated/loms/api/product/v1"
	stockssrv "github.com/igoroutine-courses/microservices.ecommerce.pkg/generated/loms/api/stocks/v1"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/igoroutine-courses/microservices.ecommerce.loms/internal/config"
	orderrepo "github.com/igoroutine-courses/microservices.ecommerce.loms/internal/repository/order"
	productrepo "github.com/igoroutine-courses/microservices.ecommerce.loms/internal/repository/product"
	"github.com/igoroutine-courses/microservices.ecommerce.loms/internal/usecase/loms"
	"github.com/igoroutine-courses/microservices.ecommerce.loms/internal/usecase/product"
	"github.com/igoroutine-courses/microservices.ecommerce.loms/internal/usecase/stocks"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

func Run(logger *zap.Logger, cfg *config.Config) {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	dbPool, err := pgxpool.New(ctx, cfg.ConstructPostgresURL())

	if err != nil {
		logger.Error("can not create pgxpool", zap.Error(err))
		os.Exit(-1)
	}

	defer dbPool.Close()

	db.SetupPostgres(dbPool, logger)

	orderRepository := orderrepo.NewPostgresRepository(dbPool)
	stocksRepository := stocksrepo.NewPostgresRepository(dbPool)

	//orderRepository := orderrepo.NewInMemoryRepository()
	productRepository := productrepo.NewInMemoryRepository()

	lomsService := loms.NewLomsService(orderRepository, stocksRepository)
	productService := product.NewProductService(productRepository)
	stocksService := stocks.NewStocksService(stocksRepository)

	lomsServer := lomsctrl.NewLomsServer(lomsService)
	productServer := productctrl.NewProductServer(productService)
	stocksServer := stocksctrl.NewStocksServer(stocksService)

	ctrl := controller.New(lomsServer, productServer, stocksServer)

	go runRest(ctx, logger, cfg)
	go runGrpc(logger, cfg, ctrl)

	<-ctx.Done()
	time.Sleep(time.Second * 3)
}

func runRest(ctx context.Context, logger *zap.Logger, cfg *config.Config) {
	mux := grpcruntime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	address := "localhost:" + cfg.GRPC.Port

	err := lomssrv.RegisterLomsHandlerFromEndpoint(ctx, mux, address, opts)

	if err != nil {
		logger.Error("can not register grpc gateway", zap.Error(err))
		os.Exit(-1)
	}

	err = productsrv.RegisterProductServiceHandlerFromEndpoint(ctx, mux, address, opts)

	if err != nil {
		logger.Error("can not register grpc gateway", zap.Error(err))
		os.Exit(-1)
	}

	err = stockssrv.RegisterStocksHandlerFromEndpoint(ctx, mux, address, opts)

	if err != nil {
		logger.Error("can not register grpc gateway", zap.Error(err))
		os.Exit(-1)
	}

	gatewayPort := ":" + cfg.GRPC.GatewayPort
	logger.Info("gateway listening at port", zap.String("port", gatewayPort))

	handler := corsHandler(mux)
	if err = http.ListenAndServe(gatewayPort, handler); err != nil {
		logger.Error("gateway listen error", zap.Error(err))
	}
}

func runGrpc(logger *zap.Logger, cfg *config.Config, ctrl *controller.API) {
	port := ":" + cfg.GRPC.Port
	lis, err := net.Listen("tcp", port)

	if err != nil {
		logger.Error("can not open tcp socket", zap.Error(err))
		os.Exit(-1)
	}

	s := grpc.NewServer()
	reflection.Register(s)

	lomssrv.RegisterLomsServer(s, ctrl)
	productsrv.RegisterProductServiceServer(s, ctrl)
	stockssrv.RegisterStocksServer(s, ctrl)

	logger.Info("grpc server listening at port", zap.String("port", port))

	if err = s.Serve(lis); err != nil {
		logger.Error("grpc server listen error", zap.Error(err))
	}
}

func corsHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if origin == "" {
			origin = "http://localhost:5173"
		}
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept")
		w.Header().Set("Access-Control-Max-Age", "86400")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}
