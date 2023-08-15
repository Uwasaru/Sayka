package context

import (
	"context"
	"fmt"

	"golang.org/x/oauth2"
)

type ContextKey string

var (
	userIdKey ContextKey = "userIdKey"
	tokenKey  ContextKey = "tokenKey"
)

func SetToken(ctx context.Context, token *oauth2.Token) context.Context {
	if token != nil {
		return context.WithValue(ctx, tokenKey, token)
	}
	return ctx
}

func GetToken(ctx context.Context) (*oauth2.Token, bool) {
	v := ctx.Value(tokenKey)
	token, ok := v.(*oauth2.Token)
	return token, ok
}

func SetUserId(ctx context.Context, userId int) context.Context {
	if userId != 0 {
		return context.WithValue(ctx, userIdKey, userId)
	}
	return ctx
}

func GetUser(ctx context.Context) (int, bool) {
	v := ctx.Value(userIdKey)
	userId, ok := v.(int)
	if !ok {
		return 0, ok
	}
	return userId, ok
}
