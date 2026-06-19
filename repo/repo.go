package repo

// Insert is a marked repo fn at depth 2 from create_task.
//
// a10n:blueprint Components.TaskRelationalStore.Dbml.Insert
func Insert() {
	execWrite()
}

// Audit is a marked repo fn at depth 2 from create_task.
//
// a10n:blueprint Components.TaskRelationalStore.AuditRow
func Audit() {
	execLog()
}

// Query is a marked repo fn at depth 2 from list_tasks.
//
// a10n:blueprint Components.TaskRelationalStore.Dbml.Query
func Query() {
	execScan()
}

// execWrite is unmarked — trace ends here at depth 3.
func execWrite() {}

// execLog is unmarked — trace ends here at depth 3.
func execLog() {}

// execScan is unmarked — trace ends here at depth 3.
func execScan() {}
