package controller

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/igoroutine-courses/the_nature_of_microservices/orders/internal/repository"
	ordersv1 "github.com/igoroutine-courses/the_nature_of_microservices/orders/pkg/api/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (a *api) CreateOrder(ctx context.Context, req *ordersv1.CreateOrderRequest) (*ordersv1.CreateOrderResponse, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "validation failed: %v", err)
	}

	now := time.Now().UTC()
	order := &ordersv1.Order{
		Id:        uuid.NewString(),
		UserId:    req.GetUserId(),
		Status:    ordersv1.OrderStatus_ORDER_STATUS_CREATED,
		CreatedAt: timestamppb.New(now),
	}

	created, err := a.ordersRepository.CreateOrder(ctx, order)

	if err != nil {
		if errors.Is(err, repository.ErrOrderAlreadyExists) {
			return nil, status.Error(codes.AlreadyExists, err.Error())
		}

		a.logger.Error("create order failed", zap.Error(err))
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &ordersv1.CreateOrderResponse{Order: created}, nil
}
