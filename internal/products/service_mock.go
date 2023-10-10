package products

type MockService struct {
	Products []Product
	Err      error
}

func (s *MockService) GetAllBySeller(sellerID string) ([]Product, error) {
	return s.Products, s.Err
}
