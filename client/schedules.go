package awx

// SchedulesService implements awx projects apis.
type ScheduleService interface {
	List(params map[string]string) ([]*Schedule, *ResultsList[Schedule], error)
	GetByID(id int, params map[string]string) (*Schedule, error)
	Create(data map[string]interface{}, params map[string]string) (*Schedule, error)
	Update(id int, data map[string]interface{}, params map[string]string) (*Schedule, error)
	Delete(id int) (*Schedule, error)
}

type scheduleServiceHTTP struct {
	AWXResourceService[Schedule]
	client *Client
}

// ListSchedulesResponse represents `List` endpoint response.
type ListSchedulesResponse struct {
	Pagination
	Results []*Schedule `json:"results"`
}

const schedulesAPIEndpoint = "/api/v2/schedules/"
