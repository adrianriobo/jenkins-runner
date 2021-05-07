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
	jobCmd.AddCommand(jobBuildCmd)
	flagSet := pflag.NewFlagSet("build", pflag.ExitOnError)
	flagSet.StringP(name, "n", "", "name of the pipeline to be built")
	flagSet.StringToStringP(parameters, "p", nil, "pipeline parameters to run the build")
	jobBuildCmd.Flags().AddFlagSet(flagSet)
}

var jobBuildCmd = &cobra.Command{
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
