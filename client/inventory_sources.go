package awx

import (
	"fmt"
)

// InventorySourcesService implements awx inventory sources apis.
type InventorySourceService interface {
	List(params map[string]string) ([]*InventorySource, *ResultsList[InventorySource], error)
	GetByID(id int, params map[string]string) (*InventorySource, error)
	Create(data map[string]interface{}, params map[string]string) (*InventorySource, error)
	Update(id int, data map[string]interface{}, params map[string]string) (*InventorySource, error)
	Delete(id int) (*InventorySource, error)

	GetInventorySource(id int, params map[string]string) (*InventorySource, error)
}

type inventorySourceServiceHTTP struct {
	AWXResourceService[InventorySource]
	client *Client
}

// ListInventorySourcesResponse represents `ListInventorySources` endpoint response.
type ListInventorySourcesResponse struct {
	Pagination
	Results []*InventorySource `json:"results"`
}

const inventorySourcesAPIEndpoint = "/api/v2/inventory_sources/"

// GetInventorySource retrives the InventorySource information from its ID or Name
func (i *inventorySourceServiceHTTP) GetInventorySource(id int, params map[string]string) (*InventorySource, error) {
	endpoint := fmt.Sprintf("%s%d", inventorySourcesAPIEndpoint, id)
	result := new(InventorySource)
	resp, err := i.client.Requester.GetJSON(endpoint, result, map[string]string{})
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}
