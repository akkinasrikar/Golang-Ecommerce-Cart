package services

import (
	"errors"
	"reflect"
	"testing"

	"github.com/akkinasrikar/ecommerce-cart/models"
	"github.com/akkinasrikar/ecommerce-cart/models/entities"
	"github.com/akkinasrikar/ecommerce-cart/models/responses"
	"github.com/akkinasrikar/ecommerce-cart/repositories"
	"github.com/akkinasrikar/ecommerce-cart/repositories/mocks"
	"github.com/akkinasrikar/ecommerce-cart/utils"
	"github.com/akkinasrikar/ecommerce-cart/validators/helper"
	redis "github.com/go-redis/redis/v8"
	"github.com/golang/mock/gomock"
)

func Test_loginService_SignUp(t *testing.T) {
	type fields struct {
		repoService repositories.RepositoryInterface
	}
	type args struct {
		req entities.SignUp
	}

	tests := []struct {
		name            string
		fields          fields
		args            args
		mockRepoService func(ctrl *gomock.Controller) *mocks.MockRepositoryInterface
		want            responses.SingUp
		want1           models.EcomError
	}{
		{
			name: "Success",
			mockRepoService: func(ctrl *gomock.Controller) *mocks.MockRepositoryInterface {
				mockRepoService := mocks.NewMockRepositoryInterface(ctrl)
				mockRepoService.EXPECT().SignUp(gomock.Any()).Return(entities.SignUp{}, models.EcomError{})
				return mockRepoService
			},
			args: args{
				req: entities.SignUp{},
			},
			want: responses.SingUp{
				Message: "User created successfully",
			},
			want1: models.EcomError{},
		},
		{
			name: "Failure at service layer",
			mockRepoService: func(ctrl *gomock.Controller) *mocks.MockRepositoryInterface {
				mockRepoService := mocks.NewMockRepositoryInterface(ctrl)
				mockRepoService.EXPECT().SignUp(gomock.Any()).Return(entities.SignUp{}, models.EcomError{
					Message: errors.New("Error while creating user"),
				})
				return mockRepoService
			},
			args: args{
				req: entities.SignUp{},
			},
			want:  responses.SingUp{},
			want1: *helper.ErrorInternalSystemError("Error while signing up : " + "Error while creating user"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			s := &loginService{
				repoService: tt.mockRepoService(ctrl),
			}
			got, got1 := s.SignUp(tt.args.req)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("loginService.SignUp() got = %v, want %v", got, tt.want)
			}
			if got1.Message != nil && (got1.Message.Error() != tt.want1.Message.Error()) {
				t.Errorf("loginService.SignUp() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_loginService_Login(t *testing.T) {
	type fields struct {
		repoService repositories.RepositoryInterface
		redisClient *redis.Client
	}
	type args struct {
		req entities.Login
	}

	tests := []struct {
		name            string
		fields          fields
		args            args
		mockRepoService func(ctrl *gomock.Controller) *mocks.MockRepositoryInterface
		mockRedisClient func(ctrl *gomock.Controller) *utils.MockClient
		want            responses.Login
		want1           models.EcomError
	}{
		{
			name: "Success",
			mockRepoService: func(ctrl *gomock.Controller) *mocks.MockRepositoryInterface {
				mockRepoService := mocks.NewMockRepositoryInterface(ctrl)
				mockRepoService.EXPECT().Login(gomock.Any()).Return(entities.Login{}, models.EcomError{})
				return mockRepoService
			},
			args: args{
				req: entities.Login{},
			},
			want:  responses.Login{
				Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODk1MDQzMjksInN1YiI6IiJ9.wDMhY7K74lEpMDbySXLWKpwtVxn0mgkSpSRdq86enBU",
			},
			want1: models.EcomError{},
		},
		{
			name: "Failure at service layer",
			mockRepoService: func(ctrl *gomock.Controller) *mocks.MockRepositoryInterface {
				mockRepoService := mocks.NewMockRepositoryInterface(ctrl)
				mockRepoService.EXPECT().Login(gomock.Any()).Return(entities.Login{}, models.EcomError{
					Message: errors.New("Error while logging in"),
				})
				return mockRepoService
			},
			args: args{
				req: entities.Login{},
			},
			want:  responses.Login{},
			want1: *helper.ErrorInternalSystemError("Error while logging in : " + "Error while logging in"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			s := &loginService{
				repoService: tt.mockRepoService(ctrl),
				redisClient: utils.InitRedisCacheTest(),
			}
			got, got1 := s.Login(tt.args.req)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("loginService.Login() got = %v, want %v", got, tt.want)
			}
			if got1.Message != nil && (got1.Message.Error() != tt.want1.Message.Error()) {
				t.Errorf("loginService.Login() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
