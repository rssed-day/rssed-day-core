package cli

import (
	"errors"
	"fmt"
	"github.com/rssed-day/rssed-day-core/utils"
	"github.com/spf13/cobra"
	"io/ioutil"
	"strings"
)

var versionCmd = &cobra.Command{
	Use:     "version",
	Short:   "TODO", // TODO
	Long:    "TODO", // TODO
	Example: "TODO", // TODO
	Run: func(cmd *cobra.Command, args []string) {
		res, err := ioutil.ReadFile(".ci/VERSION")
		if err != nil {
			utils.Error(cmd, args, err)
		}
		parts := strings.Split(string(res), "=")
		if len(parts) != 2 {
			utils.Error(cmd, args, errors.New(fmt.Sprintf(".ci/VERSION not correct")))
		}
		version := parts[1]
		utils.Result(version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
