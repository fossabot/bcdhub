package main

import (
	"fmt"
	"os"

	"github.com/baking-bad/bcdhub/internal/config"
	"github.com/baking-bad/bcdhub/internal/logger"
	"github.com/baking-bad/bcdhub/internal/models/types"
)

func main() {
	cfg, err := config.LoadDefaultConfig()
	if err != nil {
		logger.Err(err)
		return
	}

	ctx := config.NewContext(
		types.Mainnet,
		config.WithStorage(cfg.Storage, "nginx", 0, cfg.Scripts.Connections.Open, cfg.Scripts.Connections.Idle, false),
		config.WithConfigCopy(cfg),
	)
	defer ctx.Close()

	outputDir := fmt.Sprintf("%s/nginx", cfg.SharePath)
	_ = os.Mkdir(outputDir, os.ModePerm)

	env := os.Getenv("BCD_ENV")
	if env == "" {
		logger.Err(fmt.Errorf("BCD_ENV env var is empty"))
		return
	}

	nginxConfigFilename := fmt.Sprintf("%s/default.%s.conf", outputDir, env)
	if err := makeNginxConfig(nginxConfigFilename, ctx.Config.BaseURL); err != nil {
		logger.Err(err)
		return
	}

	sitemapFilename := fmt.Sprintf("%s/sitemap.%s.xml", outputDir, env)
	if err := makeSitemap(sitemapFilename, ctx.Config); err != nil {
		logger.Err(err)
		return
	}
}
