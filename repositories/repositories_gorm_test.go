package repositories

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/akkinasrikar/ecommerce-cart/database"
	dbMocks "github.com/akkinasrikar/ecommerce-cart/database/mock"
	"github.com/akkinasrikar/ecommerce-cart/models"
	"github.com/akkinasrikar/ecommerce-cart/models/entities"
	"github.com/akkinasrikar/ecommerce-cart/utils"
	"github.com/akkinasrikar/ecommerce-cart/validators/helper"
	"github.com/golang/mock/gomock"
)

func TestRepository_SignUp(t *testing.T) {
	type fields struct {
		dbStore database.DB
	}
	type args struct {
		userDetails entities.SignUp
	}
	err := errors.New("testError")
	tests := []struct {
		name              string
		fields            fields
		mockDatabaseStore func(ctrl *gomock.Controller) *dbMocks.MockDB
		args              args
		want              entities.SignUp
		want1             models.EcomError
	}{
		{
			name: "Happy Case",
			mockDatabaseStore: func(ctrl *gomock.Controller) *dbMocks.MockDB {
				mockDB := dbMocks.NewMockDB(ctrl)
				mockDB.EXPECT().Create(gomock.Any()).Return(nil)
				return mockDB
			},
			args: args{
				userDetails: entities.SignUp{},
			},
			want: entities.SignUp{},
		},
		{
			name: "Sad Case",
			mockDatabaseStore: func(ctrl *gomock.Controller) *dbMocks.MockDB {
				mockDB := dbMocks.NewMockDB(ctrl)
				mockDB.EXPECT().Create(gomock.Any()).Return(err)
				return mockDB
			},
			args: args{
				userDetails: entities.SignUp{},
			},
			want1: *helper.ErrorInternalSystemError("Error while signing up : " + err.Error()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			tt.fields.dbStore = tt.mockDatabaseStore(ctrl)
			r := &Repository{
				dbStore: tt.fields.dbStore,
			}
			got, got1 := r.SignUp(tt.args.userDetails)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Repository.SignUp() got = %v, want %v", got, tt.want)
			}
			if got1.Message != nil && (got1.Message.Error() != tt.want1.Message.Error()) {
				t.Errorf("Repository.SignUp() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestRepository_GetAllProducts(t *testing.T) {
	type fields struct {
		dbStore database.DB
	}
	tests := []struct {
		name              string
		fields            fields
		mockDatabaseStore func(ctrl *gomock.Controller) *dbMocks.MockDB
		want              []entities.Item
		want1             models.EcomError
	}{
		{
			name: "Happy Case",
			mockDatabaseStore: func(ctrl *gomock.Controller) *dbMocks.MockDB {
				mockDB := dbMocks.NewMockDB(ctrl)
				mockDB.EXPECT().Find(gomock.Any()).Return(int64(1), nil)
				return mockDB
			},
			want: []entities.Item{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			tt.fields.dbStore = tt.mockDatabaseStore(ctrl)
			r := &Repository{
				dbStore: tt.fields.dbStore,
			}
			_, got1 := r.GetAllProducts()
			if got1.Message != nil && (got1.Message.Error() != tt.want1.Message.Error()) {
				t.Errorf("Repository.GetAllProducts() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestRepository_CreateEcomAccount(t *testing.T) {
	type fields struct {
		dbStore database.DB
	}
	type args struct {
		ecomAccountDetails entities.EcomUsers
	}
	tests := []struct {
		name              string
		fields            fields
		args              args
		mockDatabaseStore func(ctrl *gomock.Controller) *dbMocks.MockDB
		want              entities.EcomUsers
		want1             models.EcomError
	}{
		{
			name: "Happy Case",
			mockDatabaseStore: func(ctrl *gomock.Controller) *dbMocks.MockDB {
				mockDB := dbMocks.NewMockDB(ctrl)
				mockDB.EXPECT().Create(gomock.Any()).Return(nil)
				return mockDB
			},
			args: args{
				ecomAccountDetails: entities.EcomUsers{},
			},
			want: entities.EcomUsers{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			tt.fields.dbStore = tt.mockDatabaseStore(ctrl)
			r := &Repository{
				dbStore: tt.fields.dbStore,
			}
			got, got1 := r.CreateEcomAccount(tt.args.ecomAccountDetails)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Repository.CreateEcomAccount() got = %v, want %v", got, tt.want)
			}
			if got1.Message != nil && (got1.Message.Error() != tt.want1.Message.Error()) {
				t.Errorf("Repository.CreateEcomAccount() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestRepository_GetUserDetails(t *testing.T) {
	type fields struct {
		dbStore database.DB
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name              string
		fields            fields
		args              args
		mockDatabaseStore func(ctrl *gomock.Controller) *dbMocks.MockDB
		want              entities.EcomUsers
		want1             models.EcomError
	}{
		{
			name: "Happy Case",
			mockDatabaseStore: func(ctrl *gomock.Controller) *dbMocks.MockDB {
				mockDB := dbMocks.NewMockDB(ctrl)
				mockDB.EXPECT().Where(gomock.Any(), gomock.Any()).Return(mockDB)
				mockDB.EXPECT().Find(gomock.Any()).Return(int64(1), nil)
				return mockDB
			},
			args: args{
				ctx: utils.SetContextWithAuthData(),
			},
			want: entities.EcomUsers{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			tt.fields.dbStore = tt.mockDatabaseStore(ctrl)
			r := &Repository{
				dbStore: tt.fields.dbStore,
			}
			got, got1 := r.GetUserDetails(tt.args.ctx)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Repository.GetUserDetails() got = %v, want %v", got, tt.want)
			}
			if got1.Message != nil && (got1.Message.Error() != tt.want1.Message.Error()) {
				t.Errorf("Repository.GetUserDetails() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestRepository_CreateCardDetails(t *testing.T) {
	type fields struct {
		dbStore database.DB
	}
	type args struct {
		cardDetails entities.CardDetails
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		mockDatabaseStore func(ctrl *gomock.Controller) *dbMocks.MockDB
		want   entities.CardDetails
		want1  models.EcomError
	}{
		{
			name: "Happy Case",
			mockDatabaseStore: func(ctrl *gomock.Controller) *dbMocks.MockDB {
				mockDB := dbMocks.NewMockDB(ctrl)
				mockDB.EXPECT().Create(gomock.Any()).Return(nil)
				return mockDB
			},
			args: args{
				cardDetails: entities.CardDetails{},
			},
			want: entities.CardDetails{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			tt.fields.dbStore = tt.mockDatabaseStore(ctrl)
			r := &Repository{
				dbStore: tt.fields.dbStore,
			}
			got, got1 := r.CreateCardDetails(tt.args.cardDetails)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Repository.CreateCardDetails() got = %v, want %v", got, tt.want)
			}
			if got1.Message != nil && (got1.Message.Error() != tt.want1.Message.Error()) {
				t.Errorf("Repository.CreateCardDetails() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestRepository_GetCardDetails(t *testing.T) {
	type fields struct {
		dbStore database.DB
	}
	type args struct {
		userDetails entities.EcomUsers
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		mockDatabaseStore func(ctrl *gomock.Controller) *dbMocks.MockDB
		want   []entities.CardDetails
		want1  models.EcomError
	}{
		{
			name: "Happy Case",
			mockDatabaseStore: func(ctrl *gomock.Controller) *dbMocks.MockDB {
				mockDB := dbMocks.NewMockDB(ctrl)
				mockDB.EXPECT().Where(gomock.Any(), gomock.Any()).Return(mockDB)
				mockDB.EXPECT().Find(gomock.Any()).Return(int64(1), nil)
				return mockDB
			},
			args: args{
				userDetails: entities.EcomUsers{},
			},
			want: []entities.CardDetails{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			tt.fields.dbStore = tt.mockDatabaseStore(ctrl)
			r := &Repository{
				dbStore: tt.fields.dbStore,
			}
			_, got1 := r.GetCardDetails(tt.args.userDetails)
			if got1.Message != nil && (got1.Message.Error() != tt.want1.Message.Error()) {
				t.Errorf("Repository.GetCardDetails() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
