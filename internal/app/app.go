package app

import (
	"go-httpframe/internal/types"
	"net/http"

	kitlog "github.com/go-kit/kit/log"
)

type Application interface {
	Name() string
	Logger() kitlog.Logger

	InitLogger()
	InitConfig()

	InitDatabase() types.Closer
	MustPrepare() types.Closer

	InitServices() http.Handler
}
