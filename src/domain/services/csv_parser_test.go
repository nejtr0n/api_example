package services

import (
	"bytes"
	"context"
	"errors"
	"github.com/nejtr0n/api_example/app/utils"
	"github.com/nejtr0n/api_example/domain/model"
	. "github.com/smartystreets/goconvey/convey"
	"io"
	"io/ioutil"
	"testing"
	"time"
)

var (
	expectedError = errors.New("test")
)

func TestCsvFileParser_Parse(t *testing.T) {
	Convey("Test csv file parser success", t, func() {
		mockTimer := new(utils.MockTimer)
		mockTimer.On("Now").Return(time.Now().AddDate(1, 0, 0))
		parser := NewCsvParser(
			NewCsvReaderFactory(),
			FetchTimeout(time.Second * 10000),
			mockTimer,
		)
		// buffer size is bigger, than lines in csv
		// so reading into pipe cache
		pipe := make(chan *model.Product, 3)
		data := ioutil.NopCloser(bytes.NewReader([]byte("test,1.254\nblah,2.666")))
		err := parser.Parse(context.Background(), data, pipe)
		So(err, ShouldBeNil)
		product := <- pipe
		So(product, ShouldNotBeNil)
		So(product, ShouldResemble, &model.Product{
			Name:  "test",
			Price: 1.254,
		})
		product = <- pipe
		So(product, ShouldResemble, &model.Product{
			Name:  "blah",
			Price: 2.666,
		})
	})

	Convey("Test csv file parser fails when reader factory return errors", t, func() {
		mockTimer := new(utils.MockTimer)
		mockTimer.On("Now").Return(time.Now().AddDate(1, 0, 0))
		mockedReaderFactory := func(data io.ReadCloser) (CsvReader, error) {
			return nil, expectedError
		}
		parser := NewCsvParser(
			mockedReaderFactory,
			FetchTimeout(time.Second * 10000),
			mockTimer,
		)
		pipe := make(chan *model.Product, 3)
		defer close(pipe)
		data := ioutil.NopCloser(bytes.NewReader([]byte("test,1.254\nblah,2.666")))
		err := parser.Parse(context.Background(), data, pipe)
		So(err, ShouldBeError)
		So(err, ShouldResemble, expectedError)
	})

	Convey("Test csv file parser fails when reader read returns error", t, func() {
		mockTimer := new(utils.MockTimer)
		mockTimer.On("Now").Return(time.Now().AddDate(1, 0, 0))

		mockedReaderFactory := func(data io.ReadCloser) (CsvReader, error) {
			mockedReader :=  new(MockCsvReader)
			mockedReader.On("Read").Return([]string{"test", "1.254"}, nil).Once()
			mockedReader.On("Read").Return(nil, expectedError).Once()
			return mockedReader, nil
		}
		parser := NewCsvParser(
			mockedReaderFactory,
			FetchTimeout(time.Second * 10000),
			mockTimer,
		)
		pipe := make(chan *model.Product, 3)
		data := ioutil.NopCloser(bytes.NewReader([]byte("test,1.254\nblah,2.666")))
		err := parser.Parse(context.Background(), data, pipe)
		So(err, ShouldBeError)
		So(err, ShouldResemble, expectedError)
		product := <- pipe
		So(product, ShouldNotBeNil)
		So(product, ShouldResemble, &model.Product{
			Name:  "test",
			Price: 1.254,
		})
	})

	Convey("Test csv file parser fails when parse is too long (context deadline exceed)", t, func() {
		mockTimer := new(utils.MockTimer)
		mockTimer.On("Now").Return(time.Now().Add(time.Duration(-60) * time.Millisecond))

		parser := NewCsvParser(
			NewCsvReaderFactory(),
			FetchTimeout(time.Millisecond * 50),
			mockTimer,
		)

		pipe := make(chan *model.Product, 3)
		data := ioutil.NopCloser(bytes.NewReader([]byte("test,1.254\nblah,2.666")))
		err := parser.Parse(context.Background(), data, pipe)
		So(err, ShouldBeError)
		So(err, ShouldResemble, context.DeadlineExceeded)
	})

	Convey("Test csv file parser fails when nobody read from pipe (context deadline exceed)", t, func() {
		mockTimer := new(utils.MockTimer)
		mockTimer.On("Now").Return(time.Now())
		parser := NewCsvParser(
			NewCsvReaderFactory(),
			FetchTimeout(30 * time.Millisecond),
			mockTimer,
		)

		pipe := make(chan *model.Product)
		data := ioutil.NopCloser(bytes.NewReader([]byte("test,1.254\nblah,2.666")))
		err := parser.Parse(context.Background(), data, pipe)
		So(err, ShouldBeError)
		So(err, ShouldResemble, context.DeadlineExceeded)
	})
}
