package db

import "sync"

var (
	instance *database
	once     sync.Once
)

type Database interface {
	GetDBName() string
	SetTableName(name string)
	GetTableName() string
}

type database struct {
	name      string
	tableName string

	sync.RWMutex
}

func GetConnect(name string) Database {
	once.Do(func() {
		instance = &database{
			name: name,
		}
	})

	return instance
}

func (d *database) GetDBName() string {
	d.RLock()
	defer d.RUnlock()

	return d.name
}

func (d *database) SetTableName(name string) {
	d.Lock()
	defer d.Unlock()

	d.tableName = name
}

func (d *database) GetTableName() string {
	d.RLock()
	defer d.RUnlock()

	return d.tableName
}
