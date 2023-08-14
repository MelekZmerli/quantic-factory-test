 package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "os"
    "encoding/json"
    _ "github.com/go-sql-driver/mysql"	
)

type Row struct {
	Arrondissement string `json:"name"`
	Count int `json:"count"`
	State string `json:"state"`
	Path string `json:"path"`
} 

type FacetField struct {
	Name string
	Facets []Row
}

type ParametersStruct struct {
	Dataset string
	Rows int
	Start int
	Facet []string
	Format string
	Timezone string
}

type Data struct {
	Nhits int `json:"nhits"`
	Parametres ParametersStruct `json:"parameters"`
	Records []string `json:"records"`
	Facet_groups []FacetField `json:"facet_groups"`
}

func extract(url string) string {
	response, err := http.Get(url)

	if err != nil {
        	fmt.Print(err.Error())
        	os.Exit(1)
    	}

    	responseData, err := ioutil.ReadAll(response.Body)
    	if err != nil {
        	log.Fatal(err)
    	}
   	return string(responseData)
}

func transform(response string) {
	var data Data
	json.Unmarshal([]byte(response), &data)
    	fmt.Printf("Data : %+v", data.Facet_groups[0].Facets)

    	for _,value := range data.Facet_groups[0].Facets{
	fmt.Print(value.Arrondissement,": ")
	fmt.Println(value.Count)
	}
}

func main() {

	var url = "https://opendata.paris.fr/api/records/1.0/search/?dataset=secteurs-des-bureaux-de-vote-en-2021&q=&rows=0&facet=arrondissement"
	transform(extract(url))

}