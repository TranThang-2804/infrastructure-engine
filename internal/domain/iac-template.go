package domain

import "context"

type IacTemplate struct {
  BaseDomainStruct
	Name        string               `json:"name" yaml:"name"`
	Id          string               `json:"id" yaml:"id"`
	Description string               `json:"description" yaml:"description"`
	Provider    string               `json:"provider" yaml:"provider"`
	Versions    []IacTemplateVersion `json:"versions" yaml:"versions"`
}

type IacTemplateVersion struct {
  BaseDomainStruct
	VersionName string `json:"versionName" yaml:"versionName"`
	GitUrl      string `json:"gitUrl" yaml:"gitUrl"`
	Path        string `json:"path" yaml:"path"`
	Branch      string `json:"branch" yaml:"branch"`
}

type IacTemplateRepository interface {
	GetAll(c context.Context) ([]IacTemplate, error)
}

type IacTemplateUsecase interface {
	GetAll(c context.Context) ([]IacTemplate, error)
}
