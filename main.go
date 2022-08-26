package main

import (
	"context"

	"github.com/rexray/gocsi"

	"github.com/jmccormick2001/jeff-csi-driver/provider"
	"github.com/jmccormick2001/jeff-csi-driver/service"
)

// main is ignored when this package is built as a go plug-in.
func main() {
	gocsi.Run(
		context.Background(),
		service.Name,
		"A description of the SP",
		"",
		provider.New())
}
