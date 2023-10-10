package products

import "errors"

type Repository interface {
	GetAllBySeller(sellerID string) ([]Product, error)
}

type repository struct {
	db []Product
}

func NewRepository(db []Product) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAllBySeller(sellerID string) ([]Product, error) {
	if sellerID == "" {
		return nil, errors.New("invalid seller id")
	}

	prodList := []Product{}

	for _, p := range r.db {
		if p.SellerID == sellerID {
			prodList = append(prodList, p)
		}
	}

	return prodList, nil
}
