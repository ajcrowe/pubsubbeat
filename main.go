package main

import (
	"os"

	"github.com/elastic/beats/libbeat/beat"

	"github.com/ajcrowe/pubsubbeat/beater"
)

func main() {
	err := beat.Run("pubsubbeat", "", beater.New)
	if err != nil {
		os.Exit(1)
	}
}
