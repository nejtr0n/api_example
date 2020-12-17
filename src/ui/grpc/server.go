package api_example

import (
	"context"
	"github.com/nejtr0n/api_example/domain"
)

func NewServer(svc domain.ApiService) ApiExampleServer {
	s := new(server)
	s.svc = svc
	return s
}

type server struct {
	svc domain.ApiService
}

func (s server) Fetch(ctx context.Context, req *FetchRequest) (*FetchResponse, error) {
	count, err := s.svc.Fetch(ctx, req.Url)
	if err != nil {
		return nil, err
	}
	return &FetchResponse{
		FetchedCount: count,
	}, nil
}

func (s server) List(ctx context.Context, req *ListRequest) (*ListResponse, error) {
	models, err := s.svc.List(ctx,
		req.Pagination.Limit,
		req.Pagination.Offset,
		req.Sorting.Field.String(),
		int(req.Sorting.Sort),
	)
	if err != nil {
		return nil, err
	}
	items, err := serializeProducts(models)
	if err != nil {
		return nil, err
	}
	return &ListResponse{
		Items:                items,
	}, nil
}






