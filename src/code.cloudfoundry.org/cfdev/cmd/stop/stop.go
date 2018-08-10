package stop

import (
	"code.cloudfoundry.org/cfdev/config"
	"code.cloudfoundry.org/cfdev/daemon"
	"code.cloudfoundry.org/cfdev/process"
	"github.com/spf13/cobra"
)

//go:generate mockgen -package mocks -destination mocks/launchd.go code.cloudfoundry.org/cfdev/cmd/stop Launchd
type Launchd interface {
	Stop(spec daemon.DaemonSpec) error
	RemoveDaemon(spec daemon.DaemonSpec) error
}

//go:generate mockgen -package mocks -destination mocks/cfdevd_client.go code.cloudfoundry.org/cfdev/cmd/stop CfdevdClient
type CfdevdClient interface {
	Uninstall() (string, error)
}

//go:generate mockgen -package mocks -destination mocks/process_manager.go code.cloudfoundry.org/cfdev/cmd/stop ProcManager
type ProcManager interface {
	SafeKill(string, string) error
}

type UI interface {
	Say(message string, args ...interface{})
}

//go:generate mockgen -package mocks -destination mocks/analytics.go code.cloudfoundry.org/cfdev/cmd/stop Analytics
type Analytics interface {
	Event(event string, data ...map[string]interface{}) error
}

//go:generate mockgen -package mocks -destination mocks/network.go code.cloudfoundry.org/cfdev/cmd/stop HostNet
type HostNet interface {
	RemoveLoopbackAliases(...string) error
}

type Stop struct {
	Config       config.Config
	Launchd      Launchd
	HyperV       *process.HyperV
	ProcManager  ProcManager
	CfdevdClient CfdevdClient
	Analytics    Analytics
	HostNet      HostNet
}

func (s *Stop) Cmd() *cobra.Command {
	return &cobra.Command{
		Use:  "stop",
		RunE: s.RunE,
	}
}
