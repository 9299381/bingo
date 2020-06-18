package provider

import (
	"github.com/9299381/bingo"
	"github.com/9299381/bingo/modules/cronjob"
)

type CronProvider struct {
}

func (c *CronProvider) Boot() {
}

func (c *CronProvider) Register() {
	bingo.Bind("cron", func(module bingo.IModule) error {
		mod := module.(*cronjob.Server)
		mod.Route("*/2 * * * * *", bingo.Handler("demo.cron"))
		return nil
	})
}
