package application_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/xXHachimanXx/product-manager/application"
	mock_application "github.com/xXHachimanXx/product-manager/application/mocks"
)

func TestProductService_Get(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	product := mock_application.NewMockProductInterface(controller)
	persistence := mock_application.NewMockProductPersistanceInterface(controller)

	persistence.EXPECT().Get(gomock.Any()).Return(product, nil).AnyTimes()

	service := application.ProductService{
		Persistence: persistence,
	}

	result, err := service.Get("123")
	require.Nil(t, err)
	require.Equal(t, product, result)
}
