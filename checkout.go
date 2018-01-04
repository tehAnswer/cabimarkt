package main

type Checkout struct{}

func NewCheckout() *Checkout {
	return &Checkout{}
}

func (*Checkout) Scan(itemCode string) error {
	return nil
}

func (*Checkout) GetTotal() float64 {
	return -1
}
