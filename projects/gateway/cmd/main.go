package main

import (
	"github.com/solo-io/solo-kit/pkg/utils/log"
	"github.com/solo-io/gloo/projects/gateway/pkg/setup"
)

func main() {
	if err := setup.Main(); err != nil {
		log.Fatalf("err in main: %v", err.Error())
	}
}
