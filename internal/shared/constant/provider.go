package constant

import ()

type Provider int

const (
	AWS ErrorCode = iota
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
