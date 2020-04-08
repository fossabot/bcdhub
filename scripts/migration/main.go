package main

import (
	"github.com/baking-bad/bcdhub/internal/jsonload"
	"github.com/baking-bad/bcdhub/scripts/migration/migrations"
)

func main() {
	var cfg migrations.Config
	if err := jsonload.StructFromFile("config.json", &cfg); err != nil {
		panic(err)
	}
	cfg.Print()

	ctx, err := migrations.NewContext(cfg)
	if err != nil {
		panic(err)
	}
	defer ctx.Close()

	migration := migrations.SetBMDTimestamp{}
	if err := migration.Do(ctx); err != nil {
		panic(err)
	}

	migration2 := migrations.SetBMDKeyStrings{}
	if err := migration2.Do(ctx); err != nil {
		panic(err)
	}
}
