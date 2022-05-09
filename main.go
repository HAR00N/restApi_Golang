package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"goapi/api"
	"goapi/models"
	"log"
)

func main() {
	godotenv.Load(".env")
	//url := os.Getenv("apiUrl") + SysKeyword.AND() + os.Getenv("apiKey")
	//api.ApiCall(url)

	db := api.DbConnection()
	err := db.Ping()
	if err != nil {
		panic(err)
	}

	mysqlQuery := `SELECT * from framework_resource Where uid = 1`

	name := models.Localhost{}
	rows, err := db.Query(mysqlQuery)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&name)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("rows : ", name)
}

//func main() {
//	type Items struct {
//		RegDt    string `json:"regDt"`
//		ThumbUrl string `json:"thumbUrl"`
//		Title    string `json:"title"`
//		ViewUrl  string `json:"viewUrl"`
//	}
//
//	type Body struct {
//		PageNo     int     `json:"pageNo"`
//		TotalCount int     `json:"totalCount"`
//		NumOfRows  int     `json:"numOfRows"`
//		Items      []Items `json:"items"`
//	}
//
//	type Header struct {
//		ResultCode int    `json:"resultCode"`
//		ResultMsg  string `json:"resultMsg"`
//	}
//
//	type ResData struct {
//		Header []Header `json:"header"`
//		Body   []Body   `json:"body"`
//	}
//
//	var items1 = Items{"2022-05-04", "http://www.mogef.go.kr/thumb", "새일센터", "www.mogef.go.kr/nw"}
//	var items2 = Items{"", "http://www.getImage.do", "TEMP", "http://www.nw_enw_s001d.do"}
//	var items3 = []Items{items1, items2}
//
//	var body = []Body{{1, 3992, 10000, items3}}
//	var header = []Header{{0, "NORMAL SERVICE."}}
//
//	var resdata = ResData{header, body}
//	res, _ := json.Marshal(resdata)
//
//	fmt.Println(string(res))
//
//}
