package main_test

import (
	"github.com/stretchr/testify/assert"
	main "github.com/tehAnswer/cabimarkt"
	"testing"
	"time"
)

var (
	lines = []*main.CheckoutLine{
		&main.CheckoutLine{
			Amount:    2,
			UnitPrice: 100.00,
			ItemCode:  "ZAR",
		},
		&main.CheckoutLine{
			Amount:    3,
			UnitPrice: 100.00,
			ItemCode:  "ZAR",
		},
		&main.CheckoutLine{
			Amount:    2,
			UnitPrice: 1000.00,
			ItemCode:  "BAR",
		},
	}
)

func TestBuyGetPromotion(t *testing.T) {
	promotion := main.BuyGetPromotion{
		ChargableAmount: 2,
		RequiredAmount:  3,
		ExpiresAt:       time.Now().Unix() + 86400,
		ItemCode:        "ZAR",
	}

	assert.Equal(t, promotion.FinalPrice(lines[0]), 200.00)
	// Valid promotion: 3 for the same price of 2.
	assert.Equal(t, promotion.FinalPrice(lines[1]), 200.00)
	// Not valid promotion, returns regular price. (WRONG ITEM)
	assert.Equal(t, promotion.FinalPrice(lines[2]), 2000.00)
}

func TestBulkPromotion(t *testing.T) {
	promotion := main.BulkPromotion{
		RequiredAmount: 3,
		BulkPrice:      10.0,
		ExpiresAt:      time.Now().Unix() + 86400,
		ItemCode:       "ZAR",
	}

	// Promoton is not valid: insufficient amount, returns
	// regular price.
	assert.Equal(t, promotion.FinalPrice(lines[0]), 200.00)
	// Valid promotion!
	assert.Equal(t, promotion.FinalPrice(lines[1]), 30.00)
	// Not valid promotion, returns regular price. (WRONG ITEM)
	assert.Equal(t, promotion.FinalPrice(lines[2]), 2000.00)
}

func TestDiscountPromotion(t *testing.T) {
	promotion := main.BulkPromotion{
		RequiredAmount: 3,
		BulkPrice:      10.0,
		ExpiresAt:      time.Now().Unix() + 86400,
		ItemCode:       "ZAR",
	}

	// Promoton is not valid: insufficient amount, returns
	// regular price.
	assert.Equal(t, promotion.FinalPrice(lines[0]), 200.00)
	// Valid promotion!
	assert.Equal(t, promotion.FinalPrice(lines[1]), 30.00)
	// Not valid promotion, returns regular price. (WRONG ITEM)
	assert.Equal(t, promotion.FinalPrice(lines[2]), 2000.00)
}

func TestExpiredPromotion(t *testing.T) {
	promotion := main.BulkPromotion{
		RequiredAmount: 3,
		BulkPrice:      10.0,
		ExpiresAt:      time.Now().Unix() - 86400,
		ItemCode:       "ZAR",
	}

	assert.Equal(t, promotion.FinalPrice(lines[0]), 200.00)
	assert.Equal(t, promotion.FinalPrice(lines[1]), 200.00)
	assert.Equal(t, promotion.FinalPrice(lines[2]), 2000.00)
}
