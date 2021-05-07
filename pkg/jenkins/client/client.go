package client

import (
	"context"

	"github.com/bndr/gojenkins"

	"github.com/adrianriobo/jkrunner/pkg/jenkins/config"
)

var Jenkins *gojenkins.Jenkins

func init() {
	// var config *JenkinsConfig
	config, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}
	Jenkins = gojenkins.CreateJenkins(nil, config.URL, config.Username, config.Password)
	if _, err = Jenkins.Init(context.Background()); err != nil {
		panic(err)
	}
}
