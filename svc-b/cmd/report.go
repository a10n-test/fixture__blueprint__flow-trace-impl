package cmd

import (
	"example.com/svcb/pipeline"
	"example.com/svcb/store"
)

// BuildReport is the svc-b entry point.
//
// a10n:blueprint Components.ReportService.Commands.build_report
func BuildReport() {
	pipeline.Run()
	store.Aggregate()
}
