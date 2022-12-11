package main

import (
	"github.com/srsad/discount-service/internal/app"
	parser "github.com/srsad/discount-service/services/parsers"
)

func main() {
  // TODO разнести по разным сервисам - все парсеры запускать кроном
  // для парсинга установить в true
	parseProduct := false
	if (parseProduct) {
    parser.Parse5Ka()
	}

	app.Run()
}
