package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/adrianriobo/jkrunner/pkg/jenkins/client"
)

const (
	jobBuildCmdName string = "build"

	name       string = "name"
	parameters string = "parameters"
	wait       string = "wait"
	output     string = "output"
)

func init() {
	jobCmd.AddCommand(jobBuildCmd)
	flagSet := pflag.NewFlagSet(jobBuildCmdName, pflag.ExitOnError)
	flagSet.StringP(name, "n", "", "name of the pipeline to be built")
	flagSet.StringToStringP(parameters, "p", nil, "pipeline parameters to run the build")
	flagSet.BoolP(wait, "w", false, "wait for build to finish")
	flagSet.StringP(output, "o", "", "output path to download artifacts and console ouput for the job when done")
	jobBuildCmd.Flags().AddFlagSet(flagSet)
}

var jobBuildCmd = &cobra.Command{
	Use:   jobBuildCmdName,
	Short: "build a job with parameters",
	Long:  "build a job wiht parameters",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := viper.BindPFlags(cmd.Flags()); err != nil {
			return err
		}
		return run()
	},
}

func run() error {
	return client.Build(
		viper.GetString(name),
		viper.GetStringMapString(parameters),
		viper.GetBool(wait),
		viper.GetString(output),
	)
}
