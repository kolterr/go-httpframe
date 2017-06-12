package login

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

func makeLoginEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return nil, nil
	}
}
