package service

import "example.com/ft/repo"

// Create is an intermediate service fn. Carries a Products.* invariant marker
// so the walk passes THROUGH it (not a stopping boundary) and descends into
// the repo layer. This exercises depth-1 node visibility in the rendered tree.
//
// a10n:blueprint Products.TaskEngine.Features.TaskLifecycle.task_created
func Create() {
	repo.Insert()
	repo.Audit()
}

// List is an intermediate service fn with a Products.* marker.
//
// a10n:blueprint Products.TaskEngine.Features.TaskLifecycle.task_listed
func List() {
	repo.Query()
}
