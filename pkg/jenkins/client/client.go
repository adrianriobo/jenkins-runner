package client

import (
	"context"
	"time"

	"github.com/adrianriobo/jkrunner/pkg/jenkins/config"
	"github.com/adrianriobo/jkrunner/pkg/jkrunner"
	"github.com/bndr/gojenkins"
)

type jenkinsClient struct {
	jenkins *gojenkins.Jenkins
	ctx     context.Context
}

var client *jenkinsClient

func Build(jobName string, params map[string]string, wait bool) error {
	build, err := build(jobName, params)
	if err != nil {
		return err
	}
	if wait {
		for build.IsRunning(client.ctx) {
			time.Sleep(jkrunner.BuildWaitInterval)
			if _, err = build.Poll(client.ctx); err != nil {
				return err
			}
		}
	}
	return nil
}

func build(jobName string, params map[string]string) (*gojenkins.Build, error) {
	if err := initalize(); err != nil {
		return nil, err
	}
	queueid, err := client.jenkins.BuildJob(client.ctx, jobName, params)
	if err != nil {
		return nil, err
	}
	return client.jenkins.GetBuildFromQueueID(client.ctx, queueid)
}

func initalize() error {
	config, err := config.LoadConfig()
	if err != nil {
		return err
	}
	client = &jenkinsClient{
		jenkins: gojenkins.CreateJenkins(nil, config.URL, config.Username, config.Password),
		ctx:     context.Background()}
	if _, err = client.jenkins.Init(client.ctx); err != nil {
		return err
	}
	return nil
}
