package dpfm_api_output_formatter

type SDC struct {
	ConnectionKey       string      `json:"connection_key"`
	RedisKey            string      `json:"redis_key"`
	Filepath            string      `json:"filepath"`
	APIStatusCode       int         `json:"api_status_code"`
	RuntimeSessionID    string      `json:"runtime_session_id"`
	BusinessPartnerID   *int        `json:"business_partner"`
	ServiceLabel        string      `json:"service_label"`
	APIType             string      `json:"api_type"`
	Message             interface{} `json:"message"`
	APISchema           string      `json:"api_schema"`
	Accepter            []string    `json:"accepter"`
	Deleted             bool        `json:"deleted"`
	SQLUpdateResult     *bool       `json:"sql_update_result"`
	SQLUpdateError      string      `json:"sql_update_error"`
	SubfuncResult       *bool       `json:"subfunc_result"`
	SubfuncError        string      `json:"subfunc_error"`
	ExconfResult        *bool       `json:"exconf_result"`
	ExconfError         string      `json:"exconf_error"`
	APIProcessingResult *bool       `json:"api_processing_result"`
	APIProcessingError  string      `json:"api_processing_error"`
}

type GooglePlacesLocationsResponseBody struct {
	Results						[]Results	`json:"results"`
	Status						string		`json:"status"`
}

type Results struct {
	PlaceID						string			`json:"place_id"`
	Description					string			`json:"name"`
	IconURI						string			`json:"icon"`
	IconBackgroundColorInHex	string			`json:"icon_background_color"`
	IconMaskBaseURI				string			`json:"icon_mask_base_uri"`
	Rating						float32			`json:"rating"`
	UserRatingsTotal			float32			`json:"user_ratings_total"`
	Vicinity					string			`json:"vicinity"`
	BusinessStatus				string			`json:"business_status"`
	Photos						[]Photos		`json:"photos"`
	Geometry					[]Geometry		`json:"geometry"`
	Types						[]Types			`json:"types"`
	OpeningHours				OpeningHours	`json:"opening_hours"`
}

type Photos struct {
	Height				int					`json:"height"`
	Width				int					`json:"width"`
	PhotoReference		string				`json:"photo_reference"`
	HtmlAttributions	[]HtmlAttributions	`json:"html_attributions"`
}

type Geometry struct {
	Location		Location	`json:"location"`
}

type Types struct {
}

type OpeningHours struct {
	OpenNow			bool	`json:"open_now"`
}

type Location struct {
	Latitude		string	`json:"lat"`
	Longitude		string	`json:"lng"`
}

type HtmlAttributions struct {
}

type Message struct {
	GooglePlacesLocationsResponse *[]GooglePlacesLocationsResponse `json:"GooglePlacesLocationsResponse"`
}

type GooglePlacesLocationsResponse struct {
	PlaceID						string			`json:"PlaceID"`
	Description					string			`json:"Description"`
	IconURI						string			`json:"IconURI"`
	IconBackgroundColorInHex	string			`json:"IconBackgroundColorInHex"`
	IconMaskBaseURI				string			`json:"IconMaskBaseURI"`
	Rating						float32			`json:"Rating"`
	UserRatingsTotal			float32			`json:"UserRatingsTotal"`
	Vicinity					string			`json:"Vicinity"`
	BusinessStatus				string			`json:"BusinessStatus"`
	Photos						[]Photos		`json:"photos"`
	Geometry					[]Geometry		`json:"geometry"`
	Types						[]Types			`json:"types"`
	OpeningHours				OpeningHours	`json:"opening_hours"`
}
