package game

import (
	"fmt"
)

type cellValue byte

const (
	A   cellValue = '+'
	B   cellValue = 'x'
	Sep cellValue = ' '
)

type cell struct {
	V     cellValue
	score int
}

type cells []*cell
type cellMatrix []cells

func Cell(in byte) (*cell, error) {
	switch v := cellValue(in); v {
	case A, B, Sep:
		return &cell{V: v, score: 0}, nil
	}
	return nil, fmt.Errorf("cell not supported char: %[1]X / %[1]d /  %[2]s ", in, string(in))
}

func LoadCells(in []byte) (cells, error) {
	res := make(cells, 0, len(in))
	for i := range in {
		c, err := Cell(in[i])
		if err != nil {
			return nil, fmt.Errorf("load [%s]cell: %w", string(in), err)
		}

		if c.V == Sep {
			continue
		}

		res = append(res, c)
	}

	return res, nil
}

func (c *cell) AddScore(v int) {
	c.score += v
}

func (c *cells) rowScore() {
	for i, consumer := range *c {
		for j, affects := range *c {
			if i == j {
				continue
			}

			consumer.giveScore(affects)
		}
	}
}

func (c cellMatrix) colScore() {
A:
	for col := 0; ; col++ {
	B:
		for col2 := 0; ; col2++ {
			if col == col2 {
				continue
			}

			for row, _ := range c {
				if col >= len(c[row]) {
					break A
				}

				if col2 >= len(c[row]) {
					break B
				}

				c[col][row].giveScore(c[col2][row])
			}
		}
	}
}

func (c cellMatrix) max() (cl, rw int) {
	var max *cell

	for col := range c {
		for row := range c[col] {
			if max == nil || max.score < c[col][row].score {
				max = c[col][row]
				cl = col
				rw = row
			}
		}
	}

	return cl, rw
}

func (c *cell) giveScore(affects *cell) {
	switch affects.V {
	case A:
		c.AddScore(1)
	case B:
		c.AddScore(-1)
	}
}
