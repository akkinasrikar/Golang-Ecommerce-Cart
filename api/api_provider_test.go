package api

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"reflect"
	"testing"

	"github.com/akkinasrikar/ecommerce-cart/api/dto"
	"github.com/akkinasrikar/ecommerce-cart/api/mocks"
	"github.com/akkinasrikar/ecommerce-cart/models"
	"github.com/akkinasrikar/ecommerce-cart/utils"
	"github.com/akkinasrikar/ecommerce-cart/validators/helper"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
)

func Test_service_GetItems(t *testing.T) {
	type fields struct {
		BaseURL string
		http    HttpCall
	}
	type args struct {
		ecomCtx context.Context
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		mockHTTPClient func(ctrl *gomock.Controller) *mocks.MockHttpCall
		want           dto.ItemsResponse
		wantErr        models.EcomError
	}{
		{
			name: "Happy Case",
			args: args{
				ecomCtx: utils.SetContext(),
			},
			mockHTTPClient: func(ctrl *gomock.Controller) *mocks.MockHttpCall {
				mockHTTPClient := mocks.NewMockHttpCall(ctrl)
				json := `[]`
				r := io.NopCloser(bytes.NewReader([]byte(json)))
				mockHTTPClient.EXPECT().Do(gomock.Any()).Return(&http.Response{StatusCode: http.StatusOK, Body: r}, nil)
				return mockHTTPClient
			},
			want:    dto.ItemsResponse{},
			wantErr: models.EcomError{},
		},
		{
			name: "Error Case",
			args: args{
				ecomCtx: utils.SetContext(),
			},
			mockHTTPClient: func(ctrl *gomock.Controller) *mocks.MockHttpCall {
				mockHTTPClient := mocks.NewMockHttpCall(ctrl)
				mockHTTPClient.EXPECT().Do(gomock.Any()).Return(nil, errors.New("Error"))
				return mockHTTPClient
			},
			want:    dto.ItemsResponse{},
			wantErr: *helper.ErrorInternalSystemError("Error in client http call"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			tt.fields.http = tt.mockHTTPClient(ctrl)
			s := &service{
				BaseURL: tt.fields.BaseURL,
				http:    tt.fields.http,
			}
			got, err := s.GetItems(tt.args.ecomCtx)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.GetItems() got = %v, want %v", got, tt.want)
			}
			if err.Message != nil && (err.Message.Error() != tt.wantErr.Message.Error()) {
				t.Errorf("service.GetItems() got1 = %v, want %v", err, tt.wantErr)
			}
		})
	}
}
