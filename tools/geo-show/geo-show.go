/*
   Primary starting point for geo-state.go

   Author: Matt Oswalt
*/

package main

import (
    "encoding/binary"
    "encoding/json"
    "flag"
    "fmt"
    "github.com/olekukonko/tablewriter"
    "io/ioutil"
    "net"
    "net/http"
    "os"
    "time"
)

func main() {

    // flag name, default value, short description
    tablePtr := flag.String("t", "none", "table to display")

    flag.Parse()

    switch *tablePtr {
    case "rip":
        getGeoripd()
    case "none":
        getHelp()
    default:
        getHelp()
    }

    // fmt.Println("svar:", svar)
    // fmt.Println("tail:", flag.Args())
}

func getHelp() {
    fmt.Println(`
usage: geo-state -t [table]  --  Outputs table from specified routing protocol
            Choices are: rip
    `)
}

func getGeoripd() {
    url := "http://127.0.0.1:8080"
    res, err := http.Get(url)
    if err != nil {
        panic(err)
    }
    defer res.Body.Close()

    body, err := ioutil.ReadAll(res.Body)
    if err != nil {
        panic(err)
    }

    var p []RipRoute

    err = json.Unmarshal(body, &p)
    if err != nil {
        panic(err)
    }

    var data [][]string

    timenow := time.Now().UTC()

    fmt.Println(`
+---------------------------+
|       RIP Database        |
+---------------------------+
    `)

    for i := range p {
        data = append(data, []string{string(convip(p[i].IpAddr)), convip(p[i].Netmask), fmt.Sprint(p[i].LastUpdate.Sub(timenow)), convip(p[i].NextHop), fmt.Sprint(p[i].Metric), convip(p[i].Sender), fmt.Sprint(p[i].RouteTag)})
    }

    table := tablewriter.NewWriter(os.Stdout)
    table.SetHeader([]string{"Network", "Mask", "Age", "Next Hop", "Metric", "Sender", "Tag"})

    for _, v := range data {
        table.Append(v)
    }
    table.Render() // Send output
}

func convip(u32 uint32) string {
    addr := net.IP{0, 0, 0, 0}
    binary.BigEndian.PutUint32(addr, u32)
    return fmt.Sprint(addr)
}

type RipRoute struct {
    LastUpdate time.Time
    Sender     uint32
    AddrFamily uint16
    RouteTag   uint16
    IpAddr     uint32
    Netmask    uint32
    NextHop    uint32
    Metric     uint32
}
