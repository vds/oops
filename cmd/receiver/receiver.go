//

package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/vds/oops/cmd/protocol"
)

func receiveOops(conn net.Conn) {
}

// Receiver Service
type Receiver struct {
	c         chan bool
	waitGroup *sync.WaitGroup
}

// Make a new Receiver
func NewReceiver() *Receiver {
	s := &Receiver{
		c:         make(chan bool),
		waitGroup: &sync.WaitGroup{},
	}
	s.waitGroup.Add(1)
	return s
}

// Close the receiver communication channel and wait for all the go routines do terminate.
func (s *Receiver) Stop() {
	close(s.c)
	s.waitGroup.Wait()
}

// Prepare to receive the oops from the dispatcher. When thce receiver communication channel is closed, closes the connection and terminate.
func (s *Receiver) Receive(listener *net.TCPListener) {
	defer s.waitGroup.Done()
	for {
		select {
		case <-s.c:
			log.Println("Stopping receiving from: %v.", listener.Addr())
			listener.Close()
			return
		default:
		}
		listener.SetDeadline(time.Now().Add(1e9))
		conn, err := listener.AcceptTCP()
		if nil != err {
			if opErr, ok := err.(*net.OpError); ok && opErr.Timeout() {
				continue
			}
			log.Println(err)
		}
		log.Println(conn.RemoteAddr(), "Connected.")
		s.waitGroup.Add(1)
		go s.receive(conn)
	}
}

// Actually receives the oops from the dispatcher, gets the length of the oops first, replies with an ack, receives the real oops, and again, replies with an ack.
func (s *Receiver) receive(conn *net.TCPConn) {
	defer conn.Close()
	defer s.waitGroup.Done()
	for {
		select {
		case <-s.c:
			log.Println("Disconnecting.", conn.RemoteAddr())
			return
		default:
		}
		conn.SetDeadline(time.Now().Add(1e9))
		for {
			log.Println("Receiving oops.")
			err := protocol.ReceiveOops(conn)
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
		log.Fatalf("resolvetcpaddr failed: %s\n", err)
	}
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Fatal(err)
	}
	receiver := NewReceiver()
	go receiver.Receive(listener)
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	log.Println(<-ch)
	receiver.Stop()
}
