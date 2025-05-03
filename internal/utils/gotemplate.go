package utils

import (
	"bytes"
	"github.com/TranThang-2804/infrastructure-engine/internal/shared/log"
	"text/template"
)

// GenerateGoTemplateOutput takes a spec map and a Go template string, executes the template with the spec data,
// and returns the resulting output as a map[string]interface{}.
func GenerateGoTemplateOutput(spec map[string]interface{}, templateStr string) (string, error) {
	// Parse the template string
	log.Logger.Debug("Parsing template", "template", templateStr)
	log.Logger.Debug("Parsing template", "spec", spec)
	tmpl, err := template.New("template").Parse(templateStr)
	if err != nil {
		return "", err
	}

	// Create a buffer to capture the template output
	var outputBuffer bytes.Buffer

	// Execute the template with the provided spec data
	if err := tmpl.Execute(&outputBuffer, spec); err != nil {
		return "", err
	}
	log.Logger.Debug("Parsing template", "output", outputBuffer.String())

	return outputBuffer.String(), nil
}
