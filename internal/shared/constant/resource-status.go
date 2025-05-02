package constant

import (
	"encoding/json"
	"errors"
	"strings"
)

type ResourceStatus int

const (
	Pending ResourceStatus = iota
	Provisioning
	Deleting
	Deleted
	Done
	Failed
)

func (r ResourceStatus) String() string {
	return [...]string{
		"Pending",
		"Provisioning",
		"Deleting",
		"Deleted",
		"Done",
		"Failed",
	}[r]
}

// MarshalJSON converts the ResourceStatus to its string representation for JSON
func (r ResourceStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.String())
}

// UnmarshalJSON converts a string representation back to a ResourceStatus for JSON
func (r *ResourceStatus) UnmarshalJSON(data []byte) error {
	var statusStr string
	if err := json.Unmarshal(data, &statusStr); err != nil {
		return err
	}

	switch strings.ToLower(statusStr) {
	case "pending":
		*r = Pending
	case "provisioning":
		*r = Provisioning
	case "deleting":
		*r = Deleting
	case "deleted":
		*r = Deleted
	case "done":
		*r = Done
	case "failed":
		*r = Failed
	default:
		return errors.New("invalid ResourceStatus value")
	}
	return nil
}

// MarshalYAML converts the ResourceStatus to its string representation for YAML
func (r ResourceStatus) MarshalYAML() (interface{}, error) {
	return r.String(), nil
}

// UnmarshalYAML converts a string representation back to a ResourceStatus for YAML
func (r *ResourceStatus) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var statusStr string
	if err := unmarshal(&statusStr); err != nil {
		return err
	}

	switch strings.ToLower(statusStr) {
	case "pending":
		*r = Pending
	case "provisioning":
		*r = Provisioning
	case "deleting":
		*r = Deleting
	case "deleted":
		*r = Deleted
	case "done":
		*r = Done
	case "failed":
		*r = Failed
	default:
		return errors.New("invalid ResourceStatus value")
	}
	return nil
}

