package domain

import "context"

type Credential struct {
  baseDomainStruct
	Name           string `json:"name"`
	Id             string `json:"id"`
	Description    string `json:"description"`
	Provider       string `json:"provider"`
	SecretProvider string `json:"secretProvider"`
}

type CredentialsUsecase interface {
	GetAll(c context.Context) ([]BluePrint, error)
}
