/*
   Primary starting point for georipd

   Author: Matt Oswalt
*/

package main

import (
	"./api"
	"./confdb"
	"./ripdb"
	"./serial"
	"fmt"
	"io/ioutil"
)

func main() {

	routelist := make(chan []ripdb.RipRoute)
	go ripdb.StartDb(routelist) //I saw some examples passing a pointer to this channel instead of the channel itself. Thoughts?

	// Need to have a loop here that goes through multiple channels and sees if something needs to be sent.
	// Will have to use buffered channels for this

	// Kick off API server async	go ripapi.Start(routelist) // TODO: Does this create a copy of the channel, or a pointer?
	go ripapi.Start(routelist) // TODO: Does this create a copy of the channel, or a pointer?

	// Start config database service
	dat, err := ioutil.ReadFile("./config.json")
	if err != nil {
		panic(err)
	}
	var config confdb.Config
	configchan := make(chan confdb.Config)
	go config.Start(dat, configchan)

	// Starr serial service
	go serial.Start(routelist)

	fmt.Println("RIP Daemon started. Press any key to exit...")
	// Press a key to exit
	var input string
	fmt.Scanln(&input)
	fmt.Println("done")

}
