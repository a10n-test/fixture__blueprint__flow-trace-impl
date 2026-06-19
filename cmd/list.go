package cmd

import "example.com/ft/service"

// ListTasks is a second marked entry point.
//
// a10n:blueprint Components.TaskService.Commands.list_tasks
func ListTasks() {
	service.List()
}
