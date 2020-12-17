package services

import (
	"context"
	"github.com/nejtr0n/api_example/app/utils"
	"github.com/nejtr0n/api_example/domain/model"
	"io"
	"strconv"
	"time"
)

type CsvParser interface {
	Parse(ctx context.Context, data io.ReadCloser, pipe chan *model.Product) error
}

func NewCsvParser(readerFactory CsvReaderFactory, readTimeout FetchTimeout, timer utils.Timer) CsvParser {
	p := new(csvFileParser)
	p.readTimeout = time.Duration(readTimeout)
	p.timer = timer
	p.readerFactory = readerFactory
	return p
}

type csvFileParser struct {
	readTimeout time.Duration
	timer utils.Timer
	readerFactory CsvReaderFactory
}

func (p csvFileParser) Parse(ctx context.Context, data io.ReadCloser, pipe chan *model.Product) error {
	ctx, cancel := context.WithDeadline(ctx, p.timer.Now().Add(p.readTimeout))
	defer cancel()

	reader, err := p.readerFactory(data)
	if err != nil {
		return err
	}
	done := make(chan error)
	go func() {
		defer close(pipe)
		for {
			row, err := reader.Read()
			if err != nil {
				done <- err
				return
			}
			price, err := strconv.ParseFloat(row[1], 64)
			if err != nil {
				done <- err
				return
			}
			pipe <- &model.Product{
				Name:  row[0],
				Price: price,
			}
		}
	}()

	select {
	case err := <- done:
		if err == io.EOF {
			return nil
		}
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}
