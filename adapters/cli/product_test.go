package cli_test

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/ruhancs/hexagonal-go/adapters/cli"
	mock_application "github.com/ruhancs/hexagonal-go/application/mocks"
	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productName := "Product 1"
	productPrice := 10.0
	productStatus := "enabled"
	productId := "abc"

	productMock := mock_application.NewMockProductInterface(ctrl)
	productMock.EXPECT().GetID().Return(productId).AnyTimes()
	productMock.EXPECT().GetName().Return(productName).AnyTimes()
	productMock.EXPECT().GetStatus().Return(productStatus).AnyTimes()
	productMock.EXPECT().GetPrice().Return(productPrice).AnyTimes()

	service := mock_application.NewMockProductServiceInterface(ctrl)
	service.EXPECT().Create(productName,productPrice).Return(productMock, nil).AnyTimes()
	service.EXPECT().Get(productId).Return(productMock, nil).AnyTimes()
	service.EXPECT().Enable(gomock.Any()).Return(productMock, nil).AnyTimes()
	service.EXPECT().Disable(gomock.Any()).Return(productMock, nil).AnyTimes()

	resultExpected := fmt.Sprintf("Product Id %s with the name %s has been created with the price %f", 
			productId,productName,productPrice)
	result,err := cli.Run(service, "create", "", productName, float64(productPrice))
	require.Nil(t,err)
	require.Equal(t, resultExpected, result)
	
	resultExpected = fmt.Sprintf("Product %s has been enabled", productName)
	result,err = cli.Run(service, "enable", productId, "", 0.0)
	require.Nil(t,err)
	require.Equal(t, resultExpected, result)
	
	resultExpected = fmt.Sprintf("Product %s has been disabled", productName)
	result,err = cli.Run(service, "disable", productId, "", 0.0)
	require.Nil(t,err)
	require.Equal(t, resultExpected, result)

	resultExpected = fmt.Sprintf("Product Id %s, name %s, price %f, status %s", 
			productId,productName,productPrice,productStatus)
			result,err = cli.Run(service, "get", productId, "", 0.0)
	require.Nil(t,err)
	require.Equal(t, resultExpected, result)

}