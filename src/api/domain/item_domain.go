package domain

import "time"

// Item es una estructura para mapear items de Mercado Libre
type Item struct {
	ID                string      `json:"id"`
	SiteID            string      `json:"site_id"`
	Title             string      `json:"title"`
	Subtitle          interface{} `json:"subtitle"`
	SellerID          int         `json:"seller_id"`
	CategoryID        string      `json:"category_id"`
	OfficialStoreID   interface{} `json:"official_store_id"`
	Price             int         `json:"price"`
	BasePrice         int         `json:"base_price"`
	OriginalPrice     int         `json:"original_price"`
	CurrencyID        string      `json:"currency_id"`
	InitialQuantity   int         `json:"initial_quantity"`
	AvailableQuantity int         `json:"available_quantity"`
	SoldQuantity      int         `json:"sold_quantity"`
	SaleTerms         []struct {
		ID          string      `json:"id"`
		Name        string      `json:"name"`
		ValueID     string      `json:"value_id"`
		ValueName   string      `json:"value_name"`
		ValueStruct interface{} `json:"value_struct"`
	} `json:"sale_terms"`
	BuyingMode      string    `json:"buying_mode"`
	ListingTypeID   string    `json:"listing_type_id"`
	StartTime       time.Time `json:"start_time"`
	StopTime        time.Time `json:"stop_time"`
	Condition       string    `json:"condition"`
	Permalink       string    `json:"permalink"`
	Thumbnail       string    `json:"thumbnail"`
	SecureThumbnail string    `json:"secure_thumbnail"`
	Pictures        []struct {
		ID        string `json:"id"`
		URL       string `json:"url"`
		SecureURL string `json:"secure_url"`
		Size      string `json:"size"`
		MaxSize   string `json:"max_size"`
		Quality   string `json:"quality"`
	} `json:"pictures"`
	VideoID      interface{} `json:"video_id"`
	Descriptions []struct {
		ID string `json:"id"`
	} `json:"descriptions"`
	AcceptsMercadopago           bool          `json:"accepts_mercadopago"`
	NonMercadoPagoPaymentMethods []interface{} `json:"non_mercado_pago_payment_methods"`
	Shipping                     struct {
		Mode        string `json:"mode"`
		FreeMethods []struct {
			ID   int `json:"id"`
			Rule struct {
				Default          bool        `json:"default"`
				FreeMode         string      `json:"free_mode"`
				FreeShippingFlag bool        `json:"free_shipping_flag"`
				Value            interface{} `json:"value"`
			} `json:"rule"`
		} `json:"free_methods"`
		Tags         []string    `json:"tags"`
		Dimensions   interface{} `json:"dimensions"`
		LocalPickUp  bool        `json:"local_pick_up"`
		FreeShipping bool        `json:"free_shipping"`
		LogisticType string      `json:"logistic_type"`
		StorePickUp  bool        `json:"store_pick_up"`
	} `json:"shipping"`
	InternationalDeliveryMode string `json:"international_delivery_mode"`
	SellerAddress             struct {
		City struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"city"`
		State struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"state"`
		Country struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"country"`
		SearchLocation struct {
			Neighborhood struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			} `json:"neighborhood"`
			City struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			} `json:"city"`
			State struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			} `json:"state"`
		} `json:"search_location"`
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
		ID        int     `json:"id"`
	} `json:"seller_address"`
	SellerContact interface{} `json:"seller_contact"`
	Location      struct {
	} `json:"location"`
	Geolocation struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	} `json:"geolocation"`
	CoverageAreas []interface{} `json:"coverage_areas"`
	Attributes    []struct {
		ID                 string      `json:"id"`
		Name               string      `json:"name"`
		ValueID            string      `json:"value_id"`
		ValueName          string      `json:"value_name"`
		ValueStruct        interface{} `json:"value_struct"`
		AttributeGroupID   string      `json:"attribute_group_id"`
		AttributeGroupName string      `json:"attribute_group_name"`
	} `json:"attributes"`
	Warnings            []interface{} `json:"warnings"`
	ListingSource       string        `json:"listing_source"`
	Variations          []interface{} `json:"variations"`
	Status              string        `json:"status"`
	SubStatus           []interface{} `json:"sub_status"`
	Tags                []string      `json:"tags"`
	Warranty            string        `json:"warranty"`
	CatalogProductID    string        `json:"catalog_product_id"`
	DomainID            string        `json:"domain_id"`
	ParentItemID        interface{}   `json:"parent_item_id"`
	DifferentialPricing interface{}   `json:"differential_pricing"`
	DealIds             []string      `json:"deal_ids"`
	AutomaticRelist     bool          `json:"automatic_relist"`
	DateCreated         time.Time     `json:"date_created"`
	LastUpdated         time.Time     `json:"last_updated"`
	Health              float64       `json:"health"`
}
