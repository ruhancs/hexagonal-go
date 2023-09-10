package application_test

import (
	"testing"

	"github.com/ruhancs/hexagonal-go/application"
	"github.com/stretchr/testify/require"
	uuid "github.com/satori/go.uuid"
)

func TestProduct_Enable(t *testing.T) {
	product := application.Product{}
	product.Name = "Hello"
	product.Status = application.DISABLED
	product.Price = 10

	err := product.Enable()
	require.Nil(t, err) // verificar se o erro é nulo

	product.Price = 0
	err = product.Enable()
	require.Equal(t, "The must be greather tha 0", err.Error())
}

func TestProduct_Disable(t *testing.T) {
	product := application.Product{}
	product.Name = "Hello"
	product.Status = application.DISABLED
	product.Price = 0

	err := product.Disable()
	require.Nil(t, err) // verificar se o erro é nulo

	product.Price = 10
	err = product.Disable()
	require.Equal(t, "The must be 0 to disable the product", err.Error())
}

func TestProduct_IsValid(t *testing.T) {
	product := application.Product{}
	product.Name = "Hello"
	product.Status= application.DISABLED
	product.Price = 10
	product.ID = uuid.NewV4().String()

	_, err := product.IsValid()
	require.Nil(t, err)

	product.Status = "INVALID"
	_,err = product.IsValid()
	require.Equal(t, "the status must be enabled or disabled", err.Error())

	product.Status = "enabled"
	_,err = product.IsValid()
	require.Nil(t, err)

	product.Price = -10
	_, err = product.IsValid()
	require.Equal(t, "the price must be greather or equal to 0", err.Error())
}

