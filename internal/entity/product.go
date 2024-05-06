package entity

import (
	"github.com/google/uuid"
)

type ProductRepository interface {
	Create(procuct *Procuct) error
	FindAll() ([]*Procuct, error)
}

type Procuct struct {
	ID    string
	Name  string
	Price float64
}

func NewProduct(name string, price float64) *Procuct {
	return &Procuct{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}
