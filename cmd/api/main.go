package main

import (
    "net/http"
    "github.com/gorilla/mux"
    "log"
)

func main() {
    router := mux.NewRouter()
    
    router.HandleFunc("/hosts", HostsHandler).Methods("GET")
    router.HandleFunc("/apps", AppsHandler).Methods("GET")
    router.HandleFunc("/cis-results", CISResultsHandler).Methods("GET")
    
    log.Fatal(http.ListenAndServe":"8080", router))
}

func HostsHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hosts endpoint"))
}

func AppsHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Apps endpoint"))
}

func CISResultsHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("CIS Results endpoint"))
}