package client

type Manager interface {
	Add(client Client)
	Get(id int) (Client, bool)
	Remove(id int)
	Update(client Client)
}

type Client struct {
	ID   int
	Name string
}

type DB map[int]Client

func NewDB() DB {
	return make(map[int]Client)
}

func (db DB) Add(client Client) {
	db[client.ID] = client
}

func (db DB) Get(id int) (Client, bool) {
	client, exists := db[id]
	return client, exists
}

func (db DB) Remove(id int) {
	delete(db, id)
}

func (db DB) Update(client Client) {
	if _, exists := db[client.ID]; exists {
		db[client.ID] = client
	}
}
