package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// https://api.freecurrencyapi.com/v1/latest?base_currency=USD&apikey=fca_live_bccACPRd97j9nzlBtCx5SRJ2GDEWRJvDfUU3xNTH
const baseURL = "https://api.freecurrencyapi.com/v1/latest"

func GetLastWithBase(baseCurrency string, sourceCurrency string) (rate float64, err error) {
	//vars := mux.Vars(r)
	//base := strings.ToUpper(vars["base"])

	sourceCurrency = strings.ToUpper(sourceCurrency)

	if baseCurrency == "" {
		baseCurrency = os.Getenv("BASE_CURRENCY")
	} else {
		baseCurrency = strings.ToUpper(baseCurrency)
	}

	log.Println("Getting last currency data with base", baseCurrency)
	params := url.Values{}
	params.Add("base_currency", baseCurrency)
	params.Add("apikey", os.Getenv("CURRENCY_API_KEY"))
	urlWithParams := fmt.Sprintf("%s?%s", baseURL, params.Encode())
	req, err := http.NewRequest("GET", urlWithParams, nil)
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	// Create an HTTP client and perform the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error making HTTP request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	log.Println(string(responseBody))

	var result map[string]interface{}
	err = json.Unmarshal([]byte(responseBody), &result)
	if err != nil {
		log.Fatalf("Error parsing JSON: %v", err)
	}
	data, ok := result["data"].(map[string]interface{})
	if !ok {
		log.Fatal("Error casting to map[string]interface{}")
	}

	conValue, ok := data[sourceCurrency].(float64)
	if !ok {
		log.Fatal("Error casting Value to float64")
	}
	log.Printf("Value: %f\n", conValue)
	return conValue, nil

}
