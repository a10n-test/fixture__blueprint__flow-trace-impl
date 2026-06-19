package pipeline

// Run is a marked pipeline fn at depth 1 from build_report.
//
// a10n:blueprint Products.ReportEngine.Features.ReportLifecycle.report_built
func Run() {
	fetch()
	transform()
}

// fetch is unmarked — trace ends here.
func fetch() {}

// transform is unmarked — trace ends here.
func transform() {}
