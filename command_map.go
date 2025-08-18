package main

import (
	"fmt"
	"io"
	"net/http"
)

func commandMap(config *config) error {
	res, err := http.Get("https://pokeapi.co/api/v2/location-area")
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

	fmt.Printf("%s", body)
	return nil
}
