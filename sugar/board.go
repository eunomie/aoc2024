package sugar

import (
	"slices"
	"strings"
)

type (
	Board struct {
		input     []rune
		diagonals string
		nbLines   int
		nbCols    int
		lenInput  int
	}

	Rune struct {
		board *Board
		rune  rune
		pos   int
		line  int
		col   int
	}
)

func NewBoard(input string) *Board {
	nbCols := strings.Index(input, "\n")
	nbLines := strings.Count(input, "\n") + 1

	return &Board{
		input:    []rune(input),
		nbLines:  nbLines,
		nbCols:   nbCols,
		lenInput: len(input),
	}
}

func (b *Board) IsolateWithRune(positions []int, r rune) *Board {
	newRunes := make([]rune, len(b.input))
	for i, r := range b.input {
		if r == '\n' {
			newRunes[i] = r
		} else {
			newRunes[i] = '.'
		}
	}
	for _, pos := range positions {
		newRunes[pos] = r
	}

	return NewBoard(string(newRunes))
}

func (b *Board) emptyRune() Rune {
	return Rune{
		board: b,
		pos:   -1,
		line:  -1,
		col:   -1,
	}
}

func (b *Board) newRune(pos int) Rune {
	if pos < 0 || pos >= b.lenInput {
		return b.emptyRune()
	}
	return Rune{
		board: b,
		pos:   pos,
		rune:  b.input[pos],
		col:   pos % (b.nbCols + 1),
		line:  pos / (b.nbCols + 1),
	}
}

func (b *Board) Input() string {
	return string(b.input)
}

func (b *Board) String() string {
	return string(b.input)
}

func (b *Board) NumLines() int {
	return b.nbLines
}

func (b *Board) NumCols() int {
	return b.nbCols
}

func (b *Board) Get(col, line int) string {
	return b.GetRune(col, line).String()
}

func (b *Board) GetRune(col, line int) Rune {
	if col < 0 || col >= b.nbCols || line < 0 || line >= b.nbLines {
		return b.emptyRune()
	}
	pos := line*(b.nbCols+1) + col
	return b.newRune(pos)
}

func (b *Board) GetRuneAt(pos int) Rune {
	return b.newRune(pos)
}

func (b *Board) GetLine(line int) string {
	return string(b.input[line*(b.nbCols+1) : (line+1)*b.nbCols])
}

func (b *Board) ForEachRunes(fn func(Rune), runes ...rune) {
	all := len(runes) == 0
	for i, c := range b.input {
		if c == '\n' {
			continue
		}
		if all || slices.Contains(runes, c) {
			fn(b.newRune(i))
		}
	}
}

func (b *Board) ForEachRunesNot(fn func(Rune), runes ...rune) {
	for i, c := range b.input {
		if c == '\n' {
			continue
		}
		if !slices.Contains(runes, c) {
			fn(b.newRune(i))
		}
	}
}

func (b *Board) Find(r rune) Rune {
	for i, c := range b.input {
		if c == r {
			return b.newRune(i)
		}
	}
	return b.emptyRune()
}

func (b *Board) CountNonOverlapping(s string) int {
	return strings.Count(b.Input(), s)
}

func (b *Board) CountNonOverlappingInDiagonal(s string) int {
	return strings.Count(b.Diagonals(), s)
}

func (b *Board) Pivot() *Board {
	p := PivotString(b.String())
	return NewBoard(p)
}

func (b *Board) Diagonals() string {
	if b.diagonals != "" {
		return b.diagonals
	}
	b.diagonals = Diagonals(b.String())
	return b.diagonals
}

func (b *Board) Write(i int, n rune) {
	b.input[i] = n
}

func (b *Board) At(pos int) Rune {
	return b.newRune(pos)
}

func (r Rune) String() string {
	if r.pos == -1 {
		return ""
	}
	return string(r.rune)
}

func (r Rune) Rune() rune {
	return r.rune
}

func (r Rune) Is(c ...rune) bool {
	if r.IsEmpty() {
		return false
	}
	return slices.Contains(c, r.rune)
}

func (r Rune) IsNot(c ...rune) bool {
	if r.IsEmpty() {
		return true
	}
	return !slices.Contains(c, r.rune)
}

func (r Rune) Move(col, line int) Rune {
	return r.board.GetRune(r.col+col, r.line+line)
}

func (r Rune) Diff(o Rune) (int, int) {
	return o.col - r.col, o.line - r.line
}

func (r Rune) Top() Rune {
	return r.board.GetRune(r.col, r.line-1)
}

func (r Rune) TopRight() Rune {
	return r.board.GetRune(r.col+1, r.line-1)
}

func (r Rune) Right() Rune {
	return r.board.GetRune(r.col+1, r.line)
}

func (r Rune) BottomRight() Rune {
	return r.board.GetRune(r.col+1, r.line+1)
}

func (r Rune) Bottom() Rune {
	return r.board.GetRune(r.col, r.line+1)
}

func (r Rune) BottomLeft() Rune {
	return r.board.GetRune(r.col-1, r.line+1)
}

func (r Rune) Left() Rune {
	return r.board.GetRune(r.col-1, r.line)
}

func (r Rune) TopLeft() Rune {
	return r.board.GetRune(r.col-1, r.line-1)
}

func (r Rune) Index() int {
	return r.pos
}

func (r Rune) Write(n rune) {
	r.board.input[r.pos] = n
}

func (r Rune) IsEmpty() bool {
	return r.pos == -1
}

func (r Rune) Col() int {
	return r.col
}

func (r Rune) Line() int {
	return r.line
}
