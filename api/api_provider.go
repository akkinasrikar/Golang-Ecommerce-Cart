package api

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/akkinasrikar/ecommerce-cart/api/dto"
	"github.com/akkinasrikar/ecommerce-cart/config"
	"github.com/akkinasrikar/ecommerce-cart/models"
	"github.com/akkinasrikar/ecommerce-cart/utils"
	"github.com/akkinasrikar/ecommerce-cart/validators/helper"
	"github.com/pkg/errors"
	"gopkg.in/gomail.v2"
)

func setHttpHeader(ecomCtx context.Context, header http.Header) http.Header {
	header.Set("Content-Type", "application/json")
	return header
}

func setHttpRequest(ecomCtx context.Context, reqD dto.Request) (*http.Request, error) {
	req, err := http.NewRequest(reqD.Method, reqD.Url, bytes.NewReader(reqD.ReqestBody))
	if err != nil {
		return nil, errors.Wrap(err, "[setHttpRequest]")
	}
	req.Header = setHttpHeader(ecomCtx, req.Header)
	return req, nil
}

func responseToByte(resp *http.Response) ([]byte, error) {
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return body, errors.Wrap(err, "[responseToByte, ioReadAll]")
	}
	return body, nil
}

func (t *service) clientHttpCall(zwCtx context.Context, client HttpCall, request *http.Request) ([]byte, int, error) {
	var response []byte
	httpResponse, err := client.Do(request)
	if err != nil {
		return response, 0, errors.Wrap(err, "[clientHttpCall, client Do]")
	}
	statusCode := httpResponse.StatusCode
	response, err = responseToByte(httpResponse)
	if err != nil {
		return response, statusCode, errors.Wrap(err, "[PostApiCall, responseToByte]")
	}
	defer httpResponse.Body.Close()
	return response, statusCode, nil
}

func (s *service) GetItems(ecomCtx context.Context) (dto.ItemsResponse, models.EcomError) {
	url, err := url.JoinPath(s.BaseURL, "products")
	if err != nil {
		return dto.ItemsResponse{}, *helper.ErrorInternalSystemError("Error in creating url")
	}

	req := dto.Request{
		Method:     http.MethodGet,
		Url:        url,
		ReqestBody: nil,
	}

	request, err := setHttpRequest(ecomCtx, req)
	if err != nil {
		return dto.ItemsResponse{}, *helper.ErrorInternalSystemError("Error in creating request")
	}

	response, statusCode, err := s.clientHttpCall(ecomCtx, s.http, request)
	if err != nil {
		return dto.ItemsResponse{}, *helper.ErrorInternalSystemError("Error in client http call")
	}

	if statusCode != http.StatusOK {
		return dto.ItemsResponse{}, *helper.ErrorDownStreamError()
	}

	var itemsResponse dto.ItemsResponse
	err = json.Unmarshal(response, &itemsResponse)
	if err != nil {
		return dto.ItemsResponse{}, *helper.ErrorInternalSystemError("Error in unmarshalling response")
	}

	return itemsResponse, models.EcomError{}
}

func (s *service) SendMail(req models.SendEmailRequest) error {
	m := gomail.NewMessage()
	m.SetHeader("From", config.FakeStore.Gmail)
	m.SetHeader("To", req.Email)
	m.SetHeader("Subject", req.Subject)
	m.SetBody("text/html", req.Message)
	attachment, err := utils.GeneratePdf(req.Message)
	if err != nil {
		return err
	}
	m.Attach("invoice.pdf", gomail.SetCopyFunc(func(w io.Writer) error {
		_, err := io.Copy(w, bytes.NewReader(attachment))
		return err
	}))
	d := gomail.NewDialer("smtp.gmail.com", 587, config.FakeStore.Gmail, config.FakeStore.MailPassword)
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
