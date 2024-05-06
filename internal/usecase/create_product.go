package usecase

import "github.com/axelgeorg/ImersaoGo_03022023/internal/entity"

type CreateProductInputDto struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type CreateProductOutput struct {
	ID    string
	Name  string
	Price float64
}

type CreateProductUseCase struct {
	ProductRepository entity.ProductRepository
}

func NewCreateProductUseCase(productRepository entity.ProductRepository) *CreateProductUseCase {
	return &CreateProductUseCase{ProductRepository: productRepository}
}

func (u *CreateProductUseCase) Execute(input CreateProductInputDto) (*CreateProductOutput, error) {
	product := entity.NewProduct(input.Name, input.Price)
	err := u.ProductRepository.Create(product)
	if err != nil {
		return nil, err
	}

	return &CreateProductOutput{
		ID:    product.ID,
		Name:  product.Name,
		Price: product.Price,
	}, nil
}
