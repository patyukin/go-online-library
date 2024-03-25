package sender

import (
	"context"
	"github.com/patyukin/go-online-library/internal/usecase/model"
)

type Sender struct {
}

func NewSender() *Sender {
	return &Sender{}
}

func (s *Sender) Send(ctx context.Context, filters []model.Filter) error {
	return nil
}
