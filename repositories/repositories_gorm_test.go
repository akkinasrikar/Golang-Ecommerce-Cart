package repositories

import (
	"errors"
	"reflect"
	"testing"

	"github.com/akkinasrikar/ecommerce-cart/database"
	dbMocks "github.com/akkinasrikar/ecommerce-cart/database/mock"
	"github.com/akkinasrikar/ecommerce-cart/models"
	"github.com/akkinasrikar/ecommerce-cart/models/entities"
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

