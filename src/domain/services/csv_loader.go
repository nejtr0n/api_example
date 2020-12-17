package services

import (
	"io"
	"net"
	"net/http"
	"time"
)

type CsvLoader interface {
	Load(string) (io.ReadCloser, error)
}

func NewCsvLoader() CsvLoader {
	l := new(csvLoader)
	l.client = &http.Client{
		Timeout: time.Second * 60,
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout:   60 * time.Second,
				KeepAlive: 60 * time.Second,
			}).DialContext,
		},
	}
	return l
}

type csvLoader struct {
	client *http.Client
}

func (l csvLoader) Load(url string) (io.ReadCloser, error) {
	resp, err := l.client.Get(url)
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}
