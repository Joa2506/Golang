package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
)

//Article data struct
type Article struct {
	Type string `json:"type"`
	HarvesterID string `json:"harvesterId,omitempty"`
	CerebroScore float32 `json:"cerebro-score,omitempty"`
	URL string `json:"url,omitempty"`
	Title string `json:"title,omitempty"`
	CleanImage string `json:"cleanImage,omitempty"`
	CommercialPartner string `json:"commercialPartner,omitempty"`
	LogoURL string `json:"logoURL,omitempty"`

	
}


type Response struct {
	Items[] Article `json:"items"`
}

type HTTPResponse struct {
	HTTPStatus int `json:"httpStatus"`
	Response Response `json:"response"`
}

//Slice of articles
var output[] Article



//Decoding from a http request
func DecodeJson(url string, target interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(target)


}
//Function to fetch json
func getArticle(url string) []Article{
	
	var resp HTTPResponse
	err :=DecodeJson(url, &resp)

	if err != nil {
		fmt.Printf("Error getting article: %s\n", err.Error())
		return resp.Response.Items
	} else {
		//fmt.Printf("%d", article.HTTPStatus)
		return resp.Response.Items
	}
}

//Used for debugging
func PrettyEncode(data interface{}, out io.Writer) error {
    enc := json.NewEncoder(out)
    enc.SetIndent("", "    ")
    if err := enc.Encode(data); err != nil {
        return err
    }
    return nil
}

//Function does most of the logic for this assignment. Was hoping to finish with a working pattern of five and fives.
func running (ArticleURL string, ContentMarketingURL string) []byte{
	var article[] Article
	var contentmarketing[] Article
	var results[] Article

	var add Article
	add.Type = "Add"
 	article = getArticle(ArticleURL)
	contentmarketing = getArticle(ContentMarketingURL)

	//Add adds to contentmarketing. Not optimal
	for i := len(contentmarketing); i < len(article); i++ {
		contentmarketing= append(contentmarketing, add)
	}

	wg := &sync.WaitGroup{}
	mut := &sync.Mutex{}
	wg.Add(2)
	//Was hoping to find a way to merge the pattern here.
	go func (wg *sync.WaitGroup, mut *sync.Mutex) {
		for i := 0; i < len(article); i++ {
			mut.Lock()
			results = append(results, article[i])
			mut.Unlock()
		}
		wg.Done()
	}(wg, mut)
	go func (wg *sync.WaitGroup, mut *sync.Mutex) {
		for i := 0; i < len(article); i++ {
			mut.Lock()
			results = append(results, contentmarketing[i])
			mut.Unlock()
		}
		wg.Done()
	}(wg, mut)
	wg.Wait()
	for i := 0; i < len(results); i++ {
		fmt.Printf("%s\n", results[i].Type)
	}
	fmt.Printf("%d\n",len(article))

	// j, _ := json.Marshal(results)
	// log.Println(string(j))
  
	j, _ := json.MarshalIndent(results, "", "  ")
	
	return j
}