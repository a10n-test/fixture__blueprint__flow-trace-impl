package service

// Validate is a marked service fn at depth 1 from create_task.
//
// a10n:blueprint Products.TaskEngine.Features.TaskLifecycle.task_validated
func Validate() {
	sanitize()
}

// sanitize is unmarked — trace ends here at depth 2.
func sanitize() {}
