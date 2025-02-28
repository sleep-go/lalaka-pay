package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"lalaka-pay/model"
	"lalaka-pay/util"
	"net/http"
)

const (
	specialCreateUrl = "/api/v3/ccss/counter/order/special_create"
	orderQueryUrl    = "/api/v3/ccss/counter/order/query"
	orderCloseUrl    = "/api/v3/ccss/counter/order/close"
	refundUrl        = "/api/v3/labs/relation/refund"
)

func newBuffer[T any](req *T) *bytes.Buffer {
	m := model.BaseReq[T]{
		ReqTime: util.GetReqTime(),
		Version: "3.0",
		ReqData: req,
	}
	data, err := json.Marshal(m)
	if err != nil {
		return nil
	}
	return bytes.NewBuffer(data)
}

// doRequest 统一请求方法
func doRequest[T any, D any](c *Client, url string, req *T) (*D, error) {
	reqStr := newBuffer[T](req)
	request, err := http.NewRequest(http.MethodPost, c.Host+url, reqStr)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	ret := reqStr.String()
	auth, err := c.GetAuthorization(ret)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Authorization", auth)
	fmt.Println(ret)
	fmt.Println(auth)
	resp, err := c.Http.Do(request)
	if err != nil {
		return nil, err
	}
	return util.ParseResp[D](resp)
}

// OrderSpecialCreate 收银台订单创建
func (c *Client) OrderSpecialCreate(req *model.SpecialCreateReq) (*model.SpecialCreateRes, error) {
	return doRequest[model.SpecialCreateReq, model.SpecialCreateRes](c, specialCreateUrl, req)
}

// OrderQuery 收银台订单查询
func (c *Client) OrderQuery(req *model.OrderQueryReq) (*model.OrderQueryRes, error) {
	return doRequest[model.OrderQueryReq, model.OrderQueryRes](c, orderQueryUrl, req)
}

// OrderClose 收银台订单关单
func (c *Client) OrderClose(req *model.OrderCloseReq) (resp *model.OrderCloseRes, err error) {
	return doRequest[model.OrderCloseReq, model.OrderCloseRes](c, orderCloseUrl, req)
}
