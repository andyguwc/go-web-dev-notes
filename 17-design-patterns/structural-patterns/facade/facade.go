/* 
Facade pattern

The OpenWeatherMap API gives lots of information, so we are going to focus on getting live
weather data in one city in some geo-located place by using its latitude and longitude
values. The following are the requirements and acceptance criteria for this design pattern:
1. Provide a single type to access the data. All information retrieved from
OpenWeatherMap service will pass through it.
2. Create a way to get the weather data for some city of some country.
3. Create a way to get the weather data for some latitude and longitude position.
4. Only second and thrird point must be visible outside of the package; everything
else must be hidden (including all connection-related data).

weatherMap := CurrentWeatherData{*apiKey}
weather, err := weatherMap.GetByCityAndCountryCode("Madrid", "ES")
if err != nil {
	t.Fatal(err)
}

*/


package openWeatherMap

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// interface for retrieving data 
type CurrentWeatherDataRetriever interface {
	GetByGeoCoordinates(lat, lon float32) (*Weather, error)
	GetByCityAndCountryCode(city, countryCode string) (*Weather, error)
}


type CurrentWeatherData struct {
	APIkey string
}

type Weather struct {
	Coord struct {
		Lon float32 `json:"lon"`
		Lat float32 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		Id          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp     float32 `json:"temp"`
		Pressure float32 `json:"pressure"`
		Humidity float32 `json:"humidity"`
		TempMin  float32 `json:"temp_min"`
		TempMax  float32 `json:"temp_max"`
	} `json:"main"`
	Wind struct {
		Speed float32 `json:"speed"`
		Deg   float32 `json:"deg"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Rain struct {
		ThreeHours float32 `json:"3h"`
	} `json:"rain"`
	Dt  uint32 `json:"dt"`
	Sys struct {
		Type    int     `json:"type"`
		ID      int     `json:"id"`
		Message float32 `json:"message"`
		Country string  `json:"country"`
		Sunrise int     `json:"sunrise"`
		Sunset  int     `json:"sunset"`
	} `json:"sys"`
	ID   int    `json:"id"`
	Name string `json:"name"`
	Cod  int    `json:"cod"`
}

const (
	commonRequestPrefix              = "http://api.openweathermap.org/data/2.5/"
	weatherByCityName                = commonRequestPrefix + "weather?q=%s,%s&APPID=%s"
	weatherByGeographicalCoordinates = commonRequestPrefix + "weather?lat=%f&lon=%f&APPID=%s"
)

//GetByGeoCoordinates returns the current weather data by passing a geographical
//coordinates (latitude and longitude) in decimal notation. It returns weather
//information or a detailed error.
//For example, to query about Madrid, Spain you do:
//	currentWeather.GetByGeoCoordinates(-3, 40)
func (c *CurrentWeatherData) GetByGeoCoordinates(lat, lon float32) (weather *Weather, err error) {
	return c.doRequest(fmt.Sprintf(weatherByGeographicalCoordinates, lat, lon, c.APIkey))
}

//GetByCityAndCountryCode returns the current weather data by passing a city name
//and an ISO country code. It returns weather information or a detailed error
//For example, to query about Madrid, Spain you do:
//	currentWeather.GetByCityAndCountryCode("Madrid", "ES)
func (c *CurrentWeatherData) GetByCityAndCountryCode(city, countryCode string) (weather *Weather, err error) {
	return c.doRequest(fmt.Sprintf(weatherByCityName, city, countryCode, c.APIkey))
}

// parse the JSON response from the REST API
func (c *CurrentWeatherData) responseParser(body io.Reader) (*Weather, error) {
	w := new(Weather)
	err := json.NewDecoder(body).Decode(w)
	if err != nil {
		return nil, err
	}

	return w, nil
}

// We start by creating an http.Client class, which will make the requests. Then, we create a request object, which will use the GET
// method, as described in the OpenWeatherMap webpage, and the URI we passed. If we were to use a different method, or more than one, they would have to be brought about by
// arguments in the signature.

func (o *CurrentWeatherData) doRequest(uri string) (weather *Weather, err error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return
	}

	if resp.StatusCode != 200 {
		byt, errMsg := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		if errMsg == nil {
			errMsg = fmt.Errorf("%s", string(byt))
		}
		err = fmt.Errorf("Status code was %d, aborting. Error message was:\n%s\n",
			resp.StatusCode, errMsg)

		return
	}

	weather, err = o.responseParser(resp.Body)
	resp.Body.Close()

	return
}
