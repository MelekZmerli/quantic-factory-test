package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "regexp"
    "bufio"
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

func check(e error) {
    if e != nil {
        log.Fatal(e)
    }
}

func get_urls(filename string) map[string]string{
    uri := "https://opendata.paris.fr/api/records/1.0/search/?dataset="
    urls := make(map[string]string)
    re := regexp.MustCompile(`20[0-2][0-9]`)
    file, err := os.Open(filename)
    check(err)
    defer file.Close()
    
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
	url := uri + scanner.Text()
        year := re.FindStringSubmatch(url)[0]
	urls[year] = url
    }
    err1 := scanner.Err()
    check(err1)

    return urls
}

func envVariable(key string) string {
	
	err := godotenv.Load(".env")

	check(err)

  return os.Getenv(key)
}

func extract(url string) string {
	response, err := http.Get(url)

	check(err)

    	responseData, err := ioutil.ReadAll(response.Body)
    	check(err)
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

func load(rows map[string]int){
    db_path := envVariable("MYSQL_USER")+":"+envVariable("MYSQL_PASSWORD")+"@tcp(127.0.0.1:"+envVariable("MYSQL_PORT")+")/"+envVariable("MYSQL_DATABASE")
    db, err := sql.Open("mysql",db_path)
    check(err)
res, err := db.Query("SHOW TABLES")
check(err)

var table string

for res.Next() {
    res.Scan(&table)
    fmt.Println(table)
}
    query := "INSERT INTO BureauxDeVote(Arrondissement, Y2021) VALUES"
    var inserts []string
    var params []interface{}
    
    for key, element := range rows {
        inserts = append(inserts, "(?, ?)")
        params = append(params, key, element)
    }
    query = query + strings.Join(inserts,",")
    /*stmt, err := db.Prepare(query)

    check(err)

    res, err := stmt.Exec(params...)   
    check(err)
    fmt.Println(res)*/
    defer db.Close()
}


func main() {
	urls := get_urls("./urls.txt")
	for _,element := range(urls){
		go transform(extract(element))
	}	
	/*data := transform(extract(url))

	load(data)*/

}