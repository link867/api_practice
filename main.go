pckage main

import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

// our main function
func main() {
    router := mux.NewRouter()
    router.HandleFunc("/movies", GetAll).Methods("GET")
    router.HandleFunc("/movies", Create).Methods("POST")
    router.HandleFunc("/movies/{id}", GetMovies).Methods("GET")
    router.HandleFunc("/movies/{id}", Update).Methods("PUT")
    router.HandleFunc("/movies/{id}", DeletePerson).Methods("DELETE")
    log.Fatal(http.ListenAndServe(":8080", router))
}




var Movies = []struct {
     ID       int
     Director string
     Name     string
     Rating   int
} {
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
