package domain

import "context"

type IacTemplate struct {
	Name        string `json:"name"`
	Version     string `json:"version"`
	Description string `json:"description"`
	Provider    string `json:"provider"`
	Path        string `json:"path"`
}

type IacTemplateUsecase interface {
	GetAll(c context.Context) ([]IacTemplate, error)
}
