package ai

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/alserok/g8s/internal/external"
	"github.com/alserok/g8s/internal/utils/errors"
)

const (
	contentWrapper = "Provide me kubernetes config yaml file for: %s, and no other additional comments"
)

func NewClient(token string) external.AIClient {
	return &client{
		token: token,
		cl:    http.DefaultClient,
	}
}

type client struct {
	token string

	cl *http.Client
}

func (c client) Prompt(ctx context.Context, prompt fmt.Stringer) (string, error) {
	b, err := json.Marshal(request{
		Model:    "deepseek/deepseek-v3-0324",
		Messages: []message{{Role: "user", Content: fmt.Sprintf("%s%s", contentWrapper, prompt.String())}},
	})
	if err != nil {
		return "", errors.New(err.Error(), errors.ErrInternal)
	}

	req, err := http.NewRequest(http.MethodPost, "https://router.huggingface.co/novita/v3/openai/chat/completions", bytes.NewReader(b))
	if err != nil {
		return "", errors.New(err.Error(), errors.ErrInternal)
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.token))
	req.Header.Set("Content-Type", "application/json")

	res, err := c.cl.Do(req)
	if err != nil {
		return "", errors.New(err.Error(), errors.ErrInternal)
	}

	b, err = io.ReadAll(res.Body)
	if err != nil {
		return "", errors.New(err.Error(), errors.ErrInternal)
	}

	var resBody response
	if err = json.Unmarshal(b, &resBody); err != nil {
		return "", errors.New(err.Error(), errors.ErrInternal)
	}

	if res.StatusCode != 200 {
		return "", errors.New(fmt.Sprintf("api response status code is not equal to 200: received %d", res.StatusCode), errors.ErrInternal)
	}

	content := strings.ReplaceAll(resBody.Choices[0].Message.Content, "```", "")
	content = strings.ReplaceAll(resBody.Choices[0].Message.Content, "yaml", "")

	return content, nil
}

type request struct {
	Model    string    `json:"model"`
	Stream   bool      `json:"stream"`
	Messages []message `json:"messages"`
}

type message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type response struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}
