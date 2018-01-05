package main

type Catalog struct {
	Products   map[string]*Item
	Promotions map[string]Promotion
}

type Item struct {
	Code   string
	Price  float64
	Hidden bool
}

var (
	instance *Catalog
	once     sync.Once
)

func GetCatalog() *Catalog {
	once.Do(func() {
		instance := &Catalog{
			Products:   make(map[string]*Item),
			Promotions: make(map[string]Promotion),
		}
		AddSeedItems(catalog)
		AddSeedPromotions(catalog)
	})
	return instance
}

func AddSeedItems(catalog *Catalog) {

}

func AddSeedPromotions(catalog *Catalog) {

}

func (catalog *Catalog) Add(code string, price Int) error {
}
func (catalog *Catalog) SetPrice(code String, price Int) error                   {}
func (catalog *Catalog) SetHidden(code String, hidden bool) error                {}
func (catalog *Catalog) Get(code String) error                                   {}
func (catalog *Catalog) AddPromotion(code String, promotion *Promotion) error    {}
func (catalog *Catalog) DeletePromotion(code String, promotion *Promotion) error {}
