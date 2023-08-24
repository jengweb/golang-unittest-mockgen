package services

import (
	"encoding/json"
	"errors"
	"main/gateways"
	"main/gateways/models"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
)

func Test_usersService_GetUsers(t *testing.T) {
	respUsers := []byte(`{
		"page": 1,
		"per_page": 6,
		"total": 12,
		"total_pages": 2,
		"data": [
			{
				"id": 1,
				"email": "george.bluth@reqres.in",
				"first_name": "George",
				"last_name": "Bluth",
				"avatar": "https://reqres.in/img/faces/1-image.jpg"
			}
		]
	}`)
	var mockRespUsers models.UsersResponse
	json.Unmarshal(respUsers, &mockRespUsers)
	type fields struct {
		mockUsersGateways *gateways.MockUsersGateways
	}
	tests := []struct {
		name     string
		fields   fields
		mockFunc func(f *fields)
		want     *models.UsersResponse
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			name: "GetUsers Success",
			mockFunc: func(f *fields) {
				gomock.InOrder(
					f.mockUsersGateways.EXPECT().GetUsers().Return(&mockRespUsers, nil),
				)
			},
			want:    &mockRespUsers,
			wantErr: false,
		},
		{
			name: "GetUsers Error",
			mockFunc: func(f *fields) {
				gomock.InOrder(
					f.mockUsersGateways.EXPECT().GetUsers().Return(nil, errors.New("GetUsers Error")),
				)
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctl := gomock.NewController(t)
			defer ctl.Finish()
			f := fields{
				mockUsersGateways: gateways.NewMockUsersGateways(ctl),
			}
			if tt.mockFunc != nil {
				tt.mockFunc(&f)
			}
			s := &usersService{
				usersGateways: f.mockUsersGateways,
			}
			got, err := s.GetUsers()
			if (err != nil) != tt.wantErr {
				t.Errorf("usersService.GetUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("usersService.GetUsers() = %v, want %v", got, tt.want)
			}
		})
	}
}
