package auth

import "context"

type ctxKey int

const (
	authToken ctxKey = iota
)

func SetAuthToken(ctx context.Context, token string) context.Context {
	return context.WithValue(ctx, authToken, token)
}

func getAuthToken(ctx context.Context) (string, error) {

}
