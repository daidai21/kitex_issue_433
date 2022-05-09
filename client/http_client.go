package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/genericclient"
	"github.com/cloudwego/kitex/pkg/generic"
	"github.com/cloudwego/kitex/pkg/klog"
	"net/http"
)

func doReq(cli genericclient.Client) error {

	userIds := []interface{}{"123"}
	userIds = append(userIds, 456)

	body := map[string]interface{}{
		"text":     "my test",
		"user_ids": userIds,
	}
	data, err := json.Marshal(body)
	if err != nil {
		klog.Fatalf("body marshal failed: %v", err)
	}
	fmt.Println(data)
	fmt.Println(string(data))

	url := "http://example.com/life/client/11?vint64=1&items=item0,item1,itme2"
	req, err := http.NewRequest(http.MethodGet, url, bytes.NewBuffer(data))
	klog.Infof("req=%v", req)
	klog.Infof("req.Body=%v", req.Body)

	if err != nil {
		klog.Fatalf("new http request failed: %v", err)
	}
	req.Header.Set("token", "1")
	customReq, err := generic.FromHTTPRequest(req)
	if err != nil {
		klog.Fatalf("convert request failed: %v", err)
	}
	resp, err := cli.GenericCall(context.Background(), "", customReq)
	if err != nil {
		klog.Fatalf("generic call failed: %v", err)
	}
	realResp := resp.(*generic.HTTPResponse)
	println(realResp)
	klog.Infof("method1 response, status code: %v, headers: %v, body: %v\n", realResp.StatusCode, realResp.Header, realResp.Body)

	return nil
}

func main() {
	path := "./http.thrift" // depends on current directory
	p, err := generic.NewThriftFileProvider(path)
	if err != nil {
		klog.Fatalf("new thrift file provider failed: %v", err)
	}
	g, err := generic.HTTPThriftGeneric(p)
	if err != nil {
		klog.Fatalf("new http thrift generic failed: %v", err)
	}
	cli, err := genericclient.NewClient("echo", g, client.WithHostPorts("0.0.0.0:8888"))
	if err != nil {
		klog.Fatalf("new http generic client failed: %v", err)
	}

	e := doReq(cli)
	println(e)
}
