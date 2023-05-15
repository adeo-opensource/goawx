package awx

// ProjectService implements awx projects apis.
type ProjectService interface {
	List(params map[string]string) ([]*Project, *ResultsList[Project], error)
	GetByID(id int, params map[string]string) (*Project, error)
	Create(data map[string]interface{}, params map[string]string) (*Project, error)
	Update(id int, data map[string]interface{}, params map[string]string) (*Project, error)
	Delete(id int) (*Project, error)
}

type projectServiceHTTP struct {
	AWXResourceService[Project]
	client *Client
}

// ListProjectsResponse represents `ListProjects` endpoint response.
type ListProjectsResponse struct {
	Pagination
	Results []*Project `json:"results"`
}

const projectsAPIEndpoint = "/api/v2/projects/"
