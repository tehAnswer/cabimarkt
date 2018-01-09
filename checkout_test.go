package main_test

import (
	"github.com/stretchr/testify/assert"
	main "github.com/tehAnswer/cabimarkt"
	"testing"
)

func TestEmptyCheckout(t *testing.T) {
	checkout := main.NewCheckout()
	total, err := checkout.GetTotal()
	if assert.NoError(t, err) {
		assert.Equal(t, total, 0.0)
	}
}

func TestSimpleOrderWithoutDiscounts(t *testing.T) {
	checkout := main.NewCheckout()
	assert.NoError(t, checkout.Scan("MUG"))
	assert.NoError(t, checkout.Scan("TSHIRT"))
	assert.NoError(t, checkout.Scan("VOUCHER"))
	total, err := checkout.GetTotal()
	if assert.NoError(t, err) {
		assert.Equal(t, total, 32.50)
	}
}

func Test2For1InVouchers(t *testing.T) {
	checkout := main.NewCheckout()
	assert.NoError(t, checkout.Scan("VOUCHER"))
	total, _ := checkout.GetTotal()
	assert.Equal(t, total, 5.00)

	assert.NoError(t, checkout.Scan("VOUCHER"))
	total, _ = checkout.GetTotal()
	assert.Equal(t, total, 5.00)

	assert.NoError(t, checkout.Scan("VOUCHER"))
	total, _ = checkout.GetTotal()
	assert.Equal(t, total, 10.00)

	assert.NoError(t, checkout.Scan("VOUCHER"))
	total, _ = checkout.GetTotal()
	assert.Equal(t, total, 10.00)

	assert.NoError(t, checkout.Scan("VOUCHER"))
	total, _ = checkout.GetTotal()
	assert.Equal(t, total, 15.00)
}

func TestBulkPriceForTShirts(t *testing.T) {
	checkout := main.NewCheckout()

	assert.NoError(t, checkout.Scan("TSHIRT"))
	total, _ := checkout.GetTotal()
	assert.Equal(t, total, 20.00)

	assert.NoError(t, checkout.Scan("TSHIRT"))
	total, _ = checkout.GetTotal()
	assert.Equal(t, total, 40.00)

	assert.NoError(t, checkout.Scan("TSHIRT"))
	total, _ = checkout.GetTotal()
	assert.Equal(t, total, 3*19.00)

	assert.NoError(t, checkout.Scan("TSHIRT"))
	total, _ = checkout.GetTotal()
	assert.Equal(t, total, 4*19.00)
}

func TestMultipleDiscounts(t *testing.T) {
	checkout := main.NewCheckout()
	assert.NoError(t, checkout.Scan("VOUCHER"))
	assert.NoError(t, checkout.Scan("VOUCHER"))
	assert.NoError(t, checkout.Scan("TSHIRT"))
	assert.NoError(t, checkout.Scan("TSHIRT"))
	assert.NoError(t, checkout.Scan("TSHIRT"))

	total, _ := checkout.GetTotal()

	assert.Equal(t, total, 3*19.00+5.00)
}

func TestIncorrectProduct(t *testing.T) {
	checkout := main.NewCheckout()
	assert.Error(t, checkout.Scan("BATTLETOADS"))
}

func TestProductGetsDeletedAfterScan(t *testing.T) {
	checkout := main.NewCheckout()
	assert.NoError(t, checkout.Catalog.Add("TONG", 0.75))
	assert.NoError(t, checkout.Scan("TONG"))
	assert.NoError(t, checkout.Catalog.Delete("TONG"))
	_, err := checkout.GetTotal()
	assert.Error(t, err)
}
