package domain

import "context"

type IacPipeline struct {
	Name        string `json:"name" yaml:"name"`
	Id          int    `json:"id" yaml:"id"`
	Action      string `json:"action" yaml:"action"`
	GitProvider string `json:"provider" yaml:"provider"`
	URL         string `json:"url" yaml:"url"`
}

type IacPipelineOutput struct {
	Status     string                 `json:"status" yaml:"status"`
	OuputValue map[string]interface{} `json:"outputValue" yaml:"outputValue"`
}

type IacPipelineRepository interface {
	Trigger(c context.Context, iacPipeline IacPipeline) (IacPipeline, error)
	GetPipelineOutputByUrl(c context.Context, iacPipeline IacPipeline) (IacPipelineOutput, error)
}

type IacPipelineUsecase interface {
	Trigger(c context.Context, iacPipeline IacPipeline) (IacPipeline, error)
	GetPipelineOutputByUrl(c context.Context, iacPipeline IacPipeline) (IacPipelineOutput, error)
}
