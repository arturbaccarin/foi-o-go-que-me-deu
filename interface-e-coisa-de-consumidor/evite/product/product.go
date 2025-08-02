package product

type Manager interface {
	Add(product Product)
	Get(id int) Product
	Remove(id int)
	Update(product Product)
}

type Product struct {
	ID   int
	Name string
}

type DB map[int]Product

func NewDB() DB {
	return make(map[int]Product)
}

func (db DB) Add(product Product) {
	db[product.ID] = product
}

func (db DB) Get(id int) Product {
	product := db[id]
	return product
}

func (db DB) Remove(id int) {
	delete(db, id)
}

func (db DB) Update(product Product) {
	if _, exists := db[product.ID]; exists {
		db[product.ID] = product
	}
}
