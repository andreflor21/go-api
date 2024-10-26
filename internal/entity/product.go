package entity

import (
	"errors"
	"time"

	"github.com/andreflor21/go-api/pkg/entity"
)


var (
	ErrIDIsReuqired = errors.New("id is required")
	ErrNameIsReuqired = errors.New("name is required")
	ErrPriceIsReuqired = errors.New("price is required")
	ErrInvalidPrice = errors.New("price is invalid")
	ErrInvalidId = errors.New("id is invalid")
)
type Product struct {
	ID		  entity.ID     `json:"id"`
	Name      string        `json:"name"`
	Price     float64       	`json:"price"`
	CreatedAt time.Time		`json:"created_at"`
}

func NewProduct(name string, price float64) (*Product, error) {
	p := &Product{
		ID: entity.NewID(),
		Name: name,
		Price: price,
		CreatedAt: time.Now(),
	}
	err := p.Validate()
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (p *Product) Validate() error {
	if p.ID.String() == "" {
		return ErrIDIsReuqired
	}
	if _,err := entity.ParseID(p.ID.String()); err != nil {
		return ErrInvalidId
	}
	if p.Name == "" {
		return ErrNameIsReuqired
	}
	if p.Price == 0 {
		return ErrPriceIsReuqired
	}
	if p.Price < 0 {
		return ErrInvalidPrice
	}
	return nil
}