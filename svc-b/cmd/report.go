package cmd

import "example.com/svcb/store"

// a10n:blueprint Components.ReportService.Commands.build_report
func BuildReport() {
	store.Aggregate()
}
