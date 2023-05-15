package awx

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// UserService implements awx Users apis.
type UserService interface {
	List(params map[string]string) ([]*User, *ResultsList[User], error)
	GetByID(id int, params map[string]string) (*User, error)
	Create(data map[string]interface{}, params map[string]string) (*User, error)
	Update(id int, data map[string]interface{}, params map[string]string) (*User, error)
	Delete(id int) (*User, error)
	ListUserRoleEntitlements(id int, params map[string]string) ([]*ApplyRole, *ListUsersEntitlementsResponse, error)
	UpdateUserRoleEntitlement(id int, data map[string]interface{}, params map[string]string) (interface{}, error)
}

type userServiceHTTP struct {
	AWXResourceService[User]
	client *Client
}

// ListUsersResponse represents `ListUsers` endpoint response.
type ListUsersResponse struct {
	Pagination
	Results []*User `json:"results"`
}

type ListUsersEntitlementsResponse struct {
	Pagination
	Results []*ApplyRole `json:"results"`
}

const usersAPIEndpoint = "/api/v2/users/"

func (u *userServiceHTTP) ListUserRoleEntitlements(id int, params map[string]string) ([]*ApplyRole, *ListUsersEntitlementsResponse, error) {
	result := new(ListUsersEntitlementsResponse)
	endpoint := fmt.Sprintf("%s%d/roles/", usersAPIEndpoint, id)
	resp, err := u.client.Requester.GetJSON(endpoint, result, params)
	if err != nil {
		return nil, result, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, result, err
	}
	return result.Results, result, nil
}

func (u *userServiceHTTP) UpdateUserRoleEntitlement(id int, data map[string]interface{}, params map[string]string) (interface{}, error) {
	result := new(interface{})
	endpoint := fmt.Sprintf("%s%d/roles/", usersAPIEndpoint, id)
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	resp, err := u.client.Requester.PostJSON(endpoint, bytes.NewReader(payload), result, nil)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}
