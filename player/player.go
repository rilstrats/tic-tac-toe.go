package player

import (
	"fmt"
	"strconv"
)

var Player string = "X"

func TakeTurn(squares [10]string) (choice string) {

    for {
        fmt.Println("")
        fmt.Printf("Player %v's Turn, select square: ", Player)
        fmt.Scan(&choice)

        if validateChoice(choice, squares) {break}
        fmt.Println("Please select a number between 1 and 9 that hasn't already been chosen.")
    }

    return choice
}

func validateChoice(choice string, squares [10]string) (isValid bool) {

    choiceInt, choiceErr := strconv.ParseInt(choice, 10, 0)
    // fmt.Println("Choice:", choiceInt, choiceErr)
    if choiceErr != nil || choiceInt < 1 || choiceInt > 9 {return false}

    squareInt, squareErr := strconv.ParseInt(squares[choiceInt], 10, 0)
    // fmt.Println("Square:", squareInt, squareErr)
    if choiceInt != squareInt || squareErr != nil {return false}

    return true
}

func ChangePlayer() {

    if Player == "X" {
        Player = "O"
    } else {
        Player = "X"
    }
}
