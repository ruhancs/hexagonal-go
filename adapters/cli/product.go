package cli

// criar cli com o cobra: cobra-cli init
// criar comando: cobra-cli add cli 
// rodar o programa com o cobra: go run main.go cli (cli é o nome do comando criado)

import (
	"fmt"

	"github.com/ruhancs/hexagonal-go/application"
)

func Run(service application.ProductServiceInterface, action string, productId string, productName string, price float64) (string, error) {
		var result = ""

		switch action {
		case "create":
			product, err := service.Create(productName,price)
			if err != nil {
				return result, err
			}
			result = fmt.Sprintf("Product Id %s with the name %s has been created with the price %f", 
			product.GetID(),product.GetName(),product.GetPrice())
		
		case "enable": 
			product, err := service.Get(productId)
			if err != nil {
				return result, err
			}
			
			res,err := service.Enable(product)
			if err != nil {
				return result, err
			}
			result = fmt.Sprintf("Product %s has been enabled", res.GetName())
		
		case "disable": 
			product, err := service.Get(productId)
			if err != nil {
				return result, err
			}
			
			res,err := service.Disable(product)
			if err != nil {
				return result, err
			}
			result = fmt.Sprintf("Product %s has been disabled", res.GetName())
		
		default:
			res, err := service.Get(productId)
			if err != nil {
				return result, err
			}
			result = fmt.Sprintf("Product Id %s, name %s, price %f, status %s", 
			res.GetID(),res.GetName(),res.GetPrice(), res.GetStatus())
		}
		return result, nil
}