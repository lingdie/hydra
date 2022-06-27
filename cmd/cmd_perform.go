package cmd

import (
	"github.com/spf13/cobra"

	"github.com/ory/x/cmdx"
)

func NewPerformCmd(root *cobra.Command) *cobra.Command {
	var cmd = &cobra.Command{
		Use:     "perform",
		Aliases: []string{"ls"},
		Short:   "Perform OAuth 2.0 Flows",
	}
	cmd.AddCommand(NewPerformClientCredentialsCmd(root))
	cmdx.RegisterHTTPClientFlags(cmd.PersistentFlags())
	cmdx.RegisterFormatFlags(cmd.PersistentFlags())
	return cmd
}