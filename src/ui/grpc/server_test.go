package api_example

import (
	"context"
	"errors"
	"github.com/nejtr0n/api_example/domain"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestServer_Fetch(t *testing.T) {
	Convey("Test grpc server fetch success", t, func() {
		mockSvc := new(domain.MockApiService)
		mockSvc.On("Fetch", mock.Anything, mock.Anything).
			Return(int64(10), nil)
		server := NewServer(mockSvc)
		result, err := server.Fetch(context.Background(), &FetchRequest{
			Url: "http://localhost/test.csv",
		})
		So(err, ShouldBeNil)
		So(result, ShouldNotBeNil)
		So(result, ShouldResemble, &FetchResponse{
			FetchedCount: 10,
		})
	})

	Convey("Test grpc server fetch fails", t, func() {
		testError := errors.New("test")
		mockSvc := new(domain.MockApiService)
		mockSvc.On("Fetch", mock.Anything, mock.Anything).
			Return(int64(-1), testError)
		server := NewServer(mockSvc)
		result, err := server.Fetch(context.Background(), &FetchRequest{
			Url: "http://localhost/test.csv",
		})
		So(result, ShouldBeNil)
		So(err, ShouldNotBeNil)
		So(err, ShouldResemble, testError)
	})
}
