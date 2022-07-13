package db

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type StorerMock struct {
	mock.Mock
}

func (m *StorerMock) GetTotalBalance(ctx context.Context, accountId string) (balance float64, err error) {
	args := m.Called(ctx, accountId)
	balance, _ = args.Get(0).(float64)
	return balance, args.Error(1)
}
