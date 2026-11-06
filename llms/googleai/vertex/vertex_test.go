package vertex

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/schema"
)

func TestSplitSystemMessages(t *testing.T) {
	t.Parallel()

	system := func(text string) llms.MessageContent {
		return llms.TextParts(schema.ChatMessageTypeSystem, text)
	}
	human := llms.TextParts(schema.ChatMessageTypeHuman, "hi")
	ai := llms.TextParts(schema.ChatMessageTypeAI, "hello")

	t.Run("no system messages", func(t *testing.T) {
		t.Parallel()
		text, rest, err := splitSystemMessages([]llms.MessageContent{human, ai})
		require.NoError(t, err)
		assert.Empty(t, text)
		assert.Equal(t, []llms.MessageContent{human, ai}, rest)
	})

	t.Run("system and human", func(t *testing.T) {
		t.Parallel()
		text, rest, err := splitSystemMessages([]llms.MessageContent{system("be brief"), human})
		require.NoError(t, err)
		assert.Equal(t, "be brief", text)
		assert.Equal(t, []llms.MessageContent{human}, rest)
	})

	t.Run("multiple system messages are concatenated", func(t *testing.T) {
		t.Parallel()
		text, rest, err := splitSystemMessages([]llms.MessageContent{system("a"), human, system("b")})
		require.NoError(t, err)
		assert.Equal(t, "a\nb", text)
		assert.Equal(t, []llms.MessageContent{human}, rest)
	})

	t.Run("only system messages", func(t *testing.T) {
		t.Parallel()
		text, rest, err := splitSystemMessages([]llms.MessageContent{system("a")})
		require.NoError(t, err)
		assert.Equal(t, "a", text)
		assert.Empty(t, rest)
	})

	t.Run("non-text system part", func(t *testing.T) {
		t.Parallel()
		_, _, err := splitSystemMessages([]llms.MessageContent{{
			Role:  schema.ChatMessageTypeSystem,
			Parts: []llms.ContentPart{llms.BinaryPart("image/png", []byte{1})},
		}})
		require.Error(t, err)
	})
}
