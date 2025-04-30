package constant

import (
	"encoding/json"
	"fmt"
)

type Provider int

const (
	AWS Provider = iota
	GCP
  AZURE
  K8S
)

func (p Provider) String() string {
	return [...]string{
		"AWS",
		"GCP",
    "AZURE",
    "K8S",
	}[p]
}

// UnmarshalYAML converts a string in YAML to the Provider type
func (p *Provider) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var providerStr string
	if err := unmarshal(&providerStr); err != nil {
		return err
	}

	switch providerStr {
	case "AWS":
		*p = AWS
	case "GCP":
		*p = GCP
	case "AZURE":
		*p = AZURE
	case "K8S":
		*p = K8S
	default:
		return fmt.Errorf("invalid provider: %s", providerStr)
	}
	return nil
}

// MarshalYAML converts the Provider type to a string in YAML
func (p Provider) MarshalYAML() (interface{}, error) {
	return p.String(), nil
}

// MarshalJSON converts the Provider type to its string representation in JSON
func (p Provider) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.String())
}

// UnmarshalJSON converts a string in JSON to the Provider type
func (p *Provider) UnmarshalJSON(data []byte) error {
	var providerStr string
	if err := json.Unmarshal(data, &providerStr); err != nil {
		return err
	}

	switch providerStr {
	case "AWS":
		*p = AWS
	case "GCP":
		*p = GCP
	case "AZURE":
		*p = AZURE
	case "K8S":
		*p = K8S
	default:
		return fmt.Errorf("invalid provider: %s", providerStr)
	}
	return nil
}
