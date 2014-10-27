package sudoku

import (
	"os"
	"fmt"
	"math"
)

type Board struct {
 size  int
 cells []*Cell
}

type Cell struct {
 // Save attributes of the cell
 row, col int
 parent   *Board
 // Allow querying and controlling of the cell
 query   chan CellCommand
 control chan CellCommand
}

type CellCommand struct {
 typ string
 arg int
 ret chan bool
 res chan int
}

func NewBoard(size int) *Board {
	 sq := int(math.Sqrt(float64(size)))
	 if sq*sq != size {
	  return nil
	 }

 b := new(Board)
 b.size = size
 b.cells = make([]*Cell, size*size)

 for r := 0; r < size; r++ {
  for c := 0; c < size; c++ {
   b.cells[r*size+c] = b.newCell(r, c)
  }
 }

 return b
}

func (b *Board) newCell(row, col int) *Cell {
	 c := new(Cell)
	 c.row, c.col = row, col
	 c.parent = b
	 c.query = make(chan CellCommand)   // for synchronous commands
	 c.control = make(chan CellCommand) // for reentrant commands
	 go c.process_queries()
	 go c.process_control()
	 return c
}

func (c *Cell) process_queries() {
	 // grab the parent because
	 b := c.parent
	 // Store state
	 charval := '.'
	 possible := b.size
	 could_be := make([]bool, b.size)
	 for i := 0; i < b.size; i++ {
	  could_be[i] = true
	 }
 for cmd := range c.query {
  switch cmd.typ {
  case "set_is":
   if cmd.arg < 0 || cmd.arg >= b.size {
    cmd.ret <- false
    continue
   }
   possible = 1
   charval = '1' + cmd.arg
   for i := 0; i < b.size; i++ {
    could_be[i] = cmd.arg == i
   }
   cmd.ret <- true
  case "set_not":
   if cmd.arg < 0 || cmd.arg >= b.size || !could_be[cmd.arg] {
    cmd.ret <- false
    continue
   }
   could_be[cmd.arg] = false
   possible--
   if possible == 1 {
    charval = '1'
    for i := 0; i < b.size; i++ {
     if could_be[i] {
      charval += i
      break
     }
    }
   } else {
    charval = '.'
   }
   cmd.ret <- true
  case "get_could_be":
   if cmd.arg < 0 || cmd.arg >= b.size {
    cmd.ret <- false
    continue
   }
   cmd.ret <- could_be[cmd.arg]
  case "get_char":
   cmd.res <- charval
  case "get_value":
   if charval == '.' {
    cmd.res <- -1
    continue
   }
   cmd.res <- charval-'1'
  default:
   fmt.Fprintln(os.Stderr, "[Unknown Query: " + cmd.typ + "]")
   if cmd.res != nil {
    cmd.res <- 0
   }
   if cmd.ret != nil {
    cmd.ret <- false
   }
  }
 }
}

func (c *Cell) process_control() {
	 for cmd := range c.control {
	  go func(cmd CellCommand) {
	   // Make channels
	   result := make(chan int)
	   status := make(chan bool)
	   // Process the command
	   switch cmd.typ {
	   case "elim":
	    c.query <- CellCommand{typ: "set_not", arg: cmd.arg, ret: status}
    if <-status { // status is true if the value was changed
     c.query <- CellCommand{typ: "get_value", res: result}
     val := <-result
     if val != -1 {
      c.set(val)
     }
    }
    cmd.ret <- true
   case "set":
    c.query <- CellCommand{typ: "set_is", arg: cmd.arg, ret: status}
    if <-status {
     c.set(cmd.arg)
    }
    cmd.ret <- true
   default:
    fmt.Println(os.Stderr, "[Unknown command: " + cmd.typ + "]")
    if cmd.res != nil {
     cmd.res <- 0
    }
    if cmd.ret != nil {
     cmd.ret <- false
    }
   }
  }(cmd)
 }
}

func (c *Cell) set(s int) {
	 b := c.parent
	 sq := int(math.Sqrt(float64(b.size)))
	 rbase := c.row / sq * sq
	 cbase := c.col / sq * sq
	 cmd := CellCommand{typ: "elim", arg: s, ret: make(chan bool)}
 for i := 0; i < b.size; i++ {
  // row
  if i != c.col {
   b.cells[c.row*b.size+i].control <- cmd
   <-cmd.ret
  }
  // col
  if i != c.row {
   b.cells[i*b.size+c.col].control <- cmd
   <-cmd.ret
  }
  // block
  rset := rbase + i/sq
  cset := cbase + i%sq
  if rset != c.row || cset != c.col {
   b.cells[rset*b.size+cset].control <- cmd
   <-cmd.ret
  }
 }
}

// Get the board cell at r,c (indexed upper left, starting at one)
func (b *Board) Set(r, c, s int) {
	 r--; c--; s--
	 cmd := CellCommand{typ: "set", arg: s, ret: make(chan bool)}
 b.cells[r*b.size+c].control <- cmd
 <-cmd.ret
}

func (c *Cell) Value() int {
	 cmd := CellCommand{typ: "get_char", res: make(chan int)}
 c.query <- cmd
 return <-cmd.res
}

func (b *Board) Print() {
	 cmd := CellCommand{typ: "get_char", res: make(chan int)}
 for r := 0; r < b.size; r++ {
  for c := 0; c < b.size; c++ {
   b.cells[r*b.size+c].query <- cmd
   fmt.Printf("%c", <-cmd.res)
   if c%3 == 2 {
    fmt.Print(" ")
   }
  }
  fmt.Println("")
  if r%3 == 2 {
   fmt.Println("")
  }
 }
}

func (b *Board) Solved() bool {
	 cmd := CellCommand{typ: "get_char", res: make(chan int)}
 for r := 0; r < b.size; r++ {
  for c := 0; c < b.size; c++ {
   b.cells[r*b.size+c].query <- cmd
   if <-cmd.res == '.' {
    return false
   }
  }
 }
 return true
}

func (b *Board) PrintAll() {
	 cmd := CellCommand{typ: "get_could_be", ret: make(chan bool)}
 for r := 0; r < b.size; r++ {
  for c := 0; c < b.size; c++ {
   for cmd.arg = 0; cmd.arg < b.size; cmd.arg++ {
    b.cells[r*b.size+c].query <- cmd
    if <-cmd.ret {
     fmt.Print(".")
    } else {
     fmt.Print(cmd.arg + 1)
    }
   }
   fmt.Print(" ")
  }
  fmt.Println("")
 }
 fmt.Println("")
}
