package main

import (
	"fmt"
	"net/http"
	// "io"
	"encoding/json"
	"io/ioutil"
	"strings"
)

type Samplestr struct {
	Name, Text string
}

func main() {
	url := "https://jsonplaceholder.typicode.com/posts/1/comments"
	result := fetchData(url)
	fmt.Println(result)
	parseUsnStruct(result)
}

func fetchData(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	// defer resp.Body.Close()
	bodyBytes, errObj := ioutil.ReadAll(resp.Body)
	if errObj != nil {
		panic(errObj)
	}

	return string(bodyBytes)
}

func parseUsnStruct(content string) []Samplestr {
	var repJson = make([]Samplestr, 0)
	var samObj Samplestr
	cd := `[ {"Name": "Ed", "Text": "Knock knock."}, {"Name": "Sam", "Text": "Who's there?"}, {"Name": "Ed", "Text": "Go fmt."}, {"Name": "Sam", "Text": "Go fmt who?"}, {"Name": "Ed", "Text": "Go fmt yourself!"} ]`
	decoder := json.NewDecoder(strings.NewReader(cd))
	decoder1, err2 := decoder.Token()
	fmt.Println("\n\n\n\n", decoder1)
	checkErrors(err2)
	for decoder.More() {
		err3 := decoder.Decode(&samObj)
		checkErrors(err3)
		repJson = append(repJson, samObj)
	}
	fmt.Println(repJson)
	return repJson
}

func checkErrors(err error) {
	if err == nil {
		fmt.Println("**** Time *****")
	}
}
