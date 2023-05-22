package awx

// InstanceGroupsService implements awx execution environments apis.
type InstanceGroupService interface {
	List(params map[string]string) ([]*InstanceGroup, *ResultsList[InstanceGroup], error)
	GetByID(id int, params map[string]string) (*InstanceGroup, error)
	Create(data map[string]interface{}, params map[string]string) (*InstanceGroup, error)
	Update(id int, data map[string]interface{}, params map[string]string) (*InstanceGroup, error)
	Delete(id int) (*InstanceGroup, error)
}

type instanceGroupServiceHTTP struct {
	AWXResourceService[InstanceGroup]
	client *Client
}

// ListInstanceGroupsResponse represents `ListInstanceGroups` endpoint response.
type ListInstanceGroupsResponse struct {
	Pagination
	Results []*InstanceGroup `json:"results"`
}

const InstanceGroupsAPIEndpoint = "/api/v2/instance_groups/"
