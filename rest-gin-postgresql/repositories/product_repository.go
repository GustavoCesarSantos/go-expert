package repositories

import (
	"database/sql"
	"fmt"
	"rest-gin-postgresql/models"
)

type ProductRepository struct {
    connection *sql.DB
}

func NewProductRepository(connection *sql.DB) ProductRepository {
    return ProductRepository{
        connection: connection,
    }
}

func (pr *ProductRepository) GetProducts() ([]models.Product, error) {
    query := "select id, name, price from product"
    rows, err := pr.connection.Query(query)
    if err != nil {
        fmt.Println(err)
        return nil, err
    }
    var list []models.Product
    var product models.Product
    for rows.Next() {
        err = rows.Scan(
            &product.ID,
            &product.Name,
            &product.Price,
            )
        if err != nil {
        fmt.Println(err)
        return nil, err
        }
        list = append(list, product)
    }
    rows.Close()
    return list, nil
}

func (pr *ProductRepository) CreateProduct(product models.Product) error {
    var id int
    query, err := pr.connection.Prepare(`
        INSERT INTO products (
            name,
            price
        )
        RETURNING
            id
        VALUES (
            $1,
            $2
        )
    `)
    if err != nil {
        fmt.Println(err)
        return err
    }
    err = query.QueryRow(product.Name, product.Price).Scan(&id)
    if err != nil {
        fmt.Println(err)
        return err
    }
    query.Close()
    return nil
}

func (pr *ProductRepository) GetProduct(id int) (*models.Product, error) {
    query, err := pr.connection.Prepare(`
        SELECT
            *
        FROM
            products
        WERE
            id = $1
    `)
    if err != nil {
        fmt.Println(err)
        return nil, err
    }
    var product models.Product
    err = query.QueryRow(id).Scan(
        &product.ID,
        &product.Name,
        &product.Price,
    )
    if err != nil {
        fmt.Println(err)
        return nil, err
    }
    query.Close()
    return &product, nil
}
