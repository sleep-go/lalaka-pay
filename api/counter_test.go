package api

import (
	"fmt"
	"testing"
	"time"
)
import "lalaka-pay/model"

func TestCreate(t *testing.T) {
	orderId := model.CreateOrderStr()
	client := NewClient(model.APPID_TEST, model.SERIAL_NO_TEST, model.KEY_PATH_TEST, model.CERT_PATH_TEST, false)
	expeirTime := time.Now().Add(24 * time.Hour).Format("20060102150405")
	req := model.SpecialCreateReq{
		OutOrderNo:         orderId,
		MerchantNo:         model.MERCHANT_NO_TEST,
		TotalAmount:        100,
		OrderEfficientTime: expeirTime,
		OrderInfo:          "保证金充值",
	}
	ret, err := client.OrderSpecialCreate(&req)

	fmt.Println(ret)
	fmt.Println(err)
}
