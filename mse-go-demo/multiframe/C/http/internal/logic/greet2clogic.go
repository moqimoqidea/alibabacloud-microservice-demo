package logic

import (
	"context"
	"fmt"
	"os"

	"github.com/aliyun/alibabacloud-microservice-demo/mse-go-demo/multiframe/C/http/internal/svc"
	"github.com/aliyun/alibabacloud-microservice-demo/mse-go-demo/multiframe/C/http/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
)

type Greet2cLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGreet2cLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Greet2cLogic {
	return &Greet2cLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Greet2cLogic) Greet2c(req *types.GreetReq) (resp *types.GreetResp, err error) {
	resp = &types.GreetResp{
		Message: fmt.Sprintf("hello %s", req.Name),
	}

	content, err := GreetD(context.Background(), req.Name)
	if err != nil {
		logx.Errorf("[Greet1cLogic.Greet1c] call d error: %v", err)
		return nil, err
	}

	tag := os.Getenv("MSE_ALICLOUD_SERVICE_TAG")
	if tag == "" {
		tag = "base"
	}
	ip := os.Getenv("KUBERNETES_POD_IP")
	resp.CallChain = fmt.Sprintf("C:%s:%s -(grpc)- %s", tag, ip, content)
	return
}