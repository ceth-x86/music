package engine

import (
	"github.com/demas/music/internal/services/datastore/repository"
)

type Engine struct {
	DataRepository *repository.Repository
}

func NewEngine(dataRepository *repository.Repository) *Engine {
	return &Engine{DataRepository: dataRepository}
}

type DownloadError struct {
	InnerError error
}

func (e *DownloadError) Error() string {
	return e.InnerError.Error()
}

func NewDownloadError(e error) *DownloadError {
	return &DownloadError{InnerError: e}
}

type StoreError struct {
	InnerError error
}

func (e *StoreError) Error() string {
	return e.InnerError.Error()
}

func NewStoreError(e error) *StoreError {
	return &StoreError{InnerError: e}
}
