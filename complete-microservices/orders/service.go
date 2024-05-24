package main

import (
	"context"
	"log"

	"github.com/GustavoCesarSantos/go-expert/complete-microservices/common"
	pb "github.com/GustavoCesarSantos/go-expert/complete-microservices/common/api"
)

type service struct {
    store OrdersStore
}

func NewService(store OrdersStore) *service {
    return &service{store}
}

func (s *service) CreateOrder(context.Context) error {
    return nil
}

func (s *service) ValidateOrder(ctx context.Context, orderRequest *pb.CreateOrderRequest) error {
    if len(orderRequest.Items) == 0 {
        return common.ErrNoItems
    }
    mergedItems := mergeItemsQuantities(orderRequest.Items)
    log.Print(mergedItems)
    //validate with the stock service
    return nil
}

func mergeItemsQuantities(items []*pb.ItemsWithQuantity) []*pb.ItemsWithQuantity {
    merged := make([]*pb.ItemsWithQuantity, 0)
    for _, item := range items {
        found := false
        for _, finalItem := range merged {
            if finalItem.ID == item.ID {
                finalItem.Quantity += item.Quantity
                found = true
                break
            }
        }
        if !found {
            merged = append(merged, item)
        }
    }
    return merged
}