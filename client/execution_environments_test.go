package awx

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_awx_GetExecutionEnvironmentByID(t *testing.T) {
	srv := GetServer(200, map[string]interface{}{
		"id":      1,
		"name":    "myEE",
		"managed": true,
		"pull":    "missing",
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
		name    string
		fields  fields
		args    args
		want    *ExecutionEnvironment
		wantErr bool
	}{
		{
			name: "Execution environment found",
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
			want: &ExecutionEnvironment{
				ID:      1,
				Name:    "myEE",
				Managed: true,
				Pull:    "missing",
			},
			wantErr: false,
		},
		{
			name: "EE doesn't exist",
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
			p := &awx{
				client: tt.fields.client,
			}
			got, err := p.GetExecutionEnvironmentByID(tt.args.id, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetExecutionEnvironmentByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetExecutionEnvironmentByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_CreateExecutionEnvironment(t *testing.T) {
	srvCreated := GetServer(201, map[string]interface{}{
		"id":      1,
		"name":    "myEE",
		"image":   "myEEImage",
		"managed": false,
		"pull":    "always",
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
		want      *ExecutionEnvironment
		wantErr   bool
		wantError error
	}{
		{
			name: "EE created",
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
					"name":  "myEE",
					"image": "myEEImage",
				},
				params: map[string]string{},
			},
			want: &ExecutionEnvironment{
				ID:      1,
				Name:    "myEE",
				Image:   "myEEImage",
				Managed: false,
				Pull:    "always",
			},
			wantErr: false,
		},
		{
			name: "EE Missing mandatory field",
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
					"name": "myEE",
				},
				params: map[string]string{},
			},
			want:      nil,
			wantErr:   true,
			wantError: fmt.Errorf("Mandatory input arguments are absent: [image]"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &awx{
				client: tt.fields.client,
			}
			got, err := p.CreateExecutionEnvironment(tt.args.data, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateExecutionEnvironment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateExecutionEnvironment() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_DeleteExecutionEnvironment(t *testing.T) {
	srv := GetServer(204, nil)
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
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *ExecutionEnvironment
		wantErr bool
	}{
		{
			name: "Execution environment deleted",
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
				id: 1,
			},
			want:    &ExecutionEnvironment{},
			wantErr: false,
		},
		{
			name: "Not found",
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
				id: 1,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &awx{
				client: tt.fields.client,
			}
			got, err := p.DeleteExecutionEnvironment(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteExecutionEnvironment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteExecutionEnvironment() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_UpdateExecutionEnvironment(t *testing.T) {
	srvUpdated := GetServer(200, map[string]interface{}{
		"id":    1,
		"name":  "myEE",
		"image": "myNewEEImage",
	})
	defer func() {
		srvUpdated.Close()
	}()
	srvNotFound := GetServer(404, map[string]interface{}{
		"detail": "Not found.",
	})
	defer func() {
		srvNotFound.Close()
	}()

	type fields struct {
		client *Client
	}
	type args struct {
		id     int
		data   map[string]interface{}
		params map[string]string
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		want      *ExecutionEnvironment
		wantErr   bool
		wantError error
	}{
		{
			name: "EE updated",
			fields: fields{
				client: &Client{
					BaseURL: srvUpdated.URL,
					Requester: &Requester{
						Base:          srvUpdated.URL,
						Authenticator: &AuthenticatorMock{},
						Client:        srvUpdated.Client(),
					},
				},
			},
			args: args{
				id: 1,
				data: map[string]interface{}{
					"image": "myNewEEImage",
				},
				params: map[string]string{},
			},
			want: &ExecutionEnvironment{
				ID:    1,
				Name:  "myEE",
				Image: "myNewEEImage",
			},
			wantErr: false,
		},
		{
			name: "EE not found",
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
				id: 10000000,
				data: map[string]interface{}{
					"image": "myNewEEImage",
				},
				params: map[string]string{},
			},
			want:      nil,
			wantErr:   true,
			wantError: fmt.Errorf("detail: Not found"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &awx{
				client: tt.fields.client,
			}
			got, err := p.UpdateExecutionEnvironment(tt.args.id, tt.args.data, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateExecutionEnvironment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateExecutionEnvironment() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awx_ListExecutionEnvironments(t *testing.T) {
	srvEmpty := GetServer(200, map[string]interface{}{
		"count":    0,
		"next":     nil,
		"previous": nil,
		"results":  []map[string]interface{}{},
	})
	defer func() {
		srvEmpty.Close()
	}()

	srv := GetServer(200, map[string]interface{}{
		"count":    2,
		"next":     nil,
		"previous": nil,
		"results": []map[string]interface{}{
			{
				"id":    1,
				"name":  "myEE",
				"image": "myEEImage:tag",
			},
			{
				"id":    2,
				"name":  "myOtherEE",
				"image": "myEEImage:tag",
			},
		},
	})
	defer func() {
		srv.Close()
	}()

	type fields struct {
		client *Client
	}
	type args struct {
		params map[string]string
	}

	eeListOk := []*ExecutionEnvironment{
		{
			ID:    1,
			Name:  "myEE",
			Image: "myEEImage:tag",
		},
		{
			ID:    2,
			Name:  "myOtherEE",
			Image: "myEEImage:tag",
		},
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*ExecutionEnvironment
		want1   *ListExecutionEnvironmentsResponse
		wantErr bool
	}{
		{
			name: "EE empty list",
			fields: fields{
				client: &Client{
					BaseURL: srvEmpty.URL,
					Requester: &Requester{
						Base:          srvEmpty.URL,
						Authenticator: &AuthenticatorMock{},
						Client:        srvEmpty.Client(),
					},
				},
			},
			args: args{
				params: map[string]string{},
			},
			want: []*ExecutionEnvironment{},
			want1: &ListExecutionEnvironmentsResponse{
				Pagination: Pagination{
					Count:    0,
					Next:     nil,
					Previous: nil,
				},
				Results: []*ExecutionEnvironment{},
			},
			wantErr: false,
		},
		{
			name: "EE list",
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
				params: map[string]string{},
			},
			want: eeListOk,
			want1: &ListExecutionEnvironmentsResponse{
				Pagination: Pagination{
					Count:    2,
					Next:     nil,
					Previous: nil,
				},
				Results: eeListOk,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &awx{
				client: tt.fields.client,
			}
			got, got1, err := p.ListExecutionEnvironments(tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListExecutionEnvironments() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListExecutionEnvironments() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ListExecutionEnvironments() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
