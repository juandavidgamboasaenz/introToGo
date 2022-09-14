package main

import (
	"fmt"
)

const (
	ld float32 = 0.1
	lq int     = 5
	md float32 = 0.2
	mq int     = 9
	hd float32 = 0.4
	hq int     = 10
	pr int     = 1100000
)

func main() {

	var quantity int

	fmt.Println("Cuantas computadoras quiere llevar?")
	fmt.Scan(&quantity)

	fmt.Printf("Usted quiere comprar %d computadoras\n", quantity)
	fmt.Printf("El precio de todas las computadoras es %f\n", float32(quantity)*float32(pr))

	discount := calculateDiscount(&quantity)

	fmt.Printf("El precio del total de su descuento es de: %f\n", discount)
	fmt.Printf("Por una oferta especial su compra es de: %f\n", ((float32(quantity) * float32(pr)) - discount))

}

func calculateDiscount(quantity *int) float32 {
	var discount float32

	if *quantity < 5 {
		discount = float32(pr) * ld * float32(*quantity)
	} else if *quantity == 5 || *quantity < 10 {
		discount = float32(pr) * md * float32(*quantity)
	} else if *quantity >= 9 {
		discount = float32(pr) * hd * float32(*quantity)
	}
	return discount
}
