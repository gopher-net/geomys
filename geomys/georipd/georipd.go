/*
   Primary starting point for georipd

   Author: Matt Oswalt
*/

package main

import (
    "./api"
    "./ripdb"
    "fmt"
    "time"
)

func main() {

    routelist := make(chan []ripdb.RipRoute)
    go ripdb.StartDb(routelist) //I saw some examples passing a pointer to this channel instead of the channel itself. Thoughts?

    // Need to have a loop here that goes through multiple channels and sees if something needs to be sent.
    // Will have to use buffered channels for this

    /* -- JUST SOME EXAMPLE DATA -- */
    // Note that I'm leaving out the timestamp. This should be the same at the serialization layer.
    var routes []ripdb.RipRoute
    routes = append(routes, ripdb.RipRoute{
        AddrFamily: 2,
        Sender:     10486276,
        RouteTag:   0,
        IpAddr:     10480001,
        Netmask:    4294967040,
        NextHop:    0,
        Metric:     1,
    })
    routes = append(routes, ripdb.RipRoute{
        AddrFamily: 2,
        Sender:     10486276,
        RouteTag:   0,
        IpAddr:     10480002,
        Netmask:    4294967040,
        NextHop:    0,
        Metric:     1,
    })
    routes = append(routes, ripdb.RipRoute{
        AddrFamily: 2,
        Sender:     10486276,
        RouteTag:   0,
        IpAddr:     10480003,
        Netmask:    4294967040,
        NextHop:    0,
        Metric:     16,
    })
    /* -- JUST SOME EXAMPLE DATA -- */

    routelist <- routes
    <-routelist

    // Kick off API server async
    go ripapi.Start(routelist) // TODO: Does this create a copy of the channel, or a pointer?
    fmt.Println("REST API for RIPv2 is running....")

    time.Sleep(10 * time.Second)
    var routestwo []ripdb.RipRoute
    routestwo = append(routestwo, ripdb.RipRoute{
        AddrFamily: 2,
        Sender:     10486276,
        RouteTag:   0,
        IpAddr:     10480003,
        Netmask:    4294967040,
        NextHop:    0,
        Metric:     16,
    })

    routelist <- routestwo
    <-routelist

    // // Press a key to exit
    var input string
    fmt.Scanln(&input)
    fmt.Println("done")

}
