package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Fact struct {
	IconURL string `json:"icon_url"`
	Id      string `json:"id"`
	Url     string `json:"url"`
	Value   string `json:"value"`
}

type SearchFact struct {
	Total  int    `json:"total"`
	Result []Fact `json:"result"`
}

func GetFact(reqUrl string) (*Fact, error) {
	resp, err := http.Get(reqUrl)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var fact Fact
	if err := json.NewDecoder(resp.Body).Decode(&fact); err != nil {
		return nil, err
	}
	return &fact, nil
}

func GetCategories() (*[]string, error) {
	reqUrl := "https://api.chucknorris.io/jokes/categories"
	resp, err := http.Get(reqUrl)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var cats []string
	if err := json.NewDecoder(resp.Body).Decode(&cats); err != nil {
		return nil, err
	}
	return &cats, nil
}

func GetRandom() (*Fact, error) {
	return GetFact("https://api.chucknorris.io/jokes/random")
}

func GetRandomByCategories(cat string) (*Fact, error) {
	return GetFact("https://api.chucknorris.io/jokes/random?category=" + cat)
}

func GetFactByQuery(query string) (*SearchFact, error) {
	resp, err := http.Get("https://api.chucknorris.io/jokes/search?query=" + query)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var sfact SearchFact
	if err := json.NewDecoder(resp.Body).Decode(&sfact); err != nil {
		return nil, err
	}
	return &sfact, nil
}

func main() {
	f, err := GetRandom()
	if err != nil {
		log.Fatalln("erreur: ", err.Error())
	}
	fmt.Println("Fact Random: ", f)

	cats, err := GetCategories()
	if err != nil {
		log.Fatalln("erreur: ", err.Error())
	}
	fmt.Println("categories:", cats)

	for _, cat := range *cats {
		f, err = GetRandomByCategories(cat)
		if err != nil {
			log.Fatalln("erreur: ", err.Error())
		}
		fmt.Println("Fact (", cat, "): ", f.Value)
	}

	query := "electric"
	sfacts, err := GetFactByQuery(query)
	if err != nil {
		log.Fatalln("erreur: ", err.Error())
	}
	fmt.Println("query: ", query)
	for _, fact := range sfacts.Result {
		fmt.Println("Fact query(", query, "): ", fact.Value)
	}
}
