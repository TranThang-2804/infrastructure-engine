package repository

import (
	"context"

	"github.com/TranThang-2804/infrastructure-engine/internal/domain"
	"github.com/TranThang-2804/infrastructure-engine/internal/infrastructure/git"
	"github.com/TranThang-2804/infrastructure-engine/internal/shared/log"
	"github.com/TranThang-2804/infrastructure-engine/internal/utils"
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
			log.BaseLogger.Error("Error unmarshalling YAML", "error", err)
			return nil, err
		}

		compositeResources = append(compositeResources, compositeResource)
	}

	log.BaseLogger.Debug("Blueprints Content", "content", compositeResources)

	return compositeResources, err
}

func (cr *compositeResourceRepository) Create(c context.Context, compositeResource domain.CompositeResource) (domain.CompositeResource, error) {
	// Validate compositeResource
	log.BaseLogger.Debug("CompositeResource", "compositeResource", compositeResource)
	err := utils.ValidateStruct(compositeResource)
	if err != nil {
		log.BaseLogger.Error("Error validating composite resource", "error", err)
		return compositeResource, err
	}

	// Convert object to YAML
	yamlString, err := convertToYaml(compositeResource)
	if err != nil {
		log.BaseLogger.Error("Error converting to YAML:", "error", err)
		return compositeResource, err
	}

	// Get filepath from metadata
	filepath := generateFilePath(compositeResource)

	err = cr.gitStore.CreateFile(
		"TranThang-2804",
		"platform-iac-resource",
		"master",
		filepath,
		yamlString,
	)
	if err != nil {
		log.BaseLogger.Error("Error creating file", "error", err)
		return compositeResource, err
	}

	return compositeResource, nil
}

func (cr *compositeResourceRepository) Update(c context.Context, compositeResource domain.CompositeResource) (domain.CompositeResource, error) {
	// Validate compositeResource
	log.BaseLogger.Debug("CompositeResource", "compositeResource", compositeResource)
	err := utils.ValidateStruct(compositeResource)
	if err != nil {
		log.BaseLogger.Error("Error validating composite resource", "error", err)
		return compositeResource, err
	}

	// Convert object to YAML
	yamlString, err := convertToYaml(compositeResource)
	if err != nil {
		log.BaseLogger.Error("Error converting to YAML:", "error", err)
		return compositeResource, err
	}

	// Get filepath from metadata
	filepath := generateFilePath(compositeResource)

	// Check if current file exists
	_, err = cr.gitStore.ReadFileContent(
		"TranThang-2804",
		"platform-iac-resource",
		"master",
		filepath,
	)
	if err != nil {
		log.BaseLogger.Error("Error reading current file content", "error", err)
		return compositeResource, err
	}

	// Update file
	err = cr.gitStore.CreateOrUpdateFile(
		"TranThang-2804",
		"platform-iac-resource",
		"master",
		filepath,
		yamlString,
	)
	if err != nil {
		log.BaseLogger.Error("Error updating file", "error", err)
		return compositeResource, err
	}

	return compositeResource, nil
}

func (cr *compositeResourceRepository) Delete(c context.Context, compositeResource domain.CompositeResource) (domain.CompositeResource, error) {
	return compositeResource, nil
}

func generateFilePath(compositeResource domain.CompositeResource) string {
	// default filepath if no metadata provided
	filepath := ""
	if compositeResource.Metadata.Project != "" {
		filepath += compositeResource.Metadata.Project + "/"
	} else {
		filepath += "default/"
	}

	if compositeResource.Metadata.Group != "" {
		filepath += compositeResource.Metadata.Group + "/"
	} else {
		filepath += "default/"
	}

	filepath += compositeResource.BluePrintId + "/" + compositeResource.Id + ".yaml"
	return filepath
}

func convertToYaml(compositeResource domain.CompositeResource) (string, error) {
	yamlBytes, err := yaml.Marshal(compositeResource)
	if err != nil {
		log.BaseLogger.Error("Error converting to YAML:", "error", err)
		return "", err
	}

	// Convert YAML bytes to string
	yamlString := string(yamlBytes)
	return yamlString, nil
}
