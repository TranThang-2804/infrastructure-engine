package repository

import (
	"context"

	"github.com/TranThang-2804/infrastructure-engine/internal/adapter/git"
	"github.com/TranThang-2804/infrastructure-engine/internal/domain"
	"github.com/TranThang-2804/infrastructure-engine/internal/shared/log"
	"gopkg.in/yaml.v3"
)

type compositeResourceRepository struct {
	gitStore git.GitStore
}

func NewCompositeResourceRepository(gitStore git.GitStore) domain.CompositeResourceRepository {
	return &compositeResourceRepository{
		gitStore: gitStore,
	}
}

func (cr *compositeResourceRepository) GetAll(c context.Context) ([]domain.CompositeResource, error) {
	var compositeResources []domain.CompositeResource

	compositeResourceContents, err := cr.gitStore.GetAllFileContentsInDirectory("TranThang-2804", "platform-iac-resource", "master", "template")

	for _, fileContent := range compositeResourceContents {
		var compositeResource domain.CompositeResource
		err = yaml.Unmarshal([]byte(fileContent), &compositeResource)
		if err != nil {
			log.Logger.Error("Error unmarshalling YAML", "error", err)
			return nil, err
		}

		compositeResources = append(compositeResources, compositeResource)
	}

	log.Logger.Debug("Blueprints Content", "content", compositeResources)

	return compositeResources, err
}

func (cr *compositeResourceRepository) Create(c context.Context, compositeResource domain.CompositeResource) (domain.CompositeResource, error) {
	// Convert object to YAML
	yamlBytes, err := yaml.Marshal(compositeResource)
	if err != nil {
    log.Logger.Error("Error converting to YAML:", "error", err)
		return compositeResource, err
	}

	// Convert YAML bytes to string
	yamlString := string(yamlBytes)
  err = cr.gitStore.CreateFile(
    "TranThang-2804", 
    "platform-iac-resource", 
    "master", 
    "template.yaml", 
    yamlString,
  )
  if err != nil {
    log.Logger.Error("Error creating file", "error", err)
    return compositeResource, err
  }

	return compositeResource, nil
}
