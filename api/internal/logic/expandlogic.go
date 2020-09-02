package logic

import (
	"context"

	"shorturl/api/internal/svc"
	"shorturl/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type ExpandLogic struct {
	ctx context.Context
	logx.Logger
}

func NewExpandLogic(ctx context.Context, svcCtx *svc.ServiceContext) ExpandLogic {
	return ExpandLogic{
		ctx:    ctx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ExpandLogic) Expand(req types.ExpandReq) (*types.ExpandResp, error) {
	return &types.ExpandResp{}, nil
}
