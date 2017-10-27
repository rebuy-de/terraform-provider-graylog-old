package main

import (
	log "github.com/Sirupsen/logrus"

	"github.com/rebuy-de/terraform-provider-graylog/cmd"
)

func main() {
	if err := cmd.NewRootCommand().Execute(); err != nil {
		log.Fatal(err)
	}
}
