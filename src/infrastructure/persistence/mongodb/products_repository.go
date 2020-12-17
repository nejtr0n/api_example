package mongodb

import (
	"context"
	"github.com/nejtr0n/api_example/domain/model"
	"github.com/nejtr0n/api_example/domain/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func NewProductsRepository(collection ProductsCollection, bulkSize ProductBulkSize) repository.ProductsRepository {
	r := new(productsRepository)
	r.collection = collection
	r.bulkSize = int64(bulkSize)
	return r
}

type productsRepository struct {
	collection ProductsCollection
	bulkSize int64
}

// todo: could be optimized with memory allocations
func (p productsRepository) SaveAll(ctx context.Context, pipe chan *model.Product) (int64, error) {
	var (
		operations []mongo.WriteModel
		count int64
		total int64
	)
	for product := range pipe {
		operation := mongo.NewUpdateOneModel().
			SetFilter(bson.D{{"name", product.Name}}).
			SetUpdate(bson.D{
			{"$set", bson.D{
				{"price", product.Price},
			}},
			{"$currentDate", bson.D{
				{"lastModified", true},
			}},
			{"$inc", bson.D{
				{"counter", 1},
			}},
		}).SetUpsert(true)
		operations = append(operations, operation)
		count++
		if count % p.bulkSize == 0 {
			modified, err := p.collectionBulkInsert(ctx, operations)
			if err != nil {
				return -1, err
			}
			total += modified
			operations = nil
		}
	}
	if len(operations) > 0 {
		modified, err := p.collectionBulkInsert(ctx, operations)
		if err != nil {
			return -1, err
		}
		total += modified
	}
	return total, nil
}

func (p productsRepository) Find(ctx context.Context, limit, offset int64, sortBy string, sortOrder int) ([]*model.Product, error) {
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{sortBy, sortOrder}})
	findOptions.SetSkip(offset)
	findOptions.SetLimit(limit)

	cur, err := p.collection.Find(ctx, bson.D{}, findOptions)
	if err != nil {
		return nil, err
	}
	var items []*model.Product
	for cur.Next(context.TODO()) {
		p := new(model.Product)
		err := cur.Decode(p)
		if err != nil {
			return nil, err
		}
		items = append(items, p)
	}
	return items, nil
}

func (p productsRepository) collectionBulkInsert(ctx context.Context, operations []mongo.WriteModel) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, 30 * time.Second)
	defer cancel()
	res, err := p.collection.BulkWrite(ctx, operations, options.BulkWrite())
	if err != nil {
		return -1, err
	}
	return  res.ModifiedCount, nil
}

