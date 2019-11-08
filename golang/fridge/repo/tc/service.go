package tc

import (
	"bufio"
	"fmt"
	"fridge/repo"
	"io"
	"log"
	"net"
	"reflect"
	"time"
)

const (
	timeout = time.Second * 10
)

type service struct {
	conn io.ReadWriteCloser
	rw   *bufio.ReadWriter
}

// New ...
// @ntw - network. default: tcp
// @addr - address. default: "84.201.186.199:1337"
func New(ntw, addr string) repo.Interface {
	conn, err := net.Dial("tcp", "84.201.186.199:1337")
	if err != nil {
		log.Fatal(err)
	}

	// i suppose time for win!
	if err = conn.SetDeadline(time.Now().Add(timeout)); err != nil {
		log.Fatal(err)
	}

	return &service{conn: conn, rw: bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))}
}

func (s *service) Close() error {
	if reflect.TypeOf(s) == nil {
		return nil
	}

	return s.conn.Close()
}

func (s *service) Hello(name string)  (err error) {
	res, err := s.rw.ReadString('\n')
	if err != nil {
		return  fmt.Errorf("read request: %w", err)
	}

	fmt.Println("=> Hello: ", res)

	name = name + "\n"

	fmt.Println("<= Hello  ", name)

	if _, err = s.rw.WriteString(name); err != nil {
		return
	}

	if err = s.rw.Flush(); err != nil {
		log.Fatal(err)
	}

	res, err = s.rw.ReadString('\n')
	if err != nil {
		return  fmt.Errorf("read request: %w", err)
	}

	fmt.Println("=> Auth Hello:  ", res)

	return
}

func (s *service) ReadPuzzle() ([][]byte, error) {
	res := make([][]byte, 0)
	scanner := bufio.NewScanner(s.conn)
	for scanner.Scan() {
		in := scanner.Bytes()
		if len(in) == 0 {
			break
		}
		res = append(res, in)
	}

	return res, nil
}

func (s *service) SendSolution(res []byte) (err error) {
	defer func() {
		if err == nil {
			err = s.rw.Flush()
		}
	}()

	_, err = s.rw.Write(res)

	return
}
