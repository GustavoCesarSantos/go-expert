package usecase

import "port-n-adapters-rest-kafka/internal/entity"

type ListProductsOutput struct {
    ID string
    Name string
    Price float64
}

type ListProductsUseCase struct {
    ProductRepository entity.IProductRepository 
}

func NewListProductsUseCase(productRepository entity.IProductRepository) *ListProductsUseCase {
    return &ListProductsUseCase{ProductRepository: productRepository}
}

func (u *ListProductsUseCase) Execute() ([]*ListProductsOutput, error) {
    products, err := u.ProductRepository.FindAll()
    if err != nil {
        return nil, err
    }
    var productsOutput []*ListProductsOutput
    for _, product := range products {
        productsOutput = append(productsOutput, &ListProductsOutput{
            ID: product.ID,        
            Name: product.Name,
            Price: product.Price,
        })
    }
    return productsOutput, nil
}
