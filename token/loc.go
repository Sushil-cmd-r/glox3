package token

import "fmt"

type Location struct {
	FilePath string
	Row, Col int
}

func (loc Location) String() string {
	return fmt.Sprintf("%s:%d:%d", loc.FilePath, loc.Row, loc.Col)
}

type Loc int

type File struct {
	path   string
	nlines []int // start of new line
}

func NewFile(path string) *File {
	nlines := make([]int, 0)
	nlines = append(nlines, 0)
	return &File{
		path:   path,
		nlines: nlines,
	}
}

func (f *File) AddLine(offset int) {
	f.nlines = append(f.nlines, offset)
}

func (f *File) LocationFor(loc Loc) Location {
	i, j := 0, len(f.nlines)
	for i < j {
		h := int(uint(i+j) >> 1) // avoid overflow when computing h
		if f.nlines[h] <= int(loc) {
			i = h + 1
		} else {
			j = h
		}
	}

	row := i
	col := int(loc) - f.nlines[i-1] + 1

	return Location{FilePath: f.path, Row: row, Col: col}
}
