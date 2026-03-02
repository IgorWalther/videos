package controller

import (
	"context"
	"errors"

	ordersv1 "github.com/igoroutine-courses/the_nature_of_microservices/orders/pkg/api/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a *api) ListUserOrders(req *ordersv1.ListUserOrdersRequest, srv ordersv1.OrderService_ListUserOrdersServer) error {
	if err := req.ValidateAll(); err != nil {
		return status.Errorf(codes.InvalidArgument, "validation failed: %v", err)
	}

	ordersCh, errCh := a.ordersRepository.ListUserOrders(srv.Context(), req.GetUserId())

	for {
		select {
		case <-srv.Context().Done():
			return srv.Context().Err()

		case err, ok := <-errCh:
			if !ok || err == nil {
				return nil
			}

			if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
				return status.Error(codes.Canceled, err.Error())
			}

			a.logger.Error("stream user orders failed", zap.Error(err))
			return status.Error(codes.Internal, "internal error")

		case o, ok := <-ordersCh:
			if !ok {
				return nil
			}

			if o == nil {
				continue
			}

			if err := srv.Send(o); err != nil {
				return err
			}
		}
	}
}
