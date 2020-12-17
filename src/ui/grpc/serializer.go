package api_example

import (
	"github.com/gogo/protobuf/types"
	"github.com/nejtr0n/api_example/domain/model"
)

func serializeProducts(items []*model.Product) ([]*Product, error) {
	var data []*Product
	for _, item := range items {
		ts, err := types.TimestampProto(item.LastModified)
		if err != nil {
			return nil, err
		}
		data = append(data, &Product{
			Id:                   item.Id,
			Name:                 item.Name,
			Price:                item.Price,
			Counter:              item.Counter,
			LastModified:         ts,
		})
	}
	return data, nil
}
