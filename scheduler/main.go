package main

import (
	"fmt"
	"os"

	"SmaSchPlugin/pkg/plugin"
	"k8s.io/component-base/logs"
)

func main() {

	command := plugin.Register()

	logs.InitLogs()
	defer logs.FlushLogs()

	if err := command.Execute(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
