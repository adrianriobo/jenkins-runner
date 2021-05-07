package jkrunner

import (
	"time"
)

const (
	Home              string        = ".jkrunner"
	LogFileName       string        = "jkrunner.log"
	ConfigFileName    string        = "config.yaml"
	BuildWaitInterval time.Duration = 5000 * time.Millisecond
)
