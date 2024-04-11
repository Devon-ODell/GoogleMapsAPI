// PlaceDetailsResult is an individual Places API Place Details result
package main


type PlaceDetailsResult struct {
	// AddressComponents is an array of separate address components used to compose a
	// given address.
	AddressComponents []AddressComponent `json:"address_components,omitempty"`
	// FormattedAddress is the human-readable address of this place.
	FormattedAddress string `json:"formatted_address,omitempty"`
	// AdrAddress is the address in the "adr" microformat.
	AdrAddress string `json:"adr_address,omitempty"`
	// BusinessStatus is a string indicating the operational status of the
	// place, if it is a business.
	BusinessStatus string `json:"business_status,omitempty"`
	// CurbsidePickup specifies if the business supports curbside pickup.
	CurbsidePickup bool `json:"curbside_pickup,omitempty"`
	// Delivery specifies if the business supports delivery.
	Delivery bool `json:"delivery,omitempty"`
	// DineIn specifies if the business supports seating options.
	DineIn bool `json:"dine_in,omitempty"`
	// EditorialSummary contains a summary of the place. A summary is comprised
	// of a textual overview, and also includes the language code for these if
	// applicable. Summary text must be presented as-is and can not be modified
	// or altered.
	EditorialSummary *PlaceEditorialSummary `json:"editorial_summary,omitempty"`
	// FormattedPhoneNumber contains the place's phone number in its local format. For
	// example, the formatted_phone_number for Google's Sydney, Australia office is
	// (02) 9374 4000.
	FormattedPhoneNumber string `json:"formatted_phone_number,omitempty"`
	// InternationalPhoneNumber contains the place's phone number in international
	// format. International format includes the country code, and is prefixed with the
	// plus (+) sign. For example, the international_phone_number for Google's Sydney,
	// Australia office is +61 2 9374 4000.
	InternationalPhoneNumber string `json:"international_phone_number,omitempty"`
	// Geometry contains geometry information about the result, generally including the
	// location (geocode) of the place and (optionally) the viewport identifying its
	// general area of coverage.
	Geometry AddressGeometry `json:"geometry,omitempty"`
	// Icon contains the URL of a recommended icon which may be displayed to the user
	// when indicating this result.
	Icon string `json:"icon,omitempty"`
	// Name contains the human-readable name for the returned result. For establishment
	// results, this is usually the business name.
	Name string `json:"name,omitempty"`
	// OpeningHours may contain whether the place is open now or not.
	OpeningHours *OpeningHours `json:"opening_hours,omitempty"`
	// CurrentOpeningHours may contain the hours of operation for the next seven
	// days (including today). The time period starts at midnight on the date of
	// the request and ends at 11:59 pm six days later. This field includes the
	// special_days subfield of all hours, set for dates that have exceptional
	// hours.
	CurrentOpeningHours *OpeningHours `json:"current_opening_hours,omitempty"`
	// SecondaryOpeningHours may contain an array of entries for the next seven
	// days including information about secondary hours of a business. Secondary
	// hours are different from a business's main hours. For example, a
	// restaurant can specify drive through hours or delivery hours as its
	// secondary hours. This field populates the type subfield, which draws from
	// a predefined list of opening hours types (such as DRIVE_THROUGH, PICKUP,
	// or TAKEOUT) based on the types of the place. This field includes the
	// special_days subfield of all hours, set for dates that have exceptional
	// hours.
	SecondaryOpeningHours []OpeningHours `json:"secondary_opening_hours,omitempty"`
	// PermanentlyClosed is a boolean flag indicating whether the place has permanently
	// shut down (value true). If the place is not permanently closed, the flag is
	// absent from the response.
	//
	// Deprecated: Use BusinessStatus instead.
	PermanentlyClosed bool `json:"permanently_closed,omitempty"`
	// Photos is an array of photo objects, each containing a reference to an image.
	Photos []Photo `json:"photos,omitempty"`
	// PlaceID is a textual identifier that uniquely identifies a place.
	PlaceID string `json:"place_id,omitempty"`
	// PriceLevel is the price level of the place, on a scale of 0 to 4.
	PriceLevel int `json:"price_level,omitempty"`
	// Rating contains the place's rating, from 1.0 to 5.0, based on aggregated user
	// reviews.
	Rating float32 `json:"rating,omitempty"`
	// Reservable specifies if the place supports reservations.
	Reservable bool `json:"reservable,omitempty"`
	// Reviews is an array of up to five reviews. If a language parameter was specified
	// in the Place Details request, the Places Service will bias the results to prefer
	// reviews written in that language.
	Reviews []PlaceReview `json:"reviews,omitempty"`
	// ServesBeer specifies if the place serves beer.
	ServesBeer bool `json:"serves_beer,omitempty"`
	// ServesBreakfast specifies if the place serves breakfast.
	ServesBreakfast bool `json:"serves_breakfast,omitempty"`
	// ServesBrunch specifies if the place serves brunch.
	ServesBrunch bool `json:"serves_brunch,omitempty"`
	// ServesDinner specifies if the place serves dinner.
	ServesDinner bool `json:"serves_dinner,omitempty"`
	// ServesLunch specifies if the place serves lunch.
	ServesLunch bool `json:"serves_lunch,omitempty"`
	// ServesVegetarianFood specifies if the place serves vegetarian food.
	ServesVegetarianFood bool `json:"serves_vegetarian_food,omitempty"`
	// ServesWine specifies if the place serves wine.
	ServesWine bool `json:"serves_wine,omitempty"`
	// Takeout specifies if the business supports takeout.
	Takeout bool `json:"takeout,omitempty"`
	// Types contains an array of feature types describing the given result.
	Types []string `json:"types,omitempty"`
	// URL contains the URL of the official Google page for this place. This will be the
	// establishment's Google+ page if the Google+ page exists, otherwise it will be the
	// Google-owned page that contains the best available information about the place.
	// Applications must link to or embed this page on any screen that shows detailed
	// results about the place to the user.
	URL string `json:"url,omitempty"`
	// UserRatingsTotal contains total number of the place's ratings
	UserRatingsTotal int `json:"user_ratings_total,omitempty"`
	// UTCOffset contains the number of minutes this place’s current timezone is offset
	// from UTC. For example, for places in Sydney, Australia during daylight saving
	// time this would be 660 (+11 hours from UTC), and for places in California outside
	// of daylight saving time this would be -480 (-8 hours from UTC).
	UTCOffset *int `json:"utc_offset,omitempty"`
	// Vicinity contains a feature name of a nearby location.
	Vicinity string `json:"vicinity,omitempty"`
	// Website lists the authoritative website for this place, such as a business'
	// homepage.
	Website string `json:"website,omitempty"`
	// WheelchairAccessibleEntrance specifies if the place has an entrance that
	// is wheelchair-accessible.
	WheelchairAccessibleEntrance bool `json:"wheelchair_accessible_entrance,omitempty"`
	// HTMLAttributions contain a set of attributions about this listing which must be
	// displayed to the user.
	HTMLAttributions []string `json:"html_attributions,omitempty"`

}


type TextSearchRequest struct {
	// Query is the text string on which to search, for example: "restaurant". The
	// Google Places service will return candidate matches based on this string and
	// order the results based on their perceived relevance.
	Query string
	// Location is the latitude/longitude around which to retrieve place information. If
	// you specify a location parameter, you must also specify a radius parameter.
	Location *LatLng
	// Radius defines the distance (in meters) within which to bias place results. The
	// maximum allowed radius is 50,000 meters. Results inside of this region will be
	// ranked higher than results outside of the search circle; however, prominent
	// results from outside of the search radius may be included.
	Radius uint
	// Language specifies the language in which to return results. Optional.
	Language string
	// MinPrice restricts results to only those places within the specified price level.
	// Valid values are in the range from 0 (most affordable) to 4 (most expensive),
	// inclusive.
	MinPrice PriceLevel
	// MaxPrice restricts results to only those places within the specified price level.
	// Valid values are in the range from 0 (most affordable) to 4 (most expensive),
	// inclusive.
	MaxPrice PriceLevel
	// OpenNow returns only those places that are open for business at the time the
	// query is sent. Places that do not specify opening hours in the Google Places
	// database will not be returned if you include this parameter in your query.
	OpenNow bool
	// Type restricts the results to places matching the specified type.
	Type PlaceType
	// PageToken returns the next 20 results from a previously run search. Setting a
	// PageToken parameter will execute a search with the same parameters used
	// previously — all parameters other than PageToken will be ignored.
	PageToken string
	// The region code, specified as a ccTLD (country code top-level domain) two-character
	// value. Most ccTLD codes are identical to ISO 3166-1 codes, with some exceptions.
	// This parameter will only influence, not fully restrict, search results. If more
	// relevant results exist outside of the specified region, they may be included. When
	// this parameter is used, the country name is omitted from the resulting formatted_address
	// for results in the specified region.
	Region string
}

// PlacesSearchResponse is the response to a Places API Search request.
type PlacesSearchResponse struct {
	// Results is the Place results for the search query
	Results []PlacesSearchResult
	// HTMLAttributions contain a set of attributions about this listing which must be
	// displayed to the user.
	HTMLAttributions []string
	// NextPageToken contains a token that can be used to return up to 20 additional
	// results.
	NextPageToken string
}



type PlacesSearchResult struct {
	// FormattedAddress is the human-readable address of this place
	FormattedAddress string `json:"formatted_address,omitempty"`
	// Geometry contains geometry information about the result, generally including the
	// location (geocode) of the place and (optionally) the viewport identifying its
	// general area of coverage.
	Name string `json:"name,omitempty"`
	// Icon contains the URL of a recommended icon which may be displayed to the user
	// when indicating this result.
	Types []string `json:"types,omitempty"
	// PermanentlyClosed is a boolean flag indicating whether the place has permanently
	// shut down.
	BusinessStatus string `json:"business_status,omitempty"`
}