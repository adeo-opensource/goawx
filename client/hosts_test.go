package awx

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_awx_GetHostByID(t *testing.T) {
	srv := GetServer(200, map[string]interface{}{
		"id":        1,
		"name":      "myHost",
		"inventory": 1,
	})
	defer func() {
		srv.Close()
	}()

	srvNotFound := GetServer(404, nil)
	defer func() {
		srvNotFound.Close()
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
		want      *Host
		wantErr   bool
		wantError error
	}{
		{
			name: "Host found",
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
				params: map[string]string{},
			},
			want: &Host{
				ID:        1,
				Name:      "myHost",
				Inventory: 1,
			},
			wantErr: false,
		},
		{
			name: "Host doesn't exist",
			fields: fields{
				client: &Client{
					BaseURL: srvNotFound.URL,
					Requester: &Requester{
						Base:          srvNotFound.URL,
						Authenticator: &AuthenticatorMock{},
						Client:        srvNotFound.Client(),
					},
				},
			},
			args: args{
				id:     100000,
				params: map[string]string{},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &awx{
				client: tt.fields.client,
			}
			got, err := h.GetHostByID(tt.args.id, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetHostByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetHostByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_CreateHost(t *testing.T) {
	srvCreated := GetServer(201, map[string]interface{}{
		"id":        1,
		"name":      "myHost",
		"inventory": 1,
		"enabled":   true,
	})
	defer func() {
		srvCreated.Close()
	}()

	srvMandatoryField := GetServer(400, nil)
	defer func() {
		srvMandatoryField.Close()
	}()

	type fields struct {
		client *Client
	}
	type args struct {
		data   map[string]interface{}
		params map[string]string
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		want      *Host
		wantErr   bool
		wantError error
	}{
		{
			name: "Host created",
			fields: fields{
				client: &Client{
					BaseURL: srvCreated.URL,
					Requester: &Requester{
						Base:          srvCreated.URL,
						Authenticator: &AuthenticatorMock{},
						Client:        srvCreated.Client(),
					},
				},
			},
			args: args{
				data: map[string]interface{}{
					"name":      "myHost",
					"inventory": 1,
				},
				params: map[string]string{},
			},
			want: &Host{
				ID:        1,
				Name:      "myHost",
				Inventory: 1,
				Enabled:   true,
			},
			wantErr: false,
		},
		{
			name: "Host missing mandatory field",
			fields: fields{
				client: &Client{
					BaseURL: srvMandatoryField.URL,
					Requester: &Requester{
						Base:          srvMandatoryField.URL,
						Authenticator: &AuthenticatorMock{},
						Client:        srvMandatoryField.Client(),
					},
				},
			},
			args: args{
				data: map[string]interface{}{
					"name": "myHost",
				},
				params: map[string]string{},
			},
			want:      nil,
			wantErr:   true,
			wantError: fmt.Errorf("Mandatory input arguments are absent: [inventory]"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &awx{
				client: tt.fields.client,
			}
			got, err := p.CreateHost(tt.args.data, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateHost() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateHost() got = %v, want %v", got, tt.want)
			}
		})
	}
}
