/*
   Serialization for RIP.
*/

package serial

import (
	"../ripdb"
	"bytes"
	"encoding/binary"
	"net"
	"time"
)

func Start(routelist chan []ripdb.RipRoute) {
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
	go SendTypeOne(conn)
	go func() {
		for {
			time.Sleep(30 * time.Second)
			SendRoutes(routelist, conn)
		}
	}()
}

func SendTypeOne(conn *net.UDPConn) {
	packet := new(bytes.Buffer)

	// Write command into new packet
	err := binary.Write(packet, binary.BigEndian, byte(1))
	if err != nil {
		panic(err)
	}

	// Write version to new packet
	err = binary.Write(packet, binary.BigEndian, byte(2))
	if err != nil {
		panic(err)
	}

	var pad16 [2]byte
	pad16[0] = 0
	pad16[1] = 0
	for i := 0; i < 10; i++ {
		err = binary.Write(packet, binary.BigEndian, pad16)
		if err != nil {
			panic(err)
		}
	}

	err = binary.Write(packet, binary.BigEndian, byte(0))
	if err != nil {
		panic(err)
	}

	err = binary.Write(packet, binary.BigEndian, byte(16))
	if err != nil {
		panic(err)
	}
	ripaddr, err := net.ResolveUDPAddr("udp", "224.0.0.9:520")
	conn.WriteToUDP(packet.Bytes(), ripaddr)
}

func SendRoutes(routelist chan []ripdb.RipRoute, conn *net.UDPConn) {
	var routes []ripdb.RipRoute
	routelist <- routes
	routes = <-routelist
	packet := new(bytes.Buffer)

	// Write command into new packet
	err := binary.Write(packet, binary.BigEndian, byte(2))
	if err != nil {
		panic(err)
	}

	// Write version to new packet
	err = binary.Write(packet, binary.BigEndian, byte(2))
	if err != nil {
		panic(err)
	}
	var pad16 [2]byte
	pad16[0] = 0
	pad16[1] = 0
	err = binary.Write(packet, binary.BigEndian, pad16)
	if err != nil {
		panic(err)
	}
	for i := range routes {
		err = binary.Write(packet, binary.BigEndian, routes[i].AddrFamily)
		if err != nil {
			panic(err)
		}
		err = binary.Write(packet, binary.BigEndian, pad16)
		if err != nil {
			panic(err)
		}
		err := binary.Write(packet, binary.BigEndian, routes[i].IpAddr)
		if err != nil {
			panic(err)
		}
		err = binary.Write(packet, binary.BigEndian, routes[i].Netmask)
		if err != nil {
			panic(err)
		}
		err = binary.Write(packet, binary.BigEndian, routes[i].NextHop)
		if err != nil {
			panic(err)
		}
		err = binary.Write(packet, binary.BigEndian, routes[i].Metric)
		if err != nil {
			panic(err)
		}
	}
	ripaddr, err := net.ResolveUDPAddr("udp", "224.0.0.9:520")
	conn.WriteToUDP(packet.Bytes(), ripaddr)
}

func parse_message(b []byte, leng int, routelist chan []ripdb.RipRoute,
	sender *net.UDPAddr, conn *net.UDPConn) {
	command := b[0]
	//version := b[1] //Leaving version here for future reference

	//Handle incoming routes
	if command == 2 {
		var routes []ripdb.RipRoute
		for i := 4; i <= leng-4; i += 20 {
			//Read bytes into protocol vars
			addr_family := b[i : i+2]
			prefix := b[i+4 : i+8]
			subnet := b[i+8 : i+12]
			nexthop := b[i+12 : i+16]
			metric := b[i+16 : i+20]

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
		SendRoutes(routelist, conn)
	}
}

func listen(conn *net.UDPConn, routelist chan []ripdb.RipRoute) {
	//Read from wire forever for RIP Packets
	for {
		b := make([]byte, 504)
		leng, sender, err := conn.ReadFromUDP(b)
		if err != nil {
			panic(err)
		}
		parse_message(b, leng, routelist, sender, conn)
	}
}
