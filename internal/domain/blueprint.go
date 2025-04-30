package domain

import (
	"context"

	"github.com/TranThang-2804/infrastructure-engine/internal/shared/constant"
)

type BluePrint struct {
	Name            string            `json:"name" yaml:"name"`
	Id              string            `json:"id" yaml:"id"`
	Description     string            `json:"description" yaml:"description"`
	Provider        constant.Provider `json:"provider" yaml:"provider"`
	IconImageUrl    string            `json:"iconImageUrl" yaml:"iconImageUrl"`
	Versions        []Version         `json:"versions" yaml:"versions"`
	Valid           bool              `json:"valid,omitempty" yaml:"valid,omitempty"`
	ValidationError string            `json:"validationError,omitempty" yaml:"validationError,omitempty"`
}

type Version struct {
	Name                 string              `json:"name" yaml:"name"`
	JsonSchema           string              `json:"jsonSchema" yaml:"jsonSchema"`
	JsonSchemaForEditing string              `json:"jsonSchemaForEditing" yaml:"jsonSchemaForEditing"`
	UiSchema             string              `json:"uiSchema,omitempty" yaml:"uiSchema,omitempty"`
	CompositeTemplate    []CompositeTemplate `json:"compositeTemplate" yaml:"compositeTemplate"`
}

type CompositeTemplate struct {
	TemplateId    string `json:"templateId" yaml:"templateId"`
	Version       string `json:"version" yaml:"version"`
	ValueTemplate string `json:"valueTemplate" yaml:"valueTemplate"`
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
