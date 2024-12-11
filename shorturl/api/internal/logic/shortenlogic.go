/*
 * @Author: codetom
 * @Date: 2024-12-10 10:38:41
 * @LastEditors: codetom
 * @LastEditTime: 2024-12-10 11:54:48
 * @FilePath: \go-zero\shorturl\api\internal\logic\shortenlogic.go
 * @Description: 这是描述
 *
 * Copyright (c) 2024 by codetom, All Rights Reserved.
 */
package logic

import (
	"context"

	"shorturl/api/internal/svc"
	"shorturl/api/internal/types"
	"shorturl/rpc/transform/transformer"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShortenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShortenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShortenLogic {
	return &ShortenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShortenLogic) Shorten(req types.ShortenReq) (types.ShortenResp, error) {
	// 手动代码开始
	rpcResp, err := l.svcCtx.Transformer.Shorten(l.ctx, &transformer.ShortenReq{
		Url: req.Url,
	})
	if err != nil {
		return types.ShortenResp{}, err
	}

	return types.ShortenResp{
		Shorten: rpcResp.Shorten,
	}, nil
	// 手动代码结束
}
