package main

import (
	"context"

	pb "github.com/GustavoCesarSantos/go-expert/complete-microservices/common/api"
)

type OrdersService interface {
    CreateOrder(context.Context) error
    ValidateOrder(context.Context, *pb.CreateOrderRequest) error
}

type OrdersStore interface {
    Create(context.Context) error
}
