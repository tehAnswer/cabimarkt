package main_test

import (
	"github.com/stretchr/testify/assert"
	main "github.com/tehAnswer/cabimarkt"
	"testing"
)

func TestAddItemToCatalog(t *testing.T) {
	catalog := main.GetCatalog()
	assert.NoError(t, catalog.Add("PIZZA", 10.00))
	assert.Error(t, catalog.Add("PIZZA", 10.00))
	item, err := catalog.Get("PIZZA")
	if assert.NoError(t, err) {
		assert.Equal(t, item.Price, 10.00)
		assert.Equal(t, item.Code, "PIZZA")
		assert.False(t, item.Hidden)
		assert.Equal(t, len(catalog.Promotions[item.Code]), 0)
	}
}

func TestDeleteItemFromCatalog(t *testing.T) {
	catalog := main.GetCatalog()
	assert.NoError(t, catalog.Add("FOO", 1.0))
	assert.NoError(t, catalog.Delete("FOO"))
	assert.Error(t, catalog.Delete("FOO"))
	item, err := catalog.Get("FOO")
	if assert.Error(t, err) {
		assert.Nil(t, item)
	}
}

func TestUpdateVisibilityItemInCatalog(t *testing.T) {
	catalog := main.GetCatalog()
	assert.NoError(t, catalog.Add("BAR", 1.0))
	assert.NoError(t, catalog.SetHidden("BAR", true))
	item, err := catalog.Get("BAR")
	if assert.NoError(t, err) {
		assert.True(t, item.Hidden)
		if assert.NoError(t, catalog.SetHidden("BAR", false)) {
			assert.False(t, item.Hidden)
		}
	}
}

func TestUpdatePriceItemInCatalog(t *testing.T) {
	catalog := main.GetCatalog()
	assert.NoError(t, catalog.Add("TAP", 1.0))
	assert.NoError(t, catalog.SetPrice("TAP", 2.0))
	item, err := catalog.Get("TAP")
	if assert.NoError(t, err) {
		assert.Equal(t, item.Price, 2.0)
	}
}

func TestAddPromotion(t *testing.T) {
	catalog := main.GetCatalog()
	assert.NoError(t, catalog.Add("LAP", 100.0))
	promotion := &main.BuyGetPromotion{ItemCode: "LAP", ID: "DANS", ChargableAmount: 3, RequiredAmount: 5, ExpiresAt: 1351700038}
	assert.NoError(t, catalog.AddPromotion(promotion))
	assert.Error(t, catalog.AddPromotion(promotion))
	assert.Equal(t, len(catalog.Promotions["LAP"]), 1)
}

func TestDeletePromotionBuyGet(t *testing.T) {
	catalog := main.GetCatalog()
	assert.NoError(t, catalog.Add("BAT", 100.0))
	promotion := &main.BuyGetPromotion{ItemCode: "BAT", ID: "MAN", ChargableAmount: 3, RequiredAmount: 5, ExpiresAt: 1351700038}
	assert.NoError(t, catalog.AddPromotion(promotion))
	assert.NoError(t, catalog.DeletePromotion("BAT", "MAN"))
	assert.Error(t, catalog.DeletePromotion("BAT", "MAN"))
}

func TestCatalogIsASingleton(t *testing.T) {
	catalog := main.GetCatalog()
	otherCatalog := main.GetCatalog()
	assert.Equal(t, &catalog, &otherCatalog)
}
