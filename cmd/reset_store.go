package cmd

import (
	"github.com/spf13/cobra"
	flag "github.com/spf13/pflag"

	"github.com/sunriselayer/sunrise-da/nodebuilder"
)

// ResetStore constructs a CLI command to reset the store of Celestia Node.
func ResetStore(fsets ...*flag.FlagSet) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "unsafe-reset-store",
		Short: "Resets the node's store to a new state. Leaves the keystore and config intact.",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()

			return nodebuilder.Reset(StorePath(ctx), NodeType(ctx))
		},
	}
	for _, set := range fsets {
		cmd.Flags().AddFlagSet(set)
	}
	return cmd
}
