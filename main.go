package main 

import (
    "fmt"
    "github.com/rilstrats/tic-tac-toe.go/board"
    "github.com/rilstrats/tic-tac-toe.go/player"
)

func main () {
    fmt.Println("Hello World")

    for i := 0; i < 9; i++ {
        board.DisplayBoard()
        player.TakeTurn()
    }
}

