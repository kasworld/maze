// Copyright 2015 SeukWon Kang (kasworld@gmail.com)
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package maze

import (
	"github.com/kasworld/rand"
)

type Cell struct {
	visited bool
	UpDoor  bool
	LtDoor  bool
}

type Maze struct {
	Xlen, Ylen int
	Cells      [][]Cell
	rnd        *rand.Rand
}

func NewMaze(x, y int) *Maze {
	m := Maze{
		Xlen:  x,
		Ylen:  y,
		Cells: make([][]Cell, x),
		rnd:   rand.New(),
	}
	for i := range m.Cells {
		m.Cells[i] = make([]Cell, y)
	}
	return &m
}

func (m *Maze) MakeMaze() *Maze {
	m.visit(m.rnd.Intn(m.Xlen), m.rnd.Intn(m.Ylen))
	return m
}

func (m *Maze) visit(x, y int) {
	m.Cells[x][y].visited = true
	for _, dir := range m.rnd.Perm(4) {
		switch dir {
		case 0: // up
			if y >= 1 && !m.Cells[x][y-1].visited {
				m.Cells[x][y].UpDoor = true // door open
				m.visit(x, y-1)
			}
		case 1: // down
			if y < m.Ylen-1 && !m.Cells[x][y+1].visited {
				m.Cells[x][y+1].UpDoor = true // door open
				m.visit(x, y+1)
			}
		case 2: // left
			if x >= 1 && !m.Cells[x-1][y].visited {
				m.Cells[x][y].LtDoor = true // door open
				m.visit(x-1, y)
			}
		case 3: // right
			if x < m.Xlen-1 && !m.Cells[x+1][y].visited {
				m.Cells[x+1][y].LtDoor = true // door open
				m.visit(x+1, y)
			}
		}
	}
}
