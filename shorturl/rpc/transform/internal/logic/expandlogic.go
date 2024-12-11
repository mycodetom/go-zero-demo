/*
 * @Author: codetom
 * @Date: 2024-12-10 10:55:30
 * @LastEditors: codetom
 * @LastEditTime: 2024-12-10 12:06:55
 * @FilePath: \go-zero\shorturl\rpc\transform\internal\logic\expandlogic.go
 * @Description: 这是描述
 * 
 * Copyright (c) 2024 by codetom, All Rights Reserved. 
 */
package logic

import (
	"context"

	"shorturl/rpc/transform/internal/svc"
	"shorturl/rpc/transform/transform"

	"github.com/zeromicro/go-zero/core/logx"
)

type ExpandLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewExpandLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExpandLogic {
	return &ExpandLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ExpandLogic) Expand(in *transform.ExpandReq) (*transform.ExpandResp, error) {
	// 手动代码开始
	res, err := l.svcCtx.Model.FindOne(l.ctx,in.Shorten)
	if err != nil {
	  return nil, err
	}
  
	return &transform.ExpandResp{
	  Url: res.Url,
	}, nil
	// 手动代码结束
  }
