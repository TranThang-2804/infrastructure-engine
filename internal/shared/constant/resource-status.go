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

type ResourceMQName string

const (
	ToPending      ResourceMQName = "composite-resource.pending"
	ToProvisioning ResourceMQName = "composite-resource.provisioning"
	ToDeleting     ResourceMQName = "composite-resource.deleting"
)
