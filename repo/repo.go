package repo

import "example.com/ft/repo/db"

// Insert is a marked repo fn at depth 2 from create_task.
//
// a10n:blueprint Components.TaskRelationalStore.Dbml.Insert
func Insert() {
	db.Write()
}

// Audit is a marked repo fn at depth 2 from create_task.
//
// a10n:blueprint Components.TaskRelationalStore.AuditRow
func Audit() {
	db.Log()
}

// Query is a marked repo fn at depth 2 from list_tasks.
//
// a10n:blueprint Components.TaskRelationalStore.Dbml.Query
func Query() {
	db.Scan()
}
