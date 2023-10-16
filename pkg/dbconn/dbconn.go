package dbconn

import (
	entities "htmx.try/m/v2/pkg/domain"
)

type InMemoryDB struct {
	data map[string]entities.InterfaceResponseFull
}

func NewInMemoryDB() *InMemoryDB {
	return &InMemoryDB{
		data: make(map[string]entities.InterfaceResponseFull),
	}
}

func (db *InMemoryDB) Get(key string) (entities.InterfaceResponseFull, bool) {
	val, ok := db.data[key]
	return val, ok
}

func (db *InMemoryDB) Set(key string, value entities.InterfaceResponseFull) {
	db.data[key] = value
}

func (db *InMemoryDB) Delete(key string) {
	delete(db.data, key)
}
