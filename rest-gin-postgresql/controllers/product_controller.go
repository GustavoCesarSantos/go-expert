package controllers

import (
	"net/http"
	"rest-gin-postgresql/models"
	"rest-gin-postgresql/usecases"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
    productUseCases usecases.ProductUseCases
}

func NewProductController(productUseCases usecases.ProductUseCases) ProductController {
    return ProductController{
        productUseCases: productUseCases,
    }
}

func (p *ProductController) GetProducts(ctx *gin.Context) {
    products, err := p.productUseCases.GetProducts()
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, err)
    }
    ctx.JSON(200, products)
}

func (p *ProductController) CreateProduct(ctx *gin.Context) {
    var product models.Product
    err := ctx.BindJSON(&product)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, err)
        return
    }
    err = p.productUseCases.CreateProduct(product)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, err)
    }
    ctx.Status(http.StatusCreated)
}

func (p *ProductController) GetProduct(ctx *gin.Context) {
    stringId := ctx.Param("id")
    if stringId == "" {
        ctx.JSON(http.StatusBadRequest, "Please provide product id")
        return
    }
    id, err := strconv.Atoi(stringId)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, err)
        return
    }
    product, err := p.productUseCases.GetProduct(id)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, err)
    }
    ctx.JSON(200, product)
}
