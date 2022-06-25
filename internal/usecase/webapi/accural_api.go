package webapi

import (
	"context"
	"log"

	"github.com/dimk00z/go-musthave-diploma/internal/entity"
	"github.com/go-resty/resty/v2"
)

func (g *GopherMartWebAPI) PostOrderInAccuralService(ctx context.Context, orderNumber string) (err error) {

	client := resty.New()
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{"order": orderNumber}).
		Post(g.cfg.API.AccrualSystemAddress + "/api/orders")
	log.Println("accural service response:", resp)
	log.Println("response status:", resp.StatusCode())

	return err
}

func (g *GopherMartWebAPI) CheckOrder(ctx context.Context, orderNumber string) (response entity.AccrualSystemResponse, err error) {
	client := resty.New()
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetResult(&response).
		Get(g.cfg.API.AccrualSystemAddress + "/api/orders/" + orderNumber)
	log.Println("accural service response:", resp)
	log.Println("response status:", resp.StatusCode())
	return
}
