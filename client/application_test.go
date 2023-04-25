package awx

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

type AuthenticatorMock struct{}

func (am *AuthenticatorMock) addAuthenticationHeaders(*http.Request) {}

func Test_awx_GetApplicationByID(t *testing.T) {
	srv := GetServer(200, map[string]interface{}{
		"name": "myApp",
	})
	defer func() {
		srv.Close()
	}()
	srvErr := GetServer(400, map[string]interface{}{
		"name": "myApp",
	})
	defer func() {
		srvErr.Close()
	}()
	srvErr100 := GetServer(http.StatusFound, map[string]interface{}{
		"name": "myApp",
	})
	defer func() {
		srvErr100.Close()
	}()

	type fields struct {
		client *Client
	}
	type args struct {
		id     int
		params map[string]string
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		want      *Application
		wantErr   bool
		wantError error
	}{
		{
			name: "Application found",
			fields: fields{
				client: &Client{
					BaseURL: srv.URL,
					Requester: &Requester{
						Base:          srv.URL,
						Authenticator: &AuthenticatorMock{},
						Client:        srv.Client(),
					},
				},
			},
			args: args{
				id:     1,
				params: nil,
			},
			want:    &Application{Name: "myApp"},
			wantErr: false,
		},
		{
			name: "Error while requesting",
			fields: fields{
				client: &Client{
					BaseURL: srvErr.URL,
					Requester: &Requester{
						Base:          srvErr.URL,
						Authenticator: &AuthenticatorMock{},
						Client:        srvErr.Client(),
					},
				},
			},
			args: args{
				id:     1,
				params: nil,
			},
			want:      nil,
			wantErr:   true,
			wantError: fmt.Errorf("Errors:\n- name: []"),
		},
		{
			name: "Response status is not 2XX",
			fields: fields{
				client: &Client{
					BaseURL: srvErr100.URL,
					Requester: &Requester{
						Base:          srvErr100.URL,
						Authenticator: &AuthenticatorMock{},
						Client:        srvErr100.Client(),
					},
				},
			},
			args: args{
				id:     1,
				params: nil,
			},
			want:      nil,
			wantErr:   true,
			wantError: fmt.Errorf("Errors:\n- name: []"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &awx{
				client: tt.fields.client,
			}
			got, err := c.GetApplicationByID(tt.args.id, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetApplicationByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetApplicationByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}
