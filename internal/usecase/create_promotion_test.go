package usecase

import (
	"context"
	"github.com/patyukin/go-online-library/internal/usecase/mocks"
	"github.com/patyukin/go-online-library/internal/usecase/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestUseCase_CreatePromotion(t *testing.T) {
	txMgrMock := &mocks.TxManager{}
	reposMock := &mocks.Repositories{}

	ctx := context.Background()
	txMgrMock.On("ReadCommitted", ctx, mock.Anything).Return(nil)

	reposMock.On("InsertPromotion", ctx, model.Promotion{}).Return(42, nil)
	reposMock.On("InsertsFilters", ctx, []model.Filter{}).Return([]int{1, 2, 3}, nil)
	reposMock.On("InsertsPromotionFilters", ctx, 42, []int{1, 2, 3}).Return(nil)
	reposMock.On("InsertsDirectories", ctx, []model.Directory{}).Return([]int{4, 5, 6}, nil)
	reposMock.On("InsertsPromotionDirectories", ctx, 42, []int{4, 5, 6}).Return(nil)

	useCase := UseCase{txManager: txMgrMock, repos: reposMock}

	err := useCase.CreatePromotion(context.Background(), model.Promotion{}, []model.Filter{}, []model.Directory{})

	assert.Nil(t, err)
}
