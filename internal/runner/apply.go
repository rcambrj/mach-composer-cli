package runner

import (
	"context"
	"github.com/mach-composer/mach-composer-cli/internal/config"
	"github.com/mach-composer/mach-composer-cli/internal/dependency"
	"github.com/mach-composer/mach-composer-cli/internal/utils"
	"github.com/rs/zerolog/log"
	"strings"
)

type ApplyOptions struct {
	Reuse       bool
	Destroy     bool
	AutoApprove bool
}

func TerraformApply(ctx context.Context, cfg *config.MachConfig, dg *dependency.Graph, opts *ApplyOptions) error {
	if opts.Reuse == false {
		if err := terraformInitAll(ctx, dg); err != nil {
			return err
		}
	} else {
		log.Info().Msgf("Reusing existing terraform state")
	}

	if err := batchRun(ctx, dg, cfg.MachComposer.Deployment.Runners, func(ctx context.Context, _ dependency.Node, tfPath string) (string, error) {
		return terraformApply(ctx, tfPath, opts)
	}); err != nil {
		return err
	}

	return nil
}

func terraformApply(ctx context.Context, path string, opt *ApplyOptions) (string, error) {
	cmd := []string{"apply"}

	if opt.Destroy {
		cmd = append(cmd, "-destroy")
	}

	if opt.AutoApprove {
		cmd = append(cmd, "-auto-approve")
	}

	// If there is a plan then we should use it.
	planFilename, err := hasTerraformPlan(path)
	if err != nil {
		return "", err
	}
	if planFilename != "" {
		cmd = append(cmd, strings.TrimPrefix(planFilename, path+"/"))
	}

	return utils.RunTerraform(ctx, false, path, cmd...)
}
