package application_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/xXHachimanXx/product-manager/application"
)

func TestPrduct_Enable(t *testing.T) {
	product := application.Product{}
	product.Name = "A"
	product.Status = application.DISABLED
	product.Price = 12.99

	err := product.Enable()
	require.Nil(t, err)

	product.Price = 0
	err = product.Enable()
	require.Equal(t, "The price must be greater than zero to enable the product", err.Error())
}
