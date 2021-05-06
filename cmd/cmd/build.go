package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	name       string = "name"
	parameters string = "parameters"
)

func init() {
	rootCmd.AddCommand(buildCmd)
	flagSet := pflag.NewFlagSet("build", pflag.ExitOnError)
	flagSet.StringP(name, "n", "", "list of brokers acting on failover")
	flagSet.StringToStringP(parameters, "p", nil, "parameters to run the build")
	buildCmd.Flags().AddFlagSet(flagSet)
}

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "build",
	Long:  "build",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := viper.BindPFlags(cmd.Flags()); err != nil {
			return err
		}
		// runStart()
		return nil
	},
}

// func runStart() {
// 	jenkins.Build(
// 		viper.GetString(name),
// 		viper.GetStringMapString(parameters)
// 	)
// }
