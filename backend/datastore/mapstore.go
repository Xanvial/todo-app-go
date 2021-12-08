package datastore

type MapStore struct {
	data map[string]bool
}

func NewMapStore() *MapStore {
	newData := make(map[string]bool, 0)

	return &MapStore{
		data: newData,
	}
}
