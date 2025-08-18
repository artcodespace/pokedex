package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func commandMapb(config *config) error {
	if config.Previous == nil {
		fmt.Println("You're on the first page!")
		return nil
	} 

	baseUrl := *config.Previous

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

