/*
   Object definitions for RIP.
   Currently only supporting RIPv2

   Author: Matt Oswalt
*/

package obj

import (
    "fmt"
    "time"
)

// Hash Table for storing all Rip Messages, keyed off of sender
var RipDB map[uint32]RipMessage

// Housekeeping, you want fluff pillow?
func init() {
    RipDB = make(map[uint32]RipMessage)
}

// Struct for RIP Messages
type RipMessage struct {
    LastUpdate time.Time
    Command    uint8
    Version    uint8
    Routes     []RipRoute
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

// Struct for Individual Networks advertised by RIP
type RipRoute struct {
    AddrFamily uint16
    RouteTag   uint16
    IpAddr     uint32
    Netmask    uint32
    NextHop    uint32
    Metric     uint32
}

// Set RipMessage for received message per sender
func UpdateNbrRoutes(cmd uint8, version uint8, sender uint32, routes []RipRoute) {
    var thisMsg RipMessage
    thisMsg.LastUpdate = time.Now().UTC()
    thisMsg.Command = cmd
    thisMsg.Version = version
    thisMsg.Routes = routes

    RipDB[sender] = thisMsg
}

// Retrieves current RIP Table
func GetRoutes() map[uint32]RipMessage {
    return RipDB
}

// Retrieves current RIP Table, but converts keys to strings first (for the REST API)
func GetRoutesStr() map[string]RipMessage {
    //TODO: This would be a good place to convert other fields to strings, for better API usability

    var RipDBStr map[string]RipMessage
    RipDBStr = make(map[string]RipMessage)

    var key uint32
    var val RipMessage

    for key, val = range RipDB {
        RipDBStr[fmt.Sprint(key)] = val
    }

    return RipDBStr
}
