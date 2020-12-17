package domain

import (
	"bytes"
	"context"
	"github.com/nejtr0n/api_example/app/utils"
	"github.com/nejtr0n/api_example/domain/model"
	"github.com/nejtr0n/api_example/domain/repository"
	"github.com/nejtr0n/api_example/domain/services"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/mock"
	"io/ioutil"
	"testing"
	"time"
)

func TestApiService_Fetch(t *testing.T) {
	Convey("Test api service fetch success", t, func() {
		mockedDownloader := new(services.MockCsvFileDownloader)
		mockedDownloader.On("Load", mock.AnythingOfType("string")).Return(
			ioutil.NopCloser(bytes.NewReader([]byte("test,1.254\nblah,2.666"))),
			nil,
		)
		parser := services.NewCsvParser(
			services.NewCsvReaderFactory(),
			services.FetchTimeout(60 * time.Second),
			utils.NewRealTimer(),
		)

		var savedItems []*model.Product
		mockedPostsRepository := new(repository.MockProductsRepository)
		mockedPostsRepository.
			On("SaveAll", mock.Anything, mock.Anything).
			Return(
				func(ctx context.Context, pipe chan *model.Product) int64 {
					for item := range pipe {
						savedItems = append(savedItems, item)
					}
					return int64(len(savedItems))
				},
				func(ctx context.Context, pipe chan *model.Product) error {
					return nil
				},
			)
		svc := NewApiService(mockedDownloader, parser, mockedPostsRepository)
		n, err := svc.Fetch(context.Background(), "http://localhost/test.csv")
		So(err, ShouldBeNil)
		So(n, ShouldEqual, 2)
		So(savedItems, ShouldResemble, []*model.Product{
			{
				Name:  "test",
				Price: 1.254,
			},
			{
				Name:  "blah",
				Price: 2.666,
			},
		})
	})
}
