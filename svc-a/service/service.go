package service

// Validate is a marked service fn at depth 1 from create_task.
//
// a10n:blueprint Components.TaskService.Commands.create_task
func Validate() {
	sanitize()
}

// sanitize is unmarked — trace ends here at depth 2.
func sanitize() {}
