package board

import (
    "fmt"
)

var squares [10]string = [10]string {"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

func DisplayBoard () {
    fmt.Println("")
    fmt.Println(squares[1], "|", squares[2], "|", squares[3])
    fmt.Println("-", "+", "-", "+", "-")
    fmt.Println(squares[4], "|", squares[5], "|", squares[6])
    fmt.Println("-", "+", "-", "+", "-")
    fmt.Println(squares[7], "|", squares[8], "|", squares[9])
}
