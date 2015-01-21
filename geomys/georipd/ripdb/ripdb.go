/*
   Object definitions for RIP.
   Currently only supporting RIPv2

   Author: Matt Oswalt
*/

package ripdb

import (
    "time"
)

// Main slice for RIP entries
var runningRipDB []RipRoute

func StartDb(routelist chan []RipRoute) {
    /*
       This function is easy. Pass it a slice of RipRoute types in order to consider them for the database.
       They will be iterated through - if an entry is already in the database, then it will have its timestamp updated.
       If it doesn't exist, then it will be inserted with a fresh timestamp.

       This function will always return the full database no matter what. So if you want to retrieve the existing database
       without making changes, just send an empty slice as an argument.

    */
    for {
        routelist <- updateDB(<-routelist)
    }
}

func updateDB(newList []RipRoute) []RipRoute {
    for i := range newList {
        updateRoute(newList[i])
    }

    return runningRipDB // Always return entire database.
}

func updateRoute(thisRoute RipRoute) {
    //var endtable bool
    endtable := false
    for i := range runningRipDB {
        if thisRoute.IpAddr == runningRipDB[i].IpAddr {
            // Duplicate - update timer
            runningRipDB[i].LastUpdate = time.Now().UTC()
            endtable = true
            break
        }
    }

    if endtable != true {
        // Not a duplicate, insert the route
        thisRoute.LastUpdate = time.Now().UTC()
        runningRipDB = append(runningRipDB, thisRoute)
    }
}

// Struct for RIP Authentication Headers
type RipAuth struct {
    AuthType     uint16
    DigestOffset uint16
    KeyID        uint8
    AuthDataLen  uint8
    SeqNum       uint32
    AuthData     string // No uint128 type, so have to do string
}

// Struct for a network advertised by RIP
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
