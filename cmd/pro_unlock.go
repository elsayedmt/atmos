package cmd

import (
	"github.com/spf13/cobra"

	e "github.com/cloudposse/atmos/internal/exec"
	u "github.com/cloudposse/atmos/pkg/utils"
)

// proUnlockCmd executes 'pro unlock' CLI command
var proUnlockCmd = &cobra.Command{
	Use:                "unlock",
	Short:              "Unlock a stack",
	Long:               `This command calls the atmos pro API and unlocks a stack`,
	FParseErrWhitelist: struct{ UnknownFlags bool }{UnknownFlags: false},
	Args:               cobra.NoArgs,
	ValidArgsFunction:  ComponentsArgCompletion,
	Run: func(cmd *cobra.Command, args []string) {
		// Check Atmos configuration
		checkAtmosConfig()

		err := e.ExecuteProUnlockCommand(cmd, args)
		if err != nil {
			u.PrintErrorMarkdownAndExit("", err, "")
		}
	},
}

func init() {
	proUnlockCmd.PersistentFlags().StringP("component", "c", "", "Specify the Atmos component to lock")
	proUnlockCmd.PersistentFlags().StringP("stack", "s", "", "Specify the Atmos stack to lock")
	AddStackCompletion(proUnlockCmd)
	proCmd.AddCommand(proUnlockCmd)
}
