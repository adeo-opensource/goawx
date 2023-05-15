package awx

// NotificationTemplatesService implements awx projects apis.
type NotificationTemplateService interface {
	List(params map[string]string) ([]*NotificationTemplate, *ResultsList[NotificationTemplate], error)
	GetByID(id int, params map[string]string) (*NotificationTemplate, error)
	Create(data map[string]interface{}, params map[string]string) (*NotificationTemplate, error)
	Update(id int, data map[string]interface{}, params map[string]string) (*NotificationTemplate, error)
	Delete(id int) (*NotificationTemplate, error)
}

type notificationTemplateServiceHTTP struct {
	AWXResourceService[NotificationTemplate]
	client *Client
}

// ListNotificationTemplatesResponse represents `List` endpoint response.
type ListNotificationTemplatesResponse struct {
	Pagination
	Results []*NotificationTemplate `json:"results"`
}

const notificationTemplatesAPIEndpoint = "/api/v2/notification_templates/"
