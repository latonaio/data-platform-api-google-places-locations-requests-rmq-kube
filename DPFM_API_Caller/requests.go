package dpfm_api_caller

import (
	dpfm_api_input_reader "data-platform-api-google-places-locations-requests-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-google-places-locations-requests-rmq-kube/DPFM_API_Output_Formatter"
	"data-platform-api-google-places-locations-requests-rmq-kube/config"
	"encoding/json"
	"fmt"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"golang.org/x/xerrors"
	"io/ioutil"
	"net/http"
)

func (c *DPFMAPICaller) GooglePlacesLocations(
	input *dpfm_api_input_reader.SDC,
	errs *[]error,
	log *logger.Logger,
	conf *config.Conf,
) *[]dpfm_api_output_formatter.GooglePlacesLocationsResponse {
	var googlePlacesLocations []dpfm_api_output_formatter.GooglePlacesLocationsResponse

	key := input.GooglePlacesLocations.Key
	language := input.GooglePlacesLocations.Language
	location := input.GooglePlacesLocations.Location
	radius := input.GooglePlacesLocations.Radius

	req, err := http.NewRequest("GET", key, language, location, radius, nil)

	if err != nil {
		*errs = append(*errs, xerrors.Errorf("NewRequest error: %d", err))
		return nil
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", key))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		*errs = append(*errs, xerrors.Errorf("Google Places Locations request error: %d", err))
		return nil
	}
	defer resp.Body.Close()

	placesLocationsBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		*errs = append(*errs, xerrors.Errorf("User info request response read error: %d", err))
		return nil
	}

	var response map[string]interface{}
	err = json.Unmarshal(placesLocationsBody, &response)
	if err != nil {
		*errs = append(*errs, xerrors.Errorf("Response response error: %d", err))
		return nil
	}

	errorObj, ok := response["error"].(map[string]interface{})
	if ok {
		code, ok := errorObj["code"].(float64)
		if ok {
			errMsg, _ := errorObj["message"].(string)
			*errs = append(*errs, xerrors.Errorf("Status code error: %v %v", code, errMsg))
			return nil
		}
	}

	var googlePlacesLocationsResponseBody dpfm_api_output_formatter.GooglePlacesLocationsResponseBody
	err = json.Unmarshal(placesLocationsBody, &googlePlacesLocationsResponseBody)
	if err != nil {
		*errs = append(*errs, xerrors.Errorf("User info request response unmarshal error: %d", err))
		return nil
	}

	placesLocations := dpfm_api_output_formatter.ConvertToGooglePlacesLocationsRequestsFromResponse(googlePlacesLocationsResponseBody)

	googlePlacesLocations = append(
		googlePlacesLocations,
		placesLocations,
	)

	return &googlePlacesLocations
}
