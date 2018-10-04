package api

import (
	"context"

	"github.com/jemmycalak/mall-tangsel/service/product"
	"github.com/jemmycalak/mall-tangsel/service/user"
)

//interface connect to service
type UserService interface {
	IsUserActive(context.Context, int64) (bool, error)
	GetUser(context.Context) user.User
	Register(context.Context, *user.User) error
	NewUser() *user.User
}

type ProductService interface {
	NewProduct() *product.Product
	CreateProduct(*product.Product) error
	Products() ([]product.Product, error)
	Product(int) (*product.Product, error)
	EditProduct(*product.Product) error
	DeleteProduct(int) error
	UpdateImage(int, []string) error
	ProductLimit(int, int) ([]product.Product, error)
}
