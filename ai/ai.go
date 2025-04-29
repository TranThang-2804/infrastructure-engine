package ai

import (
	"context"
	"encoding/json"
	"errors"
	"strings"

	"github.com/TranThang-2804/infrastructure-engine/shared/log"
	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/llms/openai"
	"github.com/tmc/langchaingo/prompts"
)

type OutputType int

const (
	Markdown OutputType = iota
	Text
)

func OutputTypeFromString(s string) (OutputType, error) {
	switch s {
	case "markdown":
		return Markdown, nil
	case "text":
		return Text, nil
	default:
		return -1, errors.New("invalid output type")
	}
}

func (o OutputType) String() string {
	return [...]string{
		"markdown",
		"text",
	}[o]
}

func (o OutputType) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.String())
}

// UnmarshalJSON custom unmarshaler for OutputType
func (o *OutputType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	outputType, err := OutputTypeFromString(s)
	if err != nil {
		return err
	}
	*o = outputType
	return nil
}

var prompt = prompts.NewChatPromptTemplate([]prompts.MessageFormatter{
	prompts.NewSystemMessagePromptTemplate(
		`You are a friendly AI Platform/DevOps/Cloud blog writer that help user generate a well-structured blog post in the 
    specified format based on user prompt. You might need to ask for more information to write a good blog. 
    You can use the previous context provided.
    `,
		nil,
	),
	prompts.NewHumanMessagePromptTemplate(
		`Prompt: {{.input}}
    Format: {{.format}}
    Previous context: {{.previousContext}}`,
		[]string{"input", "format", "previousContext"},
	),
})

var messageHistory []string

func GetAiResponseChain(input string, outputType OutputType) (string, error) {
	messageHistory = append(messageHistory, "User: "+input+"\n")
  log.Logger.Info("MessageHistory", "messageHistory", messageHistory)

	llm, err := openai.New(
		openai.WithModel("gpt-4o-mini"),
	)
	if err != nil {
		log.Logger.Fatal("Fatal Error", "error", err)
		return "", err
	}

	llmChain := chains.NewLLMChain(llm, prompt)
	ctx := context.Background()

	result, err := chains.Call(ctx, llmChain, map[string]any{
		"input":           input,
		"format":          outputType,
		"previousContext": strings.Join(messageHistory, "###"),
	})

	if err != nil {
		return "", err
	}

	if err != nil {
		log.Logger.Fatal("Fatal Error", "error", err)
		return "", err
	}

	messageHistory = append(messageHistory, "AI: "+result[llmChain.OutputKey].(string)+"\n")

  log.Logger.Info("MessageHistory", "messageHistory", messageHistory)

	return result[llmChain.OutputKey].(string), nil
}

