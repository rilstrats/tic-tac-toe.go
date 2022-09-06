package board

import (
	"fmt"
	"strconv"
)

var Squares [10]string = [10]string {"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

func DisplayBoard () {
    fmt.Println("")
    fmt.Println(Squares[1], "|", Squares[2], "|", Squares[3])
    fmt.Println("-", "+", "-", "+", "-")
    fmt.Println(Squares[4], "|", Squares[5], "|", Squares[6])
    fmt.Println("-", "+", "-", "+", "-")
    fmt.Println(Squares[7], "|", Squares[8], "|", Squares[9])
}

func UpdateBoard (choice string, player string) {
    choiceInt, _ := strconv.ParseInt(choice, 10, 0)
    Squares[choiceInt] = player  
}

func WinCheck (player string) (hasWon bool) {

    // horizontal
    if Squares[1] == player && Squares[2] == player && Squares[3] == player {return true} 
    if Squares[4] == player && Squares[5] == player && Squares[6] == player {return true} 
    if Squares[7] == player && Squares[8] == player && Squares[9] == player {return true} 

    // vertical
    if Squares[1] == player && Squares[4] == player && Squares[7] == player {return true} 
    if Squares[2] == player && Squares[5] == player && Squares[8] == player {return true} 
    if Squares[3] == player && Squares[6] == player && Squares[9] == player {return true} 

    // diagonal
    if Squares[1] == player && Squares[5] == player && Squares[9] == player {return true} 
    if Squares[3] == player && Squares[5] == player && Squares[7] == player {return true} 

    // cat
    return false
}

