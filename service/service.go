package service

import "example.com/ft/repo"

// Create is a marked service fn at depth 1 from create_task.
//
// a10n:blueprint Components.TaskService.Commands.create_task
func Create() {
	repo.Insert()
	repo.Audit()
}

// List is a marked service fn at depth 1 from list_tasks.
//
// a10n:blueprint Components.TaskService.Commands.list_tasks
func List() {
	repo.Query()
}
