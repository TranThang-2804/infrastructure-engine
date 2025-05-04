package domain

import "context"

type IacPipeline struct {
	Name        string `json:"name" yaml:"name"`
	Id          string `json:"id" yaml:"id"`
	GitProvider string `json:"provider" yaml:"provider"`
	URL         string `json:"url" yaml:"url"`
}

type IacPipelineOutput struct {
	Status     string                 `json:"status" yaml:"status"`
	OuputValue map[string]interface{} `json:"outputValue" yaml:"outputValue"`
}

type IacPipelineRepository interface {
	Trigger(c context.Context) (IacPipeline, error)
	GetPipelineOutputByUrl(c context.Context) (IacPipelineOutput, error)
}

type IacPipelineUsecase interface {
	Trigger(c context.Context) (IacPipeline, error)
	GetPipelineOutputByUrl(c context.Context) (IacPipelineOutput, error)
}
