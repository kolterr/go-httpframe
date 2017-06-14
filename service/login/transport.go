package login

import (
	"go-httpframe/internal/encoding"
	"net/http"

	kitlog "github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"

	"github.com/go-kit/kit/tracing/opentracing"
	"github.com/gorilla/mux"
	stdopentracing "github.com/opentracing/opentracing-go"
	"github.com/prometheus/client_golang/prometheus"
)

func MakeHandler(s Service, tracer stdopentracing.Tracer, logger kitlog.Logger) http.Handler {
	options := []kithttp.ServerOption{
		kithttp.ServerErrorLogger(logger),
		kithttp.ServerErrorEncoder(encoding.EncodeError),
	}
	getOptionsHandler := kithttp.NewServer(
		makeGetOptionsEndpoint(),
		decodeGetOptionsRequest,
		encodeGetOptionsResponse,
		append(options, kithttp.ServerBefore(opentracing.FromHTTPRequest(tracer, "login", logger)))...,
	)

	userLoginHandler := prometheus.InstrumentHandler("user_login", kithttp.NewServer(
		makeLoginEndpoint(),
		decodeLoginRequest,
		encodeLoginResponse,
		options...))

	m := mux.NewRouter()
	m.Handle("/v1/", getOptionsHandler).Methods("OPTIONS")  //获取可用操作
	m.Handle("/v1/login", userLoginHandler).Methods("POST") //用户登录

	return m
}
