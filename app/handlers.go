package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "io/ioutil"
    "io"
    "github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome to the Otter API!")
}

func OtterIndex(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    if error := json.NewEncoder(w).Encode(otters); error != nil {
        panic(error)
    }
}

func OtterShow(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    otterId := vars["otterId"]
    fmt.Fprintln(w, "Otter show:", otterId)
}

func OtterCreate(w http.ResponseWriter, r *http.Request) {
    var otter Otter
    body, error := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
    if error != nil {
        panic(error)
    }
    if error := r.Body.Close(); error != nil {
        panic(error)
    }
    if error := json.Unmarshal(body, &otter); error != nil {
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(422) // unprocessable entity
        if error := json.NewEncoder(w).Encode(error); error != nil {
            panic(error)
        }
    }

    o := RepoCreateOtter(otter)
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusCreated)
    if error := json.NewEncoder(w).Encode(o); error != nil {
        panic(error)
    }
}