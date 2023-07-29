package chains

import (
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/outputparser"
	"github.com/tmc/langchaingo/prompts"
	"github.com/tmc/langchaingo/schema"
)

//nolint:lll
const _conversationTemplate = `The following is a friendly conversation between a human and an AI. The AI is talkative and provides lots of specific details from its context. If the AI does not know the answer to a question, it truthfully says it does not know.

Current conversation:
{{.history}}
Human: {{.input}}
AI:`

type ConversationOptions struct {
	Template     prompts.PromptTemplate
	OutputKey    string
	OutputParser schema.OutputParser[any]
}

type ConversationOption func(opts *ConversationOptions)

func WithConversationOptionTemplate(tpl prompts.PromptTemplate) ConversationOption {
	return func(opts *ConversationOptions) {
		opts.Template = tpl
	}
}

func WithConversationOptionOutputKey(key string) ConversationOption {
	return func(opts *ConversationOptions) {
		opts.OutputKey = key
	}
}

func WithConversationOptionOutputParser(parser schema.OutputParser[any]) ConversationOption {
	return func(opts *ConversationOptions) {
		opts.OutputParser = parser
	}
}

func NewConversation(llm llms.LanguageModel, memory schema.Memory, opts ...ConversationOption) LLMChain {
	opts_ := &ConversationOptions{
		Template: prompts.NewPromptTemplate(
			_conversationTemplate,
			[]string{"history", "input"},
		),
		OutputKey:    _llmChainDefaultOutputKey,
		OutputParser: outputparser.NewSimple(),
	}

	for _, opt := range opts {
		opt(opts_)
	}

	return LLMChain{
		Prompt:       opts_.Template,
		LLM:          llm,
		Memory:       memory,
		OutputParser: opts_.OutputParser,
		OutputKey:    opts_.OutputKey,
	}
}
