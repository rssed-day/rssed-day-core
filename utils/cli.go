package utils

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func Result(res string) {
	fmt.Fprint(os.Stdout, res)
}

func Error(cmd *cobra.Command, args []string, err error) {
	fmt.Fprintf(os.Stderr, "execute cmd:%s args:%v error:%v\n", cmd.Name(), args, err)
	os.Exit(1)
}
