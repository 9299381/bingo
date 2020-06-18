package registy

import (
	"context"
	"github.com/9299381/bingo"
	"github.com/9299381/bingo/modules/grpc"
	"github.com/9299381/bingo/modules/http"
	"github.com/9299381/bingo/package/config"
	"github.com/9299381/bingo/package/consul"
	kitconsul "github.com/go-kit/kit/sd/consul"
	"github.com/spf13/viper"
	"google.golang.org/grpc/health/grpc_health_v1"
)

func init() {
	bingo.RegisterModule(new(Registy))
}

type Registy struct {
	bingo.Context
	reg map[string]*kitconsul.Registrar
}

func (r *Registy) ModuleInfo() *bingo.ModuleInfo {
	return &bingo.ModuleInfo{
		ID:      "registy",
		Type:    bingo.MODULE_SERVER,
		Version: "1.0.0",
		New: func() bingo.IModule {
			r := new(Registy)
			r.ConfigModule()
			return r
		},
	}
}

func (r *Registy) ConfigModule() {
	r.Context = bingo.NewContext(context.Background())
	r.reg = make(map[string]*kitconsul.Registrar)

	//由于grpc的服务注册需要在服务开启之前,因此在这里绑定
	bingo.Bind("grpc", func(module bingo.IModule) error {
		r.Log.Info("grpc register to consul")
		mod := module.(*grpc.Server)
		grpc_health_v1.RegisterHealthServer(mod.Server, &healthImpl{})
		r.reg["grpc"] = consul.RegistyGrpc(
			viper.GetString("config.name"),
			config.EnvString("server.grpc_host", ""),
			config.EnvString("server.grpc_port", ""),
		)
		return nil
	})

	bingo.Bind("http", func(module bingo.IModule) error {
		r.Log.Info("http register to consul")
		mod := module.(*http.Server)
		r.reg["http"] = consul.RegistyHttp(
			viper.GetString("config.name"),
			config.EnvString("server.http_host", ""),
			config.EnvString("server.http_port", ""),
		)
		mod.Get("/health", func(ctx context.Context, request interface{}) (response interface{}, err error) {
			return bingo.Success("SERVING"), nil
		})
		return nil
	})

}

func (r *Registy) Start(id string) error {
	return nil
}

func (r *Registy) Stop(id string) {
	for k, v := range r.reg {
		r.Log.Info("%s service deregister", k)
		v.Deregister()
	}
}

var (
	_ bingo.IModule       = (*Registy)(nil)
	_ bingo.IModuleServer = (*Registy)(nil)
)

type healthImpl struct{}

// Check 实现健康检查接口，这里直接返回健康状态，
// 这里也可以有更复杂的健康检查策略，比如根据服务器负载来返回
func (s *healthImpl) Check(ctx context.Context, req *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	return &grpc_health_v1.HealthCheckResponse{
		Status: grpc_health_v1.HealthCheckResponse_SERVING,
	}, nil
}
func (s *healthImpl) Watch(*grpc_health_v1.HealthCheckRequest, grpc_health_v1.Health_WatchServer) error {

	return nil
}
