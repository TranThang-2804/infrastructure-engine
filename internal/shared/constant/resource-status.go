package constant

type ResourceStatus string

const (
	Pending      ResourceStatus = "Pending"
	Provisioning ResourceStatus = "Provisioning"
	Deleting     ResourceStatus = "Deleting"
	Deleted      ResourceStatus = "Deleted"
	Done         ResourceStatus = "Done"
	Failed       ResourceStatus = "Failed"
)
