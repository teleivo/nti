package tracker

type TrackedEntity struct {
	OrgUnit           UID `json:"orgUnit"`
	TrackedEntity     UID `json:"trackedEntity"`
	TrackedEntityType UID `json:"trackedEntityType"`
}

type TrackedEntities struct {
	TrackedEntities []TrackedEntity `json:"trackedEntities"`
}

type Import struct {
	TrackedEntities
}
