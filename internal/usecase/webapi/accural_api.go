package webapi

import (
	"context"
	"log"

	"github.com/dimk00z/go-musthave-diploma/internal/entity"
	"github.com/go-resty/resty/v2"
)

func (g *GopherMartWebAPI) PostOrderInAccuralService(ctx context.Context, orderNumber string) (err error) {

	client := resty.New()
	log.Println("orderNumber:", orderNumber)
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{"order": orderNumber}).
		Post("http://" + g.cfg.API.AccrualSystemAddress + "/api/orders")
	log.Println("accural service response:", resp)
	log.Println("response status:", resp.StatusCode())

	return err
}

func (g *GopherMartWebAPI) CheckOrder(ctx context.Context, orderNumber string) (response entity.AccrualSystemResponse, err error) {
	// TODO add logic
	return
}
