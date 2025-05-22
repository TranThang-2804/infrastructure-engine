package repository

import (
	"context"
	"fmt"

	"github.com/TranThang-2804/infrastructure-engine/internal/domain"
	"github.com/TranThang-2804/infrastructure-engine/internal/infrastructure/git"
	"github.com/TranThang-2804/infrastructure-engine/internal/shared/log"
	"github.com/TranThang-2804/infrastructure-engine/internal/utils"
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

func (br *bluePrintRepository) GetAll(ctx context.Context) ([]domain.BluePrint, error) {
	logger := log.BaseLogger.FromCtx(ctx).WithFields("repository", utils.GetStructName(br))
	ctx = logger.WithCtx(ctx)

	var bluePrints []domain.BluePrint

	fileContents, err := br.gitStore.GetAllFileContentsInDirectory(
		ctx,
		"TranThang-2804",
		"platform-iac-template",
		"master",
		"blueprint",
	)

	for _, fileContent := range fileContents {
		var bluePrint domain.BluePrint
		err = yaml.Unmarshal([]byte(fileContent), &bluePrint)
		if err != nil {
			logger.Error("Error unmarshalling YAML", "error", err)
			return nil, err
		}

		bluePrints = append(bluePrints, bluePrint)
	}

	logger.Debug("Blueprints Content", "content", bluePrints)

	return bluePrints, err
}

func (br *bluePrintRepository) GetById(ctx context.Context, id string) (domain.BluePrint, error) {
	logger := log.BaseLogger.FromCtx(ctx).WithFields("repository", utils.GetStructName(br))
	ctx = logger.WithCtx(ctx)

	fileContents, err := br.gitStore.GetAllFileContentsInDirectory(
		ctx,
		"TranThang-2804",
		"platform-iac-template",
		"master",
		"blueprint",
	)

	for _, fileContent := range fileContents {
		var bluePrint domain.BluePrint
		err = yaml.Unmarshal([]byte(fileContent), &bluePrint)
		if err != nil {
			logger.Error("Error unmarshalling YAML", "error", err)
			return domain.BluePrint{}, err
		}
		if bluePrint.Id == id {
			logger.Debug("Blueprint Content Found With ID", "content", bluePrint)
			return bluePrint, nil
		}
	}

	return domain.BluePrint{}, fmt.Errorf("Blueprint not found with id: %s", id)
}

func (br *bluePrintRepository) GetByIdAndVersion(
	ctx context.Context,
	id string,
	version string,
) (domain.BluePrintVersion, error) {
	logger := log.BaseLogger.FromCtx(ctx).WithFields("repository", utils.GetStructName(br))
	ctx = logger.WithCtx(ctx)

	fileContents, err := br.gitStore.GetAllFileContentsInDirectory(
		ctx,
		"TranThang-2804",
		"platform-iac-template",
		"master",
		"blueprint",
	)

	for _, fileContent := range fileContents {
		var bluePrint domain.BluePrint
		err = yaml.Unmarshal([]byte(fileContent), &bluePrint)
		if err != nil {
			logger.Error("Error unmarshalling YAML", "error", err)
			return domain.BluePrintVersion{}, err
		}
		if bluePrint.Id == id {
			logger.Debug("Blueprint Content Found With ID", "content", bluePrint)

			for _, ver := range bluePrint.Versions {
				if ver.Name == version {
					logger.Debug("Blueprint Version Found With ID")
					return ver, nil
				}
			}
			return domain.BluePrintVersion{}, nil
		}
	}

	return domain.BluePrintVersion{}, fmt.Errorf("Blueprint not found with id: %s", id)
}
