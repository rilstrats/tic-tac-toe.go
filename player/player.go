package player

import (
    "fmt"
    "strconv"
)

var player string = "X"

func TakeTurn() (choiceReturn int64) {
    var choice string
    
    for {
        fmt.Println("")
        fmt.Printf("Player %v's Turn, select square: ", player)
        fmt.Scan(&choice)

        if validChoice(choice) {break}
    }

    choiceInt, _ := strconv.ParseInt(choice, 10, 0)
    return choiceInt

}

func validChoice(choice string) (valid bool) {
    choiceInt, choiceErr := strconv.ParseInt(choice, 10, 0)
    
    if choiceErr == nil {
        return false
    }

    if choiceInt <= 1 || choiceInt >= 9 {
        return false
    }

    return true
}
