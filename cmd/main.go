package main

import (
	"flag"
	"log"
	"os"

	"github.com/RafaelGermann/currency-converter-cli/internal"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error loading .env file: %v", err)
	}
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatal("API_KEY is not set in .env file")
	}

	from := flag.String("from", "USD", "currency to convert from")
	to := flag.String("to", "BRL", "currency to convert to")
	amount := flag.Float64("amount", 1.0, "amount to convert")
	flag.Parse()

	result, err := internal.ConvertCurrency(*from, *to, *amount, apiKey)
	if err != nil {
		log.Fatalf("error converting currency: %v", err)
	}
	log.Printf("Convert result: %f %s = %f %s", *amount, *from, result, *to)
}
