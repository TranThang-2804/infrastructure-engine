package domain

import "context"

type IacPipeline struct {
	Name              string            `json:"name"     yaml:"name"`
	Id                int               `json:"id"       yaml:"id"`
	Action            string            `json:"action"   yaml:"action"`
	GitProvider       string            `json:"provider" yaml:"provider"`
	URL               string            `json:"url"      yaml:"url"`
	IacPipelineOutput IacPipelineOutput `json:"output"   yaml:"output"`
}

type IacPipelineOutput struct {
	Status     string         `json:"status"      yaml:"status"`
	OuputValue map[string]any `json:"outputValue" yaml:"outputValue"`
}

type IacPipelineRepository interface {
	Trigger(c context.Context, iacPipeline IacPipeline) (string, error)
	GetPipelineOutputByUrl(c context.Context, iacPipeline IacPipeline) ([]byte, error)
	GetPipelineStatus(c context.Context, iacIacPipeline IacPipeline) (string, error)
	GetPipelineLog(c context.Context, iacPipeline IacPipeline) ([]byte, error)
}

type IacPipelineUsecase interface {
	Trigger(c context.Context, iacPipeline IacPipeline) (string, error)
	GetPipelineOutputByUrl(c context.Context, iacPipeline IacPipeline) ([]byte, error)
	GetPipelineStatus(c context.Context, iacIacPipeline IacPipeline) (string, error)
	GetPipelineLog(c context.Context, iacPipeline IacPipeline) ([]byte, error)
}
