package chroma

import (
	"context"

<<<<<<< HEAD
	chroma_go "github.com/amikos-tech/chroma-go"
	"github.com/tmc/langchaingo/embeddings"
)

var _ chroma_go.EmbeddingFunction = chromaGoEmbedder{} // compile-time check
=======
	chromatypes "github.com/amikos-tech/chroma-go/types"
	"github.com/tmc/langchaingo/embeddings"
)

var _ chromatypes.EmbeddingFunction = chromaGoEmbedder{} // compile-time check
>>>>>>> upstream/main

// chromaGoEmbedder adapts an 'embeddings.Embedder' to a 'chroma_go.EmbeddingFunction'.
type chromaGoEmbedder struct {
	embeddings.Embedder
}

<<<<<<< HEAD
// CreateEmbeddingWithModel passes through to `CreateEmbedding`; i.e., ignores 'model' (second) parameter.
func (e chromaGoEmbedder) CreateEmbeddingWithModel(texts []string, _ string) ([][]float32, error) {
	return e.CreateEmbedding(texts)
}

// CreateEmbedding creates the embedding using the current model.
func (e chromaGoEmbedder) CreateEmbedding(texts []string) ([][]float32, error) {
	return e.EmbedDocuments(context.TODO(), texts)
=======
func (e chromaGoEmbedder) EmbedDocuments(ctx context.Context, texts []string) ([]*chromatypes.Embedding, error) {
	_embeddings, err := e.Embedder.EmbedDocuments(ctx, texts)
	if err != nil {
		return nil, err
	}
	_chrmembeddings := make([]*chromatypes.Embedding, len(_embeddings))
	for i, emb := range _embeddings {
		_chrmembeddings[i] = chromatypes.NewEmbeddingFromFloat32(emb)
	}
	return _chrmembeddings, nil
}

func (e chromaGoEmbedder) EmbedQuery(ctx context.Context, text string) (*chromatypes.Embedding, error) {
	_embedding, err := e.Embedder.EmbedQuery(ctx, text)
	if err != nil {
		return nil, err
	}
	return chromatypes.NewEmbeddingFromFloat32(_embedding), nil
}

func (e chromaGoEmbedder) EmbedRecords(ctx context.Context, records []*chromatypes.Record, force bool) error {
	return chromatypes.EmbedRecordsDefaultImpl(e, ctx, records, force)
>>>>>>> upstream/main
}
