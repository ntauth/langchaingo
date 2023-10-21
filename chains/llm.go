package chains

import (
	"context"
	"fmt"
	"os"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/memory"
	"github.com/tmc/langchaingo/outputparser"
	"github.com/tmc/langchaingo/prompts"
	"github.com/tmc/langchaingo/schema"
)

const _llmChainDefaultOutputKey = "text"

type LLMChain struct {
	Prompt       prompts.FormatPrompter
	LLM          llms.LanguageModel
	Memory       schema.Memory
	OutputParser schema.OutputParser[any]

	OutputKey string
}

var _ Chain = &LLMChain{}

// NewLLMChain creates a new LLMChain with an llm and a prompt.
func NewLLMChain(llm llms.LanguageModel, prompt prompts.FormatPrompter) *LLMChain {
	chain := &LLMChain{
		Prompt:       prompt,
		LLM:          llm,
		OutputParser: outputparser.NewSimple(),
		Memory:       memory.NewSimple(),

		OutputKey: _llmChainDefaultOutputKey,
	}

	return chain
}

// Call formats the prompts with the input values, generates using the llm, and parses
// the output from the llm with the output parser. This function should not be called
// directly, use rather the Call or Run function if the prompt only requires one input
// value.
func (c LLMChain) Call(ctx context.Context, values map[string]any, options ...ChainCallOption) (map[string]any, error) {
	promptValue, err := c.Prompt.FormatPrompt(values)
	if err != nil {
		return nil, err
	}

	if jsonData, err := os.ReadFile("/Users/chihiro/Desktop/data.json"); err != nil {
		return nil, err
	} else {
		fmt.Println(jsonData)
	}

	result, err := c.LLM.GeneratePrompt(
		ctx,
		[]schema.PromptValue{promptValue},
		getLLMCallOptions(options...)...,
	)
	if err != nil {
		return nil, err
	}

	finalOutput, err := c.OutputParser.ParseWithPrompt(result.Generations[0][0].Text, promptValue)
	if err != nil {
		return nil, err
	}

	return map[string]any{c.OutputKey: finalOutput}, nil
}

// GetMemory returns the memory.
func (c LLMChain) GetMemory() schema.Memory { //nolint:ireturn
	return c.Memory //nolint:ireturn
}

// GetInputKeys returns the expected input keys.
func (c LLMChain) GetInputKeys() []string {
	return append([]string{}, c.Prompt.GetInputVariables()...)
}

// GetOutputKeys returns the output keys the chain will return.
func (c LLMChain) GetOutputKeys() []string {
	return []string{c.OutputKey}
}
