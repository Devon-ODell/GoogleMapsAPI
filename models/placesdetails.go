package models

type PlaceDetailsResult struct {
	// AddressComponents is an array of separate address components used to compose a
	// given address.
	AddressComponents []AddressComponent `json:"address_components,omitempty"`
	// FormattedAddress is the human-readable address of this place.
	FormattedAddress string `json:"formatted_address,omitempty"`
	// AdrAddress is the address in the "adr" microformat.
	BusinessStatus string `json:"business_status,omitempty"`
	// CurbsidePickup specifies if the business supports curbside pickup.
	Geometry []AddressGeometry `json:"geometry,omitempty"`
	// Icon contains the URL of a recommended icon which may be displayed to the user
	// when indicating this result.
	Name string `json:"name,omitempty"`
	//PermanentlyClosed is a boolean to determine if/if not permanently closed.
	PermanentlyClosed bool `json:"permanently_closed,omitempty"`
	// PriceLevel is the price level of the place, on a scale of 0 to 4.
	PriceLevel int `json:"price_level,omitempty"`
	// Types contains an array of feature types describing the given result.
	Types []string `json:"types,omitempty"`
	// HTMLAttributions contain a set of attributions about this listing which must be
	// displayed to the user.
	HTMLAttributions []string `json:"html_attributions,omitempty"`

	// //Rating contains the place's rating, from 1.0 to 5.0, based on aggregated user
	// //reviews.
	// Rating float32 `json:"rating,omitempty"`
	// //Reservable specifies if the place supports reservations.
	// Reservable bool `json:"reservable,omitempty"`
	// //ServesVegetarianFood specifies if the place serves vegetarian food.
	// ServesVegetarianFood bool `json:"serves_vegetarian_food,omitempty"`

}

type AddressComponent struct {
	streetNumber   float32
	streetName     string
	City           string
	State          string
	addressZipCode int
}
