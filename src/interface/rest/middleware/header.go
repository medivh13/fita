package middleware

import (
	"context"
	"errors"
	"net/http"

	commonError "fita/src/infra/errors"
	"fita/src/interface/rest/response"
)

const (
	XApiKey string = "x-api-key"
)

type fitaContextKey int

const (
	CtxfitaHeader fitaContextKey = iota + 1
)

type ContexfitaHeader struct {
	ApiKey *string
}

func CheckAPWebHeader(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		apiKey := r.Header.Get(XApiKey)

		if apiKey == "" {
			err := errors.New("ApiKey should exist in header")
			cerr := commonError.NewError(commonError.INVALID_HEADER_X_API_KEY, err)
			cerr.SetSystemMessage(err.Error())

			response.NewResponseClient().HttpError(w, cerr)
			return
		}

		val := ContexfitaHeader{
			ApiKey: &apiKey,
		}

		ctx = context.WithValue(ctx, CtxfitaHeader, val)

		next.ServeHTTP(w, r.WithContext(ctx))
	}

	return http.HandlerFunc(fn)
}
