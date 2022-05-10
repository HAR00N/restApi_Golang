package test

import (
	"encoding/json"
	"fmt"
	"goapi/models"
)

func TestData() models.ResData {
	var item1 = models.Item{"2022-05-04", "http://www.mogef.go.kr/thumb", "새일센터", "www.mogef.go.kr/nw"}
	var item2 = models.Item{"", "http://www.getImage.do", "TEMP", "http://www.nw_enw_s001d.do"}
	var item3 = []models.Item{item1, item2}

	var items = models.Items{item3}

	var body = []models.Body{{items, 1, 3992, 10000}}
	var header = []models.Header{{0, "NORMAL SERVICE."}}

	var resdata = models.ResData{header, body}

	res, _ := json.Marshal(resdata)

	var apiObj models.ResData

	var err = json.Unmarshal(res, &apiObj)
	if err != nil {
		fmt.Println("Failed to json.Unmarshal", err)
	}

	return apiObj

}
