package cmd

import "example.com/ft/store"

// ListTasks is a second marked entry point so whole-project tracing shows
// more than one root.
//
// a10n:blueprint Components.TaskService.Commands.list_tasks
func ListTasks() {
	store.Query()
}
