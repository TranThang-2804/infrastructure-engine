package domain

import "context"

type IacTemplate struct {
	Name        string `json:"name"`
	Version     string `json:"version"`
	Description string `json:"description"`
	Provider    string `json:"provider"`
	Url         string `json:"url"`
}

type IacTemplateRepository interface {
	GetAll(c context.Context) ([]IacTemplate, error)
}

type IacTemplateUsecase interface {
	GetAll(c context.Context) ([]IacTemplate, error)
}
