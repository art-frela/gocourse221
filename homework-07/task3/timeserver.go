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

	sema := make(chan struct{})
	input := ""
	// start canceler server
	go func() {
		for {
			fmt.Printf("For exit type EXIT\n")
			fmt.Scanln(&input)

			if strings.ToUpper(input) == "EXIT" {
				fmt.Println("typed input=", input)
				sema <- struct{}{}
			}
		}
	}()

	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		select {
		case <-sema:
			// recieved shutdown cmd
			wg.Wait()
			fmt.Println("All connection interrupted. Server shutting down. Good buy!")
			return
		default:
			conn, err := listener.Accept()
			if err != nil {
				log.Print(err)
				continue
			}
			wg.Add(1)
			go handleConn(conn, sema, &wg)
		}
	}
}

func handleConn(c net.Conn, semafor chan struct{}, wg *sync.WaitGroup) {
	defer c.Close()
	defer wg.Done()
	for {
		select {
		case <-semafor:
			semafor <- struct{}{} // send that signal to the other clients...

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
			time.Sleep(1 * time.Second)

		}
	}

}
