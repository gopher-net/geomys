/*
   Primary starting point for georipd

   Author: Matt Oswalt
*/

package main

import (
    "./api"
    "./ripdb"
    "./serial"
    "fmt"
)

func main() {

    routelist := make(chan []ripdb.RipRoute)
    go ripdb.StartDb(routelist) //I saw some examples passing a pointer to this channel instead of the channel itself. Thoughts?

    // Need to have a loop here that goes through multiple channels and sees if something needs to be sent.
    // Will have to use buffered channels for this

    // Kick off API server async
    go ripapi.Start(routelist) // TODO: Does this create a copy of the channel, or a pointer?
    fmt.Println("REST API for RIPv2 is running....")

    go serial.Start(routelist)

    // // Press a key to exit
    var input string
    fmt.Scanln(&input)
    fmt.Println("done")

}
