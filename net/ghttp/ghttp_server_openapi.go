// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package ghttp

import (
	"github.com/gogf/gf/internal/intlog"
	"github.com/gogf/gf/protocol/goai"
	"github.com/gogf/gf/text/gstr"
)

// initOpenApi generates api specification using OpenApiV3 protocol.
func (s *Server) initOpenApi() {
	if s.config.OpenApiPath == "" {
		return
	}
	var (
		err    error
		method string
	)
	for _, item := range s.GetRoutes() {
		switch item.Type {
		case HandlerTypeMiddleware, HandlerTypeHook:
			continue
		}
		method = item.Method
		if gstr.Equal(method, defaultMethod) {
			method = "POST"
		}
		if item.Handler.Info.Func == nil {
			err = s.openapi.Add(goai.AddInput{
				Path:   item.Route,
				Method: method,
				Object: item.Handler.Info.Value.Interface(),
			})
			if err != nil {
				panic(err)
			}
		}
	}
}

// openapiSpec is a build-in handler automatic producing for openapi specification json file.
func (s *Server) openapiSpec(r *Request) {
	var (
		err error
	)
	if s.config.OpenApiPath == "" {
		r.Response.Write(`OpenApi specification file producing is disabled`)
	} else {
		err = r.Response.WriteJson(s.openapi)
	}

	if err != nil {
		intlog.Error(r.Context(), err)
	}
}
