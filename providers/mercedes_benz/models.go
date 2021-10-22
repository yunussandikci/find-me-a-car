package mercedes_benz

type Provider struct {
	MinPrice uint
	MaxPrice  uint
	Classes   []Class
	FuelTypes []FuelType
}

type Class string

const (
	AClass   Class = "A-CLASS"
	BClass   Class = "B-CLASS"
	CClass   Class = "C-CLASS"
	EClass   Class = "E-CLASS"
	ClaClass Class = "CLA-CLASS"
)

type FuelType string

const (
	Bensin FuelType = "BENSIN"
	Diesel FuelType = "DIESEL"
)

type Response struct {
	Type        string `json:"type"`
	Breadcrumbs []struct {
		FacetCode      string `json:"facetCode"`
		FacetName      string `json:"facetName"`
		FacetValueCode string `json:"facetValueCode"`
		FacetValueName string `json:"facetValueName"`
		RemoveQuery    struct {
			Query struct {
				Value string `json:"value"`
			} `json:"query"`
			URL string `json:"url"`
		} `json:"removeQuery"`
		TruncateQuery struct {
			Query struct {
				Value string `json:"value"`
			} `json:"query"`
			URL string `json:"url"`
		} `json:"truncateQuery,omitempty"`
	} `json:"breadcrumbs"`
	CurrentQuery struct {
		Query struct {
			Value string `json:"value"`
		} `json:"query"`
		URL string `json:"url"`
	} `json:"currentQuery"`
	Facets []struct {
		Category         bool   `json:"category"`
		Code             string `json:"code"`
		DisplayName      string `json:"displayName"`
		FacetDisplayType string `json:"facetDisplayType,omitempty"`
		MultiSelect      bool   `json:"multiSelect"`
		Name             string `json:"name"`
		Priority         int    `json:"priority"`
		Values           []struct {
			Code  string `json:"code"`
			Count int    `json:"count"`
			Name  string `json:"name"`
			Query struct {
				Query struct {
					Value string `json:"value"`
				} `json:"query"`
				URL string `json:"url"`
			} `json:"query"`
			Selected bool `json:"selected"`
		} `json:"values"`
		Visible   bool `json:"visible"`
		Container struct {
			Code string `json:"code"`
			Name string `json:"name"`
		} `json:"container,omitempty"`
		Max  float64 `json:"max,omitempty"`
		Min  float64 `json:"min,omitempty"`
		From float64 `json:"from,omitempty"`
		To   float64 `json:"to,omitempty"`
	} `json:"facets"`
	FreeTextSearch string `json:"freeTextSearch"`
	Pagination     struct {
		CurrentPage  int    `json:"currentPage"`
		PageSize     int    `json:"pageSize"`
		Sort         string `json:"sort"`
		TotalPages   int    `json:"totalPages"`
		TotalResults int    `json:"totalResults"`
	} `json:"pagination"`
	Products []struct {
		Baumuster     string `json:"baumuster"`
		CarEquipments []struct {
			Code              string `json:"code"`
			HighlightPriority int    `json:"highlightPriority"`
			Name              string `json:"name"`
			Pack              bool   `json:"pack"`
			ParentPack        string `json:"parentPack"`
		} `json:"carEquipments"`
		Categories []struct {
			Code        string `json:"code"`
			Description string `json:"description,omitempty"`
			Name        string `json:"name,omitempty"`
			Source      string `json:"source,omitempty"`
			URL         string `json:"url"`
		} `json:"categories"`
		Classifications []struct {
			Code     string `json:"code"`
			Features []struct {
				Code          string `json:"code"`
				Comparable    bool   `json:"comparable"`
				FeatureValues []struct {
					Value string `json:"value"`
				} `json:"featureValues"`
				Name       string `json:"name"`
				Range      bool   `json:"range"`
				Visibility string `json:"visibility"`
			} `json:"features"`
		} `json:"classifications"`
		Code           string `json:"code"`
		DcpProductType struct {
			Name string `json:"name"`
		} `json:"dcpProductType"`
		Description           string  `json:"description"`
		EngineSizeInCcm       int     `json:"engineSizeInCcm"`
		FirstRegistrationDate string  `json:"firstRegistrationDate"`
		GearBox               string  `json:"gearBox"`
		Gross                 float64 `json:"gross"`
		HasPreviousDamage     bool    `json:"hasPreviousDamage"`
		HasTrialVariant       bool    `json:"hasTrialVariant"`
		Images                []struct {
			AltText        string `json:"altText"`
			Format         string `json:"format"`
			GalleryIndex   int    `json:"galleryIndex"`
			Headline       string `json:"headline"`
			ImageType      string `json:"imageType"`
			SubHeadline    string `json:"subHeadline"`
			URL            string `json:"url"`
			YouTubeVideoID string `json:"youTubeVideoId"`
		} `json:"images"`
		IsMerchantIDConfigured          bool   `json:"isMerchantIdConfigured"`
		LicenseCode                     string `json:"licenseCode"`
		LocalEmissionAndConsumptionData string `json:"localEmissionAndConsumptionData"`
		Mileage                         struct {
			Unit  string  `json:"unit"`
			Value float64 `json:"value"`
		} `json:"mileage"`
		Model            string `json:"model"`
		ModelDesignation string `json:"modelDesignation"`
		ModelYear        string `json:"modelYear"`
		Name             string `json:"name"`
		NoiseLevel       struct {
		} `json:"noiseLevel"`
		PointOfServices []struct {
			Address struct {
				BillingAddress bool `json:"billingAddress"`
				Country        struct {
					Isocode string `json:"isocode"`
				} `json:"country"`
				DefaultAddress       bool     `json:"defaultAddress"`
				Email                string   `json:"email"`
				Fax                  string   `json:"fax"`
				FormattedAddress     string   `json:"formattedAddress"`
				FormattedAddress2    []string `json:"formattedAddress2"`
				ID                   string   `json:"id"`
				Line1                string   `json:"line1"`
				LivingSince          bool     `json:"livingSince"`
				Phone                string   `json:"phone"`
				PostalCode           string   `json:"postalCode"`
				ShippingAddress      bool     `json:"shippingAddress"`
				Town                 string   `json:"town"`
				VisibleInAddressBook bool     `json:"visibleInAddressBook"`
			} `json:"address"`
			BusinessName string        `json:"businessName"`
			CompanyID    string        `json:"companyId"`
			DisplayName  string        `json:"displayName"`
			Features     []interface{} `json:"features"`
			GeoPoint     struct {
				Latitude  float64 `json:"latitude"`
				Longitude float64 `json:"longitude"`
			} `json:"geoPoint"`
			LicenseCode       string        `json:"licenseCode"`
			Name              string        `json:"name"`
			StoreImages       []interface{} `json:"storeImages"`
			TranslatedAddress struct {
				BillingAddress bool `json:"billingAddress"`
				Country        struct {
					Isocode string `json:"isocode"`
				} `json:"country"`
				DefaultAddress       bool     `json:"defaultAddress"`
				Email                string   `json:"email"`
				Fax                  string   `json:"fax"`
				FormattedAddress     string   `json:"formattedAddress"`
				FormattedAddress2    []string `json:"formattedAddress2"`
				ID                   string   `json:"id"`
				Line1                string   `json:"line1"`
				LivingSince          bool     `json:"livingSince"`
				Phone                string   `json:"phone"`
				PostalCode           string   `json:"postalCode"`
				ShippingAddress      bool     `json:"shippingAddress"`
				Town                 string   `json:"town"`
				VisibleInAddressBook bool     `json:"visibleInAddressBook"`
			} `json:"translatedAddress"`
		} `json:"pointOfServices"`
		PowerInKw int `json:"powerInKw"`
		Price     struct {
			CurrencyIso    string  `json:"currencyIso"`
			FormattedValue string  `json:"formattedValue"`
			PriceType      string  `json:"priceType"`
			Value          float64 `json:"value"`
		} `json:"price"`
		PriceRange struct {
		} `json:"priceRange"`
		PriceWithPrefix string `json:"priceWithPrefix"`
		PrimaryImage    struct {
			Format    string `json:"format"`
			ImageType string `json:"imageType"`
			URL       string `json:"url"`
		} `json:"primaryImage"`
		RunFlat      bool   `json:"runFlat"`
		SnowFlake    bool   `json:"snowFlake"`
		SoftCompound bool   `json:"softCompound"`
		Source       string `json:"source"`
		Stock        struct {
			StockLevelStatus string `json:"stockLevelStatus"`
		} `json:"stock"`
		Studded          bool   `json:"studded"`
		URL              string `json:"url"`
		VatReclaimable   bool   `json:"vatReclaimable"`
		Vin              string `json:"vin"`
		VolumePricesFlag bool   `json:"volumePricesFlag"`
		Width            struct {
		} `json:"width"`
	} `json:"products"`
	Sorts []struct {
		Code     string `json:"code"`
		Name     string `json:"name"`
		Selected bool   `json:"selected"`
	} `json:"sorts"`
}