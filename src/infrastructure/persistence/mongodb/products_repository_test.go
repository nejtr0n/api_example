package mongodb

import (
	"context"
	"github.com/nejtr0n/api_example/domain/model"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"testing"
)

func TestProductsRepository_SaveAll(t *testing.T) {
	Convey("Test products repository SaveAll success", t, func() {
		testData := make(chan *model.Product, 2)
		testData <- &model.Product{
			Name:         "test 1",
			Price:        2.66,
		}
		testData <- &model.Product{
			Name:         "test 2",
			Price:        1.333,
		}
		close(testData)
		var inserted []mongo.WriteModel
		mockedProductsCollection := new(MockProductsCollection)
		mockedProductsCollection.On("BulkWrite", mock.Anything, mock.Anything, mock.Anything).Return(
			func(ctx context.Context, models []mongo.WriteModel, opts ...*options.BulkWriteOptions) *mongo.BulkWriteResult {
				for _, item := range models {
					inserted = append(inserted, item)
				}
				return &mongo.BulkWriteResult{
					ModifiedCount: int64(len(models)),
				}
			},
			func(ctx context.Context, models []mongo.WriteModel, opts ...*options.BulkWriteOptions) error {
				return nil
			},
		)

		rep := NewProductsRepository(mockedProductsCollection, 1)
		res, err := rep.SaveAll(context.Background(), testData)
		So(err, ShouldBeNil)
		So(res, ShouldEqual, 2)
		So(inserted[0], ShouldHaveSameTypeAs, &mongo.UpdateOneModel{})
		So(inserted[0].(*mongo.UpdateOneModel).Filter, ShouldResemble, bson.D{{"name", "test 1"}})
		So(inserted[0].(*mongo.UpdateOneModel).Update, ShouldResemble, bson.D{
			{"$set", bson.D{
				{"price", 2.66},
			}},
			{"$currentDate", bson.D{
				{"lastModified", true},
			}},
			{"$inc", bson.D{
				{"counter", 1},
			}},
		})
		So(*inserted[0].(*mongo.UpdateOneModel).Upsert, ShouldBeTrue)
		So(inserted[1], ShouldHaveSameTypeAs, &mongo.UpdateOneModel{})
		So(inserted[1].(*mongo.UpdateOneModel).Filter, ShouldResemble, bson.D{{"name", "test 2"}})
		So(inserted[1].(*mongo.UpdateOneModel).Update, ShouldResemble, bson.D{
			{"$set", bson.D{
				{"price", 1.333},
			}},
			{"$currentDate", bson.D{
				{"lastModified", true},
			}},
			{"$inc", bson.D{
				{"counter", 1},
			}},
		})
		So(*inserted[1].(*mongo.UpdateOneModel).Upsert, ShouldBeTrue)
	})
}
