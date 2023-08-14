 package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "os"
    "strings"
    "database/sql"
    "encoding/json"
    _ "github.com/go-sql-driver/mysql"	
    "github.com/joho/godotenv"

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

func envVariable(key string) string {
	
	err := godotenv.Load(".env")

	if err != nil {
        	log.Fatal("Error loading .env file")
    	}

  return os.Getenv(key)
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

func transform(response string) map[string]int {
	var data Data
	m := make(map[string]int)

	json.Unmarshal([]byte(response), &data)
	for _,value := range data.Facet_groups[0].Facets{
	    m[value.Arrondissement] = value.Count
	}

    	return m
}

func load(rows map[string]int) string{
    db_path := envVariable("MYSQL_USER")+":"+envVariable("MYSQL_PASSWORD")+"@tcp(127.0.0.1:"+envVariable("MYSQL_PORT")+")/"+envVariable("MYSQL_DATABASE")
    db, err := sql.Open("mysql",db_path)
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()

    query := "INSERT INTO BureauxDeVote(Arrondissement, Y2021) VALUES"
    var inserts []string
    var params []interface{}
    for key, element := range rows {
        inserts = append(inserts, "(?, ?),")
        params = append(params, key, element)
    }  
    query = query + strings.TrimSuffix(inserts[len(inserts)-1],",")
    fmt.Println(query)
    //_, err := db.Exec(query, params)
    
    fmt.Println("Yay, values added!")

return db_path
}


func main() {

	var url = "https://opendata.paris.fr/api/records/1.0/search/?dataset=secteurs-des-bureaux-de-vote-en-2021&q=&rows=0&facet=arrondissement"
	data := transform(extract(url))
	//load(data)

}