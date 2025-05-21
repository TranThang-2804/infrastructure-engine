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

func (cr *compositeResourceRepository) GetAll(ctx context.Context) ([]domain.CompositeResource, error) {
	logger := log.BaseLogger.FromCtx(ctx).WithFields("repository", utils.GetStructName(cr))
	ctx = logger.WithCtx(ctx)

	var compositeResources []domain.CompositeResource

	compositeResourceContents, err := cr.gitStore.GetAllFileContentsInDirectory(ctx, "TranThang-2804", "platform-iac-resource", "master", "template")

	for _, fileContent := range compositeResourceContents {
		var compositeResource domain.CompositeResource
		err = yaml.Unmarshal([]byte(fileContent), &compositeResource)
		if err != nil {
			logger.Error("Error unmarshalling YAML", "error", err)
			return nil, err
		}

		compositeResources = append(compositeResources, compositeResource)
	}

	logger.Debug("Blueprints Content", "content", compositeResources)

	return compositeResources, err
}

func (cr *compositeResourceRepository) Create(ctx context.Context, compositeResource domain.CompositeResource) (domain.CompositeResource, error) {
	logger := log.BaseLogger.FromCtx(ctx).WithFields("repository", utils.GetStructName(cr))
	ctx = logger.WithCtx(ctx)

	// Validate compositeResource
	logger.Debug("CompositeResource", "compositeResource", compositeResource)
	err := utils.ValidateStruct(compositeResource)
	if err != nil {
		logger.Error("Error validating composite resource", "error", err)
		return compositeResource, err
	}

	// Convert object to YAML
	yamlString, err := convertToYaml(ctx, compositeResource)
	if err != nil {
		logger.Error("Error converting to YAML:", "error", err)
		return compositeResource, err
	}

	// Get filepath from metadata
	filepath := generateFilePath(compositeResource)

	err = cr.gitStore.CreateFile(
		ctx,
		"TranThang-2804",
		"platform-iac-resource",
		"master",
		filepath,
		yamlString,
	)
	if err != nil {
		logger.Error("Error creating file", "error", err)
		return compositeResource, err
	}

	return compositeResource, nil
}

func (cr *compositeResourceRepository) Update(ctx context.Context, compositeResource domain.CompositeResource) (domain.CompositeResource, error) {
	logger := log.BaseLogger.FromCtx(ctx).WithFields("repository", utils.GetStructName(cr))
	ctx = logger.WithCtx(ctx)

	// Validate compositeResource
	logger.Debug("CompositeResource", "compositeResource", compositeResource)
	err := utils.ValidateStruct(compositeResource)
	if err != nil {
		logger.Error("Error validating composite resource", "error", err)
		return compositeResource, err
	}

	// Convert object to YAML
	yamlString, err := convertToYaml(ctx, compositeResource)
	if err != nil {
		logger.Error("Error converting to YAML:", "error", err)
		return compositeResource, err
	}

	// Get filepath from metadata
	filepath := generateFilePath(compositeResource)

	// Check if current file exists
	_, err = cr.gitStore.ReadFileContent(
		ctx,
		"TranThang-2804",
		"platform-iac-resource",
		"master",
		filepath,
	)
	if err != nil {
		logger.Error("Error reading current file content", "error", err)
		return compositeResource, err
	}

	// Update file
	err = cr.gitStore.CreateOrUpdateFile(
		ctx,
		"TranThang-2804",
		"platform-iac-resource",
		"master",
		filepath,
		yamlString,
	)
	if err != nil {
		logger.Error("Error updating file", "error", err)
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

func convertToYaml(ctx context.Context, compositeResource domain.CompositeResource) (string, error) {
	logger := log.BaseLogger.FromCtx(ctx).WithFields("function", "convertToYaml")

	yamlBytes, err := yaml.Marshal(compositeResource)
	if err != nil {
		logger.Error("Error converting to YAML:", "error", err)
		return "", err
	}

	// Convert YAML bytes to string
	yamlString := string(yamlBytes)
	return yamlString, nil
}
