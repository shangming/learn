package main

import (
	"log"
	"net"
	"os"
	"strconv"
)

func server() {
	s, err := net.Listen("tcp", ":8888")

	if err != nil {
		log.Panic(err)
	}

	i := 0

	for {

		conn, err := s.Accept()

		if err != nil {
			log.Println(err)
			continue
		}

		i++
		log.Println(i, conn.RemoteAddr())
	}
}

func client(n int) {

	for i := 0; i < n; i++ {
		go func(i int) {
			conn, err := net.Dial("tcp", ":8888")

			if err != nil {
				log.Println(err)
				return
			}

			log.Println("connected", i, conn.LocalAddr())

			select {}

		}(i)
	}
}

func main() {
	if len(os.Args) != 2 {
		log.Panic("args error")
	}

	n, err := strconv.Atoi(os.Args[1])

	if err != nil {
		log.Panic(err)
	}

	if n > 0 {

	} else {
		server()
	}

}
