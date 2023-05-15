package awx

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// WorkflowJobTemplateService implements awx workflow job template apis.
type WorkflowJobTemplateService interface {
	List(params map[string]string) ([]*WorkflowJobTemplate, *ResultsList[WorkflowJobTemplate], error)
	GetByID(id int, params map[string]string) (*WorkflowJobTemplate, error)
	Create(data map[string]interface{}, params map[string]string) (*WorkflowJobTemplate, error)
	Update(id int, data map[string]interface{}, params map[string]string) (*WorkflowJobTemplate, error)
	Delete(id int) (*WorkflowJobTemplate, error)
	LaunchWorkflow(id int, data map[string]interface{}, params map[string]string) (*JobLaunch, error)
}

type workflowJobTemplateServiceHTTP struct {
	AWXResourceService[WorkflowJobTemplate]
	client *Client
}

// ListWorkflowJobTemplatesResponse represents `ListWorkflowJobTemplate` endpoint response.
type ListWorkflowJobTemplatesResponse struct {
	Pagination
	Results []*WorkflowJobTemplate `json:"results"`
}

const workflowJobTemplateAPIEndpoint = "/api/v2/workflow_job_templates/"

// Launch a job with the workflow job template.
func (jt *workflowJobTemplateServiceHTTP) LaunchWorkflow(id int, data map[string]interface{}, params map[string]string) (*JobLaunch, error) {
	result := new(JobLaunch)
	endpoint := fmt.Sprintf("%s%d/launch/", workflowJobTemplateAPIEndpoint, id)
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	resp, err := jt.client.Requester.PostJSON(endpoint, bytes.NewReader(payload), result, params)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}
