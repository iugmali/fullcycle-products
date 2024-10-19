package cli_test

import (
	"fmt"
	"github.com/iugmali/fullcycle-products/adapters/cli"
	"github.com/iugmali/fullcycle-products/application"
	mock_application "github.com/iugmali/fullcycle-products/application/mocks"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productName := "product"
	productPrice := int64(10_00)
	newProductPrice := int64(0)
	productStatus := application.ENABLED
	productId := "abc"

	productMock := mock_application.NewMockProductInterface(ctrl)
	productMock.EXPECT().GetID().Return(productId).AnyTimes()
	productMock.EXPECT().GetName().Return(productName).AnyTimes()
	productMock.EXPECT().GetPrice().Return(productPrice).AnyTimes()
	productMock.EXPECT().GetStatus().Return(productStatus).AnyTimes()
	productMock.EXPECT().SetPrice(newProductPrice).Return(nil).AnyTimes()

	product2Mock := mock_application.NewMockProductInterface(ctrl)
	product2Mock.EXPECT().GetID().Return(productId).AnyTimes()
	product2Mock.EXPECT().GetName().Return(productName).AnyTimes()
	product2Mock.EXPECT().GetPrice().Return(newProductPrice).AnyTimes()
	product2Mock.EXPECT().GetStatus().Return(productStatus).AnyTimes()

	serviceMock := mock_application.NewMockProductServiceInterface(ctrl)
	serviceMock.EXPECT().Create(productName, productPrice).Return(productMock, nil).AnyTimes()
	serviceMock.EXPECT().Get(productId).Return(productMock, nil).AnyTimes()
	serviceMock.EXPECT().Enable(productMock).Return(productMock, nil).AnyTimes()
	serviceMock.EXPECT().Disable(productMock).Return(productMock, nil).AnyTimes()
	serviceMock.EXPECT().SetPrice(productMock, newProductPrice).Return(product2Mock, nil).AnyTimes()

	expectedResult := fmt.Sprintf("Product ID %s with the name %s has been created with the price %d and status %s", productId, productName, productPrice, productStatus)
	result, err := cli.Run(serviceMock, "create", "", productName, productPrice)
	require.Nil(t, err)
	require.Equal(t, expectedResult, result)

	expectedResult = fmt.Sprintf("Product %s has been enabled", productName)
	result, err = cli.Run(serviceMock, "enable", productId, "", 0)
	require.Nil(t, err)
	require.Equal(t, expectedResult, result)

	expectedResult = fmt.Sprintf("Product %s has been disabled", productName)
	result, err = cli.Run(serviceMock, "disable", productId, "", 0)
	require.Nil(t, err)
	require.Equal(t, expectedResult, result)

	expectedResult = fmt.Sprintf("Product %s price has been set to %d", productName, newProductPrice)
	result, err = cli.Run(serviceMock, "setprice", productId, "", newProductPrice)
	require.Nil(t, err)
	require.Equal(t, expectedResult, result)

	expectedResult = fmt.Sprintf("ID: %s\nName: %s\nPrice: %d\nStatus: %s", productId, productName, productPrice, productStatus)
	result, err = cli.Run(serviceMock, "", productId, "", 0)
	require.Nil(t, err)
	require.Equal(t, expectedResult, result)
}
