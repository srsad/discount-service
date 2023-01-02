package parsers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
)

type FeatureCollection struct {
	Type     string    `json:"type"`
	Features []Feature `json:"features"`
}

type Feature struct {
	Properties Properties      `json:"properties"`
	Geometry   Geometry        `json:"geometry"`
	Sublease   []StoreSublease `json:"store_sublease"`
}

type Properties struct {
	CityID        int    `json:"city_id"`
	Photo         string `json:"photo"`
	OpeningDate   string `json:"opening_date"`
	Phone         string `json:"phone"`
	Address       string `json:"address"`
	CityName      string `json:"city_name"`
	WorkStartTime string `json:"work_start_time"`
	WorkEndTime   string `json:"work_end_time"`
	Is24h         bool   `json:"is_24h"`
	State         string `json:"state"`
	Type          string `json:"type"`
}

type Geometry struct {
	Type        string    `json:"type"`
	Coordinates []float32 `json:"coordinates"`
}

type StoreSublease struct {
	TypeIcon string `json:"type_icon"`
	TypeName string `json:"type_name"`
	TypeCode string `json:"type_code"`
}

type FiveAddressResponse struct {
	Data  FeatureCollection `json:"data"`
	Error string            `json:"error"`
}

func ParseAdresesX5() {
	link := "https://5ka.ru/api/v2/stores/?bbox=59.46215877074411,27.70227119726559,60.408970498503926,32.92901680273433"

	resp, err := http.Get(link)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	var responseJson string

	for true {
		bs := make([]byte, 1014)
		n, err := resp.Body.Read(bs)
		responseJson += string(bs[:n])
		if n == 0 || err != nil {
			break
		}
	}

	regex := regexp.MustCompile(`callback\((.*)\)`)
	jsonString := regex.FindStringSubmatch(responseJson)
	validJson := jsonString[1]

	targets := FiveAddressResponse{}
	_ = json.Unmarshal([]byte(validJson), &targets)

	// возвращать targets.Data.Features
	for _, value := range targets.Data.Features {
		fmt.Println("value", value)
	}
}
