package controller

import (
	"context"
	"errors"

	"github.com/igoroutine-courses/the_nature_of_microservices/orders/internal/repository"
	ordersv1 "github.com/igoroutine-courses/the_nature_of_microservices/orders/pkg/api/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a *api) GetOrder(ctx context.Context, req *ordersv1.GetOrderRequest) (*ordersv1.GetOrderResponse, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "validation failed: %v", err)
	}

	order, err := a.ordersRepository.GetOrder(ctx, req.GetOrderId())

	if err != nil {
		if errors.Is(err, repository.ErrOrderNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}

		a.logger.Error("get order failed", zap.Error(err))
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &ordersv1.GetOrderResponse{Order: order}, nil
}
