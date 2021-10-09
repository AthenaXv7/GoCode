package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type user struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
type Postdata struct {
	ID        string    `json:"ID"`
	Caption   string    `json:"caption"`
	Image_url string    `json:"url"`
	TimeStamp time.Time `json:"time"`
}
type data1 []user
type postdata1 []Postdata

func userdata(w http.ResponseWriter, r *http.Request) {
	data := data1{
		user{ID: "1", Name: "abc", Email: "abc113@gmail.com", Password: "aakg"},
		user{ID: "2", Name: "efg", Email: "efg113@gmail.com", Password: "aakg"}}
	json.NewEncoder(w).Encode(data)

}
func postsData(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	p_data := postdata1{
		Postdata{ID: "1", Caption: "First", Image_url: "xyz", TimeStamp: now},
		Postdata{ID: "2", Caption: "Second", Image_url: "zzx", TimeStamp: now}}
	json.NewEncoder(w).Encode(p_data)

}

func homepage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1> Instagram Backend API : Appointy </h1>"))
}

func HandlerRequest() {
	http.HandleFunc("/", homepage)
	http.HandleFunc("/users", userdata)
	http.HandleFunc("/posts", postsData)
	err := http.ListenAndServe(":9090", nil) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func main() {
	HandlerRequest()
}
