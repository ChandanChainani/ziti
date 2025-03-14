/*
	Copyright NetFoundry Inc.

	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

	https://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
*/

package cmd

import (
	"io"

	"github.com/openziti/ziti/common/version"
	cmdhelper "github.com/openziti/ziti/ziti/cmd/ziti/cmd/helpers"
	"github.com/openziti/ziti/ziti/cmd/ziti/cmd/templates"
	c "github.com/openziti/ziti/ziti/cmd/ziti/constants"
	"github.com/spf13/cobra"
)

var (
	upgradeZitiTunnelLong = templates.LongDesc(`
		Upgrades the Ziti Tunnel app if there is a newer release
`)

	upgradeZitiTunnelExample = templates.Examples(`
		# Upgrades the Ziti Tunnel app 
		ziti upgrade ziti-tunnel
	`)
)

// UpgradeZitiTunnelOptions the options for the upgrade ziti-tunnel command
type UpgradeZitiTunnelOptions struct {
	CommonOptions

	Version string
}

// NewCmdUpgradeZitiTunnel defines the command
func NewCmdUpgradeZitiTunnel(out io.Writer, errOut io.Writer) *cobra.Command {
	options := &UpgradeZitiTunnelOptions{
		CommonOptions: CommonOptions{
			Out: out,
			Err: errOut,
		},
	}

	cmd := &cobra.Command{
		Use:     "ziti-tunnel",
		Short:   "Upgrades the Ziti Tunnel app - if there is a new version available",
		Aliases: []string{"tunnel", "rtr", "r"},
		Long:    upgradeZitiTunnelLong,
		Example: upgradeZitiTunnelExample,
		Run: func(cmd *cobra.Command, args []string) {
			options.Cmd = cmd
			options.Args = args
			err := options.Run()
			cmdhelper.CheckErr(err)
		},
	}
	cmd.Flags().StringVarP(&options.Version, "version", "v", "", "The specific version to upgrade to")
	options.AddCommonFlags(cmd)
	return cmd
}

// Run implements the command
func (o *UpgradeZitiTunnelOptions) Run() error {
	newVersion, err := o.getLatestZitiAppVersion(version.GetBranch(), c.ZITI_TUNNEL)
	if err != nil {
		return err
	}

	newVersionStr := newVersion.String()

	if o.Version != "" {
		newVersionStr = o.Version
	}

	o.deleteInstalledBinary(c.ZITI_TUNNEL)

	return o.installZitiApp(version.GetBranch(), c.ZITI_TUNNEL, true, newVersionStr)
}
