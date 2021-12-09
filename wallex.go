package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

type (
	JsonObject = map[string]interface{}
	Market     = map[string]Symbol
)

var (
	timeout = 10 * time.Second

	ErrNotSuccessful = errors.New("operation was not successful")
)

func GetMarkets() (map[string]Market, error) {
	url := "https://wallex.ir/api/v2/markets"

	client := &http.Client{Timeout: timeout}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var r JsonObject
	err = json.Unmarshal(body, &r)
	if err != nil {
		return nil, err
	}
	if r["message"] != "The operation was successful" {
		return nil, ErrNotSuccessful
	}
	symbols := r["result"].(JsonObject)["symbols"].(JsonObject)
	ret := make(map[string]Market)
	for _, value := range symbols {
		b, err := json.Marshal(value)
		if err != nil {
			return nil, err
		}
		var s Symbol
		err = json.Unmarshal(b, &s)
		if err != nil {
			return nil, err
		}
		_, ok := ret[s.QuoteAsset]
		if !ok {
			ret[s.QuoteAsset] = make(Market)
		}
		ret[s.QuoteAsset][s.BaseAsset] = s
	}
	return ret, nil
}
