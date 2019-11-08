package game

import (
	"fmt"
	"fridge/repo"
	"fridge/usecase"
	"log"
)

type service struct {
	user string
	repo repo.Interface
}

func New(in repo.Interface, user string) usecase.Game {
	return &service{repo: in, user: user}
}

func (s *service) Play() {
	if err := s.auth(); err != nil {
		log.Panic(err)
	}

	if err := s.game(); err != nil {
		log.Panic(err)
	}
}

func (s *service) auth() error {
	err := s.repo.Hello(s.user)
	if err != nil {
		return fmt.Errorf("hello send: %w", err)
	}

	return nil
}

func (s *service) game() error {
	for i := 0; ; i++ {
		matrix, err := s.repo.ReadPuzzle()
		if err != nil {
			return fmt.Errorf("[%d] read puzzle: %w", i, err)
		}
		switch len(matrix) {
		case 0:
			return fmt.Errorf("read puzzle. unpredictable behaviour. len 0")
		case 1:
			fmt.Println("Game result: ", string(matrix[0]))
			return nil
		default:
			sol, err := handleMatrix(matrix)
			if err != nil {
				return fmt.Errorf("handle matrix: %w", err)
			}

			fmt.Println("send solution: ", string(sol))

			if err := s.repo.SendSolution(sol); err != nil {
				return fmt.Errorf("send solution: %w", err)
			}
		}
	}
}

func handleMatrix(in [][]byte) (sol []byte, err error) {
	cl, err := parseMatrix(in)
	if err != nil {
		return nil, fmt.Errorf("parse matrix: %w", err)
	}

	for i := range cl {
		fmt.Println(string(in[i]))

		cl[i].rowScore()
	}

	cl.colScore()

	col, row := cl.max()
	res := buildResp(col, row)

	return res, nil
}

func buildResp(col, row int) []byte {
	return []byte(fmt.Sprintf("%d,%d\n", row, col))
}

func parseMatrix(in [][]byte) (cellMatrix, error) {
	score := make(cellMatrix, 0, len(in))

	for i, _ := range in {
		c, err := LoadCells(in[i])
		if err != nil {
			return nil, fmt.Errorf("load row cells: %w", err)
		}

		score = append(score, c)
	}

	return score, nil
}
