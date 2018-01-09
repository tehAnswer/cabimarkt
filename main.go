package main

import (
	"fmt"
)

func main() {
	// NOTE: Use this function as you please. It would be great to have some text interface though.
	fmt.Println("Welcome to Cabimartk!")
	checkout := NewCheckout()
	checkout.Scan("VOUCHER")
	fmt.Println("*pew* VOUCHER *pew*")
	checkout.Scan("VOUCHER")
	fmt.Println("*pew* VOUCHER *pew*")
	total, _ := checkout.GetTotal()
	fmt.Printf("TOTAL: %v\n", total)
	fmt.Printf("PROMOS: %v\n", checkout.PromotionRefs)
}
