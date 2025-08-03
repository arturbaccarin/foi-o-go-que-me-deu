package receipt

import (
	"interface-e-coisa-de-consumidor/prefira/client"
	"interface-e-coisa-de-consumidor/prefira/product"
)

type clientManager interface {
	Get(id int) client.Client
}

type productManager interface {
	Get(id int) product.Product
}

type Receipt struct {
	client  clientManager
	product productManager
}

func New(clientManager clientManager, productManager productManager) *Receipt {
	return &Receipt{
		client:  clientManager,
		product: productManager,
	}
}

func (r *Receipt) Generate(clientID int, productID int) string {
	c := r.client.Get(clientID)

	p := r.product.Get(productID)

	receipt := "Receipt for " + c.Name + " for product " + p.Name
	return receipt
}
