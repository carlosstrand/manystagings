module github.com/carlosstrand/manystagings/cli/manystagings

go 1.16

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/carlosstrand/manystagings v0.0.0-20210411171433-564912bd1d15
	github.com/go-zepto/zepto v1.0.0-beta.4
	github.com/manifoldco/promptui v0.8.0
	github.com/sirupsen/logrus v1.8.1
	github.com/urfave/cli v1.22.5
	k8s.io/client-go v0.21.0
)

replace github.com/carlosstrand/manystaging => ../..
