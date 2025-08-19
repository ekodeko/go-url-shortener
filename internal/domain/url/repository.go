package url

import "errors"

type Repository interface {
	Create(e *Entity) error
	FindByCode(code string) (*Entity, error)
	IncrementClicks(code string) error
	DeleteByCode(code string) error
}

var ErrNotFound = errors.New("url not found")
