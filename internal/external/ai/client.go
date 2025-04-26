package ai

import (
	"context"

	"github.com/alserok/g8s/internal/external"
)

func NewClient() external.AIClient {
	return &client{}
}

type client struct {
}

func (c client) Prompt(ctx context.Context, prompt string) (string, error) {
	//TODO implement me
	panic("implement me")
}
