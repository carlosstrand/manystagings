package actions

import "github.com/carlosstrand/manystagings/cli/manystagings/orchestratorcli"

type Actions struct {
	orchestratorCLI orchestratorcli.OrchestratorCLI
}

type Options struct {
	OrchestratorCLI orchestratorcli.OrchestratorCLI
}

func NewActions(opts Options) *Actions {
	return &Actions{
		orchestratorCLI: opts.OrchestratorCLI,
	}
}
