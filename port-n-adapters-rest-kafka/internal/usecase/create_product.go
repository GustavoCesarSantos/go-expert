package usecase

import "port-n-adapters-rest-kafka/internal/entity"

type CreateProductInput struct {
    Name string `json:"name"`
    Price float64 `json:"price"`
}

type CreateProductOutput struct {
    ID string
    Name string
    Price float64
}

type CreateProductUseCase struct {
    ProductRepository entity.IProductRepository 
}

func NewCreateProductUseCase(productRepository entity.IProductRepository) *CreateProductUseCase {
    return &CreateProductUseCase{ProductRepository: productRepository}
}

func (u *CreateProductUseCase) Execute(input CreateProductInput) (*CreateProductOutput, error) {
    product := entity.NewProduct(input.Name, input.Price)
    err := u.ProductRepository.Create(product)
    if err != nil {
        return nil, err
    }
    return &CreateProductOutput{
        ID: product.ID,
        Name: product.Name,
        Price: product.Price,
    }, nil
}

