/*
   REST API for georipd

   Author: Matt Oswalt
*/

package ripapi

import (
    "../obj"
    "encoding/json"
    "fmt"
    "net/http"
)

var GetRipDB func() map[string]obj.RipMessage

func Start(f func() map[string]obj.RipMessage) {
    GetRipDB = f

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
    return json.MarshalIndent(GetRipDB(), "", "  ")
}
