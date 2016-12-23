package main

import "fmt"

type beans struct {
	price    int
	quantity int
}

func Ne(p, q int) *beans {
	return &beans{price: p, quantity: q}
}
func (b *beans) Price() int {
	return b.price
}
func (ben beans) String() string {
	return fmt.Sprintf("value%d%d", ben.price, ben.quantity)
}

var p = fmt.Println

func main() {
	bean := Ne(1, 2)
	p(bean)
	p(bean.Price(), bean.String())
}
