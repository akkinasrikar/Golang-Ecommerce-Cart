package services

import (
	"context"
	"reflect"
	"testing"

	"github.com/akkinasrikar/ecommerce-cart/api"
	"github.com/akkinasrikar/ecommerce-cart/config"
	"github.com/akkinasrikar/ecommerce-cart/models"
	"github.com/akkinasrikar/ecommerce-cart/models/entities"
	"github.com/akkinasrikar/ecommerce-cart/repositories"
	"github.com/akkinasrikar/ecommerce-cart/repositories/mocks"
	"github.com/golang/mock/gomock"
)

func Test_products_GetProducts(t *testing.T) {
	type fields struct {
		APIProvider api.Service
		Store       repositories.RepositoryInterface
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name            string
		fields          fields
		args            args
		mockRepoService func(ctrl *gomock.Controller) *mocks.MockRepositoryInterface
		want            []entities.Item
		want1           models.EcomError
	}{
		{
			name: "Success",
			mockRepoService: func(ctrl *gomock.Controller) *mocks.MockRepositoryInterface {
				mockRepoService := mocks.NewMockRepositoryInterface(ctrl)
				mockRepoService.EXPECT().GetAllProducts().Return([]entities.Item{}, models.EcomError{})
				return mockRepoService
			},
			args: args{
				ctx: context.Background(),
			},
			want:  []entities.Item{},
			want1: models.EcomError{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			tt.fields.Store = tt.mockRepoService(ctrl)
			p := &products{
				APIProvider: tt.fields.APIProvider,
				Store:       tt.fields.Store,
			}
			got, got1 := p.GetProducts(tt.args.ctx)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("products.GetProducts() got = %v, want %v", got, tt.want)
			}
			if got1.Message != nil && (got1.Message.Error() != tt.want1.Message.Error()) {
				t.Errorf("products.GetProducts() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_products_GetUserDetails(t *testing.T) {
	type fields struct {
		APIProvider api.Service
		Store       repositories.RepositoryInterface
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name            string
		fields          fields
		args            args
		mockRepoService func(ctrl *gomock.Controller) *mocks.MockRepositoryInterface
		want            models.EcomUsers
		want1           models.EcomError
	}{
		{
			name: "Success",
			mockRepoService: func(ctrl *gomock.Controller) *mocks.MockRepositoryInterface {
				mockRepoService := mocks.NewMockRepositoryInterface(ctrl)
				mockRepoService.EXPECT().GetUserDetails(gomock.Any()).Return(entities.EcomUsers{
					CartItems: `{"items_id":[1,20]}`,
				}, models.EcomError{})
				return mockRepoService
			},
			args: args{
				ctx: context.Background(),
			},
			want:  models.EcomUsers{
				CartItems: []int{1, 20},
			},
			want1: models.EcomError{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			tt.fields.Store = tt.mockRepoService(ctrl)
			p := &products{
				APIProvider: tt.fields.APIProvider,
				Store:       tt.fields.Store,
			}
			got, got1 := p.GetUserDetails(tt.args.ctx)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("products.GetUserDetails() got = %v, want %v", got, tt.want)
			}
			if got1.Message != nil && (got1.Message.Error() != tt.want1.Message.Error()) {
				t.Errorf("products.GetUserDetails() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_products_CardDetails(t *testing.T) {
	config.Init()
	type fields struct {
		APIProvider api.Service
		Store       repositories.RepositoryInterface
	}
	type args struct {
		ctx context.Context
		req models.CardDetails
	}
	tests := []struct {
		name            string
		fields          fields
		args            args
		mockRepoService func(ctrl *gomock.Controller) *mocks.MockRepositoryInterface
		want            models.CardDetails
		want1           models.EcomError
	}{
		{
			name: "Success",
			mockRepoService: func(ctrl *gomock.Controller) *mocks.MockRepositoryInterface {
				mockRepoService := mocks.NewMockRepositoryInterface(ctrl)
				mockRepoService.EXPECT().GetUserDetails(gomock.Any()).Return(entities.EcomUsers{}, models.EcomError{})
				mockRepoService.EXPECT().CreateCardDetails(gomock.Any()).Return(entities.CardDetails{}, models.EcomError{})
				return mockRepoService
			},
			args: args{
				ctx: context.Background(),
				req: models.CardDetails{
					CardNumber: 1234567890123456,
					ExpiryDate: "12/22",
					CVV:        123,
					Name:       "Akkina Srikar",
					CardType:   "VISA",
				},
			},
			want:  models.CardDetails{},
			want1: models.EcomError{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			tt.fields.Store = tt.mockRepoService(ctrl)
			p := &products{
				APIProvider: tt.fields.APIProvider,
				Store:       tt.fields.Store,
			}
			_, got1 := p.CardDetails(tt.args.ctx, tt.args.req)
			if got1.Message != nil && (got1.Message.Error() != tt.want1.Message.Error()) {
				t.Errorf("products.CardDetails() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_products_GetCardDetails(t *testing.T) {
	config.Init()
	type fields struct {
		APIProvider api.Service
		Store       repositories.RepositoryInterface
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name            string
		fields          fields
		args            args
		mockRepoService func(ctrl *gomock.Controller) *mocks.MockRepositoryInterface
		want            []models.CardDetails
		want1           models.EcomError
	}{
		{
			name: "Success",
			mockRepoService: func(ctrl *gomock.Controller) *mocks.MockRepositoryInterface {
				mockRepoService := mocks.NewMockRepositoryInterface(ctrl)
				mockRepoService.EXPECT().GetUserDetails(gomock.Any()).Return(entities.EcomUsers{}, models.EcomError{})
				mockRepoService.EXPECT().GetCardDetails(gomock.Any()).Return([]entities.CardDetails{
					{
						EncryptedData: "oMFFM7166y/pfxP/Pjr0ZuMXGbevVdONJJ85/Pu97LNfnCgLJl8dRuJc15IN6QvRwQo5X5ELuXK+Hdu6ryTlQ4wA2YuI/CL/e9oE7T8gurt6Yq01/NmJgh0iZiq5gGtxY5mpykyy14qrzbtP+65ewBPiH3/BPI/RutM3AYDTf5RIVHz4k06Jv8OPYIw0BO3w2/NuecBudElyTiJf3PeyYOMqLtXjRuMecvl0/NuBtWi79ABMKfdzICsdJjyoVqQyBaRE2bOkZhyWXzSIlm7rl1Dl6TPGE4KHyskYRSxLu7v8uVma5Agah+dStV2iyATFBdRQIUr/1Di9MZC+TL6Oga0/ivavYCR7BSr2Kj9hqHD3qI4cV3ZJHAdOuBbPADlFjV2nbmny7F8CZr0gGBvTTtsCVvPeUxn1HZ3z9pUOLYmZPQNnLUZ5+PUYD6j3ocItlswIFxomvowcGmfrVc+ym2zSQ4UXF9m+yFB8+kXIg2gW1T0tR7C3D4Gn80//7AxMaYAFlE+H04orBreSklNySL+zwI19uGhUYdF3f7b298s4ssqgAk9wgePtS7WTqmiwxjlsSXTk86EtghY9O5oGKVB6p+uobbMBLog+e8+83f12DifuhMQ4EkIQ+WyooB9QA1VZCjbsXJpHgTJoc4+vobvOw3zmXbQ/pLvRimZoBCY=",
						EcomId:        "1234567890123456",
					},
				}, models.EcomError{})
				return mockRepoService
			},
			args: args{
				ctx: context.Background(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			tt.fields.Store = tt.mockRepoService(ctrl)
			p := &products{
				APIProvider: tt.fields.APIProvider,
				Store:       tt.fields.Store,
			}
			_, got1 := p.GetCardDetails(tt.args.ctx)
			if got1.Message != nil && (got1.Message.Error() != tt.want1.Message.Error()) {
				t.Errorf("products.GetCardDetails() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
