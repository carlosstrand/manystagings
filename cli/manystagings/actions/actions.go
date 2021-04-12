package actions

import (
	"errors"

	"github.com/carlosstrand/manystagings/cli/manystagings/client"
	"github.com/carlosstrand/manystagings/cli/manystagings/orchestratorcli"
	"github.com/carlosstrand/manystagings/cli/manystagings/utils/msconfig"
)

var ErrAppNotFound = errors.New("app not found")

type Actions struct {
	orchestratorCLI orchestratorcli.OrchestratorCLI
	client          *client.Client
	config          *msconfig.ManyStagingsConfig
}

type Options struct {
	OrchestratorCLI orchestratorcli.OrchestratorCLI
	Client          *client.Client
	Config          *msconfig.ManyStagingsConfig
}

func NewActions(opts Options) *Actions {
	return &Actions{
		orchestratorCLI: opts.OrchestratorCLI,
		client:          opts.Client,
		config:          opts.Config,
	}
}
