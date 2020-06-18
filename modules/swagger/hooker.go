package swagger

import (
	"context"
	"github.com/9299381/bingo"
	bingo_http "github.com/9299381/bingo/modules/http"
	"net/http"
	"os/exec"
)

func init() {
	bingo.RegisterModule(new(Swagger))
}

type Swagger struct {
	bingo.Context
}

func (s *Swagger) ModuleInfo() *bingo.ModuleInfo {
	return &bingo.ModuleInfo{
		ID:      "swagger",
		Type:    bingo.MODULE_NONE,
		Version: "1.0.0",
		New: func() bingo.IModule {
			s := new(Swagger)
			s.ConfigModule()
			return s
		},
	}
}

func (s *Swagger) ConfigModule() {
	s.Context = bingo.NewContext(context.Background())
	bingo.Bind("http", func(module bingo.IModule) error {
		s.Log.Info("http swagger start")
		mod := module.(*bingo_http.Server)
		fs := http.FileServer(http.Dir("./swaggerui/"))
		mod.Methods("GET").
			PathPrefix("/swagger/").
			Handler(http.StripPrefix("/swagger/", fs))

		mod.Methods("GET").
			Path("/swagger_generate").
			HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				cmd := exec.Command(
					"swagger",
					"generate", "spec", "-o", "swaggerui/swagger.json")
				err := cmd.Run()
				response := bingo.MakeResponse("ok", err)
				_ = HttpEncodeResponse(context.Background(), w, response)
			})
		return nil
	})
}
