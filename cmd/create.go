package cmd

import "example.com/ft/service"

// CreateTask is the entry point. Calls the service layer (depth 1, marked),
// which calls the repo layer (depth 2, marked), which calls an unmarked helper
// (depth 3 — trace ends there).
//
// a10n:blueprint Components.TaskService.Commands.create_task
func CreateTask() {
	service.Create()
}
