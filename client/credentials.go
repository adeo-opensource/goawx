package awx

type CredentialService interface {
	List(params map[string]string) ([]*Credential, *ResultsList[Credential], error)
	GetByID(id int, params map[string]string) (*Credential, error)
	Create(data map[string]interface{}, params map[string]string) (*Credential, error)
	Update(id int, data map[string]interface{}, params map[string]string) (*Credential, error)
	Delete(id int) (*Credential, error)
}

type credentialServiceHTTP struct {
	AWXResourceService[Credential]
	client *Client
}

type ListCredentialsResponse struct {
	Pagination
	Results []*Credential `json:"results"`
}

const credentialsAPIEndpoint = "/api/v2/credentials/"
