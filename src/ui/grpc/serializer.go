package api_example

import "github.com/nejtr0n/api_example/domain/model"

func serializeProducts(items []*model.Product) ([]*Product, error) {
	var data []*Product
	for _, item := range items {
		data = append(data, &Product{
			Name: item.Name,
		})
	}
	return data, nil
}
