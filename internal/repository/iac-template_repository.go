package repository

import (
	"context"

	"github.com/TranThang-2804/infrastructure-engine/internal/domain"
	"github.com/TranThang-2804/infrastructure-engine/internal/infrastructure/git"
	"github.com/TranThang-2804/infrastructure-engine/internal/shared/log"
	"github.com/TranThang-2804/infrastructure-engine/internal/utils"
	"gopkg.in/yaml.v3"
)

type iacTemplateRepository struct {
	gitStore git.GitStore
}

func NewIacTemplateRepository(gitStore git.GitStore) domain.IacTemplateRepository {
	return &iacTemplateRepository{
		gitStore: gitStore,
	}
}

func (ir *iacTemplateRepository) GetAll(ctx context.Context) ([]domain.IacTemplate, error) {
	logger := log.BaseLogger.FromCtx(ctx).WithFields("repository", utils.GetStructName(ir))
	ctx = logger.WithCtx(ctx)
	var iacTemplates []domain.IacTemplate

	fileContents, err := ir.gitStore.GetAllFileContentsInDirectory(ctx, "TranThang-2804", "platform-iac-template", "master", "template")

	for _, fileContent := range fileContents {
		var iacTemplate domain.IacTemplate
		err = yaml.Unmarshal([]byte(fileContent), &iacTemplate)
		if err != nil {
			logger.Error("Error unmarshalling YAML", "error", err)
			return nil, err
		}

		iacTemplates = append(iacTemplates, iacTemplate)
	}

	logger.Debug("Blueprints Content", "content", iacTemplates)

	return iacTemplates, err
}
