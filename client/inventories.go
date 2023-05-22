package awx

import "fmt"

// InventoriesService implements awx inventories apis.
type InventoryService interface {
	List(params map[string]string) ([]*Inventory, *ResultsList[Inventory], error)
	GetByID(id int, params map[string]string) (*Inventory, error)
	Create(data map[string]interface{}, params map[string]string) (*Inventory, error)
	Update(id int, data map[string]interface{}, params map[string]string) (*Inventory, error)
	Delete(id int) (*Inventory, error)
	ListInventoryGroups(id int, params map[string]string) ([]*Group, *ListGroupsResponse, error)
}

type inventoryServiceHTTP struct {
	AWXResourceService[Inventory]
	client *Client
}

// ListInventoriesResponse represents `ListInventories` endpoint response.
type ListInventoriesResponse struct {
	Pagination
	Results []*Inventory `json:"results"`
}

const inventoriesAPIEndpoint = "/api/v2/inventories/"

func (i *inventoryServiceHTTP) ListInventoryGroups(id int, params map[string]string) ([]*Group, *ListGroupsResponse, error) {
	result := new(ListGroupsResponse)
	endpoint := fmt.Sprintf("%s%d/groups/", inventoriesAPIEndpoint, id)
	resp, err := i.client.Requester.GetJSON(endpoint, result, params)
	if err != nil {
		return nil, result, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, result, err
	}

	return result.Results, result, nil
}
