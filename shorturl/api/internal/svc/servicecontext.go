/*
 * @Author: codetom
 * @Date: 2024-12-10 10:38:41
 * @LastEditors: codetom
 * @LastEditTime: 2024-12-10 12:22:03
 * @FilePath: \go-zero\shorturl\api\internal\svc\servicecontext.go
 * @Description: 这是描述
 *
 * Copyright (c) 2024 by codetom, All Rights Reserved.
 */
package svc

import (
	"shorturl/api/internal/config"
	"shorturl/rpc/transform/transformer"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	Transformer transformer.Transformer // 手动代码
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		Transformer: transformer.NewTransformer(zrpc.MustNewClient(c.Transform)), // 手动代码
	}
}
