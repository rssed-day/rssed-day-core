package cli

import (
	"fmt"
	"github.com/rssed-day/rssed-day-core/context"
	"github.com/rssed-day/rssed-day-core/services"
	"github.com/rssed-day/rssed-day-core/utils"
	"github.com/spf13/cobra"
)

var path string

var pipeCmd = &cobra.Command{
	Use:     "pipe",
	Short:   "TODO", // TODO
	Long:    "TODO", // TODO
	Example: "TODO", // TODO
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := cmd.Flags().GetString("config")
		if err != nil {
			utils.Error(cmd, args, err)
		}
		factory := context.NewFileConfigFactory(cfg)
		if err := services.NewPipelineService().Pipe(factory); err != nil {
			utils.Error(cmd, args, err)
		}
		utils.Result(fmt.Sprintf("pipe %s done", cfg))
	},
}

func init() {
	rootCmd.AddCommand(pipeCmd)

	pipeCmd.Flags().StringVarP(&path, "config", "c", "configs/plugins.yaml",
		"plugin config file path")
}
