package mtx

import "context"

type contextKey string

var authenticatedContextKey contextKey = "isAuthenticated"

func GetIsAuthenticated(ctx context.Context) bool {
	if isAuthenticated, ok := ctx.Value(authenticatedContextKey).(bool); ok {
		return isAuthenticated
	}
	return false
}

func SetIsAuthenticated(c context.Context, value bool) context.Context {
	ctx := context.WithValue(c, authenticatedContextKey, value)
	return ctx
}
