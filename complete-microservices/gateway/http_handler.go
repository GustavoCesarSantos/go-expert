package main

import (
	"errors"
	"net/http"

	"github.com/GustavoCesarSantos/go-expert/complete-microservices/common"
	pb "github.com/GustavoCesarSantos/go-expert/complete-microservices/common/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type handler struct {
	orderClient pb.OrderServiceClient
}

func NewHandler(orderClient pb.OrderServiceClient) *handler {
 return &handler{orderClient: orderClient}
}

func (h *handler) registerRoutes(mux *http.ServeMux) {
    mux.HandleFunc("POST /api/customers/{customerID}/orders", h.HandleCreateOrder)
}

func (h *handler) HandleCreateOrder(w http.ResponseWriter, r *http.Request) {
	customerId := r.PathValue("customerID")
	var items []*pb.ItemsWithQuantity
	if err := common.ReadJSON(r, &items); err != nil {
		common.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}
	if err := validateItems(items); err != nil {
		common.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}
	order, err := h.orderClient.CreateOrder(r.Context(), &pb.CreateOrderRequest{
		CustomerID: customerId,
		Items: items,
	})
	rStatus := status.Convert(err)
	if rStatus != nil {
		if rStatus.Code() != codes.InvalidArgument {
			common.WriteError(w, http.StatusBadRequest, rStatus.Message())
			return
		}
		common.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}
	common.WriteJSON(w, http.StatusCreated, order)
}

func validateItems(items []*pb.ItemsWithQuantity) error {
	if len(items) == 0 {
		return common.ErrNoItems
	}
	for _, i := range items {
		if i.ID == "" {
			return errors.New("item ID is required")
		}
		if i.Quantity <= 0 {
			return errors.New("item must have a valid quantity")
		}
	}
	return nil
}
