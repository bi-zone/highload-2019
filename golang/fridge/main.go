package main

import (
	"bufio"
	"log"
	"net"
	"time"
)

const (
	name     = "Шеф\n"
	solution = "1,2\n"
)

func main() {
	conn, err := net.Dial("tcp", "84.201.186.199:1337")
	if err != nil {
		log.Fatal(err)
	}
	if err = conn.SetDeadline(time.Now().Add(10 * time.Second)); err != nil {
		log.Fatal(err)
	}

	rw := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))

	var text string
	if text, err = rw.ReadString('\n'); err != nil {
		log.Fatal(err)
	}
	log.Print(text)

	log.Print(name)
	if _, err = rw.WriteString(name); err != nil {
		log.Fatal(err)
	}
	rw.Flush()

	readPuzzle(rw.Reader)

	log.Print(solution)
	if _, err = rw.WriteString(solution); err != nil {
		log.Fatal(err)
	}
	rw.Flush()

	readPuzzle(rw.Reader)

	if text, err = rw.ReadString('\n'); err != nil {
		log.Fatal(err)
	}
	log.Print("Flag: " + text)
}

func readPuzzle(r *bufio.Reader) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			break
		}
		log.Println(text)
	}
}
