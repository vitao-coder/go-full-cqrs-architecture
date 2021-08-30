package database

type Data interface {
	GetCollection() string
	GetID() string
}

type Database interface {
	Insert(data Data) error
	InsertMany(data []Data) error
	Update(data Data) error
	Get(data Data) (interface{}, error)
}
