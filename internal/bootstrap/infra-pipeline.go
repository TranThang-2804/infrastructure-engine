package bootstrap

import (
	"fmt"
	"io"
	"os"

	"github.com/TranThang-2804/infrastructure-engine/internal/adapter/git"
)

type InfraPipeline struct {
	gitStore git.GitStore
}

func (ip *InfraPipeline) SettingInfraPipeline() error {
	file, err := os.Open("iac-execution/Earthfile")
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

	ip.gitStore.CreateOrUpdateFile("TranThang-2804", "platform-iac-template", "master", "Earthfile", string(content))
	return nil
}
