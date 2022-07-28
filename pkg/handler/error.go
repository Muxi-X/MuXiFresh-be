package handler

import (
	"context"
	logger "github.com/MuXiFresh-be/log"

	"github.com/micro/go-micro/client"
	errors "github.com/micro/go-micro/errors"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/server"
	"go.uber.org/zap"
)

func ServerErrorHandlerWrapper() server.HandlerWrapper {
	return func(h server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			md, ok := metadata.FromContext(ctx)
			if !ok {
				md = make(map[string]string)
			}

			err := h(ctx, req, rsp)

			if err != nil {
				logger.Error(err.Error(), zap.String("traceId", md["uber-trace-id"]))
			}
			return err
		}
	}
}

func ClientErrorHandlerWrapper() client.CallWrapper {
	return func(cf client.CallFunc) client.CallFunc {
		return func(ctx context.Context, node *registry.Node, req client.Request, rsp interface{}, opts client.CallOptions) error {
			err := cf(ctx, node, req, rsp, opts)
			if err != nil {
				e := errors.Parse(err.Error())
				return e
			}
			return err
		}
	}
}
