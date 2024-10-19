package application_test

import (
	"github.com/google/uuid"
	"github.com/iugmali/fullcycle-products/application"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestProduct_Enable(t *testing.T) {
	product := application.Product{}
	product.Name = "Hello"
	product.Status = application.DISABLED
	product.Price = 10_00

	err := product.Enable()
	require.Nil(t, err)

	product.Price = 0
	err = product.Enable()
	require.Equal(t, "the price must be greater than zero", err.Error())
}

func TestProduct_Disable(t *testing.T) {
	product := application.Product{}
	product.Name = "Hello"
	product.Status = application.ENABLED
	product.Price = 0

	err := product.Disable()
	require.Nil(t, err)

	product.Price = 10_00
	err = product.Disable()
	require.Equal(t, "the price must be zero", err.Error())
}

func TestProduct_IsValid(t *testing.T) {
	product := application.Product{}
	product.ID = uuid.NewString()
	product.Name = "Hello"
	product.Status = application.DISABLED
	product.Price = 10_00

	_, err := product.IsValid()
	require.Nil(t, err)

	product.Status = "blablabla"
	_, err = product.IsValid()
	require.Equal(t, "the status must be enabled or disabled", err.Error())

	product.Status = application.ENABLED
	product.Price = -10_00
	_, err = product.IsValid()
	require.Equal(t, "the price must be greater or equal zero", err.Error())
}

func TestProduct_SetPrice(t *testing.T) {
	product := application.Product{}
	product.ID = uuid.NewString()
	product.Name = "Hello"
	product.Status = application.DISABLED
	product.Price = 10_00

	err := product.SetPrice(20_00)
	require.Nil(t, err)
	require.Equal(t, int64(20_00), product.Price)

	err = product.SetPrice(-20_00)
	require.Equal(t, "the price must be greater or equal zero", err.Error())
}
