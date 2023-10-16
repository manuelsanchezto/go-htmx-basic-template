package dbconn

import (
	entities "htmx.try/m/v2/pkg/domain"
)

type InMemoryDB struct {
	data map[string]entities.ConversationByUsers
}

func NewInMemoryDB() *InMemoryDB {
	return &InMemoryDB{
		data: make(map[string]entities.ConversationByUsers),
	}
}

func (db *InMemoryDB) Get(key string) (entities.ConversationByUsers, bool) {
	val, ok := db.data[key]
	return val, ok
}

func (db *InMemoryDB) Set(key string, value entities.ConversationByUsers) {
	db.data[key] = value
}

func (db *InMemoryDB) Delete(key string) {
	delete(db.data, key)
}
