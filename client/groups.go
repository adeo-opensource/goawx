package awx

type GroupService interface {
	List(params map[string]string) ([]*Group, *ResultsList[Group], error)
	GetByID(id int, params map[string]string) (*Group, error)
	Create(data map[string]interface{}, params map[string]string) (*Group, error)
	Update(id int, data map[string]interface{}, params map[string]string) (*Group, error)
	Delete(id int) (*Group, error)
}

type groupServiceHTTP struct {
	AWXResourceService[Group]
	client *Client
}

// ListGroupsResponse represents `ListGroups` endpoint response.
type ListGroupsResponse struct {
	Pagination
	Results []*Group `json:"results"`
}

const groupsAPIEndpoint = "/api/v2/groups/"
