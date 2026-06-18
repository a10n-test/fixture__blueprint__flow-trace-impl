package cmd

import (
	"example.com/ft/scheduler"
	"example.com/ft/store"
)

// CreateTask is the marked entry point. It calls an unmarked helper
// (validateInput → a [no marker] stop) and a marked store op (Insert →
// accountability-transfer stop). scheduler is imported but only referenced as
// a value, never CALLED — a dead import.
//
// a10n:blueprint Components.TaskService.Commands.create_task
func CreateTask() {
	validateInput()
	store.Insert()
	_ = scheduler.Name
}

func validateInput() {}
