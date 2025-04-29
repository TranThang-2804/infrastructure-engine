package dto

import ()

type AITextGenerationRequestMessage struct {
	Role          string `json:"role"`
	Content       string `json:"content"`
	Logprobs      string `json:"logprobs"`
	Finish_reason string `json:"finish_reason"`
}

type Function struct {
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Parameters  map[string]interface{} `json:"parameters"`
}

type CallFunction struct {
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Arguments   map[string]interface{} `json:"arguments"`
}

type Tool struct {
	Type     string   `json:"type"`
	Function Function `json:"function"`
}

type ToolCall struct {
	Id       string       `json:"id"`
	Type     string       `json:"type"`
	Function CallFunction `json:"function"`
}

type AITextGenerationRequestMessages struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type AITextGenerationRequest struct {
	Model       string                            `json:"model"`
	Messages    []AITextGenerationRequestMessages `json:"messages"`
	MaxTokens   int32                             `json:"max_tokens"`
	Stream      bool                              `json:"stream"`
	Temperature float32                           `json:"temperature"`
	Tools       []Tool                            `json:"tools"`
}

type AITextGenerationResponseMessages struct {
	Role      string     `json:"role"`
	Content   string     `json:"content"`
	ToolCalls []ToolCall `json:"tool_calls"`
}

type AITextGenerationResponseChoices struct {
	Index   int32                            `json:"index"`
	Message AITextGenerationResponseMessages `json:"message"`
}

type AITextGenerationResponse struct {
	Object            string                            `json:"object"`
	Id                string                            `json:"id"`
	Created           int32                             `json:"created"`
	Model             string                            `json:"model"`
	SystemFingerprint string                            `json:"system_fingerprint"`
	Choices           []AITextGenerationResponseChoices `json:"choices"`
	Usage             map[string]interface{}            `json:"usage"`
}
