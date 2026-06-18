package store

// Insert is a marked store op; it calls an unmarked helper that the trace
// could descend into at greater depth.
//
// a10n:blueprint Components.TaskRelationalStore.Dbml.Insert
func Insert() { helper() }

// a10n:blueprint Components.TaskRelationalStore.Dbml.Query
func Query() {}

func helper() {}
