package recommendations

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
)

func (recommendationsApi *recommendationsApiImpl) GetRecommendations(request *web.GetRecommendationsRequest) (*web.GetRecommendationsResponse, error) {
  recommendationsApiEndpoint := recommendationsApi.RecommendationsApiConfig.ApiEndpoint

  jsonRequest, err := json.Marshal(request)
  if err != nil {
    return nil, err
  }
  
  response, err := http.Post(recommendationsApiEndpoint, "application/json", bytes.NewBuffer(jsonRequest))
  if err != nil {
    return nil, err
  }
  defer response.Body.Close()

  // Parse the response and extract recommendations
  if response.StatusCode != http.StatusOK {
    return nil, fmt.Errorf("unexpected status code: %d", response.StatusCode)
  }

  // Read the response body
  responseBody, err := ioutil.ReadAll(response.Body)
  if err != nil {
    return nil, err
  }

  // Check for invalid characters in the response
  if !json.Valid(responseBody) {
    return nil, fmt.Errorf("invalid character found in JSON response")
  }

  var recommendationsResponse web.GetRecommendationsResponse
  err = json.Unmarshal(responseBody, &recommendationsResponse)
  if err != nil {
    return nil, err
  }

  return &recommendationsResponse, nil
}
