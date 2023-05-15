package awx

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// OrganizationsService implements awx organizations apis.
type OrganizationService interface {
	List(params map[string]string) ([]*Organization, *ResultsList[Organization], error)
	GetByID(id int, params map[string]string) (*Organization, error)
	Create(data map[string]interface{}, params map[string]string) (*Organization, error)
	Update(id int, data map[string]interface{}, params map[string]string) (*Organization, error)
	Delete(id int) (*Organization, error)
	DisAssociateGalaxyCredentials(id int, data map[string]interface{}, params map[string]string) (*Organization, error)
	AssociateGalaxyCredentials(id int, data map[string]interface{}, params map[string]string) (*Organization, error)
}

type organizationServiceHTTP struct {
	AWXResourceService[Organization]
	client *Client
}

// ListOrganizationsResponse represents `ListOrganizations` endpoint response.
type ListOrganizationsResponse struct {
	Pagination
	Results []*Organization `json:"results"`
}

const organizationsAPIEndpoint = "/api/v2/organizations/"

// DisAssociateGalaxyCredentials remove Credentials form an awx job template
func (p *organizationServiceHTTP) DisAssociateGalaxyCredentials(id int, data map[string]interface{}, params map[string]string) (*Organization, error) {
	result := new(Organization)
	endpoint := fmt.Sprintf("%s%d/galaxy_credentials/", organizationsAPIEndpoint, id)
	data["disassociate"] = true
	mandatoryFields := []string{"id", "disassociate"}
	validate, status := ValidateParams(data, mandatoryFields)
	if !status {
		err := fmt.Errorf("mandatory input arguments are absent: %s", validate)
		return nil, err
	}
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	resp, err := p.client.Requester.PostJSON(endpoint, bytes.NewReader(payload), result, nil)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}

// AssociateGalaxyCredentials adding credentials to Organization.
func (p *organizationServiceHTTP) AssociateGalaxyCredentials(id int, data map[string]interface{}, params map[string]string) (*Organization, error) {
	result := new(Organization)

	endpoint := fmt.Sprintf("%s%d/galaxy_credentials/", organizationsAPIEndpoint, id)
	data["associate"] = true
	mandatoryFields := []string{"id"}
	validate, status := ValidateParams(data, mandatoryFields)
	if !status {
		err := fmt.Errorf("mandatory input arguments are absent: %s", validate)
		return nil, err
	}
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	resp, err := p.client.Requester.PostJSON(endpoint, bytes.NewReader(payload), result, nil)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}
