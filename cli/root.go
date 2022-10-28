package cli

import (
	"errors"
	"fmt"
	"github.com/rssed-day/rssed-day-core/utils"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "rsseday",
	Short:   "TODO", // TODO
	Long:    "TODO", // TODO
	Example: "TODO", // TODO
	Run: func(cmd *cobra.Command, args []string) {
		utils.Error(cmd, args, errors.New(fmt.Sprintf("%s: command not found", cmd.Name())))
	},
}
