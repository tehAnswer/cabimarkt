package main

import (
	"fmt"
	"time"
)

type Checkout struct {
	Catalog       *Catalog
	Lines         []*CheckoutLine
	PromotionRefs []string
}

type CheckoutLine struct {
	ItemCode  string
	Amount    int
	UnitPrice float64
}

var (
	TimeoutErr = fmt.Errorf("The operation timed out. Please try again.")
)

func NewCheckout() *Checkout {
	return &Checkout{
		// Mock of a database connection, so to speak
		Catalog: GetCatalog(),
	}
}

func (checkout *Checkout) Scan(itemCode string) error {
	if _, err := checkout.Catalog.Get(itemCode); err != nil {
		return err
	} else {
		if line := checkout.GetLineFor(itemCode); line != nil {
			line.Amount = line.Amount + 1
		} else {
			checkout.AddNewLine(itemCode, 1)
		}
	}

	return nil
}

func (checkout *Checkout) GetTotal() (float64, error) {
	if len(checkout.Lines) == 0 {
		return 0.0, nil
	}

	checkout.PromotionRefs = []string{}
	subtotalChannel := make(chan *Subtotal)
	errChannel := make(chan error)

	defer close(subtotalChannel)
	defer close(errChannel)

	for _, line := range checkout.Lines {
		go Dispatch(line, subtotalChannel, errChannel)
	}

	subtotals, err := checkout.CollectSubtotals(len(checkout.Lines), subtotalChannel, errChannel)
	if err == nil {
		var total float64
		for _, subtotal := range subtotals {
			total = total + subtotal.FinalPrice
			if subtotal.PromotionRef != "" {
				checkout.PromotionRefs = append(checkout.PromotionRefs, subtotal.PromotionRef)
			}
		}
		return total, nil
	} else {
		return -1, err
	}
}

func (checkout *Checkout) CollectSubtotals(amountOfSubtotals int, subtotalChannel chan *Subtotal, errChannel chan error) ([]*Subtotal, error) {
	subtotals := []*Subtotal{}
	for {
		select {
		case subtotal := <-subtotalChannel:
			subtotals = append(subtotals, subtotal)
			if len(subtotals) == amountOfSubtotals {
				return subtotals, nil
			}
		case err := <-errChannel:
			return []*Subtotal{}, err
		case <-time.After(time.Second * 2):
			return []*Subtotal{}, TimeoutErr
		}
	}
}

func (checkout *Checkout) GetLineFor(itemCode string) *CheckoutLine {
	for _, line := range checkout.Lines {
		if line.ItemCode == itemCode {
			return line
		}
	}
	return nil
}

func (checkout *Checkout) AddNewLine(itemCode string, amount int) {
	newLine := &CheckoutLine{ItemCode: itemCode, Amount: amount}
	checkout.Lines = append(checkout.Lines, newLine)
}

func Dispatch(line *CheckoutLine, subtotalChannel chan *Subtotal, errChannel chan error) {
	NewHandler().Call(line, subtotalChannel, errChannel)
}
