package main 

import (
    "fmt"
    "github.com/rilstrats/tic-tac-toe.go/board"
    "github.com/rilstrats/tic-tac-toe.go/player"
)

func main () {

    for i := 0; i < 9; i++ {
        
        board.DisplayBoard()
        choice := player.TakeTurn(board.Squares)
        board.UpdateBoard(choice, player.Player)
        
        if board.WinCheck(player.Player) {

            board.DisplayBoard()
            fmt.Println("")
            fmt.Println("Player", player.Player, "wins!")

            return
        }

        player.ChangePlayer()
    }

    fmt.Println("Cat. No one wins.")
}

