package dbconn

import (
	"htmx.try/m/v2/pkg/domain"
)

type InMemoryDB struct {
	data      map[string]domain.InterfaceResponseFull
	responses map[string][]domain.Response
}

func NewInMemoryDB() *InMemoryDB {
	return &InMemoryDB{
		data:      make(map[string]domain.InterfaceResponseFull),
		responses: make(map[string][]domain.Response),
	}
}

func (db *InMemoryDB) GetData(key string) (domain.InterfaceResponseFull, bool) {
	val, ok := db.data[key]
	return val, ok
}

func (db *InMemoryDB) SetData(key string, value domain.InterfaceResponseFull) {
	db.data[key] = value
}

func (db *InMemoryDB) DeleteData(key string) {
	delete(db.data, key)
}

func (db *InMemoryDB) GetResponses(key string) ([]domain.Response, bool) {
	val, ok := db.responses[key]
	return val, ok
}

func (db *InMemoryDB) SetResponse(key string, value domain.Response) {
	db.responses[key] = append(db.responses[key], value)
}

func (db *InMemoryDB) DeleteResponses(key string) {
	delete(db.responses, key)
}
