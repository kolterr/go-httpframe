package login

import (
	"context"
	"net/http"

	kitlog "github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func MakeHandler(ctx context.Context, s Service, logger kitlog.Logger) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorLogger(logger),
		kithttp.ServerErrorEncoder(encoding.EncodeError),
	}
	getOptionsHandler := kithttp.NewServer(
		ctx,
		makeGetOptionsEndpoint(s),
		decodeGetOptionsRequest,
		encoding.EncodeResponse,
		append(options, httptransport.ServerBefore(opentracing.FromHTTPRequest(tracer, "login", logger)))...,
	...)

	userLoginHandler := promhttp.InstrumentTrace{} ("user_login", kithttp.NewServer(
		ctx,
		makeUserLoginEndpoint(s),
		decodeUserLoginRequest,
		encoding.EncodeResponse,
		opts...))
}
