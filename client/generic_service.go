package awx

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type AWXResourceService[T any] struct {
	client          *Client
	basePath        string
	mandatoryFields []string
}

func NewAWXResourceService[T any](client *Client, basepath string, mandatoryFields []string) AWXResourceService[T] {
	return AWXResourceService[T]{
		client:          client,
		basePath:        basepath,
		mandatoryFields: mandatoryFields,
	}
}

type ResultsList[T any] struct {
	Pagination
	Results []*T `json:"results"`
}

func (rs *AWXResourceService[T]) List(params map[string]string) ([]*T, *ResultsList[T], error) {
	result := new(ResultsList[T])
	resp, err := rs.client.Requester.GetJSON(rs.basePath, result, params)
	if err != nil {
		return nil, result, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, result, err
	}

	return result.Results, result, nil
}

func (rs *AWXResourceService[T]) GetByID(id int, params map[string]string) (*T, error) {
	result := new(T)
	endpoint := fmt.Sprintf("%s%d/", rs.basePath, id)
	resp, err := rs.client.Requester.GetJSON(endpoint, result, params)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}
func (rs *AWXResourceService[T]) Create(data map[string]interface{}, params map[string]string) (*T, error) {
	validate, status := ValidateParams(data, rs.mandatoryFields)

	if !status {
		err := fmt.Errorf("mandatory input arguments are absent: %s", validate)
		return nil, err
	}

	result := new(T)
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	resp, err := rs.client.Requester.PostJSON(rs.basePath, bytes.NewReader(payload), result, params)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}

func (rs *AWXResourceService[T]) Update(id int, data map[string]interface{}, params map[string]string) (*T, error) {
	result := new(T)
	endpoint := fmt.Sprintf("%s%d", rs.basePath, id)
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	resp, err := rs.client.Requester.PatchJSON(endpoint, bytes.NewReader(payload), result, nil)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}

func (rs *AWXResourceService[T]) Delete(id int) (*T, error) {
	result := new(T)
	endpoint := fmt.Sprintf("%s%d", rs.basePath, id)

	resp, err := rs.client.Requester.Delete(endpoint, result, nil)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}
