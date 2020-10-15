package main

import (
  "encoding/json"
"fmt"
"log"
"net/http"
"strings"
"time"
)
type Article struct {
	ID                string    `json:"id"`
	Title             string    `json:"title"`
	Subtitle          string    `json:"subtitle"`
	Content           string    `json:"content"`
  CreationTime      time.Time `json:"time"`
}


var Articles []Article

//handlers

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

//Create Article
func return_All_Articles(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}


//Creating Article by Id
func return_one_article(w http.ResponseWriter, r *http.Request) {

  var Urls =  strings.Split(r.URL.Path, "/")
  var Temp_Article = Articles[0]
  //Calling Id's
  Temp_Article.ID = "NOTFOUND"
  for i := 0; i < len(Articles); i++{
      //loop for calling Ids
      if Urls[2]==Articles[i].ID {
        Temp_Article = Articles[i]
      }

    }
  if Temp_Article.ID == "NOTFOUND" {

    w.WriteHeader(http.StatusNotFound)
    return
  }

  if len(Urls) != 3 {

    w.WriteHeader(http.StatusNotFound)
    return
  }
  json.NewEncoder(w).Encode(Temp_Article)

}

func create(w http.ResponseWriter, r *http.Request) {
  var temp_article Article

    // Try to decode the request body into the struct. If there is an error,
    // respond to the client with the error message and a 400 status code.
    err := json.NewDecoder(r.Body).Decode(&temp_article)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    temp_article.CreationTime = time.Now()
    // Do something with the Person struct...
    Articles = append(Articles,temp_article)
    json.NewEncoder(w).Encode(temp_article)
}

// Search function
func Search(w http.ResponseWriter, r *http.Request) {

    var Found_Articles  []Article
    var q1 = r.URL.Query().Get("q")

    for i := 0; i < len(Articles); i++{
        contains_title := strings.Contains(Articles[i].Title,q1)
        contains_subtitle := strings.Contains(Articles[i].Subtitle,q1)
        contains_content := strings.Contains(Articles[i].Content,q1)
        present := contains_title || contains_subtitle || contains_content
        if present {
          Found_Articles = append(Found_Articles,Articles[i])
        }

      }
  json.NewEncoder(w).Encode(Found_Articles)
}






//Create Request
func Create_or_return(w http.ResponseWriter, r *http.Request) {
  if r.Method == "POST" {
    create(w,r)
    return
  }

  if r.Method == "GET" {
    return_All_Articles(w,r)
    return
  }


}

//url function
func handleRequests() {
    http.HandleFunc("/", homePage)
    http.HandleFunc("/articles",Create_or_return)
    http.HandleFunc("/articles/",return_one_article)
    http.HandleFunc("/articles/search",Search)
    log.Fatal(http.ListenAndServe(":10000", nil))
}


func main() {
  Articles = []Article{
  Article{ID: "1", Title: "Hello", Subtitle: "Article Description", Content: "Article Content"},
  Article{ID: "2", Title: "Hello 2", Subtitle: "Article Description", Content: "Article Content"},
}

    handleRequests()

}
