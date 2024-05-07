package api

import(
	"context"
	"json"
	"fmt"
	"net/https"
) 

func main(){
	
	apiKey, zipCode := readConfig("requirements/config.txt")
	client, err := NewClient(WithApiKey(api_key.txt)
	if err != nil {
		fmt.Println(err)
		return
		}

	// Defines Request Parameters
	req := &TextSearchRequest {
		Query : "hotels"
		Location: &LatLng{
			Lat: x,   // Set these variables
			Lng: y,   // Set these variables
		},
		Radius: 5000  // in meters
		Type: PlaceTypeHotel,
		Region: "us",
		PageToken: "",
	}

	resp, err := client.TextSearch(context.Background(), req)
	if err != nil {
		fmt.Println ("Error performing test search:", err)
	}

	for _, result := range resp.Results {
		fmt.Println(result.name, result.FormattedAddress)
	}





}

func (c *Client) TextSearch(ctx context.Context, r *TextSearchRequest) (PlacesSearchResponse, error) {

	if r.Query == "" && r.PageToken == "" && r.Type == "" {
		return PlacesSearchResponse{}, errors.New("maps: Query, PageToken and Type are all missing")
	}

	if r.Location != nil && r.Radius == 0 {
		return PlacesSearchResponse{}, errors.New("maps: Radius missing, required with Location")
	}

	var response struct {
		Results          []PlacesSearchResult `json:"results,omitempty"`
		HTMLAttributions []string             `json:"html_attributions,omitempty"`
		NextPageToken    string               `json:"next_page_token,omitempty"`
		commonResponse
	}

	if err := c.getJSON(ctx, placesTextSearchAPI, r, &response); err != nil {
		return PlacesSearchResponse{}, err
	}

	if err := response.StatusError(); err != nil {
		return PlacesSearchResponse{}, err
	}

	return PlacesSearchResponse{response.Results, response.HTMLAttributions, response.NextPageToken}, nil
}


func (r *TextSearchRequest) params() url.Values {
	q := make(url.Values)

	q.Set("query", r.Query)

	if r.Location != nil {
		q.Set("location", r.Location.String())
	}

	if r.Radius != 0 {
		q.Set("radius", fmt.Sprint(r.Radius))
	}

	if r.Language != "" {
		q.Set("language", r.Language)
	}

	if r.MinPrice != "" {
		q.Set("minprice", string(r.MinPrice))
	}

	if r.MaxPrice != "" {
		q.Set("maxprice", string(r.MaxPrice))
	}

	if r.OpenNow {
		q.Set("opennow", "true")
	}

	if r.Type != "" {
		q.Set("type", string(r.Type))
	}

	if r.PageToken != "" {
		q.Set("pagetoken", r.PageToken)
	}

	if r.Region != "" {
		q.Set("region", r.Region)
	}

	return q
}




var placeDetailsAPI = &apiConfig{
	host:            "https://maps.googleapis.com",
	path:            "/maps/api/place/details/json",
	acceptsClientID: true,
}




func readConfig(config_filepath, zipcode_filepath string) (string, string, error) {
	// Read the contents of the config file(TOTALLY NOT THE API KEY!!) new line cmd at the end
	configData, err := os.ReadFile(config_filepath)
	if err != nil {
		return "", "", fmt.Errorf("Error reading API Key file:", err)
	}
	config := strings.TrimSpace(string(configData), "\n")

	// Read the contents of the zipcode file
	zipcodeData, err := os.ReadFile(zipcode_filepath)
	if err != nil {
		return "", "", fmt.Errorf("Error reading API Key file:", err)
	}
	return apiKey, zipCode, nil

}