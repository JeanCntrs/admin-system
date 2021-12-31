package utils

import (
	"errors"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
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
func SendRequest(method, url string) string {
	request, requestErr := http.NewRequest(method, url, nil)
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

	return string(bytes)
}

var funcsMap = template.FuncMap{"Welcome": Welcome}
var allTemplates = template.Must(template.New("T").Funcs(funcsMap).ParseGlob("./html/**/*.html"))
var errTemplate = template.Must(template.ParseFiles("./html/error/error.html"))

// RenderTemplate generates templates with optional data
func RenderTemplate(w http.ResponseWriter, pageName string, data interface{}) {
	w.Header().Set("Content-Type", "text/html")

	err := allTemplates.ExecuteTemplate(w, pageName, data)
	if err != nil {
		w.WriteHeader(500)
		errTemplate.Execute(w, nil)
	}
}

func Welcome(name string) string {
	return "Welcome to the page " + name
}

func RequiredField(value, name string) error {
	if value == "" {
		return errors.New(name + " field must be mandatory")
	}

	return nil
}

func MaxLength(value, name string, maxLength int) error {
	if len(value) > maxLength {
		return errors.New(name + " field exceeds maximum length " + strconv.Itoa(maxLength))
	}

	return nil

}

func MinLength(value, name string, minLength int) error {
	if len(value) < minLength {
		return errors.New(name + " field exceeds minimum length " + strconv.Itoa(minLength))
	}

	return nil
}
