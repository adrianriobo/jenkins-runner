package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/adrianriobo/jkrunner/pkg/jenkins/config"
)

const (
	configCmdName string = "config"
	jenkinsUrl    string = "jenkins-url"
	username      string = "username"
	password      string = "password"
)

func init() {
	rootCmd.AddCommand(configCmd)
	flagSet := pflag.NewFlagSet(configCmdName, pflag.ExitOnError)
	flagSet.StringP(jenkinsUrl, "j", "", "jenkins url")
	flagSet.StringP(username, "u", "", "username")
	flagSet.StringP(password, "p", "", "password")
	configCmd.Flags().AddFlagSet(flagSet)
}

var configCmd = &cobra.Command{
	Use:   configCmdName,
	Short: "config",
	Long:  "config",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := viper.BindPFlags(cmd.Flags()); err != nil {
			return err
		}
		createConfig()
		return nil
	},
}

func createConfig() {
	if err := config.CreateConfig(
		viper.GetString(jenkinsUrl),
		viper.GetString(username),
		viper.GetString(password)); err != nil {
		panic(err)
	}
}
