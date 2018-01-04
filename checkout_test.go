package main_test

import (
	"github.com/stretchr/testify/assert"
	main "github.com/tehAnswer/cabimarkt"
	"testing"
)

func TestEmptyCheckout(t *testing.T) {
	checkout := main.NewCheckout()
	assert.Equal(t, checkout.GetTotal(), 0.0)
}

func TestSimpleOrderWithoutDiscounts(t *testing.T) {
	checkout := main.NewCheckout()
	assert.NoError(t, checkout.Scan("MUG"))
	assert.NoError(t, checkout.Scan("TSHIRT"))
	assert.NoError(t, checkout.Scan("VOUCHER"))
	assert.Equal(t, checkout.GetTotal(), 32.50)
}

func Test2For1InVouchers(t *testing.T) {
	checkout := main.NewCheckout()
	assert.NoError(t, checkout.Scan("VOUCHER"))
	assert.Equal(t, checkout.GetTotal(), 5.00)
	assert.NoError(t, checkout.Scan("VOUCHER"))
	assert.Equal(t, checkout.GetTotal(), 5.00)
	assert.NoError(t, checkout.Scan("VOUCHER"))
	assert.Equal(t, checkout.GetTotal(), 10.00)
	assert.NoError(t, checkout.Scan("VOUCHER"))
	assert.Equal(t, checkout.GetTotal(), 10.00)
	assert.NoError(t, checkout.Scan("VOUCHER"))
	assert.Equal(t, checkout.GetTotal(), 15.00)
}

func TestBulkPriceForTShirts(t *testing.T) {
	checkout := main.NewCheckout()
	assert.NoError(t, checkout.Scan("TSHIRT"))
	assert.Equal(t, checkout.GetTotal(), 20.00)
	assert.NoError(t, checkout.Scan("TSHIRT"))
	assert.Equal(t, checkout.GetTotal(), 40.00)
	assert.NoError(t, checkout.Scan("TSHIRT"))
	assert.Equal(t, checkout.GetTotal(), 3*19.00)
	assert.NoError(t, checkout.Scan("TSHIRT"))
	assert.Equal(t, checkout.GetTotal(), 4*19.00)
}

func TestMultipleDiscounts(t *testing.T) {
	checkout := main.NewCheckout()
	assert.NoError(t, checkout.Scan("VOUCHER"))
	assert.NoError(t, checkout.Scan("VOUCHER"))
	assert.NoError(t, checkout.Scan("TSHIRT"))
	assert.NoError(t, checkout.Scan("TSHIRT"))
	assert.NoError(t, checkout.Scan("TSHIRT"))
	assert.Equal(t, checkout.GetTotal(), 3*19.00+5.00)
}

func TestIncorrectProduct(t *testing.T) {
	checkout := main.NewCheckout()
	assert.Error(t, checkout.Scan("BATTLETOADS"))
}
