/*CopyPast from teaching material
Добавьте для time-сервера возможность его корректного завершения при вводе команды exit.
*/
package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup

	abort := make(chan struct{}) // signal channel for gracefull shutdown
	conns := make(chan net.Conn) // conn channel for transmit network connections for handlers
	input := ""
	// start canceler server
	go func() {
		for {
			fmt.Printf("For exit type EXIT >\n")
			fmt.Scanln(&input)
			if strings.ToUpper(input) == "EXIT" {
				close(abort)
				break
			}
		}
	}()

	// open and listening the tcp port
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	// start of handlers - that approach for demo of goroutines starts by anonim func with copy of variables  wg *sync.WaitGroup, sema chan struct{}
	go func() {
		for conn := range conns {
			wg.Add(1)
			go handleConn(conn, abort, &wg)
		}
	}()

	// start acceptor of network connection
	go acceptor(listener, conns)

	// waiter for cancel
	select {
	case <-abort:
		// recieved shutdown signal
		wg.Wait() // wait for all handlers close connections, print goodbuy and shutdown server
		fmt.Println("All connection interrupted. Server shutting down. Good buy!")
		return
	}
}

// acceptor - accept network connection and send conn to the channel for handler
func acceptor(l net.Listener, tohandler chan<- net.Conn) {
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		tohandler <- conn
	}
}

// handleConn - handler for network connection.
// Can interrupt connection by signal
func handleConn(c net.Conn, semafor chan struct{}, wg *sync.WaitGroup) {
	defer c.Close()
	defer wg.Done()
	for {
		select {
		case <-semafor: // triggered when semafor channel have been closed or recieve struct{}{}
			_, err := io.WriteString(c, "Server will interrupt connection through 3sec\n")
			if err != nil {
				return
			}
			time.Sleep(3 * time.Second)
			c.Close()
		default:
			_, err := io.WriteString(c, time.Now().Format("15:04:05\n\r"))
			if err != nil {
				return
			}
			time.Sleep(time.Second)
		}
	}

}
