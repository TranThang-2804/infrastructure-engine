package repository

import (
	"context"
	"fmt"

	"github.com/TranThang-2804/infrastructure-engine/internal/domain"
	"github.com/TranThang-2804/infrastructure-engine/internal/infrastructure/git"
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

func (br *bluePrintRepository) GetById(c context.Context, id string) (domain.BluePrint, error) {
	fileContents, err := br.gitStore.GetAllFileContentsInDirectory("TranThang-2804", "platform-iac-template", "master", "blueprint")

	for _, fileContent := range fileContents {
		var bluePrint domain.BluePrint
		err = yaml.Unmarshal([]byte(fileContent), &bluePrint)
		if err != nil {
			log.Logger.Error("Error unmarshalling YAML", "error", err)
			return domain.BluePrint{}, err
		}
		if bluePrint.Id == id {
			log.Logger.Debug("Blueprint Content Found With ID", "content", bluePrint, "id", id)
			return bluePrint, nil
		}
	}

	return domain.BluePrint{}, fmt.Errorf("Blueprint not found with id: %s", id)
}

func (br *bluePrintRepository) GetByIdAndVersion(c context.Context, id string, version string) (domain.BluePrintVersion, error) {
	fileContents, err := br.gitStore.GetAllFileContentsInDirectory("TranThang-2804", "platform-iac-template", "master", "blueprint")

	for _, fileContent := range fileContents {
		var bluePrint domain.BluePrint
		err = yaml.Unmarshal([]byte(fileContent), &bluePrint)
		if err != nil {
			log.Logger.Error("Error unmarshalling YAML", "error", err)
			return domain.BluePrintVersion{}, err
		}
		if bluePrint.Id == id {
			log.Logger.Debug("Blueprint Content Found With ID", "content", bluePrint, "id", id)

			for _, ver := range bluePrint.Versions {
				if ver.Name == version {
					log.Logger.Debug("Blueprint Version Found With ID", "content", version, "id", id)
					return ver, nil
				}
			}
			return domain.BluePrintVersion{}, nil
		}
	}

	return domain.BluePrintVersion{}, fmt.Errorf("Blueprint not found with id: %s", id)
}
