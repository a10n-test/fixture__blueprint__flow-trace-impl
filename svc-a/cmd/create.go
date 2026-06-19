package cmd

import (
	"example.com/svca/service"
	"example.com/svca/store"
)

// CreateTask is the svc-a entry point.
//
// a10n:blueprint Components.TaskService.Commands.create_task
func CreateTask() {
	service.Validate()
	store.Insert()
}

// helper is unmarked — trace ends here.
func helper() {}
