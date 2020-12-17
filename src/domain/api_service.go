package domain

import (
	"context"
	"github.com/nejtr0n/api_example/domain/model"
	"github.com/nejtr0n/api_example/domain/repository"
	"github.com/nejtr0n/api_example/domain/services"
	"golang.org/x/sync/errgroup"
)

type ApiService interface {
	Fetch(ctx context.Context, url string) (int64, error)
	List(ctx context.Context, limit, offset int64, sortBy string, sortOrder int) ([]*model.Product, error)
}

func NewApiService(downloader services.CsvLoader, parser services.CsvParser, repository repository.ProductsRepository) ApiService {
	s := new(apiService)
	s.downloader = downloader
	s.parser = parser
	s.repository = repository
	return s
}

type apiService struct {
	downloader services.CsvLoader
	parser services.CsvParser
	repository repository.ProductsRepository
}

func (a apiService) Fetch(ctx context.Context, url string) (int64, error) {
	data, err := a.downloader.Load(url)
	if err != nil {
		return -1, err
	}
	defer data.Close()

	g, ctx := errgroup.WithContext(ctx)
	pipe := make(chan *model.Product)
	g.Go(func() error {
		return a.parser.Parse(ctx, data, pipe)
	})
	var n int64
	g.Go(func() error {
		n, err = a.repository.SaveAll(ctx, pipe)
		if err != nil {
			return err
		}
		return nil
	})

	return n, g.Wait()
}

func (a apiService) List(ctx context.Context, limit, offset int64, sortBy string, sortOrder int) ([]*model.Product, error) {

	return a.repository.Find(ctx, limit, offset, sortBy, sortOrder)
}


