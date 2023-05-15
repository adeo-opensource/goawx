package awx

type CredentialTypeService interface {
	List(params map[string]string) ([]*CredentialType, *ResultsList[CredentialType], error)
	Create(data map[string]interface{}, params map[string]string) (*CredentialType, error)
	GetByID(id int, params map[string]string) (*CredentialType, error)
	Update(id int, data map[string]interface{}, params map[string]string) (*CredentialType, error)
	Delete(id int) (*CredentialType, error)
}

type credentialTypeServiceHTTP struct {
	AWXResourceService[CredentialType]
	client *Client
}

type ListCredentialTypeResponse struct {
	Pagination
	Results []*CredentialType `json:"results"`
}

const credentialTypesAPIEndpoint = "/api/v2/credential_types/"
