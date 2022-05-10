package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"goapi/api"
	"goapi/config"
	SysKeyword "goapi/global"
	"os"
	"strings"
)

func main() {
	godotenv.Load(".env")
	url := os.Getenv("apiUrl") + SysKeyword.AND() + os.Getenv("apiKey")

	database := config.DbConnection()
	defer database.Close()

	apiData := api.ApiCall(url)
	var itemData = apiData.Body[0].Items.Item

	//for i := 0; i < 2; i++ {
	for i := 0; i < len(itemData); i++ {
		_, err := database.Exec(`insert into mogefapi (regDt, thumbUrl, title, viewUrl) value ("` + itemData[i].RegDt + `","` + itemData[i].ThumbUrl + `","` + strings.ReplaceAll(itemData[i].Title, `"`, ``) + `","` + itemData[i].ViewUrl + `")`)
		if err != nil {
			panic(err)
		}
	}

}

func main2() {
	database := config.DbConnection()
	defer database.Close()

	result, err := database.Exec(`insert into framework_resource value (4, "MJH", "dain")`)
	if err != nil {
		panic(err)
	}
	nRow, err := result.RowsAffected()
	fmt.Println("insert Count : ", nRow)
}

func mai22n() {
	godotenv.Load(".env")
	url := os.Getenv("apiUrl") + SysKeyword.AND() + os.Getenv("apiKey")

	apiData := api.ApiCall(url)
	var itemData = apiData.Body[0].Items.Item

	//for i := 0; i < len(itemData); i++ {
	//for i := 0; i <= 5; i++ {
	//fmt.Println("\nRegDt : ", itemData[i].RegDt, "\nThumbUrl : ", itemData[i].ThumbUrl, "\nTitle : ", itemData[i].Title, "\nViewUrl : ", itemData[i].ViewUrl)
	//fmt.Println(itemData[i])
	//}
	fmt.Println(itemData[1014])

}
