/*
   REST API for georipd

   Author: Matt Oswalt
*/

package ripapi

import (
    "../ripdb"
    "encoding/json"
    "fmt"
    "net/http"
)

var ripchan chan []ripdb.RipRoute

func Start(c chan []ripdb.RipRoute) {
    ripchan = c

    http.HandleFunc("/", serveRest)
    http.ListenAndServe("localhost:8080", nil)
}

func serveRest(w http.ResponseWriter, r *http.Request) {
    response, err := createJsonResponse()
    if err != nil {
        panic(err)
    }

    fmt.Fprint(w, string(response))
}

func createJsonResponse() ([]byte, error) {
    var tempslice []ripdb.RipRoute
    ripchan <- tempslice
    retdb := <-ripchan
    return json.MarshalIndent(retdb, "", "  ")
}
