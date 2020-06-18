package http

import (
	"context"
	"github.com/9299381/bingo"
	"github.com/9299381/bingo/package/config"
	"github.com/go-kit/kit/endpoint"
	"github.com/gorilla/mux"
	httpserver "net/http"
	"strings"
)

func init() {
	bingo.RegisterModule(new(Server))
}

type Server struct {
	bingo.Context
	*mux.Router
	Host string
	Port string
}

func (s *Server) ModuleInfo() *bingo.ModuleInfo {
	return &bingo.ModuleInfo{
		ID:      "http",
		Type:    bingo.MODULE_SERVER,
		Version: "1.0",
		New: func() bingo.IModule {
			server := new(Server)
			server.ConfigModule()
			return server
		},
	}
}
func (s *Server) ConfigModule() {
	s.Context = bingo.NewContext(context.Background())
	s.Router = mux.NewRouter()
	s.Host = config.EnvString("server.http_host", "0.0.0.0")
	s.Port = config.EnvString("server.http_port", "8341")
}

func (s *Server) Start(id string) error {
	address := strings.Join([]string{s.Host, s.Port}, ":")
	s.Log.Infof("%s start at %s", id, address)
	return httpserver.ListenAndServe(address, s.Router)

}

func (s *Server) Route(method string, path string, endpoint endpoint.Endpoint) {
	s.Methods(method).
		Path(path).
		Handler(NewHTTP(endpoint))
}

func (s *Server) Get(path string, endpoint endpoint.Endpoint) {
	s.Methods("GET").
		Path(path).
		Handler(NewHTTP(endpoint))
}
func (s *Server) Post(path string, endpoint endpoint.Endpoint) {
	s.Methods("POST").
		Path(path).
		Handler(NewHTTP(endpoint))
}

func (s *Server) WeChatNotify(path string, endpoint endpoint.Endpoint) {
	s.Methods("POST").
		Path(path).
		Handler(NewWeChatNotify(endpoint))
}

func (s *Server) Upload(path string, endpoint endpoint.Endpoint) {
	s.Methods("POST").
		Path(path).
		Handler(NewUpload(endpoint))
}

func (s *Server) Stop(id string) {
	s.Log.Infof("%s stop now", id)
}

var (
	_ bingo.IModule       = (*Server)(nil)
	_ bingo.IModuleServer = (*Server)(nil)
)
