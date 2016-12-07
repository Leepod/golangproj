package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	result, err := http.Get("http://barno.info")
	if err != nil {
		log.Fatalln(err)
	}
	//status, _ := ioutil.ReadAll(result.Status)
	//fmt.Println(string(status))
	body, _ := ioutil.ReadAll(result.Body)
	fmt.Println(string(body))
}
