// Code generated by goctl. DO NOT EDIT!
// Source: transform.proto

//go:generate mockgen -destination ./transformer_mock.go -package transformer -source $GOFILE

package transformer

import (
	"context"

	transform "hello-shorturl-go/rpc/transform/pb"

	"github.com/tal-tech/go-zero/core/jsonx"
	"github.com/tal-tech/go-zero/rpcx"
)

type (
	Transformer interface {
		Expand(ctx context.Context, in *ExpandReq) (*ExpandResp, error)
		Shorten(ctx context.Context, in *ShortenReq) (*ShortenResp, error)
	}

	defaultTransformer struct {
		cli rpcx.Client
	}
)

func NewTransformer(cli rpcx.Client) Transformer {
	return &defaultTransformer{
		cli: cli,
	}
}

func (m *defaultTransformer) Expand(ctx context.Context, in *ExpandReq) (*ExpandResp, error) {
	var request transform.ExpandReq
	bts, err := jsonx.Marshal(in)
	if err != nil {
		return nil, errJsonConvert
	}

	err = jsonx.Unmarshal(bts, &request)
	if err != nil {
		return nil, errJsonConvert
	}

	client := transform.NewTransformerClient(m.cli.Conn())
	resp, err := client.Expand(ctx, &request)
	if err != nil {
		return nil, err
	}

	var ret ExpandResp
	bts, err = jsonx.Marshal(resp)
	if err != nil {
		return nil, errJsonConvert
	}

	err = jsonx.Unmarshal(bts, &ret)
	if err != nil {
		return nil, errJsonConvert
	}

	return &ret, nil
}

func (m *defaultTransformer) Shorten(ctx context.Context, in *ShortenReq) (*ShortenResp, error) {
	var request transform.ShortenReq
	bts, err := jsonx.Marshal(in)
	if err != nil {
		return nil, errJsonConvert
	}

	err = jsonx.Unmarshal(bts, &request)
	if err != nil {
		return nil, errJsonConvert
	}

	client := transform.NewTransformerClient(m.cli.Conn())
	resp, err := client.Shorten(ctx, &request)
	if err != nil {
		return nil, err
	}

	var ret ShortenResp
	bts, err = jsonx.Marshal(resp)
	if err != nil {
		return nil, errJsonConvert
	}

	err = jsonx.Unmarshal(bts, &ret)
	if err != nil {
		return nil, errJsonConvert
	}

	return &ret, nil
}
