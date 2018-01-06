package main

import (
	"fmt"
	"time"
)

type Checkout struct {
	Catalog *Catalog
	Lines   []*CheckoutLine
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
	if checkout.Catalog.Contains(itemCode) {
		return fmt.Errorf("There's not a product with \"%v\" code.", itemCode)
	} else {
		if line := checkout.GetLineFor(itemCode); line != nil {
			line.Amount = line.Amount + 1
		} else {
			checkout.AddNewLine(itemCode, 1)
		}
	}
}

func (checkout *Checkout) GetTotal() (float64, error) {
	subtotalChannel := make(chan float64)
	errChannel := make(chan error)

	defer close(subtotalChannel)
	defer close(errChannel)

	for _, line := range checkout.Lines {
		go dispatch(line)
	}

	subtotals, err := checkout.CollectSubtotals(len(checkout.Lines), subtotalChannel, errChannel)
	if err == nil {
		var total float64
		for _, subtotal := range subtotals {
			total = total + subtotal
		}
		return total, nil
	} else {
		return -1, err
	}
}

func (checkout *Checkout) CollectSubtotals(amountOfSubtotals int, subtotalChannel chan float64, errChannel chan error) ([]float64, error) {
	subtotals := []float64{}
	for {
		select {
		case subtotal := <-subtotalChannel:
			subtotals = append(subtotals, subtotal)
			if len(subtotals) == amountOfSubtotals {
				return subtotals, nil
			}
		case err := <-errChannel:
			return []float64{}, err
		case <-time.After(time.Second * 2):
			return []float64{}, TimeoutErr
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

func (checkout *Checkout) AddNewLine(itemCode string, amount int, unitPrice float64) {
	newLine := &CheckoutLine{ItemCode: itemCode, Amount: amount}
	checkout.Lines = append(checkout.Lines, newLine)
}
