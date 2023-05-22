package awx

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

// JobTemplateService implements awx job template apis.
type JobTemplateService interface {
	List(params map[string]string) ([]*JobTemplate, *ResultsList[JobTemplate], error)
	GetByID(id int, params map[string]string) (*JobTemplate, error)
	Create(data map[string]interface{}, params map[string]string) (*JobTemplate, error)
	Update(id int, data map[string]interface{}, params map[string]string) (*JobTemplate, error)
	Delete(id int) (*JobTemplate, error)

	LaunchJob(id int, data map[string]interface{}, params map[string]string) (*JobLaunch, error)
	DisAssociateCredentials(id int, data map[string]interface{}, params map[string]string) (*JobTemplate, error)
	AssociateCredentials(id int, data map[string]interface{}, params map[string]string) (*JobTemplate, error)
}

type jobTemplateServiceHTTP struct {
	AWXResourceService[JobTemplate]
	client *Client
}

// ListJobTemplatesResponse represents `ListJobTemplates` endpoint response.
type ListJobTemplatesResponse struct {
	Pagination
	Results []*JobTemplate `json:"results"`
}

const jobTemplatesAPIEndpoint = "/api/v2/job_templates/"

// Launch lauchs a job with the job template.
func (jt *jobTemplateServiceHTTP) LaunchJob(id int, data map[string]interface{}, params map[string]string) (*JobLaunch, error) {
	result := new(JobLaunch)
	endpoint := fmt.Sprintf("%s%d/launch/", jobTemplatesAPIEndpoint, id)
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

	// in case invalid job id return
	if result.Job == 0 {
		return nil, errors.New("invalid job id 0")
	}

	return result, nil
}

// DisAssociateCredentials remove Credentials form an awx job template
func (jt *jobTemplateServiceHTTP) DisAssociateCredentials(id int, data map[string]interface{}, params map[string]string) (*JobTemplate, error) {
	result := new(JobTemplate)
	endpoint := fmt.Sprintf("%s%d/credentials/", jobTemplatesAPIEndpoint, id)
	data["disassociate"] = true
	mandatoryFields := []string{"id"}
	validate, status := ValidateParams(data, mandatoryFields)
	if !status {
		err := fmt.Errorf("Mandatory input arguments are absent: %s", validate)
		return nil, err
	}
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	resp, err := jt.client.Requester.PostJSON(endpoint, bytes.NewReader(payload), result, nil)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}

// AssociateCredentials  adding credentials to JobTemplate.
func (jt *jobTemplateServiceHTTP) AssociateCredentials(id int, data map[string]interface{}, params map[string]string) (*JobTemplate, error) {
	result := new(JobTemplate)

	endpoint := fmt.Sprintf("%s%d/credentials/", jobTemplatesAPIEndpoint, id)
	data["associate"] = true
	mandatoryFields := []string{"id"}
	validate, status := ValidateParams(data, mandatoryFields)
	if !status {
		err := fmt.Errorf("Mandatory input arguments are absent: %s", validate)
		return nil, err
	}
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	resp, err := jt.client.Requester.PostJSON(endpoint, bytes.NewReader(payload), result, nil)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}
