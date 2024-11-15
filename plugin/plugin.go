package plugin

import (
	"github.com/hashicorp/go-hclog"
	"os/exec"
)

type Config struct {
	Cmd    *exec.Cmd
	Logger hclog.Logger
}
