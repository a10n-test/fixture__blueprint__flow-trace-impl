package service

import "example.com/ft/repo"

// Create is an intermediate service fn. Carries a Products.* invariant marker
// so the walk passes THROUGH it (not a stopping boundary) and descends into
// the repo layer. Also calls validate() — an unmarked leaf that shows as a
// ··· trace ends stop at depth 2 (still in service/).
//
// a10n:blueprint Products.TaskEngine.Features.TaskLifecycle.task_created
func Create() {
	validate()
	repo.Insert()
	repo.Audit()
}

// validate is an unmarked helper — trace ends here at depth 2.
func validate() {}

// List is an intermediate service fn with a Products.* marker.
//
// a10n:blueprint Products.TaskEngine.Features.TaskLifecycle.task_listed
func List() {
	preflight()
	repo.Query()
}

// preflight is an unmarked helper — trace ends here at depth 2.
func preflight() {}
