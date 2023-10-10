package products

type MockRepository struct {
	Products []Product
	Err      error
}

func (r *MockRepository) GetAllBySeller(sellerID string) ([]Product, error) {
	return r.Products, r.Err
}
