package repository

import (
	"context"

	"github.com/TranThang-2804/infrastructure-engine/internal/domain"
	"github.com/TranThang-2804/infrastructure-engine/internal/infrastructure/git"
	"github.com/TranThang-2804/infrastructure-engine/internal/shared/log"
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

func (ir *iacTemplateRepository) GetAll(c context.Context) ([]domain.IacTemplate, error) {
	var iacTemplates []domain.IacTemplate

	fileContents, err := ir.gitStore.GetAllFileContentsInDirectory("TranThang-2804", "platform-iac-template", "master", "template")

	for _, fileContent := range fileContents {
		var iacTemplate domain.IacTemplate
		err = yaml.Unmarshal([]byte(fileContent), &iacTemplate)
		if err != nil {
			log.Logger.Error("Error unmarshalling YAML", "error", err)
			return nil, err
		}

		iacTemplates = append(iacTemplates, iacTemplate)
	}

	log.Logger.Debug("Blueprints Content", "content", iacTemplates)

	return iacTemplates, err
}
