package runner

import (
	"context"

	"github.com/labd/mach-composer-go/config"
)

func TerraformApply(cfg *config.MachConfig, locations map[string]string, reuse bool) {
	ctx := context.Background()

	for i := range cfg.Sites {
		site := cfg.Sites[i]
		TerraformApplySite(ctx, cfg, &site, locations[site.Identifier], reuse)
	}
}

func TerraformApplySite(ctx context.Context, cfg *config.MachConfig, site *config.Site, path string, reuse bool) {

	if !reuse {
		RunTerraform(ctx, path, "init")
	}

	cmd := []string{"apply"}

	RunTerraform(ctx, path, cmd...)
}
