package main

import (
	"fmt"
	"net/http"
)

func getRespnonse(url string)(*http.Response, error){
	response, err := http.Get(url)
	if err != nil{
		return  nil, fmt.Errorf("http get reequest error: %w", err)
	}
	return response, nil
}