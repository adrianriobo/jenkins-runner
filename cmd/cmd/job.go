package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func init() {
	rootCmd.AddCommand(jobCmd)
	flagSet := pflag.NewFlagSet("job", pflag.ExitOnError)
	jobCmd.Flags().AddFlagSet(flagSet)
}

var jobCmd = &cobra.Command{
	Use:   "job",
	Short: "job",
	Long:  "job",
}
