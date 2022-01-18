package tracker

func NewImport() (*Import, error) {
	const ou UID = "O6uvpzGd5pu"  // ??
	const tet UID = "nEenWmSyUEp" // person
	te := TrackedEntity{OrgUnit: ou, TrackedEntityType: tet}
	uid, err := NewUID()
	if err != nil {
		return nil, err
	}
	te.TrackedEntity = uid
	im := &Import{}
	im.TrackedEntities.TrackedEntities = append(im.TrackedEntities.TrackedEntities, te)
	return im, nil
}
