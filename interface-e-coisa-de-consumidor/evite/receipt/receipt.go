package receipt

import (
	"interface-e-coisa-de-consumidor/evite/client"
	"interface-e-coisa-de-consumidor/evite/product"
)

type Receipt struct {
	client  client.Manager
	product product.Manager
}

func New(clientManager client.Manager, productManager product.Manager) *Receipt {
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
