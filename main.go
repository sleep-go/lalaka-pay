package main

import (
	"fmt"
	"lalaka-pay/api"
	"lalaka-pay/model"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	client := api.NewClient("", "", "", "", true)
	fmt.Println(client)
	api.Buffer[model.SpecialCreateReq](&model.SpecialCreateReq{
		OutOrderNo:           "",
		MerchantNo:           "",
		VposID:               "",
		ChannelID:            "",
		TotalAmount:          0,
		OrderEfficientTime:   "",
		NotifyURL:            "",
		SupportCancel:        0,
		SupportRefund:        0,
		SupportRepeatPay:     0,
		OutUserID:            "",
		CallbackURL:          "",
		OrderInfo:            "",
		TermNo:               "",
		SplitMark:            "",
		SettleType:           "",
		OutSplitInfo:         nil,
		CounterParam:         "",
		CounterRemark:        "",
		BusiTypeParam:        "",
		SgnInfo:              nil,
		ProductID:            "",
		GoodsMark:            "",
		GoodsField:           "",
		OrderSceneField:      "",
		AgeLimit:             "",
		RepeatPayAutoRefund:  "",
		RepeatPayNotify:      "",
		CloseOrderAutoRefund: "",
		ShopName:             "",
		InteRouting:          "",
	})
}
