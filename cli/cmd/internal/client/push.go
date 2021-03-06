package client

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/suse/carrier/cli/paas"
)

var ()

// CmdPush implements the carrier orgs command
var CmdPush = &cobra.Command{
	Use:   "push NAME",
	Short: "Push an application from the current working directory",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		client, cleanup, err := paas.NewCarrierClient(cmd.Flags(), nil)
		defer func() {
			if cleanup != nil {
				cleanup()
			}
		}()

		if err != nil {
			return errors.Wrap(err, "error initializing cli")
		}

		// TODO - add a parameter for path
		err = client.Push(args[0], ".")
		if err != nil {
			return errors.Wrap(err, "error pushing app")
		}

		return nil
	},
	SilenceErrors: true,
	SilenceUsage:  true,
}
