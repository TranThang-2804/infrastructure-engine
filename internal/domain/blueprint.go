package domain

import (
	"context"

	"github.com/TranThang-2804/infrastructure-engine/internal/shared/constant"
)

type BluePrint struct {
	Name            string            `json:"name" yaml:"Name"`
	Description     string            `json:"description" yaml:"Description"`
	Provider        constant.Provider `json:"provider" yaml:"Provider"`
	IconImageUrl    string            `json:"iconImageUrl" yaml:"IconImageUrl"`
	Versions        []VersionInfo     `json:"versions" yaml:"versions"` // Changed to a slice to match the format
	Valid           bool              `json:"valid,omitempty" yaml:"valid,omitempty"`
	ValidationError string            `json:"validationError,omitempty" yaml:"validationError,omitempty"`
}

type VersionInfo struct {
	Name                 string      `json:"name" yaml:"name"`
	JsonSchema           string      `json:"jsonSchema" yaml:"jsonSchema"`
	JsonSchemaForEditing string      `json:"jsonSchemaForEditing" yaml:"jsonSchemaForEditing"`
	UISchema             interface{} `json:"uiSchema" yaml:"uiSchema"` // Use interface{} for null or other types
	Template             Template    `json:"template" yaml:"template"`
}

type Template struct {
	TemplateRef []TemplateRef `json:"templateRef" yaml:"template-ref"`
}

type TemplateRef struct {
	Template string `json:"template" yaml:"template"`
	Version  string `json:"version" yaml:"version"`
	Value    string `json:"value" yaml:"value"`
}

type BluePrintRepository interface {
	GetAll(c context.Context) ([]BluePrint, error)
}

type GetBluePrintRequest struct {
}

type GetBluePrintResponse struct {
}

type BluePrintUsecase interface {
	GetAll(c context.Context) ([]BluePrint, error)
}
