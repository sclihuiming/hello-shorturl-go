package logic

import (
	"context"

	"shorturl/api/internal/svc"
	"shorturl/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type ShortenLogic struct {
	ctx context.Context
	logx.Logger
}

func NewShortenLogic(ctx context.Context, svcCtx *svc.ServiceContext) ShortenLogic {
	return ShortenLogic{
		ctx:    ctx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ShortenLogic) Shorten(req types.ShortenReq) (*types.ShortenResp, error) {
	return &types.ShortenResp{}, nil
}
