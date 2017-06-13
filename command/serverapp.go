package command

import (
	"go-httpframe/internal/app"
	"go-httpframe/internal/util"
	"go-httpframe/service/login"
	"net/http"

	"github.com/go-kit/kit/log"
	stdopentracing "github.com/opentracing/opentracing-go"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/sony/sonyflake"
	"github.com/spf13/viper"
)

type serverApp struct {
	*app.Base
}

func NewServerApp(n, c string) *serverApp {
	var st sonyflake.Settings
	sf := sonyflake.NewSonyflake(st)
	if sf == nil {
		return nil
	}

	return &serverApp{
		&app.Base{
			AppName:    n,
			ConfigPath: c,
			Sonyflake:  sf,
		},
	}
}

func (s *serverApp) buildDSN(db string) string {
	var prefix = "database." + db

	var (
		host     = viper.GetString(prefix + ".host")
		port     = viper.GetInt(prefix + ".port")
		username = viper.GetString(prefix + ".username")
		password = viper.GetString(prefix + ".password")
		dbname   = viper.GetString(prefix + ".dbname")
		args     = viper.GetString(prefix + ".args")
	)

	return util.BuildDSN(username, password, host, port, dbname, args)
}

func (s *serverApp) MustService(srv interface{}, srvName string) {
	if srv == nil {
		panic("no service for " + srvName)
	}
	s.Logger().Log("service", srvName, "msg", "running")
}

func (s *serverApp) InitServices() http.Handler {
	var (
		mux        = http.NewServeMux()
		httpLogger = log.With(s.Logger(), "component", "http")
		webDir     = viper.GetString("webpage_dir")
		tracer     stdopentracing.Tracer
	)

	mux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("success"))
	})

	mux.Handle("/user", login.MakeLoginHandler(httpLogger, tracer))

	//prometheus
	mux.Handle("/metrics", stdprometheus.Handler())
	//static page
	mux.Handle("/web/", http.StripPrefix("/web/", http.FileServer(http.Dir(webDir))))

	return mux
}
