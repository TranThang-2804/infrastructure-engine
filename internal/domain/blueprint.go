package domain

import (
	"context"

	"github.com/TranThang-2804/infrastructure-engine/internal/shared/constant"
)

type BluePrint struct {
	Name            string             `json:"name" yaml:"name"`
	Id              string             `json:"id" yaml:"id"`
	Description     string             `json:"description" yaml:"description"`
	Provider        constant.Provider  `json:"provider" yaml:"provider"`
	IconImageUrl    string             `json:"iconImageUrl" yaml:"iconImageUrl"`
	Versions        []BluePrintVersion `json:"versions" yaml:"versions"`
	Valid           bool               `json:"valid,omitempty" yaml:"valid,omitempty"`
	ValidationError string             `json:"validationError,omitempty" yaml:"validationError,omitempty"`
}

type BluePrintVersion struct {
	Name                 string                       `json:"versionName" yaml:"versionName"`
	JsonSchema           string                       `json:"jsonSchema" yaml:"jsonSchema"`
	JsonSchemaForEditing string                       `json:"jsonSchemaForEditing" yaml:"jsonSchemaForEditing"`
	UiSchema             string                       `json:"uiSchema,omitempty" yaml:"uiSchema,omitempty"`
	CompositeTemplate    []BluePrintCompositeTemplate `json:"compositeTemplate" yaml:"compositeTemplate"`
}

type BluePrintCompositeTemplate struct {
	TemplateId    string `json:"templateId" yaml:"templateId"`
	Version       string `json:"version" yaml:"version"`
	ValueTemplate string `json:"valueTemplate" yaml:"valueTemplate"`
}

type GetBluePrintRequest struct {
}

type GetBluePrintResponse struct {
}

type BluePrintRepository interface {
	GetAll(c context.Context) ([]BluePrint, error)
	GetById(c context.Context, id string) (BluePrint, error)
}

type BluePrintUsecase interface {
	GetAll(c context.Context) ([]BluePrint, error)
	GetById(c context.Context, id string) (BluePrint, error)
}
