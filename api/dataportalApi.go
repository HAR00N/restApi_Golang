package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func ApiCall(url string) {
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	jsonBytes, _ := json.Marshal(string(data))
	fmt.Println(string(jsonBytes))
	//jsonString := string(jsonBytes)
	//
	//var apiObj = models.ResData{}
	//
	//err = json.Unmarshal([]byte(jsonString), &apiObj)
	//
	//if err != nil {
	//	fmt.Println("Failed to json.Unmarshal", err)
	//}
	//return apiObj
}
