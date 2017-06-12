package app

import (
	"fmt"
	"go-httpframe/internal/types"
	"go-httpframe/internal/util"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	kitlog "github.com/go-kit/kit/log"
	"github.com/sony/sonyflake"
	"github.com/spf13/viper"
)

type Base struct {
	AppName    string
	ConfigPath string
	logger     kitlog.Logger
	Sonyflake  *sonyflake.Sonyflake
}

func (b *Base) Name() string {
	return b.AppName
}

func (b *Base) Logger() kitlog.Logger {
	return b.logger
}

func (b *Base) InitLogger() {
	// context log
	var logger kitlog.Logger
	logger = kitlog.NewLogfmtLogger(os.Stdout)
	logger = kitlog.With(logger, "ts", kitlog.DefaultTimestamp)
	logger = kitlog.With(logger, "caller", kitlog.Valuer(util.CallSite))
	b.logger = logger
}

func (b *Base) InitConfig() {
	if b.ConfigPath == "" || !util.FileExists(b.ConfigPath) {
		panic("no config file for lilin service frame")
	}

	viper.SetConfigType("toml")
	viper.SetConfigFile(b.ConfigPath)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func (b *Base) MustPrepare() types.Closer {
	return func() {}
}

func (b *Base) InitDatabase() types.Closer {
	return func() {}
}

func (b *Base) InitServices() http.Handler {
	panic("InitServices must override")
}

func Run(app Application) {
	app.InitLogger()
	app.InitConfig()

	dbCloser := app.InitDatabase()
	defer dbCloser()

	closer := app.MustPrepare()
	defer closer()

	mux := app.InitServices()

	// cross domain support
	mux = util.AccessControl(mux)

	var (
		errs      = make(chan error, 2)
		addr      = viper.GetString("webserver.addr")
		cert      = viper.GetString("webserver.certificates.cert")
		key       = viper.GetString("webserver.certificates.key")
		enableSSL = viper.GetBool("webserver.enable_ssl")
	)
	go func() {
		app.Logger().Log(app.Name(), "http", "address", addr, "msg", "listening", "enable_ssl", enableSSL)

		if enableSSL {
			errs <- http.ListenAndServeTLS(addr, cert, key, mux)
		} else {
			errs <- http.ListenAndServe(addr, mux)
		}
	}()
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	app.Logger().Log("terminated", <-errs)
}
