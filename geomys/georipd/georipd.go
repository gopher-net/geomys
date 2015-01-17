/*
   Primary starting point for georipd

   Author: Matt Oswalt
*/

package main

import (
    "./api"
    "./obj"
    "fmt"
)

func main() {

    /* -- JUST SOME EXAMPLE DATA -- */
    var routes []obj.RipRoute
    r1 := obj.RipRoute{
        AddrFamily: 2,
        RouteTag:   0,
        IpAddr:     0,
        Netmask:    0,
        NextHop:    0,
        Metric:     1,
    }
    routes = append(routes, r1)
    r2 := obj.RipRoute{
        AddrFamily: 2,
        RouteTag:   0,
        IpAddr:     10486272,
        Netmask:    4294967040,
        NextHop:    0,
        Metric:     1,
    }
    routes = append(routes, r2)
    r3 := obj.RipRoute{
        AddrFamily: 2,
        RouteTag:   0,
        IpAddr:     167772672,
        Netmask:    4294967040,
        NextHop:    0,
        Metric:     16,
    }
    routes = append(routes, r3)
    obj.UpdateNbrRoutes(2, 2, 168558593, routes)
    /* -- JUST SOME EXAMPLE DATA -- */

    GetRipDB := func() map[string]obj.RipMessage {
        return obj.GetRoutesStr()
    }

    // Kick off API server async
    go ripapi.Start(GetRipDB)
    fmt.Println("REST API for RIPv2 is running....")

    // Press a key to exit
    var input string
    fmt.Scanln(&input)
    fmt.Println("done")

}
