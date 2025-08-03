package main

import (
	"fmt"
	"interface-e-coisa-de-consumidor/prefira/client"
	"interface-e-coisa-de-consumidor/prefira/product"
	"interface-e-coisa-de-consumidor/prefira/receipt"
)

func main() {
	clientDB := client.NewDB()
	productDB := product.NewDB()

	clientDB.Add(client.Client{ID: 1, Name: "John Doe"})
	productDB.Add(product.Product{ID: 1, Name: "Laptop"})

	receipt := receipt.New(clientDB, productDB)

	receiptText := receipt.Generate(1, 1)
	fmt.Println(receiptText)
}
