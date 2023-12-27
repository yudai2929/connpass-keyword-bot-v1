package external

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/yudai2929/connpass-keyword-bot-v1/pkg/config"
	"github.com/yudai2929/connpass-keyword-bot-v1/pkg/domain/entity"
	"github.com/yudai2929/connpass-keyword-bot-v1/pkg/domain/repository"
	"github.com/yudai2929/connpass-keyword-bot-v1/pkg/domain/valueobject"
	"github.com/yudai2929/connpass-keyword-bot-v1/pkg/errors"
)

type LocationRepositoryImpl struct {
	BaseURL  string
	ClientID string
	Client   *http.Client
}

func NewLocationRepository() repository.LocationRepository {
	return &LocationRepositoryImpl{
		BaseURL:  "https://map.yahooapis.jp",
		ClientID: config.Env.YahooClientID,
		Client:   http.DefaultClient,
	}
}

func (repo *LocationRepositoryImpl) SearchByCoordinate(coordinate valueobject.Coordinate) (entity.Location, error) {
	url := repo.BaseURL + "/geoapi/V1/reverseGeoCoder?lat=" + strconv.FormatFloat(coordinate.Latitude, 'f', -1, 64) + "&lon=" + strconv.FormatFloat(coordinate.Longitude, 'f', -1, 64) + "&appid=" + repo.ClientID + "&output=json"
	resp, err := repo.Client.Get(url)
	if err != nil {
		return entity.Location{}, errors.Wrap(err, "failed to get location")
	}
	defer resp.Body.Close()

	response := LocationResponse{}

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return entity.Location{}, errors.Wrap(err, "failed to decode response")
	}

	if response.ResultInfo.Count == 0 {
		return entity.Location{}, errors.New("location not found")
	}

	address := convertAddress(response)

	return entity.NewLocation(coordinate, address), nil

}

func convertAddress(response LocationResponse) valueobject.Address {
	future := response.Feature[0]
	prefecture := future.Property.AddressElement[0].Name
	city := future.Property.AddressElement[1].Name
	return valueobject.NewAddress(prefecture, city)

}

type LocationResponse struct {
	ResultInfo struct {
		Count        int     `json:"Count"`
		Total        int     `json:"Total"`
		Start        int     `json:"Start"`
		Latency      float64 `json:"Latency"`
		Status       int     `json:"Status"`
		Description  string  `json:"Description"`
		Copyright    string  `json:"Copyright"`
		CompressType string  `json:"CompressType"`
	} `json:"ResultInfo"`
	Feature []struct {
		Geometry struct {
			Type        string `json:"Type"`
			Coordinates string `json:"Coordinates"`
		} `json:"Geometry"`
		Property struct {
			Country struct {
				Code string `json:"Code"`
				Name string `json:"Name"`
			} `json:"Country"`
			Address        string `json:"Address"`
			AddressElement []struct {
				Name  string `json:"Name"`
				Kana  string `json:"Kana"`
				Level string `json:"Level"`
				Code  string `json:"Code,omitempty"`
			} `json:"AddressElement"`
			Building []struct {
				ID    string `json:"Id"`
				Name  string `json:"Name"`
				Floor string `json:"Floor"`
				Area  string `json:"Area"`
			} `json:"Building"`
		} `json:"Property"`
	} `json:"Feature"`
}
