package repository

import (
	"context"

	"github.com/TranThang-2804/infrastructure-engine/internal/adapter/git"
	"github.com/TranThang-2804/infrastructure-engine/internal/domain"
	"github.com/TranThang-2804/infrastructure-engine/internal/shared/log"
)

type bluePrintRepository struct {
	gitStore git.GitStore
}

func NewBluePrintRepository(gitStore git.GitStore) domain.BluePrintRepository {
	return &bluePrintRepository{
		gitStore: gitStore,
	}
}

func (br *bluePrintRepository) GetAll(c context.Context) ([]domain.BluePrint, error) {
	var bluePrint []domain.BluePrint = []domain.BluePrint{}

	fileContent, err := br.gitStore.ReadFileContent("blueprint-config-info.yaml", "TranThang-2804", "platform-iac-template", "master")

	log.Logger.Info("File Content", "content", fileContent)

	return bluePrint, err
}
