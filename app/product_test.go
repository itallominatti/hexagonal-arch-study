package app_test

import (
	"testing"

	"github.com/google/uuid"

	"github.com/itallominatti/hexagonal-arch-study/app"
	"github.com/stretchr/testify/require"
)

func TestProduct_Enable(t *testing.T) {
	product := app.Product{}
	product.ID = uuid.New().String()
	product.Name = "Hello"
	product.Status = app.DISABLED
	product.Price = 10

	err := product.Enable()
	require.Nil(t, err)

	product.Price = 0
	err = product.Enable()
	require.Equal(t, "The price must be grater than zero to enable the product", err.Error())
}

func TestProductDisable(t *testing.T) {
	product := app.Product{}
	product.Name = "Hello"
	product.Status = app.ENABLED
	product.Price = 0

	err := product.Disable()
	require.Nil(t, err)

	product.Price = 10
	err = product.Disable()
	require.Equal(t, "The price must be zero to enable the product", err.Error())
}

func TestProduct_IsValid(t *testing.T) {
	product := app.Product{}
	product.ID = uuid.New().String()
	product.Name = "Hello"
	product.Status = app.ENABLED
	product.Price = 10

	valid, err := product.IsValid()
	require.Nil(t, err)
	require.True(t, valid)
}
