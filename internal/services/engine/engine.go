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
	Cause error
}

func (e *DownloadError) Error() string {
	return e.Cause.Error()
}

type StoreError struct {
	Cause error
}

func (e *StoreError) Error() string {
	return e.Cause.Error()
}
