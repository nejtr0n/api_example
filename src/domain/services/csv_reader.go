package services

import (
	"encoding/csv"
	"io"
)

type CsvReader interface {
	Read() (record []string, err error)
}

type CsvReaderFactory func(io.ReadCloser) (CsvReader, error)
func NewCsvReaderFactory() CsvReaderFactory {
	return func(data io.ReadCloser) (CsvReader, error) {
		return csv.NewReader(data), nil
	}
}