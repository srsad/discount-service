package main

import (
	parser "github.com/srsad/discount-service/services/parsers"
)

// go run github.com/srsad/discount-service/cmd/parsers.go
// go build -o ./tmp/parsers ./cmd/parsers.go

func main() {
	parser.Parse5Ka()
}
