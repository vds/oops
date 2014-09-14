//

package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"

	"github.com/vds/oops/cmd/protocol"
)

func receiveOops(conn net.Conn) {
	for {
		log.Println("Receiving length.")
		l, err := protocol.ReceiveOopsLength(conn)
		if err != nil {
			log.Println(err)
			if err == io.EOF {
				return
			} else {
				continue
			}
		}
		log.Println("Sending ack.")
		err = protocol.SendAck(conn)
		if err != nil {
			log.Println(err)
			if err == io.EOF {
				return
			} else {
				continue
			}
		}
		log.Println("Receiving oops")
		err = protocol.ReceiveOops(conn, l)
		if err != nil {
			log.Println(err)
			if err == io.EOF {
				return
			} else {
				continue
			}
		}
		log.Println("Sending ack")
		err = protocol.SendAck(conn)
		if err != nil {
			log.Println(err)
			if err == io.EOF {
				return
			} else {
				continue
			}
		}
	}
}

func main() {
	server := flag.String("server", "localhost", "The address of the receiver.")
	port := flag.Int("port", 5678, "The port of the receiver.")
	flag.Parse()

	servAddr := fmt.Sprintf("%s:%d", *server, *port)
	tcpAddr, err := net.ResolveTCPAddr("tcp", servAddr)
	if err != nil {
		log.Fatalf("ResolveTCPAddr failed: %s\n", err)
	}
	l, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go receiveOops(conn)
	}
}
