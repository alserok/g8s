package service

import "context"

type Service interface {
	GetPods(ctx context.Context)
}

func New() Service {
	return &service{}
}

type service struct {
}

func (s service) GetPods(ctx context.Context) {
	//TODO implement me
	panic("implement me")
}
