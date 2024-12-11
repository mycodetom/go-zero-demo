/*
 * @Author: codetom
 * @Date: 2024-12-10 10:55:30
 * @LastEditors: codetom
 * @LastEditTime: 2024-12-10 12:07:38
 * @FilePath: \go-zero\shorturl\rpc\transform\internal\logic\shortenlogic.go
 * @Description: 这是描述
 *
 * Copyright (c) 2024 by codetom, All Rights Reserved.
 */
package logic

import (
	"context"

	"shorturl/rpc/transform/internal/svc"
	"shorturl/rpc/transform/model"
	"shorturl/rpc/transform/transform"

	"github.com/zeromicro/go-zero/core/hash"
	"github.com/zeromicro/go-zero/core/logx"
)

type ShortenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewShortenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShortenLogic {
	return &ShortenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ShortenLogic) Shorten(in *transform.ShortenReq) (*transform.ShortenResp, error) {
	// 手动代码开始，生成短链接
	key := hash.Md5Hex([]byte(in.Url))[:6]
	_, err := l.svcCtx.Model.Insert(l.ctx, &model.Shorturl{
		Shorten: key,
		Url:     in.Url,
	})
	if err != nil {
		return nil, err
	}

	return &transform.ShortenResp{
		Shorten: key,
	}, nil
	// 手动代码结束
}
