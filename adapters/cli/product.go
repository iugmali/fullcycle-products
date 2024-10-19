package cli

import (
	"fmt"
	"github.com/iugmali/fullcycle-products/application"
)

func Run(service application.ProductServiceInterface, action string, productId string, productName string, productPrice int64) (string, error) {
	var result string = ""
	switch action {
	case "list":
		products, err := service.GetAll()
		if err != nil {
			return result, err
		}
		result += "\n----------------------------------------\nPRODUCT LIST\n----------------------------------------\n"
		for _, product := range products {
			result += fmt.Sprintf("ID: %s\nName: %s\nPrice: %d\nStatus: %s\n----------------------------------------\n", product.GetID(), product.GetName(), product.GetPrice(), product.GetStatus())
		}
	case "create":
		product, err := service.Create(productName, productPrice)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product ID %s with the name %s has been created with the price %d and status %s", product.GetID(), product.GetName(), product.GetPrice(), product.GetStatus())
	case "enable":
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}
		res, err := service.Enable(product)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product %s has been enabled", res.GetName())
	case "disable":
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}
		res, err := service.Disable(product)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product %s has been disabled", res.GetName())
	case "setprice":
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}
		res, err := service.SetPrice(product, productPrice)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product %s price has been set to %d", res.GetName(), res.GetPrice())
	default:
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("ID: %s\nName: %s\nPrice: %d\nStatus: %s", product.GetID(), product.GetName(), product.GetPrice(), product.GetStatus())
	}
	return result, nil
}
