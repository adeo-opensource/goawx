package awx

type ApplicationService interface {
	List(params map[string]string) ([]*Application, *ResultsList[Application], error)
	GetByID(id int, params map[string]string) (*Application, error)
	Create(data map[string]interface{}, params map[string]string) (*Application, error)
	Update(id int, data map[string]interface{}, params map[string]string) (*Application, error)
	Delete(id int) (*Application, error)
}

type applicationServiceHTTP struct {
	AWXResourceService[Application]
	client *Client
}

type ListApplicationResponse struct {
	Pagination
	Results []*Application `json:"results"`
}

const applicationsAPIEndpoint = "/api/v2/applications/"
