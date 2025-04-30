package domain

import "context"

type BluePrint struct {
	Name              string `json:"name"`
	Id                string `json:"id"`
	Version           string `json:"version"`
	OpenAPI           string `json:"openapi"`
	OpenAPIForEditing string `json:"openapiForEditing"`
	UiSchema          string `json:"uiSchema"`
	Template          string `json:"template"`
	Description       string `json:"description"`
	Provider          string `json:"provider"`
	IconImageUrl      string `json:"iconImageUrl"`
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
