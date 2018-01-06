package main

import (
	"time"
)

type Promotion interface {
	Active() bool
	// Items that are present of the promotion
	ItemCode() string

	FinalPrice(line *CheckoutLine) int
}

type BuyGetPromotion struct {
	ChargableAmount int
	RequiredAmount  int
	ExpiresAt       int64
	ItemCode        string
}

func (promotion *BuyGetPromotion) IsValid(line *CheckoutLine) bool {
	return promotion.Active() && promotion.MeetsRequirements(line)
}

func (promotion *BuyGetPromotion) Active() bool {
	return promotion.ExpiresAt < time.Now().Unix()
}

func (promotion *BuyGetPromotion) MeetsRequirements(line *CheckoutLine) bool {
	return promotion.RequiredAmount == line.Amount && promotion.ItemCode == line.ItemCode
}

func (promotion *BuyGetPromotion) GetItemCode() string { return promotion.ItemCode }

func (promotion *BuyGetPromotion) FinalPrice(line *CheckoutLine) float64 {
	if promotion.IsValid(line) {
		return float64(promotion.ChargableAmount) * line.UnitPrice
	} else {
		return float64(line.Amount) * line.UnitPrice
	}
}

type BulkPromotion struct {
	RequiredAmount int
	BulkPrice      float64
	ExpiresAt      int64
	ItemCode       string
}

func (promotion *BulkPromotion) IsValid(line *CheckoutLine) bool {
	return promotion.Active() && promotion.MeetsRequirements(line)
}

func (promotion *BulkPromotion) Active() bool {
	return promotion.ExpiresAt < time.Now().Unix()
}

func (promotion *BulkPromotion) MeetsRequirements(line *CheckoutLine) bool {
	return promotion.RequiredAmount >= line.Amount && promotion.ItemCode == line.ItemCode
}

func (promotion *BulkPromotion) GetItemCode() string { return promotion.ItemCode }

func (promotion *BulkPromotion) FinalPrice(line *CheckoutLine) float64 {
	if promotion.IsValid(line) && promotion.Active() {
		return float64(line.Amount) * promotion.BulkPrice
	} else {
		return float64(line.Amount) * line.UnitPrice
	}
}

type DiscountPromotion struct {
	DiscountPrice float64
	ExpiresAt     int64
	ItemCode      string
}

func (promotion *DiscountPromotion) IsValid(line *CheckoutLine) bool {
	return promotion.Active() && promotion.MeetsRequirements(line)
}

func (promotion *DiscountPromotion) Active() bool {
	return promotion.ExpiresAt < time.Now().Unix()
}

func (promotion *DiscountPromotion) MeetsRequirements(line *CheckoutLine) bool {
	return promotion.ItemCode == line.ItemCode
}

func (promotion *DiscountPromotion) GetItemCode() string { return promotion.ItemCode }

func (promotion *DiscountPromotion) FinalPrice(line *CheckoutLine) float64 {
	if promotion.IsValid(line) && promotion.Active() {
		return float64(line.Amount) * promotion.DiscountPrice
	} else {
		return float64(line.Amount) * line.UnitPrice
	}
}
