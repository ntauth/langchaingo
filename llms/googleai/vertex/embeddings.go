package vertex

import (
	"context"
	"fmt"

	"google.golang.org/genai"
)

// CreateEmbedding creates embeddings from texts.
func (g *Vertex) CreateEmbedding(ctx context.Context, texts []string) ([][]float32, error) {
	var contents []*genai.Content
	for _, text := range texts {
		contents = append(contents, genai.NewContentFromText(text, genai.RoleUser))
	}
	res, err := g.client.Models.EmbedContent(
		ctx,
		g.opts.defaultEmbeddingModel,
		contents,
		&genai.EmbedContentConfig{},
	)
	if err != nil {
		return [][]float32{}, err
	}

	embeddingValues := make([][]float32, 0, len(res.Embeddings))
	for _, embedding := range res.Embeddings {
		embeddingValues = append(embeddingValues, embedding.Values)
	}

	if len(texts) != len(embeddingValues) {
		return embeddingValues, fmt.Errorf("returned %d embeddings for %d texts", len(embeddingValues), len(texts))
	}

	return embeddingValues, nil
}
