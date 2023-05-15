package awx

type CredentialInputSourceService interface {
	List(params map[string]string) ([]*CredentialInputSource, *ResultsList[CredentialInputSource], error)
	GetByID(id int, params map[string]string) (*CredentialInputSource, error)
	Create(data map[string]interface{}, params map[string]string) (*CredentialInputSource, error)
	Update(id int, data map[string]interface{}, params map[string]string) (*CredentialInputSource, error)
	Delete(id int) (*CredentialInputSource, error)
}

type credentialInputSourceServiceHTTP struct {
	AWXResourceService[CredentialInputSource]
	client *Client
}

type ListCredentialInputSourceResponse struct {
	Pagination
	Results []*CredentialInputSource `json:"results"`
}

const credentialInputSourceAPIEndpoint = "/api/v2/credential_input_sources/"
