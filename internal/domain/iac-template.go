package domain

import "context"

type IacTemplate struct {
	Name        string               `json:"name"`
  Id          string               `json:"id"`
	Description string               `json:"description"`
	Provider    string               `json:"provider"`
	Versions    []IacTemplateVersion `json:"versions"`
}

type IacTemplateVersion struct {
	VersionName string `json:"versionName"`
	GitUrl      string `json:"gitUrl"`
	Path        string `json:"path"`
	Branch      string `json:"branch"`
}

type IacTemplateRepository interface {
	GetAll(c context.Context) ([]IacTemplate, error)
}

type IacTemplateUsecase interface {
	GetAll(c context.Context) ([]IacTemplate, error)
}
