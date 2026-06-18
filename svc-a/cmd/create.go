package cmd

import "example.com/svca/store"

// a10n:blueprint Components.TaskService.Commands.create_task
func CreateTask() {
	store.Insert()
}
