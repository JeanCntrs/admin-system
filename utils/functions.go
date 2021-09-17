package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// GenerateURL generates a url dynamically
func GenerateURL(uri, host, protocol string, urlParams map[string]string) string {
	url, _ := url.Parse(uri)
	url.Host = host
	url.Scheme = protocol
	mapFunction := url.Query()

	for key, value := range urlParams {
		mapFunction.Add(key, value)
	}

	url.RawQuery = mapFunction.Encode()

	return url.String()
}

// SendRequest send a request dynamically
func SendRequest() {
	url := GenerateURL("/products", "localhost:8000", "http", nil)
	request, requestErr := http.NewRequest("GET", url, nil)
	if requestErr != nil {
		panic("There was a problem with the request")
	}

	client := &http.Client{}
	response, responseErr := client.Do(request)
	if responseErr != nil {
		fmt.Println("responseErr", responseErr)
		panic("There was a problem with the client")
	}

	bytes, readAllErr := ioutil.ReadAll(response.Body)
	if readAllErr != nil {
		panic("There was a problem reading the request body")
	}

	fmt.Println("Status:", response.Status)
	fmt.Println("Body:", string(bytes))
}
