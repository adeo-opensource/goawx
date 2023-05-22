package awx

import (
	"fmt"
	"net/http"
)

// AWX represents awx api endpoints with services, and using
// client to communicate with awx server.
type AWX struct {
	ApplicationService                              ApplicationService
	CredentialService                               CredentialService
	CredentialTypeService                           CredentialTypeService
	CredentialInputSourceService                    CredentialInputSourceService
	ExecutionEnvironmentService                     ExecutionEnvironmentService
	GroupService                                    GroupService
	HostService                                     HostService
	InstanceGroupService                            InstanceGroupService
	InventoryService                                InventoryService
	InventorySourceService                          InventorySourceService
	JobService                                      JobService
	JobTemplateService                              JobTemplateService
	JobTemplateNotificationTemplatesService         JobTemplateNotificationTemplateService
	NotificationTemplatesService                    NotificationTemplateService
	OrganizationService                             OrganizationService
	PingService                                     PingService
	ProjectService                                  ProjectService
	ProjectUpdatesService                           ProjectUpdateService
	TeamService                                     TeamService
	ScheduleService                                 ScheduleService
	SettingService                                  SettingService
	UserService                                     UserService
	WorkflowJobTemplateService                      WorkflowJobTemplateService
	WorkflowJobTemplateNodeService                  WorkflowJobTemplateNodeService
	WorkflowJobTemplateNodeStepService              WorkflowJobTemplateNodeStepService
	WorkflowJobTemplateScheduleService              WorkflowJobTemplateScheduleService
	WorkflowJobTemplateNotificationTemplatesService WorkflowJobTemplateNotificationTemplateService
}

// Client implement http client.
type Client struct {
	BaseURL   string
	Requester *Requester
}

// CheckResponse do http response check, and return err if not in [200, 300).
func CheckResponse(resp *http.Response) error {
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return nil
	}

	return fmt.Errorf("responsed with %d, resp: %v", resp.StatusCode, resp)
}

// ValidateParams is to validate the input to use the services.
func ValidateParams(data map[string]interface{}, mandatoryFields []string) (notfound []string, status bool) {
	status = true
	for _, key := range mandatoryFields {
		_, exists := data[key]

		if !exists {
			notfound = append(notfound, key)
			status = false
		}
	}
	return notfound, status
}

// NewAWX news an awx handler with basic auth support, you could customize the http
// transport by passing custom client.
func NewAWX(baseURL, userName, passwd string, client *http.Client) (*AWX, error) {
	r := &Requester{Base: baseURL, Authenticator: &BasicAuth{Username: userName, Password: passwd}, Client: client} // pragma: allowlist secret
	if r.Client == nil {
		r.Client = http.DefaultClient
	}

	awxClient := &Client{
		BaseURL:   baseURL,
		Requester: r,
	}

	newAWX := newAWX(awxClient)

	// test the connection and return and error if there's an issue
	_, err := newAWX.PingService.Ping()
	if err != nil {
		return nil, err
	}

	return newAWX, nil
}

// NewAWXToken creates an AWX handler with token support.
func NewAWXToken(baseURL, token string, client *http.Client) (*AWX, error) {
	r := &Requester{Base: baseURL, Authenticator: &TokenAuth{Token: token}, Client: client}
	if r.Client == nil {
		r.Client = http.DefaultClient
	}

	awxClient := &Client{
		BaseURL:   baseURL,
		Requester: r,
	}

	newAWX := newAWX(awxClient)
	//newAWX.TeamService.List(xxxx)
	// test the connection and return and error if there's an issue
	_, err := newAWX.PingService.Ping()
	if err != nil {
		return nil, err
	}

	return newAWX, nil
}

func newAWX(c *Client) *AWX {
	return &AWX{
		ApplicationService: &applicationServiceHTTP{
			AWXResourceService: NewAWXResourceService[Application](c, applicationsAPIEndpoint, []string{"name", "client_type", "authorization_grant_type", "organization"}),
			client:             c,
		},
		CredentialService: &credentialServiceHTTP{
			AWXResourceService: NewAWXResourceService[Credential](c, credentialsAPIEndpoint, []string{}),
			client:             c,
		},
		CredentialInputSourceService: &credentialInputSourceServiceHTTP{
			AWXResourceService: NewAWXResourceService[CredentialInputSource](c, credentialInputSourceAPIEndpoint, []string{}),
			client:             c,
		},
		CredentialTypeService: &credentialTypeServiceHTTP{
			AWXResourceService: NewAWXResourceService[CredentialType](c, credentialTypesAPIEndpoint, []string{}),
			client:             c,
		},
		ExecutionEnvironmentService: &executionEnvironmentServiceHTTP{
			AWXResourceService: NewAWXResourceService[ExecutionEnvironment](c, executionEnvironmentsAPIEndpoint, []string{"name", "inventory"}),
			client:             c,
		},
		GroupService: &groupServiceHTTP{
			AWXResourceService: NewAWXResourceService[Group](c, groupsAPIEndpoint, []string{"name", "image"}),
			client:             c,
		},
		HostService: &hostServiceHTTP{
			AWXResourceService: NewAWXResourceService[Host](c, hostsAPIEndpoint, []string{"name", "inventory"}),
			client:             c,
		},
		InstanceGroupService: &instanceGroupServiceHTTP{
			AWXResourceService: NewAWXResourceService[InstanceGroup](c, InstanceGroupsAPIEndpoint, []string{"name"}),
			client:             c,
		},
		InventoryService: &inventoryServiceHTTP{
			AWXResourceService: NewAWXResourceService[Inventory](c, inventoriesAPIEndpoint, []string{"name", "organization"}),
			client:             c,
		},
		InventorySourceService: &inventorySourceServiceHTTP{
			AWXResourceService: NewAWXResourceService[InventorySource](c, inventorySourcesAPIEndpoint, []string{"name", "inventory"}),
			client:             c,
		},
		JobService: &jobServiceHTTP{
			client: c,
		},
		JobTemplateService: &jobTemplateServiceHTTP{
			AWXResourceService: NewAWXResourceService[JobTemplate](c, jobTemplatesAPIEndpoint, []string{"name", "job_type", "inventory", "project"}),
			client:             c,
		},
		JobTemplateNotificationTemplatesService: &jobTemplateNotificationTemplateServiceHTTP{
			client: c,
		},
		NotificationTemplatesService: &notificationTemplateServiceHTTP{
			AWXResourceService: NewAWXResourceService[NotificationTemplate](c, notificationTemplatesAPIEndpoint, []string{"name", "organization", "notification_type"}),
			client:             c,
		},
		OrganizationService: &organizationServiceHTTP{
			AWXResourceService: NewAWXResourceService[Organization](c, organizationsAPIEndpoint, []string{"name"}),
			client:             c,
		},
		PingService: &pingServiceHTTP{
			client: c,
		},
		ProjectService: &projectServiceHTTP{
			AWXResourceService: NewAWXResourceService[Project](c, projectsAPIEndpoint, []string{"name", "organization", "scm_type"}),
			client:             c,
		},
		ProjectUpdatesService: &projectUpdateServiceHTTP{
			client: c,
		},
		TeamService: &teamServiceHTTP{
			AWXResourceService: NewAWXResourceService[Team](c, teamsAPIEndpoint, []string{"name", "organization"}),
			client:             c,
		},
		ScheduleService: &scheduleServiceHTTP{
			AWXResourceService: NewAWXResourceService[Schedule](c, schedulesAPIEndpoint, []string{"name", "rrule", "unified_job_template"}),
			client:             c,
		},
		SettingService: &settingServiceHTTP{
			client: c,
		},
		UserService: &userServiceHTTP{
			AWXResourceService: NewAWXResourceService[User](c, usersAPIEndpoint, []string{"username", "password", "first_name", "last_name", "email"}),
			client:             c,
		},
		WorkflowJobTemplateService: &workflowJobTemplateServiceHTTP{
			AWXResourceService: NewAWXResourceService[WorkflowJobTemplate](c, workflowJobTemplateAPIEndpoint, []string{"name"}),
			client:             c,
		},
		WorkflowJobTemplateNodeService: &workflowJobTemplateNodeServiceHTTP{
			AWXResourceService: NewAWXResourceService[WorkflowJobTemplateNode](c, workflowJobTemplateNodeAPIEndpoint, []string{"workflow_job_template", "unified_job_template", "identifier"}),
			client:             c,
		},
		WorkflowJobTemplateNodeStepService: &workflowJobTemplateNodeStepServiceHTTP{
			client: c,
		},
		WorkflowJobTemplateScheduleService: &workflowJobTemplateScheduleServiceHTTP{
			client: c,
		},
		WorkflowJobTemplateNotificationTemplatesService: &workflowJobTemplateNotificationTemplateServiceHTTP{
			client: c,
		},
	}
}
