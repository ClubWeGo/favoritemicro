// Code generated by Kitex v0.4.4. DO NOT EDIT.

package favoriteservice

import (
	"context"
	favorite "github.com/ClubWeGo/favoritemicro/kitex_gen/favorite"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return favoriteServiceServiceInfo
}

var favoriteServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "FavoriteService"
	handlerType := (*favorite.FavoriteService)(nil)
	methods := map[string]kitex.MethodInfo{
		"FavoriteMethod":     kitex.NewMethodInfo(favoriteMethodHandler, newFavoriteServiceFavoriteMethodArgs, newFavoriteServiceFavoriteMethodResult, false),
		"FavoriteListMethod": kitex.NewMethodInfo(favoriteListMethodHandler, newFavoriteServiceFavoriteListMethodArgs, newFavoriteServiceFavoriteListMethodResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "favorite",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.4.4",
		Extra:           extra,
	}
	return svcInfo
}

func favoriteMethodHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*favorite.FavoriteServiceFavoriteMethodArgs)
	realResult := result.(*favorite.FavoriteServiceFavoriteMethodResult)
	success, err := handler.(favorite.FavoriteService).FavoriteMethod(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newFavoriteServiceFavoriteMethodArgs() interface{} {
	return favorite.NewFavoriteServiceFavoriteMethodArgs()
}

func newFavoriteServiceFavoriteMethodResult() interface{} {
	return favorite.NewFavoriteServiceFavoriteMethodResult()
}

func favoriteListMethodHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*favorite.FavoriteServiceFavoriteListMethodArgs)
	realResult := result.(*favorite.FavoriteServiceFavoriteListMethodResult)
	success, err := handler.(favorite.FavoriteService).FavoriteListMethod(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newFavoriteServiceFavoriteListMethodArgs() interface{} {
	return favorite.NewFavoriteServiceFavoriteListMethodArgs()
}

func newFavoriteServiceFavoriteListMethodResult() interface{} {
	return favorite.NewFavoriteServiceFavoriteListMethodResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) FavoriteMethod(ctx context.Context, request *favorite.FavoriteReq) (r *favorite.FavoriteResp, err error) {
	var _args favorite.FavoriteServiceFavoriteMethodArgs
	_args.Request = request
	var _result favorite.FavoriteServiceFavoriteMethodResult
	if err = p.c.Call(ctx, "FavoriteMethod", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) FavoriteListMethod(ctx context.Context, request *favorite.FavoriteListReq) (r *favorite.FavoriteListResp, err error) {
	var _args favorite.FavoriteServiceFavoriteListMethodArgs
	_args.Request = request
	var _result favorite.FavoriteServiceFavoriteListMethodResult
	if err = p.c.Call(ctx, "FavoriteListMethod", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}