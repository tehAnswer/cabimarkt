package main

import (
	"fmt"
	"sync"
)

type Catalog struct {
	Products   map[string]*Item
	Promotions map[string][]Promotion
}

type Item struct {
	Code   string
	Price  float64
	Hidden bool
}

var (
	instance                 *Catalog
	once                     sync.Once
	ItemDoesNotExistErr      = fmt.Errorf("The item does not exist.")
	ItemExistsErr            = fmt.Errorf("The item already exist.")
	PromotionDoesNotExistErr = fmt.Errorf("The promotion does not exist.")
	PromotionExistsErr       = fmt.Errorf("The promotion already exist for such item.")
)

func GetCatalog() *Catalog {
	once.Do(func() {
		instance = &Catalog{
			Products:   make(map[string]*Item),
			Promotions: make(map[string][]Promotion),
		}
		AddSeedItems(instance)
		AddSeedPromotions(instance)
	})
	return instance
}

func AddSeedItems(catalog *Catalog) {
	catalog.Add("MUG", 7.5)
	catalog.Add("VOUCHER", 5.00)
	catalog.Add("TSHIRT", 20.00)
}

func AddSeedPromotions(catalog *Catalog) {
	twoForOneInVouchers := &BuyGetPromotion{
		ChargableAmount: 1,
		RequiredAmount:  2,
		ItemCode:        "VOUCHER",
		// 33658-09-27 03:46:39 +0200
		ExpiresAt: 999999999999,
		ID:        "241",
	}
	catalog.AddPromotion(twoForOneInVouchers)

	bulkTShirtPrice := &BulkPromotion{
		BulkPrice:      19.00,
		RequiredAmount: 3,
		ItemCode:       "TSHIRT",
		ExpiresAt:      999999999999,
		ID:             "DAMN",
	}

	catalog.AddPromotion(bulkTShirtPrice)
}

func (catalog *Catalog) Add(code string, price float64) error {
	if _, err := catalog.Get(code); err != nil {
		catalog.Products[code] = &Item{
			Code:   code,
			Price:  price,
			Hidden: false,
		}
		catalog.Promotions[code] = []Promotion{}
		return nil
	} else {
		return ItemExistsErr
	}
}

func (catalog *Catalog) Delete(code string) error {
	if _, err := catalog.Get(code); err == nil {
		delete(catalog.Promotions, code)
		delete(catalog.Products, code)
		return nil
	} else {
		return err
	}
}

func (catalog *Catalog) SetPrice(code string, price float64) error {
	if _, err := catalog.Get(code); err == nil {
		catalog.Products[code].Price = price
		return nil
	} else {
		return err
	}
}

func (catalog *Catalog) SetHidden(code string, hidden bool) error {
	if _, err := catalog.Get(code); err == nil {
		catalog.Products[code].Hidden = hidden
		return nil
	} else {
		return err
	}
}

func (catalog *Catalog) Get(code string) (*Item, error) {
	if item := catalog.Products[code]; item != nil {
		return item, nil
	} else {
		return nil, ItemDoesNotExistErr
	}
}
func (catalog *Catalog) AddPromotion(promotion Promotion) error {
	if _, err := catalog.Get(promotion.GetItemCode()); err == nil {
		if catalog.HasPromotion(promotion.GetItemCode(), promotion.GetRef()) {
			return PromotionExistsErr
		} else {
			promotions := catalog.Promotions[promotion.GetItemCode()]
			catalog.Promotions[promotion.GetItemCode()] = append(promotions, promotion)
			return nil
		}
	} else {
		return err
	}
}

func (catalog *Catalog) HasPromotion(itemCode string, promotionCode string) bool {
	for _, promotion := range catalog.Promotions[itemCode] {
		if promotion.GetRef() == promotionCode {
			return true
		}
	}
	return false
}

func (catalog *Catalog) DeletePromotion(itemCode string, promotionCode string) error {
	for i, promotion := range catalog.Promotions[itemCode] {
		if promotion.GetRef() == promotionCode {
			catalog.Promotions[itemCode][i] = catalog.Promotions[itemCode][len(catalog.Promotions[itemCode])-1]
			catalog.Promotions[itemCode] = catalog.Promotions[itemCode][:len(catalog.Promotions[itemCode])-1]
			return nil
		}
	}
	return PromotionDoesNotExistErr
}

func (catalog *Catalog) Contains(itemCode string) bool {
	_, err := catalog.Get(itemCode)
	return err != nil
}
