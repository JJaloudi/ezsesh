package ezsesh

type EzSesh struct {
	store *EzStore
}

func CreateEZSesh(store *EzStore) *EzSesh {
	return &EzSesh{
		store: store,
	}
}

func (sesh *EzSesh) SetStore(store *EzStore) {
	sesh.store = store
}

func (sesh *EzSesh) GetStore() *EzStore {
	return sesh.store
}
