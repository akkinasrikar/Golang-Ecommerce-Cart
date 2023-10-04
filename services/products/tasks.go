package services

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"strconv"

	"github.com/akkinasrikar/ecommerce-cart/api"
	"github.com/akkinasrikar/ecommerce-cart/constants"
	"github.com/akkinasrikar/ecommerce-cart/kafka"
	"github.com/akkinasrikar/ecommerce-cart/models"
	"github.com/akkinasrikar/ecommerce-cart/models/entities"
	"github.com/akkinasrikar/ecommerce-cart/repositories"
	"github.com/akkinasrikar/ecommerce-cart/utils"
	"github.com/hibiken/asynq"
)

type AsynqPublisher interface {
	EnqueueContext(ctx context.Context, task *asynq.Task, opts ...asynq.Option) (*asynq.TaskInfo, error)
}

type ProducAsynqService interface {
	ProductImageResize(context.Context, int) error
	ImageResize(context.Context, models.ImageResize) error
	UpdateDeliveryStatus(context.Context) error
	SubscribeDataFromKafka(context.Context) error
	SendAnEmail(context.Context, models.OrderDetailsEmail) error
}

type productAsynqImpl struct {
	Store       repositories.RepositoryInterface
	APIProvider api.Service
	asynqClient AsynqPublisher
	Producer    kafka.Producer
}

func NewAsynqService(store repositories.RepositoryInterface, asynqClient AsynqPublisher, apiProvider api.Service, producer kafka.Producer) ProducAsynqService {
	return &productAsynqImpl{
		Store:       store,
		asynqClient: asynqClient,
		APIProvider: apiProvider,
		Producer:    producer,
	}
}

func (p *productAsynqImpl) ProductImageResize(ctx context.Context, id int) error {
	var item entities.Item
	item, ecomErr := p.Store.GetProductById(id)
	if ecomErr.Message != nil {
		return errors.New("error while fetching product from db")
	}
	imageBytes, err := utils.ReadImageFromUrl(item.ItemImage)
	if err != nil {
		return errors.New("error while reading image from url")
	}
	resizeImageData := models.ImageResize{
		Id:         id,
		ImageBytes: imageBytes,
	}

	data, err := json.Marshal(resizeImageData)
	if err != nil {
		return errors.New("error while marshalling image data")
	}

	task := asynq.NewTask(constants.ProcessTasks.IMAGERESIZE, data, asynq.TaskID(strconv.Itoa(int(id))))
	_, err = p.asynqClient.EnqueueContext(ctx, task)
	if err != nil {
		return errors.New("error while enqueueing task")
	}
	return nil
}

func (p *productAsynqImpl) ImageResize(ctx context.Context, resizeImageData models.ImageResize) error {
	imageBytes, err := utils.ResizeImage(resizeImageData.ImageBytes)
	if err != nil {
		return errors.New("error while resizing image")
	}
	imageBase64 := base64.StdEncoding.EncodeToString(imageBytes)
	var item entities.Item
	item.ItemID = resizeImageData.Id
	item.ImageBase64 = imageBase64
	_, ecomErr := p.Store.UpdateProductByID(item.ItemID, item)
	if ecomErr.Message != nil {
		return errors.New("error while updating product image")
	}
	return nil
}

func (p *productAsynqImpl) UpdateDeliveryStatus(ctx context.Context) error {
	orders, ecomErr := p.Store.GetAllOrders()
	if ecomErr.Message != nil {
		return errors.New("error while fetching orders from db")
	}
	for _, order := range orders {
		if order.DeliveryDate == utils.GenerateCurrentDate() {
			order.DeliveryStatus = "Delivered"
			_, ecomErr := p.Store.UpdateOrderByID(order.OrderID, order)
			if ecomErr.Message != nil {
				return errors.New("error while updating order status")
			}
			userDetails, err := p.Store.GetUserDetailsById(order.UsersID)
			if err.Message != nil {
				return errors.New("error while fetching user details from db")
			}
			req := models.SendEmailRequest{
				Email:   userDetails.EmailID,
				Subject: "Order Delivery Status",
				Message: order.OrderID + " is Successfully Delivered, Thank you for shopping with us",
			}
			errEcom := p.APIProvider.SendMail(req)
			if errEcom != nil {
				return errEcom
			}
		}
	}
	return nil
}

func (p *productAsynqImpl) SubscribeDataFromKafka(ctx context.Context) error {
	p.Producer.Consumer(ctx)
	return nil
}

func (p *productAsynqImpl) SendAnEmail(ctx context.Context, orderDetailsEmail models.OrderDetailsEmail) error {
	var itemDetails entities.Item
	var orderDetails entities.Order
	var err models.EcomError
	orderDetails, err = p.Store.GetOrderByID(orderDetailsEmail.OrderID)
	if err.Message != nil {
		return errors.New("error while fetching order from db")
	}
	itemDetails, err = p.Store.GetProductById(orderDetails.ItemID)
	if err.Message != nil {
		return errors.New("error while fetching product from db")
	}
	userDetails, err := p.Store.GetUserDetailsById(orderDetails.UsersID)
	if err.Message != nil {
		return errors.New("error while fetching user details from db")
	}
	req := models.SendEmailRequest{
		Email:   userDetails.EmailID,
		Subject: "Order Confirmation Details",
		Message: utils.GenerateHtmlResponse2(itemDetails, orderDetails),
	}
	errEcom := p.APIProvider.SendMail(req)
	if errEcom != nil {
		return errEcom
	}
	return nil
}
