/*
    Serialization for RIP.
*/

package serial

import (
    "fmt"
    "net"
    "../ripdb"
    "encoding/binary"
)

func Start(routelist chan []ripdb.RipRoute){
    //Attach to RIP multicast stream
    ripaddr, err := net.ResolveUDPAddr("udp", "224.0.0.9:520")
    if err != nil {
        panic(err)
    }
    conn, err := net.ListenMulticastUDP("udp", nil, ripaddr)
    if err != nil {
        panic(err)
    }
    //Create listener instance
    go listen(conn, routelist)
}

func listen(conn *net.UDPConn, routelist chan []ripdb.RipRoute) {
    //Read from wire forever for RIP Packets
    fmt.Println("Attached to RIP Multicast IP....")
    for {
        b := make([]byte, 504)
        leng, sender, err := conn.ReadFromUDP(b)
        if err != nil {
            panic(err)
        }
        command := b[0]
        //version := b[1] //Leaving version here for future reference

        //Handle incoming routes
        if command == 2 {
            var routes []ripdb.RipRoute
            for i := 4; i <= leng-4; i+=20 {
                //Read bytes into protocol vars
                addr_family := b[i:i+2]
                prefix := b[i+4:i+8]
                subnet := b[i+8:i+12]
                nexthop := b[i+12:i+16]
                metric := b[i+16:i+20]

                //Set up route object and append to route slice
                routes = append(routes, ripdb.RipRoute{
                    AddrFamily: binary.BigEndian.Uint16(addr_family),
                    Sender:     binary.BigEndian.Uint32(sender.IP),
                    RouteTag:   0,
                    IpAddr:     binary.BigEndian.Uint32(prefix),
                    Netmask:    binary.BigEndian.Uint32(subnet),
                    NextHop:    binary.BigEndian.Uint32(nexthop),
                    Metric:     binary.BigEndian.Uint32(metric),
                })
            }

            //Send recieved routes to handler and dump return
            routelist <- routes
            <-routelist
        } else {
            //TODO: Handle incoming type 1 packet and recieve current DB state
            fmt.Printf("Perform DB send.\n")
        }
    }
}
