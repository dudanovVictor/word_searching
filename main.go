package main

import (
	"sync"
)

type Way struct {
	steps map[Coord]bool
	max_r int
	max_c int
}
type Coord struct {
	r int
	c int
}

func NewWay(r int, c int) *Way {
	return &Way{steps: make(map[Coord]bool), max_r: r, max_c: c}
}
func (w *Way) Add(r int, c int, v bool) {
	w.steps[Coord{r: r, c: c}] = v
}
func (w *Way) Exist(r int, c int) bool {
	if r < 0 || c < 0 || r >= w.max_r || c >= w.max_c {
		return true
	}
	v, ok := w.steps[Coord{r: r, c: c}]
	return ok && v
}

func existA(board [][]byte, word string) bool {
	if board == nil || len(board) == 0 || len(board[0]) == 0 || len(word) == 0 {
		return false
	}
	for r := range board {
		for c := range board[r] {
			if board[r][c] == word[0] {
				way := NewWay(len(board), len(board[0]))
				if exist0(r, c, way, board, word) {
					return true
				}
			}
		}
	}
	return false
}

func exist0(r int, c int, way *Way, board [][]byte, word string) bool {
	if len(word) == 0 {
		return true
	}
	if way.Exist(r, c) {
		return false
	}
	if board[r][c] != word[0] {
		return false
	}
	way.Add(r, c, true)
	res := exist0(r, c-1, way, board, word[1:]) ||
		exist0(r, c+1, way, board, word[1:]) ||
		exist0(r-1, c, way, board, word[1:]) ||
		exist0(r+1, c, way, board, word[1:])
	if !res {
		way.Add(r, c, false)
	}
	return res
}

func existB(board [][]byte, word string) bool {
	if board == nil || len(word) == 0 {
		return false
	}
	var wg sync.WaitGroup
	w := len(board[0])
	calls := len(board) * w
	results := make([]bool, calls)
	for r := range board {
		for c := range board[r] {
			if board[r][c] == word[0] {
				wg.Add(1)
				go func(r int, c int) {
					defer wg.Done()
					way := NewWay(len(board), len(board[0]))
					results[int(r)*w+c] = exist0(r, c, way, board, word)
				}(r, c)
			}
		}
	}
	wg.Wait()
	for _, v := range results {
		if v {
			return true
		}
	}
	return false
}
func main() {

}
