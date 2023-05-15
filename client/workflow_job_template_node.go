package awx

// WorkflowJobTemplateNodeService implements awx job template node apis.
type WorkflowJobTemplateNodeService interface {
	List(params map[string]string) ([]*WorkflowJobTemplateNode, *ResultsList[WorkflowJobTemplateNode], error)
	GetByID(id int, params map[string]string) (*WorkflowJobTemplateNode, error)
	Create(data map[string]interface{}, params map[string]string) (*WorkflowJobTemplateNode, error)
	Update(id int, data map[string]interface{}, params map[string]string) (*WorkflowJobTemplateNode, error)
	Delete(id int) (*WorkflowJobTemplateNode, error)
}

type workflowJobTemplateNodeServiceHTTP struct {
	AWXResourceService[WorkflowJobTemplateNode]
	client *Client
}

// ListWorkflowJobTemplateNodesResponse represents `ListWorkflowJobTemplateNodes` endpoint response.
type ListWorkflowJobTemplateNodesResponse struct {
	Pagination
	Results []*WorkflowJobTemplateNode `json:"results"`
}

const workflowJobTemplateNodeAPIEndpoint = "/api/v2/workflow_job_template_nodes/"
