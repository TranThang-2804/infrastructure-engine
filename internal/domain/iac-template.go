package domain

import "context"

type IacTemplate struct {
	Name        string `json:"name"`
	Version     string `json:"version"`
	Description string `json:"description"`
	Provider    string `json:"provider"`
	Url         string `json:"url"`
}

type IacTemplateUsecase interface {
	GetAll(c context.Context) ([]IacTemplate, error)
}
