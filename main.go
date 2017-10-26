package main

import (
    "encoding/json"
    "io/ioutil"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "fmt"
    "strings"
    "strconv"
)

type Movie struct {
     ID       int
     Director string
     Name     string
     Rating   int
}

var Movies = []Movie  {
    {
        ID: 1,
        Director: "Steven Spielberg",
        Name: "E.T. the Extra-Terrestrial",
        Rating: 5,
    },
    {
        ID: 2,
        Director: "Howard Hawks",
        Name: "Rio Bravo",
        Rating: 5,
    },
    {
        ID: 3,
        Director: "Tony Leondis",
        Name: "The Emoji Movie",
        Rating: 1,
    },
    {
        ID: 4,
        Director: "Michael Bay",
        Name: "Transformers 3",
        Rating: 3,
    },
}





func GetAll(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Returning all movies")

	json.NewEncoder(w).Encode(Movies)

	return
}


func GetMovie(w http.ResponseWriter, r *http.Request) {

	u := strings.Split(r.URL.Path, "/")
	i, _ := strconv.Atoi(u[2])
	m := Movies[i - 1]
	json.NewEncoder(w).Encode(m)
}


func Create(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
	//	panic(err)
		fmt.Println("Error readying body: %v", err)
	}

	var m Movie
	fmt.Println("body:", body)
	err = json.Unmarshal(body, &m)
}

func Update(w http.ResponseWriter, r *http.Request) {
}

func Delete(w http.ResponseWriter, r *http.Request) {
}



func main() {
    router := mux.NewRouter()
    router.HandleFunc("/movies", GetAll).Methods("GET")
    router.HandleFunc("/movies", Create).Methods("POST")
    router.HandleFunc("/movies/{id}", GetMovie).Methods("GET")
    router.HandleFunc("/movies/{id}", Update).Methods("PUT")
    router.HandleFunc("/movies/{id}", Delete).Methods("DELETE")
    log.Fatal(http.ListenAndServe(":8080", router))
}
