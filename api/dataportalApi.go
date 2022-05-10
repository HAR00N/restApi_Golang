package api

import (
	"encoding/json"
	"fmt"
	"goapi/models"
	"io/ioutil"
	"net/http"
)

func ApiCall(url string) models.ResData {
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	data, _ := ioutil.ReadAll(res.Body)
	var apiObj models.ResData
	err = json.Unmarshal(data, &apiObj)
	if err != nil {
		fmt.Println("Failed to json.Unmarshal : ", err)
	}

	return apiObj //JSONObject
}
