// Code generated by Kitex v0.4.4. DO NOT EDIT.

package favoriteservice

import (
	"context"
	favorite "github.com/ClubWeGo/favoritemicro/kitex_gen/favorite"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	FavoriteMethod(ctx context.Context, request *favorite.FavoriteReq, callOptions ...callopt.Option) (r *favorite.FavoriteResp, err error)
	FavoriteListMethod(ctx context.Context, request *favorite.FavoriteListReq, callOptions ...callopt.Option) (r *favorite.FavoriteListResp, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kFavoriteServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kFavoriteServiceClient struct {
	*kClient
}

func (p *kFavoriteServiceClient) FavoriteMethod(ctx context.Context, request *favorite.FavoriteReq, callOptions ...callopt.Option) (r *favorite.FavoriteResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.FavoriteMethod(ctx, request)
}

func (p *kFavoriteServiceClient) FavoriteListMethod(ctx context.Context, request *favorite.FavoriteListReq, callOptions ...callopt.Option) (r *favorite.FavoriteListResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.FavoriteListMethod(ctx, request)
}