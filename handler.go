package main

import (
	"fmt"
)

type Subtotal struct {
	Items        map[string]int
	PromotionRef string
	FinalPrice   float64
}

type Handler interface {
	Call(line *CheckoutLine, subtotalChannel chan float64, errChannel chan error)
}

type MinPriceHandler struct {
	Catalog *Catalog
}

var (
	ItemCantBeCheckoutErr = fmt.Errorf("The product is not longer in stock.")
)

func NewHandler() *MinPriceHandler {
	return &MinPriceHandler{
		Catalog: GetCatalog(),
	}
}

func (handler *MinPriceHandler) Call(line *CheckoutLine, subtotalChannel chan *Subtotal, errChannel chan error) {
	promotions := handler.Catalog.Promotions[line.ItemCode]
	if item, err := handler.Catalog.Get(line.ItemCode); err == nil {
		subtotalChannel <- handler.BestPriceFor(item, line, promotions)
	} else {
		errChannel <- ItemCantBeCheckoutErr
	}
}

func (handler *MinPriceHandler) BestPriceFor(item *Item, line *CheckoutLine, promotions []Promotion) *Subtotal {
	bestSubtotal := &Subtotal{
		Items: map[string]int{
			item.Code: line.Amount,
		},
		FinalPrice: item.Price * float64(line.Amount),
	}
	// Ensure that last-time changes in prices are synced.
	line.UnitPrice = item.Price

	for _, promotion := range promotions {
		if promo := promotion.Call(line); promo != nil && promo.FinalPrice < bestSubtotal.FinalPrice {
			bestSubtotal = promo
		}
	}

	return bestSubtotal
}
