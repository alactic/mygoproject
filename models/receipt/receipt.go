package receipt

import (
	"github.com/alactic/mygoproject/models/customers"
	"github.com/alactic/mygoproject/models/product"
)

type Customer = customers.Customer
type Product = product.Product

type Receipt struct {
	Id       string    `json:"id, omitempty"`
	Type     string    `json:"type"`
	Customer Customer  `json:"customer"`
	Product  []Product `json:"products"`
	Total    float32   `json:"total`
}
