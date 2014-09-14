//
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"path"

	"github.com/vds/oops/cmd/protocol"
)

var (
	oopsDir *string
)

func dispatch(conn *net.TCPConn, oopsFiles []os.FileInfo) (err error) {
	log.Println("Start dispatching.")
	for _, fi := range oopsFiles {
		log.Printf("Dispatching: %s", fi.Name())
		oopsPath := path.Join(*oopsDir, fi.Name())
		encodedOops, err := ioutil.ReadFile(oopsPath)
		if err != nil {
			log.Printf("Error reading oops file: %v\n", err)
			return err
		}

		l := int64(len(encodedOops))
		log.Println("sending length")
		err = protocol.SendOopsLength(conn, l)
		if err != nil {
			log.Println(err)
			return err
		}
		log.Println("receiving ack")
		protocol.ReceiveAck(conn)
		if err != nil {
			log.Println(err)
			return err
		}
		log.Println("sending oops")
		protocol.SendOops(conn, encodedOops, l)
		if err != nil {
			log.Println(err)
			return err
		}
		log.Println("receiving ack")
		protocol.ReceiveAck(conn)
		if err != nil {
			log.Println(err)
			return err
		}
		log.Printf("Finished: %s", fi.Name())
	}
	return
}

func main() {
	oopsDir = flag.String("oopsdir", "/home/vds/oops", "The directory that contains the oopses.")
	server := flag.String("server", "localhost", "The address of the receiver.")
	port := flag.Int("port", 5678, "The port of the receiver.")
	flag.Parse()

	// Connecting to the server

	oopsFiles, err := ioutil.ReadDir(*oopsDir)
	if err != nil {
		log.Fatalf("Error reading oops directory: %v\n", err)
	}
	if len(oopsFiles) == 0 {
		log.Println("No oops found")
		return
	}

	servAddr := fmt.Sprintf("%s:%d", *server, *port)
	tcpAddr, err := net.ResolveTCPAddr("tcp", servAddr)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Fatalf("Dial failed: %s\n", err)
	}
	err = dispatch(conn, oopsFiles)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Shutting down.")
	conn.Close()
}
