package awx

type ExecutionEnvironmentService interface {
	List(params map[string]string) ([]*ExecutionEnvironment, *ResultsList[ExecutionEnvironment], error)
	GetByID(id int, params map[string]string) (*ExecutionEnvironment, error)
	Create(data map[string]interface{}, params map[string]string) (*ExecutionEnvironment, error)
	Update(id int, data map[string]interface{}, params map[string]string) (*ExecutionEnvironment, error)
	Delete(id int) (*ExecutionEnvironment, error)
}

type executionEnvironmentServiceHTTP struct {
	AWXResourceService[ExecutionEnvironment]
	client *Client
}

// ListExecutionEnvironmentsResponse represents `ListExecutionEnvironments` endpoint response.
type ListExecutionEnvironmentsResponse struct {
	Pagination
	Results []*ExecutionEnvironment `json:"results"`
}

const executionEnvironmentsAPIEndpoint = "/api/v2/execution_environments/"
