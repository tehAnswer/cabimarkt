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
		&main.CheckoutLine{
			Amount:    7,
			UnitPrice: 100.00,
			ItemCode:  "ZAR",
		},
	}
)

func TestBuyGetPromotion(t *testing.T) {
	promotion := main.BuyGetPromotion{
		ChargableAmount: 2,
		RequiredAmount:  3,
		ExpiresAt:       time.Now().Unix() + 86400,
		ItemCode:        "ZAR",
		ID:              "PAR",
	}

	assert.Nil(t, promotion.Call(lines[0]))
	// Valid promotion: 3 for the same price of 2.
	assert.Equal(t, promotion.Call(lines[1]).FinalPrice, 200.00)
	assert.Equal(t, promotion.Call(lines[1]).PromotionRef, promotion.ID)
	// Not valid promotion, returns regular price. (WRONG ITEM)
	assert.Nil(t, promotion.Call(lines[2]))
	assert.Equal(t, promotion.Call(lines[3]).FinalPrice, 500.00)
	assert.Equal(t, promotion.Call(lines[3]).PromotionRef, promotion.ID)
}

func TestBulkPromotion(t *testing.T) {
	promotion := main.BulkPromotion{
		RequiredAmount: 3,
		BulkPrice:      10.0,
		ExpiresAt:      time.Now().Unix() + 86400,
		ItemCode:       "ZAR",
		ID:             "DIN",
	}

	// Promoton is not valid: insufficient amount
	assert.Nil(t, promotion.Call(lines[0]))
	// Valid promotion!
	assert.Equal(t, promotion.Call(lines[1]).FinalPrice, 30.00)
	assert.Equal(t, promotion.Call(lines[1]).PromotionRef, promotion.ID)
	// Not valid promotion, returns regular price. (WRONG ITEM)
	assert.Nil(t, promotion.Call(lines[2]))
}

func TestDiscountPromotion(t *testing.T) {
	promotion := main.DiscountPromotion{
		DiscountPrice: 10.0,
		ExpiresAt:     time.Now().Unix() + 86400,
		ItemCode:      "ZAR",
		ID:            "AAA",
	}

	assert.Equal(t, promotion.Call(lines[0]).FinalPrice, 20.00)
	assert.Equal(t, promotion.Call(lines[0]).PromotionRef, promotion.ID)
	// Valid promotion!
	assert.Equal(t, promotion.Call(lines[1]).FinalPrice, 30.00)
	assert.Equal(t, promotion.Call(lines[1]).PromotionRef, promotion.ID)
	// Not valid promotion, returns regular price. (WRONG ITEM)
	assert.Nil(t, promotion.Call(lines[2]))
}

func TestExpiredPromotion(t *testing.T) {
	promotion := main.BulkPromotion{
		RequiredAmount: 3,
		BulkPrice:      10.0,
		ExpiresAt:      time.Now().Unix() - 86400,
		ItemCode:       "ZAR",
		ID:             "BAR",
	}

	assert.Nil(t, promotion.Call(lines[0]))
	assert.Nil(t, promotion.Call(lines[1]))
	assert.Nil(t, promotion.Call(lines[2]))
}
