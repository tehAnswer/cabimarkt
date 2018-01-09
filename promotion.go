package main

import (
	"time"
)

type Promotion interface {
	Active() bool
	// Items that are present of the promotion
	GetItemCode() string
	GetRef() string
	Call(line *CheckoutLine) *Subtotal
}

type BuyGetPromotion struct {
	ChargableAmount int
	RequiredAmount  int
	ExpiresAt       int64
	ItemCode        string
	ID              string
}

func (promotion *BuyGetPromotion) IsValid(line *CheckoutLine) bool {
	return promotion.Active() && promotion.MeetsRequirements(line)
}

func (promotion *BuyGetPromotion) Active() bool {
	return promotion.ExpiresAt > time.Now().Unix()
}

func (promotion *BuyGetPromotion) MeetsRequirements(line *CheckoutLine) bool {
	return line.Amount >= promotion.RequiredAmount && promotion.ItemCode == line.ItemCode
}

func (promotion *BuyGetPromotion) GetItemCode() string { return promotion.ItemCode }
func (promotion *BuyGetPromotion) GetRef() string      { return promotion.ID }

func (promotion *BuyGetPromotion) Call(line *CheckoutLine) *Subtotal {
	if promotion.IsValid(line) {
		// You can use this promotion more than one time (e.g 2x1 is 4x2 two times).
		promotionItems := float64(line.Amount / promotion.RequiredAmount * promotion.ChargableAmount)
		remainingItems := float64(line.Amount % promotion.RequiredAmount)
		finalPrice := promotionItems*line.UnitPrice + remainingItems*line.UnitPrice
		return &Subtotal{
			Items: map[string]int{
				line.ItemCode: line.Amount,
			},
			PromotionRef: promotion.ID,
			FinalPrice:   finalPrice,
		}
	} else {
		return nil
	}
}

type BulkPromotion struct {
	RequiredAmount int
	BulkPrice      float64
	ExpiresAt      int64
	ItemCode       string
	ID             string
}

func (promotion *BulkPromotion) IsValid(line *CheckoutLine) bool {
	return promotion.Active() && promotion.MeetsRequirements(line)
}

func (promotion *BulkPromotion) Active() bool {
	return promotion.ExpiresAt > time.Now().Unix()
}

func (promotion *BulkPromotion) MeetsRequirements(line *CheckoutLine) bool {
	return line.Amount >= promotion.RequiredAmount && promotion.ItemCode == line.ItemCode
}

func (promotion *BulkPromotion) GetItemCode() string { return promotion.ItemCode }
func (promotion *BulkPromotion) GetRef() string      { return promotion.ID }

func (promotion *BulkPromotion) Call(line *CheckoutLine) *Subtotal {
	if promotion.IsValid(line) {
		return &Subtotal{
			Items: map[string]int{
				line.ItemCode: line.Amount,
			},
			PromotionRef: promotion.ID,
			FinalPrice:   float64(line.Amount) * promotion.BulkPrice,
		}
	} else {
		return nil
	}
}

type DiscountPromotion struct {
	DiscountPrice float64
	ExpiresAt     int64
	ItemCode      string
	ID            string
}

func (promotion *DiscountPromotion) IsValid(line *CheckoutLine) bool {
	return promotion.Active() && promotion.MeetsRequirements(line)
}

func (promotion *DiscountPromotion) Active() bool {
	return promotion.ExpiresAt > time.Now().Unix()
}

func (promotion *DiscountPromotion) MeetsRequirements(line *CheckoutLine) bool {
	return promotion.ItemCode == line.ItemCode
}

func (promotion *DiscountPromotion) GetItemCode() string { return promotion.ItemCode }
func (promotion *DiscountPromotion) GetRef() string      { return promotion.ID }

func (promotion *DiscountPromotion) Call(line *CheckoutLine) *Subtotal {
	if promotion.IsValid(line) && promotion.Active() {
		return &Subtotal{
			Items: map[string]int{
				line.ItemCode: line.Amount,
			},
			PromotionRef: promotion.ID,
			FinalPrice:   float64(line.Amount) * promotion.DiscountPrice,
		}
	} else {
		return nil
	}
}
