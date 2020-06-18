package grpc

import (
	"context"
	"github.com/9299381/bingo"
	"github.com/9299381/bingo/modules/grpc/protobuf"
	"github.com/9299381/bingo/package/config"
	"github.com/go-kit/kit/endpoint"
	GrpcTransport "github.com/go-kit/kit/transport/grpc"
	grpcserver "google.golang.org/grpc"
	"net"
	"strings"
)

func init() {
	bingo.RegisterModule(new(Server))
}

type Server struct {
	*grpcserver.Server
	bingo.Context
	Host string
	Port string
}

func (s *Server) ModuleInfo() *bingo.ModuleInfo {
	return &bingo.ModuleInfo{
		ID:      "grpc",
		Type:    bingo.MODULE_SERVER,
		Version: "1.0.0",
		New: func() bingo.IModule {
			server := new(Server)
			server.ConfigModule()
			return server
		},
	}
}

func (s *Server) ConfigModule() {
	s.Context = bingo.NewContext(context.Background())
	s.Server = grpcserver.NewServer()
	s.Host = config.EnvString("server.grpc_host", "0.0.0.0")
	s.Port = config.EnvString("server.grpc_port", "9341")
}

func (s *Server) Route(name string, endpoint endpoint.Endpoint) {

	sd := s.getServiceDesc(name)
	service := &grpcService{
		handler: NewGRPC(endpoint),
	}
	s.RegisterService(&sd, service)
}

func (s *Server) Start(id string) error {
	address := strings.Join([]string{s.Host, s.Port}, ":")
	s.Log.Infof("%s start at %s", id, address)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	return s.Serve(lis)
}

func (s *Server) Stop(id string) {
	s.Log.Infof("%s stop now", id)
}

//------------------------------//
func (s *Server) getServiceDesc(name string) grpcserver.ServiceDesc {
	var serviceDesc = grpcserver.ServiceDesc{
		ServiceName: "protobuf." + name,
		HandlerType: (*protobuf.ServiceServer)(nil),
		Methods: []grpcserver.MethodDesc{
			{
				MethodName: "Handle",
				Handler:    s.serviceHandleHandler,
			},
		},
		Streams:  []grpcserver.StreamDesc{},
		Metadata: "message.proto",
	}
	return serviceDesc
}

func (s *Server) serviceHandleHandler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpcserver.UnaryServerInterceptor) (interface{}, error) {
	in := new(protobuf.Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	//interceptor 拦截器
	if interceptor == nil {
		return srv.(protobuf.ServiceServer).Handle(ctx, in)
	}
	info := &grpcserver.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.Service/Handle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(protobuf.ServiceServer).Handle(ctx, req.(*protobuf.Request))
	}
	return interceptor(ctx, in, info, handler)
}

type grpcService struct {
	handler GrpcTransport.Handler
}

func (s *grpcService) Handle(ctx context.Context, req *protobuf.Request) (*protobuf.Response, error) {
	_, rsp, err := s.handler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rsp.(*protobuf.Response), err
}

//---------------

var (
	_ bingo.IModule       = (*Server)(nil)
	_ bingo.IModuleServer = (*Server)(nil)
)
