package main

import (
	"bufio"
	"fridge/repo/tc"
	"fridge/usecase/game"
	"io"
	"log"
)

const (
	name     = "Шеф"
	solution = "1,2"
	netw     = "tcp"
	addr     = "84.201.186.199:1337"
)

func main() {
	r := tc.New(netw, addr)
	defer r.Close()

	g := game.New(r, name)
	g.Play()

	/*hello(rw)
	readPuzzle(rw.Reader)
	sendSolution(rw)
	readPuzzle(rw.Reader)
	printFlag(rw.Reader)*/
}

func hello(rw *bufio.ReadWriter) {
	text, err := rw.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	log.Print(text)

	name := name + "\n"
	log.Print(name)
	if _, err = rw.WriteString(name); err != nil {
		log.Fatal(err)
	}

	if err = rw.Flush(); err != nil {
		log.Fatal(err)
	}
}

func readPuzzle(r io.Reader) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			break
		}
		log.Println(text)
	}
}

func sendSolution(rw *bufio.ReadWriter) {
	solution := solution + "\n"
	log.Print(solution)

	if _, err := rw.WriteString(solution); err != nil {
		log.Fatal(err)
	}

	if err := rw.Flush(); err != nil {
		log.Fatal(err)
	}
}

func printFlag(r *bufio.Reader) {
	text, err := r.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Flag: " + text)
}
