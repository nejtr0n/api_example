package model

import "time"

type Product struct {
	Id string `bson:"_id"`
	Name string `bson:"name"`
	Price float64 `bson:"price"`
	LastModified time.Time `bson:"lastModified"`
	Counter int64 `bson:"counter"`
}
