package runner

import (
	"context"
	"os"
	"os/exec"
	"strings"

	"github.com/sirupsen/logrus"
)

func RunTerraform(ctx context.Context, cwd string, args ...string) {
	logrus.Debugf("Running: terraform %s\n", strings.Join(args, " "))
	cmd := exec.CommandContext(
		ctx,
		"terraform",
		args...,
	)
	cmd.Dir = cwd
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()

	if err != nil {
		panic(err)
	}
}
