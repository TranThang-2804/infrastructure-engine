package repository

import (
	"context"

	"github.com/TranThang-2804/infrastructure-engine/internal/adapter/git"
	"github.com/TranThang-2804/infrastructure-engine/internal/domain"
	"github.com/TranThang-2804/infrastructure-engine/internal/shared/log"
	"gopkg.in/yaml.v3"
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
	var bluePrints []domain.BluePrint

	fileContents, err := br.gitStore.GetAllFileContentsInDirectory("TranThang-2804", "platform-iac-template", "master", "blueprint")

	for _, fileContent := range fileContents {
    var bluePrint domain.BluePrint
		err = yaml.Unmarshal([]byte(fileContent), &bluePrint)
		if err != nil {
			log.Logger.Error("Error unmarshalling YAML", "error", err)
			return nil, err
		}

    bluePrints = append(bluePrints, bluePrint)
	}

	log.Logger.Debug("Blueprints Content", "content", bluePrints)

	return bluePrints, err
}
