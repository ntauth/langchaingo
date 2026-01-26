// package vertex implements a langchaingo provider for Google Vertex AI LLMs,
// including the new Gemini models.
// See https://cloud.google.com/vertex-ai for more details.
package vertex

import (
	"context"

	"github.com/tmc/langchaingo/callbacks"
	"github.com/tmc/langchaingo/llms"
	"google.golang.org/genai"
)

type Vertex struct {
	CallbacksHandler callbacks.Handler
	client           *genai.Client
	opts             options
}

var _ llms.Model = &Vertex{}

// New creates a new Vertex client.
func New(ctx context.Context, opts ...Option) (*Vertex, error) {
	clientOptions := defaultOptions()
	for _, opt := range opts {
		opt(&clientOptions)
	}

	config := &genai.ClientConfig{
		Project:  clientOptions.cloudProject,
		Location: clientOptions.cloudLocation,
	}

	client, err := genai.NewClient(ctx, config)
	if err != nil {
		return nil, err
	}

	v := &Vertex{
		opts:   clientOptions,
		client: client,
	}
	return v, nil
}
