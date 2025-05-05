package repository

import (
	"context"
	"encoding/json"

	"github.com/TranThang-2804/infrastructure-engine/internal/domain"
	"github.com/TranThang-2804/infrastructure-engine/internal/infrastructure/git"
	"github.com/TranThang-2804/infrastructure-engine/internal/shared/log"
)

type iacPipelineRepository struct {
	gitStore git.GitStore
}

func NewIacPipelineRepository(gitStore git.GitStore) domain.IacPipelineRepository {
	return &iacPipelineRepository{
		gitStore: gitStore,
	}
}

func (ir *iacPipelineRepository) Trigger(c context.Context, iacPipeline domain.IacPipeline) (string, error) {
	pipelineData := map[string]interface{}{
		"action":   iacPipeline.Action,
		"filepath": iacPipeline.Name,
	}
	pipelinePayload, err := json.Marshal(pipelineData)
	if err != nil {
		return "", err
	}

	// Trigger pipeline
	pipelineUrl, err := ir.gitStore.TriggerPipeline(
		"TranThang-2804",
		"platform-iac-resource",
		pipelinePayload,
	)
	if err != nil {
		log.Logger.Error("Error Triggering pipeline", "error", err)
		return "", err
	}

	// Get Pipeline URL and attach it
	return pipelineUrl, nil
}

func (ir *iacPipelineRepository) GetPipelineOutputByUrl(c context.Context, iacPipeline domain.IacPipeline) ([]byte, error) {
	return []byte{}, nil
}

func (ir *iacPipelineRepository) GetPipelineStatus(c context.Context, iacIacPipeline domain.IacPipeline) (string, error) {
	return "Running", nil
}

func (ir *iacPipelineRepository) GetPipelineLog(c context.Context, iacPipeline domain.IacPipeline) ([]byte, error) {
	return []byte{}, nil
}
