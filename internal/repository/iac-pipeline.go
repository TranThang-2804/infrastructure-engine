package repository

import (
	"context"

	"github.com/TranThang-2804/infrastructure-engine/internal/domain"
	"github.com/TranThang-2804/infrastructure-engine/internal/infrastructure/git"
	"github.com/TranThang-2804/infrastructure-engine/internal/shared/log"
	"github.com/TranThang-2804/infrastructure-engine/internal/utils"
)

type iacPipelineRepository struct {
	gitStore git.GitStore
}

func NewIacPipelineRepository(gitStore git.GitStore) domain.IacPipelineRepository {
	return &iacPipelineRepository{
		gitStore: gitStore,
	}
}

func (ir *iacPipelineRepository) Trigger(ctx context.Context, iacPipeline domain.IacPipeline) (string, error) {
	logger := log.BaseLogger.FromCtx(ctx).WithFields("repository", utils.GetStructName(ir))
	ctx = logger.WithCtx(ctx)

	pipelineData := map[string]any{
		"action":   iacPipeline.Action,
		"filepath": iacPipeline.Name,
	}

	// Trigger pipeline
	pipelineUrl, err := ir.gitStore.TriggerPipeline(
		ctx,
		"dev2die",
		"platform-iac-resource",
		pipelineData,
	)
	if err != nil {
		logger.Error("Error Triggering pipeline", "error", err)
		return "", err
	}

	// Get Pipeline URL and attach it
	return pipelineUrl, nil
}

func (ir *iacPipelineRepository) GetPipelineOutputByUrl(ctx context.Context, iacPipeline domain.IacPipeline) ([]byte, error) {
	return []byte{}, nil
}

func (ir *iacPipelineRepository) GetPipelineStatus(ctx context.Context, iacIacPipeline domain.IacPipeline) (string, error) {
	return "Running", nil
}

func (ir *iacPipelineRepository) GetPipelineLog(ctx context.Context, iacPipeline domain.IacPipeline) ([]byte, error) {
	return []byte{}, nil
}
