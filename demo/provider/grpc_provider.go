package provider

import (
	"github.com/9299381/bingo"
	"github.com/9299381/bingo/modules/grpc"
)

type GrpcProvider struct {
}

func (g *GrpcProvider) Boot() {

}

func (g *GrpcProvider) Register() {
	bingo.Bind("grpc", func(module bingo.IModule) error {
		mod := module.(*grpc.Server)
		mod.Route("demo.one", bingo.Handler("demo.one"))
		return nil
	})
}
