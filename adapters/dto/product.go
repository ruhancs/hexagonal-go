package dto

import "github.com/ruhancs/hexagonal-go/application"

type Product struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Price float64 `json:"price"`
	Status string `json:"status"`
}

func NewProduct() *Product {
	return &Product{}
}

//fazer a tranferencia do dto para a entidade Product
func (p *Product) Bind(product *application.Product)( *application.Product, error) {
	if p.Id != "" {
		product.ID = p.Id
	}
	product.Name = p.Name
	product.Price = p.Price
	product.Status = p.Status

	_,err := product.IsValid()
	if err != nil {
		return &application.Product{}, err
	}
	return product,nil
}