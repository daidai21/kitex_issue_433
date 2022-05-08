package main

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/daidai21/kitex_issue_433/kitex_gen/http"
	"github.com/daidai21/kitex_issue_433/kitex_gen/http/bizservice"
	"log"
)

type BizServiceImpl struct{}

func (s *BizServiceImpl) BizMethod1(ctx context.Context, req *http.BizRequest) (resp *http.BizResponse, err error) {
	klog.Infof("BizMethod1 called, request: %#v", req)
	return &http.BizResponse{HttpCode: 200, Text: "Method1 response"}, nil
}

func main() {
	svr := bizservice.NewServer(new(BizServiceImpl))

	err := svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
