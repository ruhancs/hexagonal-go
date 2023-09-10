// gerar mock de product:
//mockgen -destination=application/mocks/application.go -source=application/product.go application

package application_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/ruhancs/hexagonal-go/application"
	mock_application "github.com/ruhancs/hexagonal-go/application/mocks"
	"github.com/stretchr/testify/require"
)

func TestProductService_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish() // espera todos processos de ctrl finalizar para finalizar o ctrl

	//criar mock de product e persistence
	product := mock_application.NewMockProductInterface(ctrl)
	persistence := mock_application.NewMockProductPersistenceInterface(ctrl)

	//todas as vezes que o metodo get for chamado independente do valor retorna um product e nil
	persistence.EXPECT().Get(gomock.Any()).Return(product, nil).AnyTimes()

	service := application.ProductService{
		Persistence: persistence,
	}

	result,err := service.Get("123")
	require.Nil(t,err)
	require.Equal(t,product,result)
}

func TestProductService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish() // espera todos processos de ctrl finalizar para finalizar o ctrl

	//criar mock de product e persistence
	product := mock_application.NewMockProductInterface(ctrl)
	persistence := mock_application.NewMockProductPersistenceInterface(ctrl)

	//todas as vezes que o metodo get for chamado independente do valor retorna um product e nil
	persistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()

	service := application.ProductService{Persistence: persistence}

	result,err := service.Create("P1", 10)
	require.Nil(t,err)
	require.Equal(t,product,result)
}

func TestProductService_EnableDisable(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish() // espera todos processos de ctrl finalizar para finalizar o ctrl

	//criar mock de product e persistence
	product := mock_application.NewMockProductInterface(ctrl)
	//simula que o enable retorna erro nil(nulo)
	product.EXPECT().Enable().Return(nil)
	product.EXPECT().Disable().Return(nil)

	persistence := mock_application.NewMockProductPersistenceInterface(ctrl)

	//todas as vezes que o metodo get for chamado independente do valor retorna um product e nil
	persistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()

	service := application.ProductService{Persistence: persistence}

	result,err := service.Enable(product)
	require.Nil(t,err)
	require.Equal(t,product,result)
	
	result,err = service.Disable(product)
	require.Nil(t,err)
	require.Equal(t,product,result)
}