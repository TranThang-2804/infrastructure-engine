package bootstrap

import (
	"fmt"
	"io"
	"os"

	"github.com/TranThang-2804/infrastructure-engine/internal/infrastructure/git"
	"github.com/TranThang-2804/infrastructure-engine/internal/shared/env"
	"github.com/TranThang-2804/infrastructure-engine/internal/shared/log"
)

type InfraPipeline struct {
	gitStore git.GitStore
}

type PipelineFileMapping struct {
	filePath       string
	remoteFilePath string
}

type GitPipelineFileConfig struct {
	pipelineConfigurationFiles []PipelineFileMapping
}

var githubPipelineFileConfig = GitPipelineFileConfig{
	pipelineConfigurationFiles: []PipelineFileMapping{
		{
			filePath:       "iac-execution/Earthfile",
			remoteFilePath: "Earthfile",
		},
		{
			filePath:       "iac-execution/github/.github/workflows/action.yml",
			remoteFilePath: ".github/workflows/action.yml",
		},
	},
}

var gitlabPipelineFileConfig = GitPipelineFileConfig{
	pipelineConfigurationFiles: []PipelineFileMapping{
		{
			filePath:       "iac-execution/Earthfile",
			remoteFilePath: "Earthfile",
		},
		{
			filePath:       "iac-execution/gitlab/.gitlab-ci.yml",
			remoteFilePath: ".gitlab-ci.yml",
		},
	},
}

func NewInfraPipeline(gitStore git.GitStore) InfraPipeline {
	return InfraPipeline{
		gitStore: gitStore,
	}
}

func (ip *InfraPipeline) SettingInfraPipeline() error {
	log.BaseLogger.Info("Setting up infrastructure pipeline...")

	var pipelineFileConfig GitPipelineFileConfig

	switch env.Env.CI {
	case "github":
		pipelineFileConfig = githubPipelineFileConfig
	case "gitlab":
		pipelineFileConfig = gitlabPipelineFileConfig
	default:
		log.BaseLogger.Fatal("Unsupported CI/CD platform: ", env.Env.CI)
	}

	for _, fileMapping := range pipelineFileConfig.pipelineConfigurationFiles {
		file, err := os.Open(fileMapping.filePath)
		if err != nil {
			fmt.Printf("Error opening file: %v\n", err)
			return err
		}
		defer file.Close()

		content, err := io.ReadAll(file)
		if err != nil {
			fmt.Printf("Error reading file: %v\n", err)
			return nil
		}

		ip.gitStore.CreateOrUpdateFile("TranThang-2804", "platform-iac-resource", "master", fileMapping.remoteFilePath, string(content))
	}
	return nil
}
