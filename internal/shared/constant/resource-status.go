package constant

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
		"Running",
		"Error",
	}[r]
}
