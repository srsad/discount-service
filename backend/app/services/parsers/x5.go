package parser

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type FiveResultCurrentPrices struct {
	Price_reg__min 		float32 `json:price_reg__min`
	price_promo__min	float32 `json:price_promo__min`
}

type FiveResultPromo struct {
	Id							uint		`json:id`
	Date_begin			string	`json:bate_begin`
	Date_end				string	`json:date_end`
	Type						string 	`json:type`
	Description			string 	`json:description`
	Kind						string 	`json:kind`
	Expired_at			uint		`json:expired_at`
}

type FiveResultResponse struct {
	Id							uint										`json.id`
	Name 						string									`json:name`
	Img_link 				string									`json:img_link`
	Plu							int 										`json:plu`
	Promo						FiveResultPromo					`json:promo`
	Current_prices 	FiveResultCurrentPrices	`json:current_prices`
	Store_name			string									`json:store_name`
}

type FiveResponse struct {
	Next						string								`json.next`
	Previous				string 								`json.previous`
	Results					[]FiveResultResponse 	`json.results`
	Store_address 	string								`json.store_address`
}

/**
 * TODO:
 * 	- разобраться с параметрами по умолчанию
 *  - найти список магазинов (store) и парсить все
 *  - проверить возможность удалять по FiveResponse.Results.Promo.Date_end
 *  - или по FiveResponse.Results.Promo.Expired_at
 */
func Parse5Ka(args ...string) {
	link := "https://5ka.ru/api/v2/special_offers/?store=31Z6&records_per_page=15&page=1"

   if len(args) > 0 {
		link = args[0]
	}

	fmt.Println("link", link)
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
		Parse5Ka(targets.Next)
	}
}
