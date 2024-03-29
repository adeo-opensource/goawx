package awx

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type HostService interface {
	List(params map[string]string) ([]*Host, *ResultsList[Host], error)
	GetByID(id int, params map[string]string) (*Host, error)
	Create(data map[string]interface{}, params map[string]string) (*Host, error)
	Update(id int, data map[string]interface{}, params map[string]string) (*Host, error)
	Delete(id int) (*Host, error)
	AssociateGroup(id int, data map[string]interface{}, params map[string]string) (*Host, error)
	DisAssociateGroup(id int, data map[string]interface{}, params map[string]string) (*Host, error)
}

type hostServiceHTTP struct {
	AWXResourceService[Host]
	client *Client
}

// AssociateGroup implement the awx group association request
type AssociateGroup struct {
	ID        int  `json:"id"`
	Associate bool `json:"associate"`
}

// ListHostsResponse represents `ListHosts` endpoint response.
type ListHostsResponse struct {
	Pagination
	Results []*Host `json:"results"`
}

const hostsAPIEndpoint = "/api/v2/hosts/"

// AssociateGroup update an awx Host
func (h *hostServiceHTTP) AssociateGroup(id int, data map[string]interface{}, params map[string]string) (*Host, error) {
	result := new(Host)
	endpoint := fmt.Sprintf("%s%d/groups/", hostsAPIEndpoint, id)
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
	resp, err := h.client.Requester.PostJSON(endpoint, bytes.NewReader(payload), result, nil)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}

// DisAssociateGroup update an awx Host
func (h *hostServiceHTTP) DisAssociateGroup(id int, data map[string]interface{}, params map[string]string) (*Host, error) {
	result := new(Host)
	endpoint := fmt.Sprintf("%s%d/groups/", hostsAPIEndpoint, id)
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
	resp, err := h.client.Requester.PostJSON(endpoint, bytes.NewReader(payload), result, nil)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}
