package huggingfaceclient

import (
	"context"
	"errors"
	"fmt"
	"strings"
)

var (
	ErrInvalidToken  = errors.New("invalid token")
	ErrEmptyResponse = errors.New("empty response")
)

type Client struct {
	Token string
	Model string
	url   string
}

func New(token string, model string, url string) (*Client, error) {
	if token == "" {
		return nil, ErrInvalidToken
	}
	if url == "" {
		url = hfInferenceAPI
	}

	return &Client{
		Token: token,
		Model: model,
		url:   url,
	}, nil
}

type InferenceRequest struct {
	Model             string        `json:"repositoryId"`
	Prompt            string        `json:"prompt"`
	Task              InferenceTask `json:"task"`
	Temperature       float64       `json:"temperature,omitempty"`
	MaxNewTokens      int           `json:"max_new_tokens"`
	DoSample          bool          `json:"do_sample"`
	TopP              float64       `json:"top_p,omitempty"`
	TopK              int           `json:"top_k,omitempty"`
	MinLength         int           `json:"min_length,omitempty"`
	MaxLength         int           `json:"max_length,omitempty"`
	RepetitionPenalty float64       `json:"repetition_penalty,omitempty"`
	Seed              int           `json:"seed,omitempty"`
	StopWords         []string      `json:"-"`
}

type InferenceResponse struct {
	Text string `json:"generated_text"`
}

func (c *Client) RunInference(ctx context.Context, request *InferenceRequest) (*InferenceResponse, error) {
	payload := &inferencePayload{
		Model:  request.Model,
		Inputs: request.Prompt,
		Parameters: parameters{
			Temperature:       request.Temperature,
			MaxNewTokens:      request.MaxNewTokens,
			DoSample:          request.DoSample,
			TopP:              request.TopP,
			TopK:              request.TopK,
			MinLength:         request.MinLength,
			MaxLength:         request.MaxLength,
			RepetitionPenalty: request.RepetitionPenalty,
			Seed:              request.Seed,
		},
	}
	resp, err := c.runInference(ctx, payload)
	if err != nil {
		return nil, fmt.Errorf("failed to run inference: %w", err)
	}
	if len(resp) == 0 {
		return nil, ErrEmptyResponse
	}
	text := resp[0].Text

	// Emulate stop words
	for _, stopWord := range payload.Parameters.StopWords {
		splits := strings.Split(text, stopWord)
		if len(splits) > 0 {
			text = splits[0]
			break
		}
	}

	// TODO: Add response cleaning based on Model.
	// e.g., for gpt2, text = text[len(request.Prompt)+1:]
	return &InferenceResponse{
		Text: text,
	}, nil
}
