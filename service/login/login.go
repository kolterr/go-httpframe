package login

import (
	"go-httpframe/internal/encoding"
	"go-httpframe/protocol"
	"net/http"

	"go-httpframe/internal/errutil"

	kitlog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/opentracing"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	stdopentracing "github.com/opentracing/opentracing-go"
)

type Service interface {
	UserLogin(*protocol.UserLoginRequest) (*protocol.UserLoginResponse, error)           // 用户登录
	ThirdUserLogin(*protocol.ThirdUserLoginRequest) (*protocol.UserLoginResponse, error) // 三方登录登录
	CheckUser(token string) (*protocol.UserInfo, error)                                  // 登录验证
	UserLogout(token string) error                                                       // 用户注销
	GetOptions() string                                                                  // 获取支持操作
}

var supportOptions = `{
	"POST": "/v1/user/login",
	"POST": "/v1/user/check",
	"POST": "/v1/user/logout",
}`

type service struct {
}

func NewService() Service {
	return &service{}
}

func (s *service) UserLogin(r *protocol.UserLoginRequest) (*protocol.UserLoginResponse, error) {

	return nil, errutil.ErrNotImplemented
}

func (s *service) ThirdUserLogin(r *protocol.ThirdUserLoginRequest) (*protocol.UserLoginResponse, error) {
	return nil, errutil.ErrNotImplemented
}
func (s *service) CheckUser(token string) (*protocol.UserInfo, error) {
	return &protocol.UserInfo{}, nil
	//um, err := tokenutil.UserMeta(token)
	//if err != nil {
	//	return nil, err
	//}
	//app, err := model.QueryApp(um.AppID)
	//if err != nil {
	//	return nil, err
	//}
	//
	//uuid, err := model.QueryUUID(um.Uid, um.AppID)
	//if err != nil && err != errutil.ErrUUIDNotFound {
	//	println(err.Error())
	//	return nil, err
	//}
	///*
	//  此uuid当且仅当在第一次执行check操作时生成
	//*/
	//if err == errutil.ErrUUIDNotFound {
	//	uuid, err = model.InsertUUID(um.Uid, um.AppID)
	//	if err != nil {
	//		println(err.Error())
	//		return nil, err
	//	}
	//}
	//buffer := bytes.Buffer{}
	//buffer.WriteString("uuid=" + uuid)
	//sign, err := algoutil.Sign(buffer.Bytes(), app.AppSecret)
	//if err != nil {
	//	return nil, err
	//}
	//
	//return &protocol.UserInfo{UUID: uuid, Sign: sign}, nil
}

func (s *service) UserLogout(token string) error {
	return nil
	//if !cache.Exists(token) {
	//	return errutil.ErrTokenNotFound
	//}
	//um := &types.UserMeta{}
	//err := cache.Struct(token, um)
	//if err != nil {
	//	return err
	//}
	//u, err := model.QueryUser(um.Uid)
	//if err != nil {
	//	return err
	//}
	//u.IsOnline = model.UserOffline
	//model.UpdateUser(u)
	//model.UpdateLogoutTime(u.Uid)
	//cache.Delete(token)
	return nil
}

func (*service) GetOptions() string {
	return supportOptions
}

func MakeLoginHandler(logger kitlog.Logger, tracer stdopentracing.Tracer) http.Handler {
	options := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(encoding.EncodeError),
		httptransport.ServerErrorLogger(logger),
	}

	m := mux.NewRouter()
	m.Handle("/login", httptransport.NewServer(
		makeLoginEndpoint(),
		decodeLoginRequest,
		encodeLoginResponse,
		append(options, httptransport.ServerBefore(opentracing.FromHTTPRequest(tracer, "login", logger)))...,
	))

	return m
}
