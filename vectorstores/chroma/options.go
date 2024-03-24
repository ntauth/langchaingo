package chroma

import (
	"errors"
	"fmt"
	"os"

<<<<<<< HEAD
	chromago "github.com/amikos-tech/chroma-go"
=======
	chromatypes "github.com/amikos-tech/chroma-go/types"
>>>>>>> upstream/main
	"github.com/tmc/langchaingo/embeddings"
)

const (
<<<<<<< HEAD
	OpenAiAPIKeyEnvVarName = "OPENAI_API_KEY" // #nosec G101
	ChromaURLKeyEnvVarName = "CHROMA_URL"
	DefaultNameSpace       = "langchain"
	DefaultNameSpaceKey    = "nameSpace"
	DefaultDistanceFunc    = chromago.L2
=======
	OpenAIAPIKeyEnvVarName = "OPENAI_API_KEY" // #nosec G101
	OpenAIOrgIDEnvVarName  = "OPENAI_ORGANIZATION"
	ChromaURLKeyEnvVarName = "CHROMA_URL"
	DefaultNameSpace       = "langchain"
	DefaultNameSpaceKey    = "nameSpace"
	DefaultDistanceFunc    = chromatypes.L2
>>>>>>> upstream/main
)

// ErrInvalidOptions is returned when the options given are invalid.
var ErrInvalidOptions = errors.New("invalid options")

// Option is a function type that can be used to modify the client.
type Option func(p *Store)

// WithNameSpace sets the nameSpace used to upsert and query the vectors from.
func WithNameSpace(nameSpace string) Option {
	return func(p *Store) {
		p.nameSpace = nameSpace
	}
}

// WithChromaURL is an option for specifying the Chroma URL. Must be set.
func WithChromaURL(chromaURL string) Option {
	return func(p *Store) {
		p.chromaURL = chromaURL
	}
}

// WithEmbedder is an option for setting the embedder to use.
func WithEmbedder(e embeddings.Embedder) Option {
	return func(p *Store) {
		p.embedder = e
	}
}

<<<<<<< HEAD
// WithDistanceFunction specifies the distance function which will be used
// see: https://github.com/amikos-tech/chroma-go/blob/d0087270239eccdb2f4f03d84b18d875c601ad6b/main.go#L96
func WithDistanceFunction(distanceFunction chromago.DistanceFunction) Option {
=======
// WithDistanceFunction specifies the distance function which will be used (default is L2)
// see: https://github.com/amikos-tech/chroma-go/blob/ab1339d0ee1a863be7d6773bcdedc1cfd08e3d77/types/types.go#L22
func WithDistanceFunction(distanceFunction chromatypes.DistanceFunction) Option {
>>>>>>> upstream/main
	return func(p *Store) {
		p.distanceFunction = distanceFunction
	}
}

// WithIncludes is an option for setting the includes to query the vectors.
<<<<<<< HEAD
func WithIncludes(includes []chromago.QueryEnum) Option {
=======
func WithIncludes(includes []chromatypes.QueryEnum) Option {
>>>>>>> upstream/main
	return func(p *Store) {
		p.includes = includes
	}
}

<<<<<<< HEAD
// WithOpenAiAPIKey is an option for setting the OpenAI api key. If the option is not set
// the api key is read from the OPENAI_API_KEY environment variable. If the
// variable is not present, an error will be returned.
func WithOpenAiAPIKey(openAiAPIKey string) Option {
=======
// WithOpenAIAPIKey is an option for setting the OpenAI api key. If the option is not set
// the api key is read from the OPENAI_API_KEY environment variable. If the
// variable is not present, an error will be returned.
func WithOpenAIAPIKey(openAiAPIKey string) Option {
>>>>>>> upstream/main
	return func(p *Store) {
		p.openaiAPIKey = openAiAPIKey
	}
}

<<<<<<< HEAD
func applyClientOptions(opts ...Option) (Store, error) {
	o := &Store{
		nameSpace:        DefaultNameSpace,
		nameSpaceKey:     DefaultNameSpaceKey,
		distanceFunction: DefaultDistanceFunc,
		openaiAPIKey:     os.Getenv(OpenAiAPIKeyEnvVarName),
=======
// WithOpenAIOrganization is an option for setting the OpenAI organization id.
func WithOpenAIOrganization(openAiOrganization string) Option {
	return func(p *Store) {
		p.openaiOrganization = openAiOrganization
	}
}

func applyClientOptions(opts ...Option) (Store, error) {
	o := &Store{
		nameSpace:          DefaultNameSpace,
		nameSpaceKey:       DefaultNameSpaceKey,
		distanceFunction:   DefaultDistanceFunc,
		openaiAPIKey:       os.Getenv(OpenAIAPIKeyEnvVarName),
		openaiOrganization: os.Getenv(OpenAIOrgIDEnvVarName),
>>>>>>> upstream/main
	}

	for _, opt := range opts {
		opt(o)
	}

	if o.chromaURL == "" {
		o.chromaURL = os.Getenv(ChromaURLKeyEnvVarName)
		if o.chromaURL == "" {
			return Store{}, fmt.Errorf(
				"%w: missing chroma URL. Pass it as an option or set the %s environment variable",
				ErrInvalidOptions, ChromaURLKeyEnvVarName)
		}
	}

	// a embedder or an openai api key must be provided
	if o.openaiAPIKey == "" && o.embedder == nil {
		return Store{}, fmt.Errorf("%w: missing embedder or openai api key", ErrInvalidOptions)
	}

	return *o, nil
}
