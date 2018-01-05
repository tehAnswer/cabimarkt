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
		assert.Equal(t, item.Price(), 10.00)
		assert.Equal(t, item.Code(), "PIZZA")
		assert.False(t, item.IsHidden())
		assert.Equal(t, len(item.Promotions), 0)
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
	if assert.NoError(err) {
		assert.True(t, item.IsHidden())
		if assert.NoError(t, catalog.SetHidden("BAR", false)) {
			assert.False(t, item.IsHidden())
		}
	}
}

func TestUpdatePriceItemInCatalog(t *testing.T) {
	catalog := main.GetCatalog()
	assert.NoError(t, catalog.Add("TAP", 1.0))
	assert.NoError(t, catalog.SetPrice("TAP", 2.0))
	item, err := catalog.Get("TAP")
	if assert.NoError(err) {
		assert.Equal(t, item.Price(), 2.0)
	}
}

func TestAddPromotion(t *testing.T) {
	catalog := main.GetCatalog()
	assert.NoError(t, catalog.Add("LAP", 100.0))
	promotion := &main.BuyGetPromotion{Code: "DANS", BuyAmount: 3, GetAmount: 5, Expires: 1351700038}
	assert.NoError(t, catalog.AddPromotion("LAP", promotion))
	assert.Error(t, catalog.AddPromotion("LAP", promotion))
	item, err := catalog.Get("LAP")
	if assert.NoError(err) {
		assert.Equal(t, len(item.Promotions), 1)
	}
}

func TestDeletePromotionBuyGet(t *testing.T) {
	catalog := main.GetCatalog()
	assert.NoError(t, catalog.Add("BAT", 100.0))
	promotion := &main.BuyGetPromotion{Code: "MAN", BuyAmount: 3, getAmount: 5, expires: 1351700038}
	assert.NoError(t, catalog.AddPromotion("BAT", promotion))
	assert.NoError(t, catalog.DeletePromotion("BAT", "MAN"))
	assert.Error(t, catalog.DeletePromotion("BAT", "MAN"))
}

func TestCatalogIsASingleton(t *testing.T) {
	catalog := main.GetCatalog()
	otherCatalog := main.GetCatalog()
	assert.Equal(t, &catalog, &otherCatalog)
}
