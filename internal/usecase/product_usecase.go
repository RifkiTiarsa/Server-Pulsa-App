package usecase

import (
	"fmt"
	"server-pulsa-app/internal/entity"
	"server-pulsa-app/internal/repository"
)

type ProductUseCase interface {
	CreateNewProduct(Product entity.Product) (entity.Product, error)
	FindAllProduct() ([]entity.Product, error)
	FindProductpyId(id string) (entity.Product, error)
	UpdateProduct(Product entity.Product) (entity.Product, error)
	DeleteProduct(id string) error
}

type productUseCase struct {
	repo repository.ProductRepository
}

func (p *productUseCase) CreateNewProduct(Product entity.Product) (entity.Product, error) {
	return p.repo.Create(Product)
}

func (p *productUseCase) FindAllProduct() ([]entity.Product, error) {
	return p.repo.List()
}

func (p *productUseCase) FindProductpyId(id string) (entity.Product, error) {
	return p.repo.Get(id)
}

func (p *productUseCase) UpdateProduct(product entity.Product) (entity.Product, error) {
	_, err := p.repo.Get(product.IdProduct)
	if err != nil {
		return entity.Product{}, fmt.Errorf("product with ID %s not found", product.IdProduct)
	}
	return p.repo.Update(product.IdProduct, product)
}

func (p *productUseCase) DeleteProduct(id string) error {
	_, err := p.repo.Get(id)
	if err != nil {
		return fmt.Errorf("product with ID %s not found", id)
	}
	return p.repo.Delete(id)
}

func NewProductUseCase(repo repository.ProductRepository) ProductUseCase {
	return &productUseCase{repo: repo}
}
