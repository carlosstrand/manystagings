package orchestratorcli

type OrchestratorCLI interface {
	ProxyApp(namespace string, app string) error
}
