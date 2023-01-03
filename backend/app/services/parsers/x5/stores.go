package parsers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type FiveResultCurrentPrices struct {
	PriceRegMin    float32 `json:"price_reg__min"`
	Price_promoMin float32 `json:"price_promo__min"`
}

type FiveResultPromo struct {
	Id          uint   `json:"id"`
	DateBegin   string `json:"bate_begin"`
	DateEnd     string `json:"date_end"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Kind        string `json:"kind"`
	ExpiredAt   uint   `json:"expired_at"`
}

type FiveResultResponse struct {
	Id            uint                    `json:"id"`
	Name          string                  `json:"name"`
	ImgLink       string                  `json:"img_link"`
	Plu           int                     `json:"plu"`
	Promo         FiveResultPromo         `json:"promo"`
	CurrentPrices FiveResultCurrentPrices `json:"current_prices"`
	StoreName     string                  `json:"store_name"`
}

type FiveResponse struct {
	Next         string               `json:"next"`
	Previous     string               `json:"previous"`
	Results      []FiveResultResponse `json:"results"`
	StoreAddress string               `json:"store_address"` // адрес магаина
}

/**
 * Парсинг данных пятерочки
 * TODO:
 * 	- разобраться с параметрами по умолчанию
 *  - найти список магазинов (store) и парсить все
 *  - проверить возможность удалять по FiveResponse.Results.Promo.Date_end
 *  - или по FiveResponse.Results.Promo.Expired_at
 */
func ParseStoreX5(args ...string) {
	link := "https://5ka.ru/api/v2/special_offers/?store=31Z6&records_per_page=15&page=1"

	if len(args) > 0 {
		link = args[0]
	}

	resp, err := http.Get(link)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	var responseJson string

	// хз насколько это плохо?!
	for true {
		bs := make([]byte, 1014)
		n, err := resp.Body.Read(bs)
		responseJson += string(bs[:n])
		if n == 0 || err != nil {
			break
		}
	}

	targets := FiveResponse{}
	_ = json.Unmarshal([]byte(responseJson), &targets)

	for _, value := range targets.Results {
		fmt.Println("value", value)
	}

	if targets.Next != "" {
		ParseStoreX5(targets.Next)
	}
}
