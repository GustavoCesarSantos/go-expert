package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"port-n-adapters-rest-kafka/internal/infra/akafka"
	"port-n-adapters-rest-kafka/internal/infra/repository"
	"port-n-adapters-rest-kafka/internal/infra/web"
	"port-n-adapters-rest-kafka/internal/usecase"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/go-chi/chi/v5"
    _ "github.com/go-sql-driver/mysql"
)

func main() {
    db, err := sql.Open("mysql", "root:root@tcp(host.docker.internal:3306)/products")
    if err != nil {
        panic(err)
    }
    defer db.Close()
    repository := repository.NewProductRepositoryMysql(db)
    createProductUseCase := usecase.NewCreateProductUseCase(repository)
    listProductsUseCase := usecase.NewListProductsUseCase(repository)
    productHandler := web.NewProductHandler(createProductUseCase, listProductsUseCase)
    r := chi.NewRouter()
    r.Get("/products", productHandler.ListProductsHandler)
    r.Post("/products", productHandler.CreateProductHandler)
    go http.ListenAndServe(":8000", r)
    msgChannel := make(chan *kafka.Message)
    go akafka.Consume([]string{"products"}, "host.docker.internal:9094", msgChannel)
    for msg := range msgChannel {
        dto := usecase.CreateProductInput{}
        err := json.Unmarshal(msg.Value, &dto)
        if err != nil {
            continue
        }
        _, err = createProductUseCase.Execute(dto)
    }
}
