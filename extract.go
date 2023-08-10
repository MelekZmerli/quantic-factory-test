package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "os"
    "encoding/json"
)

func main() {
    response, err := http.Get("https://opendata.paris.fr/api/records/1.0/search/?dataset=secteurs-des-bureaux-de-vote-en-2021&q=&rows=0&facet=arrondissement")

    if err != nil {
        fmt.Print(err.Error())
        os.Exit(1)
    }

    responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }
   // fmt.Println(string(responseData))

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

    var data Data
    json.Unmarshal([]byte(string(responseData)), &data)
    fmt.Printf("Data : %+v", data.Facet_groups[0].Facets)
    //data = data.Facet_groups[0].Facets
    for _,value := range data.Facet_groups[0].Facets{
	fmt.Print(value.Arrondissement,": ")
	fmt.Println(value.Count)

	
}
}