package usecases

import (
	"rest-gin-postgresql/models"
	"rest-gin-postgresql/repositories"
)

type ProductUseCases struct {
    productRepository repositories.ProductRepository
}

func NewProductUseCase(productRepository repositories.ProductRepository) ProductUseCases {
    return ProductUseCases{
        productRepository: productRepository,
    }
}

func (pu *ProductUseCases) GetProducts() ([]models.Product, error) {
    return pu.productRepository.GetProducts()
}

func (pu *ProductUseCases) CreateProduct(product models.Product) error {
    return pu.productRepository.CreateProduct(product)
}

func (pu *ProductUseCases) GetProduct(id int) (*models.Product, error) {
    return pu.productRepository.GetProduct(id)
}
