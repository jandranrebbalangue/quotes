package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"model"
	"net/http"
	"os"
)

func main() {
	getRandom()
}

// getRandom quotes
func getRandom() {
	response, err := http.Get("https://api.quotable.io/random/")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	responseData,
		err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Print(err)
	}
	var responseObject model.Quotes
	json.Unmarshal(responseData, &responseObject)
	fmt.Println(responseObject.Content)
}
