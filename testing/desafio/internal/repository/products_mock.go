package repository

import (
	"app/internal"

	"github.com/stretchr/testify/mock"
)

// ProductsRepositoryMock is a mock implementation of the ProductsRepository interface.
type ProductsRepositoryMock struct {
	mock.Mock
}

func (m *ProductsRepositoryMock) SearchProducts(query internal.ProductQuery) (p map[int]internal.Product, err error) {
	args := m.Called(query)
	return args.Get(0).(map[int]internal.Product), args.Error(1)
}
