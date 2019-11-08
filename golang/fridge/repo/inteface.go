package repo

type Interface interface {
	Close() error

	Hello(name string) ( error)
	ReadPuzzle() ([][]byte, error)
	SendSolution([]byte) error
}
