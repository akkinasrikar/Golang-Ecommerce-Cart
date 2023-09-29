package services

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"strconv"

	"github.com/akkinasrikar/ecommerce-cart/constants"
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
	SubscribeDataFromKafka(context.Context, entities.Consume) error
}

type productAsynqImpl struct {
	Store       repositories.RepositoryInterface
	asynqClient AsynqPublisher
}

func NewAsynqService(store repositories.RepositoryInterface, asynqClient AsynqPublisher) ProducAsynqService {
	return &productAsynqImpl{
		Store:       store,
		asynqClient: asynqClient,
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
		}
	}
	return nil
}

func (p *productAsynqImpl) SubscribeDataFromKafka(ctx context.Context, data entities.Consume) error {
	_, ecomErr := p.Store.ConsumeKafkaData(ctx, data)
	if ecomErr.Message != nil {
		return errors.New("error while consuming data from kafka")
	}
	return nil
}
