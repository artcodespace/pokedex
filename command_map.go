package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type DataResult struct {
	Name string
	Url string
}

type Data struct {
	Count int
	Next *string
	Previous *string
	Results []DataResult
}

func commandMap(config *config) error {
	baseUrl := "https://pokeapi.co/api/v2/location-area"

	if config.Next != nil {
		baseUrl = *config.Next
	} 

	res, err := http.Get(baseUrl)

	if err != nil {
		return err
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		return fmt.Errorf("Response failed with status code %v", res.StatusCode)
	}
	if err != nil {
		return err
	}

	data := Data{}
	err = json.Unmarshal(body, &data) 
	if err != nil {
		return fmt.Errorf("Unable to marshal data from %v", body)
	}

	if data.Next != nil {
		config.Next = data.Next
	}

	if data.Previous != nil {
		config.Previous = data.Previous
	}
		
	for _, v := range(data.Results) {
		fmt.Println(v.Name)
	}

	return nil
}
